// Copyright (c) 2018 Uber Technologies, Inc.
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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/xichen2020/eventdb/generated/proto/encodingpb/string.proto

package encodingpb // import "github.com/xichen2020/eventdb/generated/proto/encodingpb"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StringMeta struct {
	Encoding             EncodingType    `protobuf:"varint,1,opt,name=encoding,proto3,enum=encodingpb.EncodingType" json:"encoding,omitempty"`
	Compression          CompressionType `protobuf:"varint,2,opt,name=compression,proto3,enum=encodingpb.CompressionType" json:"compression,omitempty"`
	UseBlocks            bool            `protobuf:"varint,3,opt,name=use_blocks,json=useBlocks,proto3" json:"use_blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StringMeta) Reset()         { *m = StringMeta{} }
func (m *StringMeta) String() string { return proto.CompactTextString(m) }
func (*StringMeta) ProtoMessage()    {}
func (*StringMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_string_d9c3b2c2b38ee610, []int{0}
}
func (m *StringMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StringMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMeta.Merge(dst, src)
}
func (m *StringMeta) XXX_Size() int {
	return m.Size()
}
func (m *StringMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMeta.DiscardUnknown(m)
}

var xxx_messageInfo_StringMeta proto.InternalMessageInfo

func (m *StringMeta) GetEncoding() EncodingType {
	if m != nil {
		return m.Encoding
	}
	return EncodingType_UNKNOWN_ENCODING
}

func (m *StringMeta) GetCompression() CompressionType {
	if m != nil {
		return m.Compression
	}
	return CompressionType_UNKNOWN_COMPRESSION
}

func (m *StringMeta) GetUseBlocks() bool {
	if m != nil {
		return m.UseBlocks
	}
	return false
}

type StringArray struct {
	Data                 []string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringArray) Reset()         { *m = StringArray{} }
func (m *StringArray) String() string { return proto.CompactTextString(m) }
func (*StringArray) ProtoMessage()    {}
func (*StringArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_string_d9c3b2c2b38ee610, []int{1}
}
func (m *StringArray) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringArray.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StringArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringArray.Merge(dst, src)
}
func (m *StringArray) XXX_Size() int {
	return m.Size()
}
func (m *StringArray) XXX_DiscardUnknown() {
	xxx_messageInfo_StringArray.DiscardUnknown(m)
}

var xxx_messageInfo_StringArray proto.InternalMessageInfo

func (m *StringArray) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*StringMeta)(nil), "encodingpb.StringMeta")
	proto.RegisterType((*StringArray)(nil), "encodingpb.StringArray")
}
func (m *StringMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Encoding != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintString(dAtA, i, uint64(m.Encoding))
	}
	if m.Compression != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintString(dAtA, i, uint64(m.Compression))
	}
	if m.UseBlocks {
		dAtA[i] = 0x18
		i++
		if m.UseBlocks {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *StringArray) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringArray) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, s := range m.Data {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintString(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StringMeta) Size() (n int) {
	var l int
	_ = l
	if m.Encoding != 0 {
		n += 1 + sovString(uint64(m.Encoding))
	}
	if m.Compression != 0 {
		n += 1 + sovString(uint64(m.Compression))
	}
	if m.UseBlocks {
		n += 2
	}
	return n
}

func (m *StringArray) Size() (n int) {
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, s := range m.Data {
			l = len(s)
			n += 1 + l + sovString(uint64(l))
		}
	}
	return n
}

func sovString(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozString(x uint64) (n int) {
	return sovString(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StringMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowString
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StringMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encoding", wireType)
			}
			m.Encoding = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Encoding |= (EncodingType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compression", wireType)
			}
			m.Compression = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Compression |= (CompressionType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseBlocks", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UseBlocks = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipString(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthString
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StringArray) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowString
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StringArray: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringArray: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipString(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthString
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipString(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowString
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowString
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowString
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthString
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowString
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipString(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthString = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowString   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/xichen2020/eventdb/generated/proto/encodingpb/string.proto", fileDescriptor_string_d9c3b2c2b38ee610)
}

var fileDescriptor_string_d9c3b2c2b38ee610 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xaf, 0xc8, 0x4c, 0xce, 0x48, 0xcd, 0x33, 0x32, 0x30,
	0x32, 0xd0, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x49, 0x49, 0xd2, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a,
	0x2c, 0x49, 0x4d, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xcd, 0x4b, 0xce, 0x4f, 0xc9,
	0xcc, 0x4b, 0x2f, 0x48, 0xd2, 0x2f, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0xd7, 0x03, 0x8b, 0x0b, 0x71,
	0x21, 0x24, 0xa4, 0x5c, 0xc8, 0x36, 0xb2, 0xa4, 0xb2, 0x20, 0xb5, 0x18, 0x62, 0x22, 0x05, 0xa6,
	0x24, 0xe5, 0xe4, 0x27, 0x67, 0x43, 0x4c, 0x51, 0x5a, 0xc0, 0xc8, 0xc5, 0x15, 0x0c, 0x76, 0xa8,
	0x6f, 0x6a, 0x49, 0xa2, 0x90, 0x09, 0x17, 0x07, 0x4c, 0xa1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x9f,
	0x91, 0x84, 0x1e, 0x42, 0xa7, 0x9e, 0x2b, 0x94, 0x19, 0x52, 0x59, 0x90, 0x1a, 0x04, 0x57, 0x29,
	0x64, 0xcb, 0xc5, 0x9d, 0x9c, 0x9f, 0x5b, 0x50, 0x94, 0x5a, 0x5c, 0x9c, 0x99, 0x9f, 0x27, 0xc1,
	0x04, 0xd6, 0x28, 0x8d, 0xac, 0xd1, 0x19, 0x21, 0x0d, 0xd6, 0x8b, 0xac, 0x5e, 0x48, 0x96, 0x8b,
	0xab, 0xb4, 0x38, 0x35, 0x1e, 0xec, 0xac, 0x62, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0xce,
	0xd2, 0xe2, 0x54, 0x27, 0xb0, 0x80, 0x92, 0x22, 0x17, 0x37, 0xc4, 0x85, 0x8e, 0x45, 0x45, 0x89,
	0x95, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41,
	0x60, 0xb6, 0x93, 0xd7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x38, 0xe1, 0xb1, 0x1c, 0x43, 0x94, 0x05, 0xb9, 0xa1, 0x93, 0xc4, 0x06, 0x16, 0x31, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0x33, 0x31, 0x48, 0xf9, 0x01, 0x00, 0x00,
}