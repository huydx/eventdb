package fs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/xichen2020/eventdb/digest"
	"github.com/xichen2020/eventdb/document/field"
	"github.com/xichen2020/eventdb/generated/proto/infopb"
	"github.com/xichen2020/eventdb/index"
	indexfield "github.com/xichen2020/eventdb/index/field"
	"github.com/xichen2020/eventdb/index/segment"
	"github.com/xichen2020/eventdb/persist/schema"
	"github.com/xichen2020/eventdb/values"
	"github.com/xichen2020/eventdb/values/encoding"
	xbytes "github.com/xichen2020/eventdb/x/bytes"
)

// segmentWriter is responsible for writing segments to filesystem.
type segmentWriter interface {
	// Start starts persisting a segment.
	Start(opts writerStartOptions) error

	// Finish finishes persisting a segment and performs cleanups as necessary.
	Finish() error

	// WriteFields writes a list of segment fields.
	WriteFields(fields []indexfield.DocsField) error
}

// writerStartOptions provide a set of options for opening a writer.
type writerStartOptions struct {
	Namespace   []byte
	SegmentMeta segment.Metadata
}

type writer struct {
	filePathPrefix     string
	newFileMode        os.FileMode
	newDirectoryMode   os.FileMode
	fieldPathSeparator string
	timestampPrecision time.Duration

	fdWithDigestWriter digest.FdWithDigestWriter
	info               *infopb.SegmentInfo
	segmentDir         string
	numDocuments       int32
	buf                []byte
	bytesBuf           bytes.Buffer

	bw     encoding.BoolEncoder
	iw     encoding.IntEncoder
	dw     encoding.DoubleEncoder
	sw     encoding.BytesEncoder
	tw     encoding.TimeEncoder
	values valuesUnion

	err error
}

// newSegmentWriter creates a new segment writer.
// TODO(xichen): Add encoding hints when encoding raw docs.
// TODO(xichen): Investigate the benefit of writing a single field file.
func newSegmentWriter(opts *Options) segmentWriter {
	w := &writer{
		filePathPrefix:     opts.FilePathPrefix(),
		newFileMode:        opts.NewFileMode(),
		newDirectoryMode:   opts.NewDirectoryMode(),
		fieldPathSeparator: string(opts.FieldPathSeparator()),
		timestampPrecision: opts.TimestampPrecision(),
		fdWithDigestWriter: digest.NewFdWithDigestWriter(opts.WriteBufferSize()),
		info:               &infopb.SegmentInfo{},

		bw: encoding.NewBoolEncoder(),
		iw: encoding.NewIntEncoder(),
		dw: encoding.NewDoubleEncoder(),
		sw: encoding.NewBytesEncoder(),
		tw: encoding.NewTimeEncoder(),
	}
	return w
}

func (w *writer) Start(opts writerStartOptions) error {
	var (
		namespace   = opts.Namespace
		segmentMeta = opts.SegmentMeta
	)

	segmentDir := segmentDirPath(w.filePathPrefix, namespace, segmentMeta)
	if err := os.MkdirAll(segmentDir, w.newDirectoryMode); err != nil {
		return err
	}
	w.segmentDir = segmentDir
	w.numDocuments = segmentMeta.NumDocs
	w.err = nil

	w.info.Reset()
	w.info.Version = schema.SegmentVersion
	w.info.MinTimestampNanos = segmentMeta.MinTimeNanos
	w.info.MaxTimestampNanos = segmentMeta.MaxTimeNanos
	w.info.NumDocuments = segmentMeta.NumDocs
	return w.writeInfoFile(segmentDir, w.info)
}

func (w *writer) WriteFields(fields []indexfield.DocsField) error {
	for _, field := range fields {
		if err := w.writeField(field); err != nil {
			return err
		}
	}
	return nil
}

func (w *writer) Finish() error {
	if w.err != nil {
		return w.err
	}
	w.err = w.writeCheckpointFile(w.segmentDir)
	return w.err
}

func (w *writer) writeInfoFile(
	segmentDir string,
	info *infopb.SegmentInfo,
) error {
	path := infoFilePath(segmentDir)
	f, err := w.openWritable(path)
	if err != nil {
		return err
	}
	w.fdWithDigestWriter.Reset(f)
	defer w.fdWithDigestWriter.Close()

	msgSize := info.Size()
	payloadSize := maxMessageSizeInBytes + msgSize
	w.ensureBufferSize(payloadSize)
	size := binary.PutVarint(w.buf, int64(msgSize))
	n, err := info.MarshalTo(w.buf[size:])
	if err != nil {
		return err
	}
	size += n
	_, err = w.fdWithDigestWriter.Write(w.buf[:size])
	if err != nil {
		return err
	}
	return w.fdWithDigestWriter.Flush()
}

