package fs

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"

	"github.com/xichen2020/eventdb/digest"
	"github.com/xichen2020/eventdb/encoding"
	"github.com/xichen2020/eventdb/event/field"
	"github.com/xichen2020/eventdb/generated/proto/infopb"
	"github.com/xichen2020/eventdb/persist/schema"
	xbytes "github.com/xichen2020/eventdb/x/bytes"

	"github.com/pilosa/pilosa/roaring"
)

// segmentWriter is responsible for writing segments to filesystem.
// TODO(xichen): Make docIDs an interface with `WriteTo()` to abstract the details
// of how the doc IDs are encoded.
// TODO(xichen): Encapsulate type-specific write functions for writing compound files.
type segmentWriter interface {
	// Open opens the writer.
	Open(opts writerOpenOptions) error

	// WriteNullField writes a field with docID set and null values.
	WriteNullField(fieldPath []string, docIDs *roaring.Bitmap) error

	// WriteBoolField writes a field with docID set and bool values.
	WriteBoolField(fieldPath []string, docIDs *roaring.Bitmap, vals []bool) error

	// WriteIntField writes a field with docID set and int values.
	WriteIntField(fieldPath []string, docIDs *roaring.Bitmap, vals []int) error

	// WriteDoubleField writes a field with docID set and double values.
	WriteDoubleField(fieldPath []string, docIDs *roaring.Bitmap, vals []float64) error

	// WriteStringField writes a field with docID set and string values.
	WriteStringField(fieldPath []string, docIDs *roaring.Bitmap, vals []string) error

	// Close closes the writer.
	Close() error
}

var (
	errEmptyDocIDSet = errors.New("empty document ID set")
)

// writerOpenOptions provide a set of options for opening a writer.
type writerOpenOptions struct {
	Namespace    []byte
	Shard        uint32
	SegmentID    string
	MinTimeNanos int64
	MaxTimeNanos int64
	NumDocuments int32
}

type writer struct {
	filePathPrefix     string
	newFileMode        os.FileMode
	newDirectoryMode   os.FileMode
	fieldPathSeparator string

	fdWithDigestWriter digest.FdWithDigestWriter
	info               *infopb.SegmentInfo
	segmentDir         string
	buf                []byte
	bytesBuf           bytes.Buffer

	bw encoding.BoolEncoder
	iw encoding.IntEncoder
	dw encoding.DoubleEncoder
	sw encoding.StringEncoder

	boolIt   encoding.RewindableBoolIterator
	intIt    encoding.RewindableIntIterator
	doubleIt encoding.RewindableDoubleIterator
	stringIt encoding.RewindableStringIterator
	valueIt  valueIteratorUnion

	err error
}

// newSegmentWriter creates a new segment writer.
// TODO(xichen): Initialize the type-specific encoders.
// TODO(xichen): Encode timestamp field and source field differently.
// TODO(xichen): Encode timestamp with configurable precision.
// TODO(xichen): Encode ALL doc IDs differently for values.
// TODO(xichen): Investigate the benefit of writing a single field file.
func newSegmentWriter(opts *Options) segmentWriter {
	w := &writer{
		filePathPrefix:     opts.FilePathPrefix(),
		newFileMode:        opts.NewFileMode(),
		newDirectoryMode:   opts.NewDirectoryMode(),
		fieldPathSeparator: string(opts.FieldPathSeparator()),
		fdWithDigestWriter: digest.NewFdWithDigestWriter(opts.WriteBufferSize()),
		info:               &infopb.SegmentInfo{},
		boolIt:             encoding.NewArrayBasedBoolIterator(nil),
		intIt:              encoding.NewArrayBasedIntIterator(nil),
		doubleIt:           encoding.NewArrayBasedDoubleIterator(nil),
		stringIt:           encoding.NewArrayBasedStringIterator(nil),
	}
	w.valueIt = valueIteratorUnion{
		boolIt:   w.boolIt,
		intIt:    w.intIt,
		doubleIt: w.doubleIt,
		stringIt: w.stringIt,
	}
	return w
}

func (w *writer) Open(opts writerOpenOptions) error {
	var (
		namespace    = opts.Namespace
		shard        = opts.Shard
		minTimeNanos = opts.MinTimeNanos
		maxTimeNanos = opts.MaxTimeNanos
	)

	shardDir := shardDataDirPath(w.filePathPrefix, namespace, shard)
	segmentDir := segmentDirPath(shardDir, minTimeNanos, maxTimeNanos, opts.SegmentID)
	if err := os.MkdirAll(segmentDir, w.newDirectoryMode); err != nil {
		return err
	}
	w.segmentDir = segmentDir
	w.err = nil

	w.info.Reset()
	w.info.Version = schema.SegmentVersion
	w.info.MinTimestampNanos = opts.MinTimeNanos
	w.info.MaxTimestampNanos = opts.MaxTimeNanos
	w.info.NumDocuments = opts.NumDocuments
	return w.writeInfoFile()
}

