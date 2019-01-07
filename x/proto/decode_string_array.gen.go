// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package proto

import (
	"encoding/binary"

	"fmt"

	"io"

	"github.com/xichen2020/eventdb/generated/proto/encodingpb"
	"github.com/xichen2020/eventdb/x/bytes"

	xio "github.com/xichen2020/eventdb/x/io"
)

// *encodingpb.StringArray is a generic gogo protobuf generated Message type.

// DecodeStringArray decodes a *encodingpb.StringArray message from a reader.
func DecodeStringArray(
	reader xio.Reader,
	msg *encodingpb.StringArray,
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

// DecodeStringArrayRaw decodes raw bytes into a protobuf message, returning the
// number of bytes read and any errors encountered.
func DecodeStringArrayRaw(
	data []byte,
	msg *encodingpb.StringArray,
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
