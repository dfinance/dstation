// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dfinance/namespace/namespace.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Whois to sell and buy.
type Whois struct {
	// ID.
	ID github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"ID" yaml:"id"`
	// Creator address.
	Creator string `protobuf:"bytes,2,opt,name=Creator,proto3" json:"Creator,omitempty" yaml:"creator"`
	// Value (domain name).
	Value string `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty" yaml:"value"`
	// Price.
	Price github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=Price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"Price" yaml:"price"`
}

func (m *Whois) Reset()         { *m = Whois{} }
func (m *Whois) String() string { return proto.CompactTextString(m) }
func (*Whois) ProtoMessage()    {}
func (*Whois) Descriptor() ([]byte, []int) {
	return fileDescriptor_f890c34d0c54e517, []int{0}
}
func (m *Whois) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Whois) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Whois.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Whois) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Whois.Merge(m, src)
}
func (m *Whois) XXX_Size() int {
	return m.Size()
}
func (m *Whois) XXX_DiscardUnknown() {
	xxx_messageInfo_Whois.DiscardUnknown(m)
}

var xxx_messageInfo_Whois proto.InternalMessageInfo

func (m *Whois) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Whois) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Whois) GetPrice() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Price
	}
	return nil
}

func init() {
	proto.RegisterType((*Whois)(nil), "dfinance.namespace.v1beta1.Whois")
}

func init() {
	proto.RegisterFile("dfinance/namespace/namespace.proto", fileDescriptor_f890c34d0c54e517)
}

var fileDescriptor_f890c34d0c54e517 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x6d, 0xcb, 0xe3, 0xbd, 0xd0, 0x67, 0x8c, 0x69, 0x5c, 0x14, 0x16, 0x2d, 0xe9, 0x42, 0x59,
	0x68, 0x27, 0xe8, 0xce, 0xb8, 0x30, 0xc0, 0x42, 0x12, 0x17, 0xa6, 0x89, 0x9a, 0xb8, 0x31, 0xd3,
	0x76, 0x28, 0x13, 0xe9, 0xdc, 0xda, 0x19, 0x88, 0xfc, 0x85, 0x5f, 0xe0, 0x07, 0xf8, 0x25, 0x2c,
	0x59, 0x1a, 0x17, 0xd5, 0xc0, 0x1f, 0xf0, 0x05, 0xa6, 0x9d, 0x22, 0x2c, 0x5d, 0xf5, 0xdc, 0x7b,
	0xce, 0x3d, 0x37, 0xe7, 0x76, 0x74, 0x27, 0x1c, 0x50, 0x86, 0x59, 0x40, 0x10, 0xc3, 0x31, 0xe1,
	0x09, 0xde, 0x46, 0x6e, 0x92, 0x82, 0x00, 0xa3, 0xb1, 0xd6, 0xb8, 0x1b, 0x66, 0xd2, 0xf6, 0x89,
	0xc0, 0xed, 0x86, 0x1d, 0x01, 0x44, 0x23, 0x82, 0x0a, 0xa5, 0x3f, 0x1e, 0x20, 0x41, 0x63, 0xc2,
	0x05, 0x8e, 0x13, 0x39, 0xdc, 0xd8, 0x8f, 0x20, 0x82, 0x02, 0xa2, 0x1c, 0x95, 0xdd, 0x7a, 0x00,
	0x3c, 0x06, 0xfe, 0x20, 0x09, 0x59, 0x94, 0x94, 0x25, 0x2b, 0xe4, 0x63, 0x4e, 0x50, 0xb9, 0x06,
	0x05, 0x40, 0x99, 0xe4, 0x9d, 0x57, 0x4d, 0xaf, 0xde, 0x0d, 0x81, 0x72, 0xe3, 0x4a, 0xd7, 0xfa,
	0x3d, 0x53, 0x6d, 0xaa, 0xad, 0x5a, 0xe7, 0x7c, 0x96, 0xd9, 0xca, 0x47, 0x66, 0x1f, 0x46, 0x54,
	0x0c, 0xc7, 0xbe, 0x1b, 0x40, 0x5c, 0xda, 0x96, 0x9f, 0x63, 0x1e, 0x3e, 0x22, 0x31, 0x4d, 0x08,
	0x77, 0x6f, 0x28, 0x13, 0xab, 0xcc, 0xae, 0x4d, 0x71, 0x3c, 0x3a, 0x73, 0x68, 0xe8, 0x78, 0x5a,
	0xbf, 0x67, 0x1c, 0xe9, 0xff, 0xba, 0x29, 0xc1, 0x02, 0x52, 0x53, 0x2b, 0x2c, 0x8d, 0x55, 0x66,
	0xef, 0x4a, 0x4d, 0x20, 0x09, 0xc7, 0x5b, 0x4b, 0x8c, 0x03, 0xbd, 0x7a, 0x8b, 0x47, 0x63, 0x62,
	0x56, 0x0a, 0xed, 0xde, 0x2a, 0xb3, 0x77, 0xa4, 0x76, 0x92, 0xb7, 0x1d, 0x4f, 0xd2, 0xc6, 0x93,
	0x5e, 0xbd, 0x4e, 0x69, 0x40, 0xcc, 0x3f, 0xcd, 0x4a, 0xeb, 0xff, 0x49, 0xdd, 0x2d, 0xb3, 0xe6,
	0xe9, 0xd6, 0x47, 0x74, 0xbb, 0x40, 0x59, 0xe7, 0x22, 0x4f, 0xb0, 0xb1, 0x49, 0xf2, 0x29, 0xe7,
	0xed, 0xd3, 0x6e, 0xfd, 0x22, 0x51, 0x6e, 0xc0, 0x3d, 0xb9, 0xa9, 0x73, 0x39, 0x5b, 0x58, 0xea,
	0x7c, 0x61, 0xa9, 0x5f, 0x0b, 0x4b, 0x7d, 0x59, 0x5a, 0xca, 0x7c, 0x69, 0x29, 0xef, 0x4b, 0x4b,
	0xb9, 0x77, 0xb7, 0xac, 0x7e, 0xfe, 0x7b, 0xc8, 0x05, 0x16, 0x14, 0x18, 0x7a, 0xde, 0x7a, 0x02,
	0x85, 0xad, 0xff, 0xb7, 0xb8, 0xf8, 0xe9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0xca, 0x04,
	0x74, 0x25, 0x02, 0x00, 0x00,
}

func (m *Whois) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Whois) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Whois) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNamespace(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNamespace(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.ID.Size()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNamespace(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintNamespace(dAtA []byte, offset int, v uint64) int {
	offset -= sovNamespace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Whois) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovNamespace(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovNamespace(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovNamespace(uint64(l))
		}
	}
	return n
}

func sovNamespace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNamespace(x uint64) (n int) {
	return sovNamespace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Whois) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNamespace
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
			return fmt.Errorf("proto: Whois: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Whois: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNamespace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNamespace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
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
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNamespace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNamespace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNamespace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNamespace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNamespace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNamespace
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
func skipNamespace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNamespace
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
					return 0, ErrIntOverflowNamespace
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
					return 0, ErrIntOverflowNamespace
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
				return 0, ErrInvalidLengthNamespace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNamespace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNamespace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNamespace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNamespace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNamespace = fmt.Errorf("proto: unexpected end of group")
)
