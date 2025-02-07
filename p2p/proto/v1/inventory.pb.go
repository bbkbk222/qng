// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inventory.proto

package qitmeer_p2p_v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Inventory struct {
	Invs                 []*InvVect `protobuf:"bytes,1,rep,name=invs,proto3" json:"invs,omitempty" ssz-max:"2000"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Inventory) Reset()         { *m = Inventory{} }
func (m *Inventory) String() string { return proto.CompactTextString(m) }
func (*Inventory) ProtoMessage()    {}
func (*Inventory) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{0}
}
func (m *Inventory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Inventory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Inventory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Inventory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inventory.Merge(m, src)
}
func (m *Inventory) XXX_Size() int {
	return m.Size()
}
func (m *Inventory) XXX_DiscardUnknown() {
	xxx_messageInfo_Inventory.DiscardUnknown(m)
}

var xxx_messageInfo_Inventory proto.InternalMessageInfo

func (m *Inventory) GetInvs() []*InvVect {
	if m != nil {
		return m.Invs
	}
	return nil
}

type InvVect struct {
	Type                 uint32   `protobuf:"varint,100,opt,name=type,proto3" json:"type,omitempty"`
	Hash                 *Hash    `protobuf:"bytes,101,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvVect) Reset()         { *m = InvVect{} }
func (m *InvVect) String() string { return proto.CompactTextString(m) }
func (*InvVect) ProtoMessage()    {}
func (*InvVect) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{1}
}
func (m *InvVect) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InvVect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InvVect.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InvVect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvVect.Merge(m, src)
}
func (m *InvVect) XXX_Size() int {
	return m.Size()
}
func (m *InvVect) XXX_DiscardUnknown() {
	xxx_messageInfo_InvVect.DiscardUnknown(m)
}

var xxx_messageInfo_InvVect proto.InternalMessageInfo

func (m *InvVect) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *InvVect) GetHash() *Hash {
	if m != nil {
		return m.Hash
	}
	return nil
}

func init() {
	proto.RegisterType((*Inventory)(nil), "qitmeer.p2p.v1.Inventory")
	proto.RegisterType((*InvVect)(nil), "qitmeer.p2p.v1.InvVect")
}

func init() { proto.RegisterFile("inventory.proto", fileDescriptor_7173caedb7c6ae96) }

var fileDescriptor_7173caedb7c6ae96 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xcc, 0x2b, 0x4b,
	0xcd, 0x2b, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0xcc, 0x2c,
	0xc9, 0x4d, 0x4d, 0x2d, 0xd2, 0x2b, 0x30, 0x2a, 0xd0, 0x2b, 0x33, 0x94, 0xd2, 0x4d, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x4b,
	0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x5d, 0x8a, 0x37, 0x37, 0xb5, 0xb8, 0x38,
	0x31, 0x3d, 0x15, 0xc2, 0x55, 0xf2, 0xe1, 0xe2, 0xf4, 0x84, 0x59, 0x20, 0x64, 0xcf, 0xc5, 0x92,
	0x99, 0x57, 0x56, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xae, 0x87, 0x6a, 0x93, 0x9e,
	0x67, 0x5e, 0x59, 0x58, 0x6a, 0x72, 0x89, 0x93, 0xd0, 0xa7, 0x7b, 0xf2, 0x7c, 0xc5, 0xc5, 0x55,
	0xba, 0xb9, 0x89, 0x15, 0x56, 0x4a, 0x46, 0x06, 0x06, 0x06, 0x4a, 0x41, 0x60, 0x8d, 0x4a, 0xee,
	0x5c, 0xec, 0x50, 0x45, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0x29, 0x0a, 0x8c,
	0x1a, 0xbc, 0x41, 0x60, 0xb6, 0x90, 0x06, 0x17, 0x4b, 0x46, 0x62, 0x71, 0x86, 0x44, 0xaa, 0x02,
	0xa3, 0x06, 0xb7, 0x91, 0x08, 0xba, 0xf9, 0x1e, 0x89, 0xc5, 0x19, 0x41, 0x60, 0x15, 0x4e, 0x02,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72,
	0x0c, 0x49, 0x6c, 0x60, 0xf7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x84, 0x7b, 0xcb,
	0x10, 0x01, 0x00, 0x00,
}

func (m *Inventory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Inventory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Inventory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Invs) > 0 {
		for iNdEx := len(m.Invs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Invs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintInventory(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *InvVect) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InvVect) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InvVect) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Hash != nil {
		{
			size, err := m.Hash.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInventory(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xaa
	}
	if m.Type != 0 {
		i = encodeVarintInventory(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xa0
	}
	return len(dAtA) - i, nil
}

func encodeVarintInventory(dAtA []byte, offset int, v uint64) int {
	offset -= sovInventory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Inventory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Invs) > 0 {
		for _, e := range m.Invs {
			l = e.Size()
			n += 1 + l + sovInventory(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InvVect) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 2 + sovInventory(uint64(m.Type))
	}
	if m.Hash != nil {
		l = m.Hash.Size()
		n += 2 + l + sovInventory(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovInventory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInventory(x uint64) (n int) {
	return sovInventory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Inventory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInventory
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
			return fmt.Errorf("proto: Inventory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Inventory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Invs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInventory
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
				return ErrInvalidLengthInventory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInventory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Invs = append(m.Invs, &InvVect{})
			if err := m.Invs[len(m.Invs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInventory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInventory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInventory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InvVect) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInventory
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
			return fmt.Errorf("proto: InvVect: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InvVect: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 100:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInventory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 101:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInventory
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
				return ErrInvalidLengthInventory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInventory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Hash == nil {
				m.Hash = &Hash{}
			}
			if err := m.Hash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInventory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInventory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInventory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInventory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInventory
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
					return 0, ErrIntOverflowInventory
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
					return 0, ErrIntOverflowInventory
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
				return 0, ErrInvalidLengthInventory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInventory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInventory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInventory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInventory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInventory = fmt.Errorf("proto: unexpected end of group")
)