func (w *writer) WriteNullField(fieldPath []string, docIDs *roaring.Bitmap) error {
	w.valueIt.valueType = field.NullType
	return w.writeFieldDataFile(fieldPath, docIDs, w.valueIt)
}

func (w *writer) WriteBoolField(fieldPath []string, docIDs *roaring.Bitmap, vals []bool) error {
	w.boolIt.Reset(vals)
	w.valueIt.valueType = field.BoolType
	return w.writeFieldDataFile(fieldPath, docIDs, w.valueIt)
}

func (w *writer) WriteIntField(fieldPath []string, docIDs *roaring.Bitmap, vals []int) error {
	w.intIt.Reset(vals)
	w.valueIt.valueType = field.IntType
	return w.writeFieldDataFile(fieldPath, docIDs, w.valueIt)
}

func (w *writer) WriteDoubleField(fieldPath []string, docIDs *roaring.Bitmap, vals []float64) error {
	w.doubleIt.Reset(vals)
	w.valueIt.valueType = field.DoubleType
	return w.writeFieldDataFile(fieldPath, docIDs, w.valueIt)
}

func (w *writer) WriteStringField(fieldPath []string, docIDs *roaring.Bitmap, vals []string) error {
	w.stringIt.Reset(vals)
	w.valueIt.valueType = field.StringType
	return w.writeFieldDataFile(fieldPath, docIDs, w.valueIt)
}

func (w *writer) Close() error {
	if w.err != nil {
		return w.err
	}
	// NB(xichen): only write out the checkpoint file if there are no errors
	// encountered between calling writer.Open() and writer.Close().
	if err := w.writeCheckpointFile(); err != nil {
		w.err = err
		return err
	}
	return nil
}

func (w *writer) writeInfoFile() error {
	path := infoFilePath(w.segmentDir)
	f, err := w.openWritable(path)
	if err != nil {
		return err
	}
	w.fdWithDigestWriter.Reset(f)
	defer w.fdWithDigestWriter.Close()

	msgSize := w.info.Size()
	payloadSize := maxMessageSizeInBytes + msgSize
	w.ensureBufferSize(payloadSize)
	size := binary.PutVarint(w.buf, int64(msgSize))
	n, err := w.info.MarshalTo(w.buf[size:])
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

func (w *writer) writeFieldDataFile(
	fieldPath []string,
	docIDs *roaring.Bitmap,
	valueIt valueIteratorUnion,
) error {
	if w.err != nil {
		return w.err
	}
	w.err = w.writeFieldDataFileInternal(fieldPath, docIDs, valueIt)
	return w.err
}

func (w *writer) writeFieldDataFileInternal(
	fieldPath []string,
	docIDs *roaring.Bitmap,
	valueIt valueIteratorUnion,
) error {
	if docIDs.Count() == 0 {
		return errEmptyDocIDSet
	}
	path := fieldDataFilePath(w.segmentDir, fieldPath, w.fieldPathSeparator, &w.bytesBuf)
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

	// Write doc ID set.
	// TODO(xichen): Precompute the size of the encoded bitmap to avoid the memory
	// cost of writing the encoded bitmap to a byte buffer then to file.
	w.bytesBuf.Reset()
	_, err = docIDs.WriteTo(&w.bytesBuf)
	if err != nil {
		return err
	}
	w.ensureBufferSize(maxMessageSizeInBytes)
	size := binary.PutVarint(w.buf, int64(w.bytesBuf.Len()))
	_, err = w.fdWithDigestWriter.Write(w.buf[:size])
	if err != nil {
		return err
	}
	_, err = w.fdWithDigestWriter.Write(w.bytesBuf.Bytes())
	if err != nil {
		return err
	}

	// Write values.
	// TODO(xichen): Use a streaming encoder to directly encode values into an bufio.Writer
	// instead of writing it to in-memory byte buffer and then to file.
	switch valueIt.valueType {
	case field.NullType:
		break
	case field.BoolType:
		w.bw.Reset()
		err = w.bw.Encode(w.fdWithDigestWriter, valueIt.boolIt)
	case field.IntType:
		w.iw.Reset()
		err = w.iw.Encode(w.fdWithDigestWriter, valueIt.intIt)
	case field.DoubleType:
		w.dw.Reset()
		err = w.dw.Encode(w.fdWithDigestWriter, valueIt.doubleIt)
	case field.StringType:
		w.sw.Reset()
		err = w.sw.Encode(w.fdWithDigestWriter, valueIt.stringIt)
	}
	if err != nil {
		return err
	}

	// Flush.
	return w.fdWithDigestWriter.Flush()
}

func (w *writer) writeCheckpointFile() error {
	path := checkpointFilePath(w.segmentDir)
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

type valueIteratorUnion struct {
	valueType field.ValueType
	boolIt    encoding.RewindableBoolIterator
	intIt     encoding.RewindableIntIterator
	doubleIt  encoding.RewindableDoubleIterator
	stringIt  encoding.RewindableStringIterator
}
