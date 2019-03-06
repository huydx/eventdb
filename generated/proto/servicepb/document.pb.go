// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/xichen2020/eventdb/generated/proto/servicepb/document.proto

/*
	Package servicepb is a generated protocol buffer package.

	It is generated from these files:
		github.com/xichen2020/eventdb/generated/proto/servicepb/document.proto
		github.com/xichen2020/eventdb/generated/proto/servicepb/query.proto
		github.com/xichen2020/eventdb/generated/proto/servicepb/service.proto
		github.com/xichen2020/eventdb/generated/proto/servicepb/value.proto

	It has these top-level messages:
		Field
		Document
		RawQuery
		GroupedQuery
		TimeBucketQuery
		FilterList
		Filter
		Calculation
		OrderBy
		OptionalTimeUnit
		OptionalInt64
		OptionalFilterCombinator
		OptionalString
		OptionalCalculationOp
		OptionalSortOrder
		HealthRequest
		HealthResult
		WriteRequest
		WriteResults
		RawQueryResults
		SingleKeyGroupQueryResult
		SingleKeyGroupQueryResults
		MultiKeyGroupQueryResult
		MultiKeyGroupQueryResults
		GroupedQueryResults
		TimeBucketQueryResult
		TimeBucketQueryResults
		FieldValue
		OptionalFieldValue
		FilterValue
		OptionalFilterValue
		CalculationValue
*/
package servicepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

type Field struct {
	Path  []string   `protobuf:"bytes,1,rep,name=path" json:"path,omitempty"`
	Value FieldValue `protobuf:"bytes,2,opt,name=value" json:"value"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptorDocument, []int{0} }

func (m *Field) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Field) GetValue() FieldValue {
	if m != nil {
		return m.Value
	}
	return FieldValue{}
}

type Document struct {
	Id        []byte  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TimeNanos int64   `protobuf:"varint,2,opt,name=time_nanos,json=timeNanos,proto3" json:"time_nanos,omitempty"`
	RawData   []byte  `protobuf:"bytes,3,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	Fields    []Field `protobuf:"bytes,4,rep,name=fields" json:"fields"`
}

func (m *Document) Reset()                    { *m = Document{} }
func (m *Document) String() string            { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()               {}
func (*Document) Descriptor() ([]byte, []int) { return fileDescriptorDocument, []int{1} }

func (m *Document) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Document) GetTimeNanos() int64 {
	if m != nil {
		return m.TimeNanos
	}
	return 0
}

