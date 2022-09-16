// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/types/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	types1 "github.com/cosmos/ibc-go/v5/modules/core/03-connection/types"
	types2 "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	types3 "github.com/cosmos/ibc-go/v5/modules/core/28-wasm/types"
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

// GenesisState defines the ibc module's genesis state.
type GenesisState struct {
	// ICS002 - Clients genesis state
	ClientGenesis types.GenesisState `protobuf:"bytes,1,opt,name=client_genesis,json=clientGenesis,proto3" json:"client_genesis" yaml:"client_genesis"`
	// ICS003 - Connections genesis state
	ConnectionGenesis types1.GenesisState `protobuf:"bytes,2,opt,name=connection_genesis,json=connectionGenesis,proto3" json:"connection_genesis" yaml:"connection_genesis"`
	// ICS004 - Channel genesis state
	ChannelGenesis types2.GenesisState `protobuf:"bytes,3,opt,name=channel_genesis,json=channelGenesis,proto3" json:"channel_genesis" yaml:"channel_genesis"`
	// ICS028 - Wasm Light Client genesis state
	WasmGenesis types3.GenesisState `protobuf:"bytes,4,opt,name=wasm_genesis,json=wasmGenesis,proto3" json:"wasm_genesis" yaml:"wasm_genesis"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a49c5663e6fc59, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetClientGenesis() types.GenesisState {
	if m != nil {
		return m.ClientGenesis
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetConnectionGenesis() types1.GenesisState {
	if m != nil {
		return m.ConnectionGenesis
	}
	return types1.GenesisState{}
}

func (m *GenesisState) GetChannelGenesis() types2.GenesisState {
	if m != nil {
		return m.ChannelGenesis
	}
	return types2.GenesisState{}
}

func (m *GenesisState) GetWasmGenesis() types3.GenesisState {
	if m != nil {
		return m.WasmGenesis
	}
	return types3.GenesisState{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.core.types.v1.GenesisState")
}

func init() { proto.RegisterFile("ibc/core/types/v1/genesis.proto", fileDescriptor_b9a49c5663e6fc59) }

var fileDescriptor_b9a49c5663e6fc59 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xb1, 0x4e, 0xe3, 0x30,
	0x1c, 0xc6, 0x93, 0xeb, 0xe9, 0x86, 0xb4, 0xd7, 0x53, 0x73, 0x80, 0xa0, 0x08, 0xb7, 0x8d, 0x3a,
	0xb0, 0x60, 0xab, 0x20, 0x16, 0xc6, 0x2e, 0x30, 0x87, 0x8d, 0x01, 0x94, 0x18, 0x93, 0x1a, 0x25,
	0x76, 0x55, 0xbb, 0x41, 0x7d, 0x0b, 0x9e, 0x87, 0x27, 0xe8, 0xd8, 0x91, 0xa9, 0x42, 0xed, 0x1b,
	0xf0, 0x04, 0x28, 0xb1, 0xeb, 0xa4, 0xf2, 0x96, 0xfc, 0xbf, 0x9f, 0xbf, 0x9f, 0x6d, 0xd9, 0xeb,
	0xd1, 0x18, 0x23, 0xcc, 0x67, 0x04, 0xc9, 0xc5, 0x94, 0x08, 0x94, 0x8f, 0x50, 0x42, 0x18, 0x11,
	0x54, 0xc0, 0xe9, 0x8c, 0x4b, 0xee, 0x77, 0x68, 0x8c, 0x61, 0x01, 0xc0, 0x12, 0x80, 0xf9, 0xa8,
	0x7b, 0x90, 0xf0, 0x84, 0x97, 0x29, 0x2a, 0xbe, 0x14, 0xd8, 0xed, 0x9b, 0x26, 0x9c, 0x52, 0xc2,
	0xa4, 0x55, 0xd5, 0x1d, 0x56, 0x04, 0x67, 0x8c, 0x60, 0x49, 0x39, 0xb3, 0xa9, 0x41, 0x45, 0x4d,
	0x22, 0xc6, 0x48, 0x6a, 0x23, 0xc0, 0x20, 0x6f, 0x91, 0xc8, 0xac, 0x3c, 0xf8, 0x68, 0x78, 0xad,
	0x5b, 0x35, 0xb9, 0x97, 0x91, 0x24, 0xfe, 0x8b, 0xd7, 0x56, 0x9b, 0x7a, 0xd2, 0xe0, 0xb1, 0xdb,
	0x77, 0xcf, 0x9b, 0x97, 0x7d, 0x68, 0x4e, 0xa7, 0x72, 0x98, 0x8f, 0x60, 0x7d, 0xe5, 0xf8, 0x6c,
	0xb9, 0xee, 0x39, 0xdf, 0xeb, 0xde, 0xe1, 0x22, 0xca, 0xd2, 0x9b, 0x60, 0xbf, 0x25, 0x08, 0xff,
	0xaa, 0x81, 0x5e, 0xe2, 0xe7, 0x9e, 0x5f, 0x1d, 0xcd, 0xb8, 0x7e, 0x95, 0xae, 0x61, 0xcd, 0x65,
	0x18, 0xcb, 0x37, 0xd0, 0xbe, 0x13, 0xed, 0xb3, 0xda, 0x82, 0xb0, 0x53, 0x0d, 0x77, 0xde, 0x57,
	0xef, 0x9f, 0xbe, 0x2c, 0x23, 0x6d, 0x94, 0xd2, 0x41, 0x4d, 0xaa, 0x00, 0xcb, 0x08, 0xb4, 0xf1,
	0x48, 0x1b, 0xf7, 0x7b, 0x82, 0xb0, 0xad, 0x27, 0x3b, 0xd7, 0xa3, 0xd7, 0x2a, 0x6e, 0xdd, 0x88,
	0x7e, 0x97, 0x22, 0x50, 0x89, 0x8a, 0xd4, 0xb2, 0x9c, 0x6a, 0xcb, 0x7f, 0x65, 0xa9, 0x37, 0x04,
	0x61, 0xb3, 0xf8, 0xd5, 0xf8, 0xf8, 0x6e, 0xb9, 0x01, 0xee, 0x6a, 0x03, 0xdc, 0xaf, 0x0d, 0x70,
	0xdf, 0xb7, 0xc0, 0x59, 0x6d, 0x81, 0xf3, 0xb9, 0x05, 0xce, 0x03, 0x4c, 0xa8, 0x9c, 0xcc, 0x63,
	0x88, 0x79, 0x86, 0x30, 0x17, 0x19, 0x17, 0x88, 0xc6, 0xf8, 0x22, 0xe1, 0x28, 0xbf, 0x46, 0x19,
	0x7f, 0x9e, 0xa7, 0x44, 0xd4, 0xde, 0x72, 0xfc, 0xa7, 0x7c, 0x0d, 0x57, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0x3e, 0x80, 0xf7, 0xe4, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.WasmGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.ChannelGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ConnectionGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ClientGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClientGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ConnectionGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ChannelGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.WasmGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClientGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConnectionGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChannelGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WasmGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
