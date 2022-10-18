// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coreum/asset/v1/event.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EventFungibleTokenIssued is emitted on MsgIssueFungibleToken.
type EventFungibleTokenIssued struct {
	Denom         string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Issuer        string                                 `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Symbol        string                                 `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Description   string                                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Recipient     string                                 `protobuf:"bytes,5,opt,name=recipient,proto3" json:"recipient,omitempty"`
	InitialAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=initial_amount,json=initialAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initial_amount"`
}

func (m *EventFungibleTokenIssued) Reset()         { *m = EventFungibleTokenIssued{} }
func (m *EventFungibleTokenIssued) String() string { return proto.CompactTextString(m) }
func (*EventFungibleTokenIssued) ProtoMessage()    {}
func (*EventFungibleTokenIssued) Descriptor() ([]byte, []int) {
	return fileDescriptor_aede4b64fdc52aa3, []int{0}
}
func (m *EventFungibleTokenIssued) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFungibleTokenIssued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFungibleTokenIssued.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFungibleTokenIssued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFungibleTokenIssued.Merge(m, src)
}
func (m *EventFungibleTokenIssued) XXX_Size() int {
	return m.Size()
}
func (m *EventFungibleTokenIssued) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFungibleTokenIssued.DiscardUnknown(m)
}

var xxx_messageInfo_EventFungibleTokenIssued proto.InternalMessageInfo

func (m *EventFungibleTokenIssued) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventFungibleTokenIssued) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *EventFungibleTokenIssued) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *EventFungibleTokenIssued) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventFungibleTokenIssued) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func init() {
	proto.RegisterType((*EventFungibleTokenIssued)(nil), "coreum.asset.v1.EventFungibleTokenIssued")
}

func init() { proto.RegisterFile("coreum/asset/v1/event.proto", fileDescriptor_aede4b64fdc52aa3) }

var fileDescriptor_aede4b64fdc52aa3 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x67, 0xd4, 0x16, 0x1a, 0x51, 0x61, 0x28, 0x12, 0x54, 0xa6, 0xc5, 0x85, 0xb8, 0x31,
	0xb1, 0xf8, 0x04, 0x56, 0x2c, 0xd4, 0x65, 0xd1, 0x8d, 0x1b, 0x99, 0x9f, 0xcb, 0x18, 0xda, 0xe4,
	0x0e, 0x93, 0xa4, 0xd8, 0xb7, 0xf0, 0xb1, 0xba, 0xec, 0x52, 0x5c, 0x14, 0x69, 0x1f, 0xc2, 0xad,
	0x24, 0x33, 0x60, 0x57, 0xc9, 0xfd, 0xbe, 0x9b, 0x43, 0x38, 0xe4, 0x3c, 0xc3, 0x0a, 0xac, 0xe4,
	0x89, 0xd6, 0x60, 0xf8, 0x7c, 0xc0, 0x61, 0x0e, 0xca, 0xb0, 0xb2, 0x42, 0x83, 0xd1, 0x49, 0x2d,
	0x99, 0x97, 0x6c, 0x3e, 0x38, 0xeb, 0x16, 0x58, 0xa0, 0x77, 0xdc, 0xdd, 0xea, 0xb5, 0xcb, 0xdf,
	0x90, 0xd0, 0x47, 0xf7, 0x6c, 0x64, 0x55, 0x21, 0xd2, 0x19, 0x3c, 0xe3, 0x14, 0xd4, 0x58, 0x6b,
	0x0b, 0x79, 0xd4, 0x25, 0xad, 0x1c, 0x14, 0x4a, 0x1a, 0xf6, 0xc3, 0xeb, 0xce, 0xa4, 0x1e, 0xa2,
	0x53, 0xd2, 0x16, 0xce, 0x57, 0x74, 0xcf, 0xe3, 0x66, 0x72, 0x5c, 0x2f, 0x64, 0x8a, 0x33, 0xba,
	0x5f, 0xf3, 0x7a, 0x8a, 0xfa, 0xe4, 0x30, 0x07, 0x9d, 0x55, 0xa2, 0x34, 0x02, 0x15, 0x3d, 0xf0,
	0x72, 0x17, 0x45, 0x17, 0xa4, 0x53, 0x41, 0x26, 0x4a, 0x01, 0xca, 0xd0, 0x96, 0xf7, 0xff, 0x20,
	0x7a, 0x21, 0xc7, 0x42, 0x09, 0x23, 0x92, 0xd9, 0x5b, 0x22, 0xd1, 0x2a, 0x43, 0xdb, 0x6e, 0x65,
	0xc8, 0x96, 0xeb, 0x5e, 0xf0, 0xbd, 0xee, 0x5d, 0x15, 0xc2, 0xbc, 0xdb, 0x94, 0x65, 0x28, 0x79,
	0x86, 0x5a, 0xa2, 0x6e, 0x8e, 0x1b, 0x9d, 0x4f, 0xb9, 0x59, 0x94, 0xa0, 0xd9, 0x58, 0x99, 0xc9,
	0x51, 0x93, 0x72, 0xef, 0x43, 0x86, 0x4f, 0xcb, 0x4d, 0x1c, 0xae, 0x36, 0x71, 0xf8, 0xb3, 0x89,
	0xc3, 0xcf, 0x6d, 0x1c, 0xac, 0xb6, 0x71, 0xf0, 0xb5, 0x8d, 0x83, 0xd7, 0xdb, 0x9d, 0xc0, 0x07,
	0xdf, 0xe2, 0x08, 0xad, 0xca, 0x13, 0xf7, 0x57, 0xde, 0x74, 0xfe, 0xd1, 0xb4, 0xee, 0xe3, 0xd3,
	0xb6, 0x2f, 0xf3, 0xee, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x7c, 0xdb, 0x05, 0x92, 0x01, 0x00,
	0x00,
}

func (m *EventFungibleTokenIssued) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFungibleTokenIssued) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFungibleTokenIssued) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InitialAmount.Size()
		i -= size
		if _, err := m.InitialAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventFungibleTokenIssued) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = m.InitialAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventFungibleTokenIssued) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventFungibleTokenIssued: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFungibleTokenIssued: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
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
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)