func (w *writer) writeField(df indexfield.DocsField) error {
	path := df.Metadata().FieldPath

	// Write null values.
	if nullField, exists := df.NullField(); exists {
		docIDSet := nullField.DocIDSet()
		w.values.valueType = field.NullType
		if err := w.writeFieldDataFile(w.segmentDir, path, docIDSet, w.values); err != nil {
			return err
		}
	}

	// Write boolean values.
	if boolField, exists := df.BoolField(); exists {
		docIDSet := boolField.DocIDSet()
		w.values.valueType = field.BoolType
		w.values.boolValues = boolField.Values()
		if err := w.writeFieldDataFile(w.segmentDir, path, docIDSet, w.values); err != nil {
			return err
		}
	}

	// Write int values.
	if intField, exists := df.IntField(); exists {
		docIDSet := intField.DocIDSet()
		w.values.valueType = field.IntType
		w.values.intValues = intField.Values()
		if err := w.writeFieldDataFile(w.segmentDir, path, docIDSet, w.values); err != nil {
			return err
		}
	}

	// Write double values.
	if doubleField, exists := df.DoubleField(); exists {
		docIDSet := doubleField.DocIDSet()
		w.values.valueType = field.DoubleType
		w.values.doubleValues = doubleField.Values()
		if err := w.writeFieldDataFile(w.segmentDir, path, docIDSet, w.values); err != nil {
			return err
		}
	}

	// Write string values.
	if bytesField, exists := df.BytesField(); exists {
		docIDSet := bytesField.DocIDSet()
		w.values.valueType = field.BytesType
		w.values.bytesValues = bytesField.Values()
		if err := w.writeFieldDataFile(w.segmentDir, path, docIDSet, w.values); err != nil {
			return err
		}
	}

	// Write time values.
	if timeField, exists := df.TimeField(); exists {
		docIDSet := timeField.DocIDSet()
		w.values.valueType = field.TimeType
		w.values.timeValues = timeField.Values()
		if err := w.writeFieldDataFile(w.segmentDir, path, docIDSet, w.values); err != nil {
			return err
		}
	}

	return nil
}

func (w *writer) writeFieldDataFile(
	segmentDir string,
	fieldPath []string,
	docIDSet index.DocIDSet,
	values valuesUnion,
) error {
	if w.err != nil {
		return w.err
	}
	path := fieldDataFilePath(segmentDir, fieldPath, values.valueType, w.fieldPathSeparator, &w.bytesBuf)

	f, err := w.openWritable(path)
	if err != nil {
		return err
	}
	w.fdWithDigestWriter.Reset(f)
	defer w.fdWithDigestWriter.Close()

	// Write header.
	_, err = w.fdWithDigestWriter.Write(magicHeader)
	if err != nil {
		return err
	}
	if err = w.writeDocIDSet(w.fdWithDigestWriter, docIDSet); err != nil {
		return err
	}
	if err = w.writeValues(w.fdWithDigestWriter, values); err != nil {
		return err
	}

	return w.fdWithDigestWriter.Flush()
}

func (w *writer) writeDocIDSet(
	writer digest.FdWithDigestWriter,
	docIDSet index.DocIDSet,
) error {
	return docIDSet.WriteTo(writer, &w.bytesBuf)
}

func (w *writer) writeValues(
	writer digest.FdWithDigestWriter,
	values valuesUnion,
) error {
	switch values.valueType {
	case field.NullType:
		return nil
	case field.BoolType:
		return w.bw.Encode(values.boolValues, writer)
	case field.IntType:
		return w.iw.Encode(values.intValues, writer)
	case field.DoubleType:
		return w.dw.Encode(values.doubleValues, writer)
	case field.BytesType:
		return w.sw.Encode(values.bytesValues, writer)
	case field.TimeType:
		return w.tw.Encode(values.timeValues, writer, encoding.EncodeTimeOptions{Resolution: w.timestampPrecision})
	default:
		return fmt.Errorf("unknown value type: %v", values.valueType)
	}
}

func (w *writer) writeCheckpointFile(segmentDir string) error {
	path := checkpointFilePath(segmentDir)
	f, err := w.openWritable(path)
	if err != nil {
		return err
	}
	return f.Close()
}

func (w *writer) openWritable(filePath string) (*os.File, error) {
	return openWritable(filePath, w.newFileMode)
}

func (w *writer) ensureBufferSize(targetSize int) {
	w.buf = xbytes.EnsureBufferSize(w.buf, targetSize, xbytes.DontCopyData)
}

type valuesUnion struct {
	valueType    field.ValueType
	boolValues   values.BoolValues
	intValues    values.IntValues
	doubleValues values.DoubleValues
	bytesValues  values.BytesValues
	timeValues   values.TimeValues
}