func (m *Document) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func (m *Document) GetFields() []Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*Field)(nil), "servicepb.Field")
	proto.RegisterType((*Document)(nil), "servicepb.Document")
}
func (m *Field) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Field) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		for _, s := range m.Path {
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
	dAtA[i] = 0x12
	i++
	i = encodeVarintDocument(dAtA, i, uint64(m.Value.Size()))
	n1, err := m.Value.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *Document) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Document) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDocument(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.TimeNanos != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDocument(dAtA, i, uint64(m.TimeNanos))
	}
	if len(m.RawData) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDocument(dAtA, i, uint64(len(m.RawData)))
		i += copy(dAtA[i:], m.RawData)
	}
	if len(m.Fields) > 0 {
		for _, msg := range m.Fields {
			dAtA[i] = 0x22
			i++
			i = encodeVarintDocument(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintDocument(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Field) Size() (n int) {
	var l int
	_ = l
	if len(m.Path) > 0 {
		for _, s := range m.Path {
			l = len(s)
			n += 1 + l + sovDocument(uint64(l))
		}
	}
	l = m.Value.Size()
	n += 1 + l + sovDocument(uint64(l))
	return n
}

func (m *Document) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDocument(uint64(l))
	}
	if m.TimeNanos != 0 {
		n += 1 + sovDocument(uint64(m.TimeNanos))
	}
	l = len(m.RawData)
	if l > 0 {
		n += 1 + l + sovDocument(uint64(l))
	}
	if len(m.Fields) > 0 {
		for _, e := range m.Fields {
			l = e.Size()
			n += 1 + l + sovDocument(uint64(l))
		}
	}
	return n
}

func sovDocument(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDocument(x uint64) (n int) {
	return sovDocument(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Field) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDocument
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
			return fmt.Errorf("proto: Field: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Field: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
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
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = append(m.Path, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDocument(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDocument
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
func (m *Document) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDocument
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
			return fmt.Errorf("proto: Document: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Document: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeNanos", wireType)
			}
			m.TimeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawData = append(m.RawData[:0], dAtA[iNdEx:postIndex]...)
			if m.RawData == nil {
				m.RawData = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDocument
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDocument
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fields = append(m.Fields, Field{})
			if err := m.Fields[len(m.Fields)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDocument(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDocument
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
func skipDocument(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDocument
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
					return 0, ErrIntOverflowDocument
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
					return 0, ErrIntOverflowDocument
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
				return 0, ErrInvalidLengthDocument
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDocument
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
				next, err := skipDocument(dAtA[start:])
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
	ErrInvalidLengthDocument = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDocument   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/xichen2020/eventdb/generated/proto/servicepb/document.proto", fileDescriptorDocument)
}

var fileDescriptorDocument = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4e, 0xc2, 0x30,
	0x1c, 0xc6, 0x29, 0x03, 0x84, 0x62, 0x8c, 0x69, 0x62, 0x32, 0x49, 0x9c, 0x0b, 0xa7, 0x5d, 0x5c,
	0x71, 0x1e, 0xbc, 0x23, 0x21, 0xf1, 0xc2, 0x61, 0x07, 0x0f, 0x5e, 0x48, 0xb7, 0xfe, 0x19, 0x4d,
	0xa0, 0x25, 0x5b, 0x37, 0x7c, 0x00, 0x1f, 0xc0, 0xc7, 0xe2, 0xe8, 0x13, 0x18, 0x33, 0x5f, 0xc4,
	0xac, 0x5b, 0x88, 0xf1, 0xc8, 0xed, 0xdf, 0xff, 0xf7, 0x7d, 0xbf, 0x7e, 0x2d, 0x9e, 0x27, 0x42,
	0xaf, 0xf3, 0xc8, 0x8f, 0xd5, 0x96, 0xbe, 0x89, 0x78, 0x0d, 0x32, 0x98, 0x04, 0x13, 0x0a, 0x05,
	0x48, 0xcd, 0x23, 0x9a, 0x80, 0x84, 0x94, 0x69, 0xe0, 0x74, 0x97, 0x2a, 0xad, 0x68, 0x06, 0x69,
	0x21, 0x62, 0xd8, 0x45, 0x94, 0xab, 0x38, 0xdf, 0x82, 0xd4, 0xbe, 0x11, 0xc8, 0xe0, 0xa8, 0x8c,
	0xee, 0xfe, 0x20, 0x13, 0x95, 0xa8, 0x3a, 0x1a, 0xe5, 0x2b, 0x73, 0xaa, 0x39, 0xd5, 0x54, 0x27,
	0x47, 0x4f, 0xa7, 0x36, 0x28, 0xd8, 0x26, 0x87, 0x1a, 0x32, 0x5e, 0xe0, 0xee, 0x5c, 0xc0, 0x86,
	0x13, 0x82, 0x3b, 0x3b, 0xa6, 0xd7, 0x36, 0x72, 0x2d, 0x6f, 0x10, 0x9a, 0x99, 0xdc, 0xe3, 0xae,
	0xf1, 0xda, 0x6d, 0x17, 0x79, 0xc3, 0xe0, 0xca, 0x3f, 0x32, 0x7c, 0x13, 0x7a, 0xa9, 0xc4, 0x69,
	0xe7, 0xf0, 0x75, 0xdb, 0x0a, 0x6b, 0xe7, 0xf8, 0x1d, 0xe1, 0xfe, 0xac, 0x79, 0x21, 0xb9, 0xc0,
	0x6d, 0xc1, 0x6d, 0xe4, 0x22, 0xef, 0x3c, 0x6c, 0x0b, 0x4e, 0x6e, 0x30, 0xd6, 0x62, 0x0b, 0x4b,
	0xc9, 0xa4, 0xca, 0x0c, 0xd4, 0x0a, 0x07, 0xd5, 0x66, 0x51, 0x2d, 0xc8, 0x35, 0xee, 0xa7, 0x6c,
	0xbf, 0xe4, 0x4c, 0x33, 0xdb, 0x32, 0xa1, 0xb3, 0x94, 0xed, 0x67, 0x4c, 0x33, 0xe2, 0xe3, 0xde,
	0xaa, 0xba, 0x31, 0xb3, 0x3b, 0xae, 0xe5, 0x0d, 0x83, 0xcb, 0xff, 0x55, 0x9a, 0x16, 0x8d, 0x6b,
	0xfa, 0x7c, 0x28, 0x1d, 0xf4, 0x59, 0x3a, 0xe8, 0xbb, 0x74, 0xd0, 0xc7, 0x8f, 0xd3, 0x7a, 0x7d,
	0x3c, 0xf1, 0xb7, 0xa2, 0x9e, 0x59, 0x3c, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xff, 0x50, 0x7e,
	0x84, 0xf1, 0x01, 0x00, 0x00,
}
