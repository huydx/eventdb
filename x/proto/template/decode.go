package template

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/mauricelam/genny/generic"
	"github.com/xichen2020/eventdb/x/bytes"
	xio "github.com/xichen2020/eventdb/x/io"
)

// GenericDecodeProtoMessage is a generic gogo protobuf generated Message type.
type GenericDecodeProtoMessage interface {
	generic.Type

	Unmarshal(dst []byte) error
}

// DecodeValue decodes a GenericDecodeProtoMessage message from a reader.
func DecodeValue(
	reader xio.Reader,
	msg GenericDecodeProtoMessage,
	extBuf *[]byte, // extBuf is an external byte buffer for memory re-use.
) (int, error) {
	protoSizeBytes, err := binary.ReadVarint(reader)
	if err != nil {
		return 0, err
	}
	// Compute the number of bytes read from the reader for the proto size bytes.
	bytesRead := xio.VarintBytes(protoSizeBytes)

	var buf []byte
	if extBuf == nil {
		buf = make([]byte, protoSizeBytes)
	} else {
		*extBuf = bytes.EnsureBufferSize(*extBuf, int(protoSizeBytes), bytes.DontCopyData)
		buf = *extBuf
	}
	if _, err := io.ReadFull(reader, buf[:protoSizeBytes]); err != nil {
		return 0, err
	}
	if err := msg.Unmarshal(buf[:protoSizeBytes]); err != nil {
		return 0, err
	}
	bytesRead += int(protoSizeBytes)
	return bytesRead, nil
}

// DecodeValueRaw decodes raw bytes into a protobuf message, returning the
// number of bytes read and any errors encountered.
func DecodeValueRaw(
	data []byte,
	msg GenericDecodeProtoMessage,
) (int, error) {
	protoSizeBytes, bytesRead, err := xio.ReadVarint(data)
	if err != nil {
		return 0, err
	}
	remainder := data[bytesRead:]
	if int(protoSizeBytes) > len(remainder) {
		return 0, fmt.Errorf("decoded message size %d exceeds available buffer size %d", protoSizeBytes, len(remainder))
	}
	marshalledBytes := remainder[:protoSizeBytes]
	if err := msg.Unmarshal(marshalledBytes); err != nil {
		return 0, err
	}
	bytesRead += int(protoSizeBytes)
	return bytesRead, nil
}
