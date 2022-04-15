// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: documentservice/contract.proto

package types

import (
	fmt "fmt"
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

type Contract struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id           uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	ContractHash string `protobuf:"bytes,3,opt,name=contractHash,proto3" json:"contractHash,omitempty"`
	State        string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Seller       string `protobuf:"bytes,5,opt,name=seller,proto3" json:"seller,omitempty"`
	Buyer        string `protobuf:"bytes,6,opt,name=buyer,proto3" json:"buyer,omitempty"`
	CreateDate   string `protobuf:"bytes,7,opt,name=createDate,proto3" json:"createDate,omitempty"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e05e308744b7503, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return m.Size()
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Contract) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Contract) GetContractHash() string {
	if m != nil {
		return m.ContractHash
	}
	return ""
}

func (m *Contract) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Contract) GetSeller() string {
	if m != nil {
		return m.Seller
	}
	return ""
}

func (m *Contract) GetBuyer() string {
	if m != nil {
		return m.Buyer
	}
	return ""
}

func (m *Contract) GetCreateDate() string {
	if m != nil {
		return m.CreateDate
	}
	return ""
}

func init() {
	proto.RegisterType((*Contract)(nil), "cosmonaut.documentservice.documentservice.Contract")
}

func init() { proto.RegisterFile("documentservice/contract.proto", fileDescriptor_6e05e308744b7503) }

var fileDescriptor_6e05e308744b7503 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe3, 0xd0, 0xa6, 0x70, 0x85, 0x18, 0x2c, 0x84, 0x3c, 0x59, 0x55, 0xa7, 0xb2, 0x24,
	0x03, 0x2b, 0x13, 0x30, 0x30, 0x77, 0x60, 0x60, 0x73, 0x9c, 0x2b, 0x1a, 0xa9, 0xc9, 0xad, 0xec,
	0x1b, 0x44, 0xdf, 0x82, 0x07, 0xe2, 0x01, 0x18, 0x3b, 0x32, 0xa2, 0xe4, 0x45, 0x50, 0xdd, 0x06,
	0x41, 0x3a, 0x7e, 0xe7, 0x7c, 0xfe, 0xd1, 0x01, 0x5d, 0x90, 0x6d, 0x2a, 0xac, 0xd9, 0xa3, 0x7b,
	0x2d, 0x2d, 0x66, 0x96, 0x6a, 0x76, 0xc6, 0x72, 0xba, 0x76, 0xc4, 0x24, 0xaf, 0x2d, 0xf9, 0x8a,
	0x6a, 0xd3, 0x70, 0x3a, 0x30, 0x87, 0x3c, 0xfb, 0x10, 0x70, 0x7a, 0x7f, 0x38, 0x2d, 0x15, 0x4c,
	0xac, 0x43, 0xc3, 0xe4, 0x94, 0x98, 0x8a, 0xf9, 0xd9, 0xa2, 0x47, 0x79, 0x01, 0x71, 0x59, 0xa8,
	0x78, 0x2a, 0xe6, 0xa3, 0x45, 0x5c, 0x16, 0x72, 0x06, 0xe7, 0xfd, 0x9b, 0x8f, 0xc6, 0x2f, 0xd5,
	0x49, 0xd0, 0xff, 0x65, 0xf2, 0x12, 0xc6, 0x9e, 0x0d, 0xa3, 0x1a, 0x85, 0x72, 0x0f, 0xf2, 0x0a,
	0x12, 0x8f, 0xab, 0x15, 0x3a, 0x35, 0x0e, 0xf1, 0x81, 0x76, 0x76, 0xde, 0x6c, 0xd0, 0xa9, 0x64,
	0x6f, 0x07, 0x90, 0x1a, 0x20, 0x7c, 0x01, 0x1f, 0x76, 0x17, 0x4d, 0x42, 0xf5, 0x27, 0xb9, 0x7b,
	0xfa, 0x6c, 0xb5, 0xd8, 0xb6, 0x5a, 0x7c, 0xb7, 0x5a, 0xbc, 0x77, 0x3a, 0xda, 0x76, 0x3a, 0xfa,
	0xea, 0x74, 0xf4, 0x7c, 0xfb, 0x52, 0xf2, 0xb2, 0xc9, 0x53, 0x4b, 0x55, 0xf6, 0x3b, 0x47, 0x36,
	0x1c, 0xee, 0xed, 0x28, 0xe1, 0xcd, 0x1a, 0x7d, 0x9e, 0x84, 0x21, 0x6f, 0x7e, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x22, 0x59, 0xad, 0x44, 0x6a, 0x01, 0x00, 0x00,
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CreateDate) > 0 {
		i -= len(m.CreateDate)
		copy(dAtA[i:], m.CreateDate)
		i = encodeVarintContract(dAtA, i, uint64(len(m.CreateDate)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Buyer) > 0 {
		i -= len(m.Buyer)
		copy(dAtA[i:], m.Buyer)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Buyer)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintContract(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ContractHash) > 0 {
		i -= len(m.ContractHash)
		copy(dAtA[i:], m.ContractHash)
		i = encodeVarintContract(dAtA, i, uint64(len(m.ContractHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovContract(uint64(m.Id))
	}
	l = len(m.ContractHash)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Buyer)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.CreateDate)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	return n
}

func sovContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContract(x uint64) (n int) {
	return sovContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
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
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buyer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreateDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContract
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
func skipContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
				return 0, ErrInvalidLengthContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContract = fmt.Errorf("proto: unexpected end of group")
)
