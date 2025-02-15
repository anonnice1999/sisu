// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sisu/params.proto

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

type TransferOutParams struct {
	// Id of the chain
	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	// The number of maximum transfer in a single transaction.
	MaxBatching int32 `protobuf:"varint,2,opt,name=maxBatching,proto3" json:"maxBatching,omitempty"`
}

func (m *TransferOutParams) Reset()         { *m = TransferOutParams{} }
func (m *TransferOutParams) String() string { return proto.CompactTextString(m) }
func (*TransferOutParams) ProtoMessage()    {}
func (*TransferOutParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7faea80042f6b5b, []int{0}
}
func (m *TransferOutParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferOutParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferOutParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferOutParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferOutParams.Merge(m, src)
}
func (m *TransferOutParams) XXX_Size() int {
	return m.Size()
}
func (m *TransferOutParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferOutParams.DiscardUnknown(m)
}

var xxx_messageInfo_TransferOutParams proto.InternalMessageInfo

func (m *TransferOutParams) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *TransferOutParams) GetMaxBatching() int32 {
	if m != nil {
		return m.MaxBatching
	}
	return 0
}

type Params struct {
	MajorityThreshold int32                `protobuf:"varint,1,opt,name=majority_threshold,json=majorityThreshold,proto3" json:"majority_threshold,omitempty"`
	TransferOutParams []*TransferOutParams `protobuf:"bytes,2,rep,name=transfer_out_params,json=transferOutParams,proto3" json:"transfer_out_params,omitempty"`
	SupportedChains   []string             `protobuf:"bytes,3,rep,name=supported_chains,json=supportedChains,proto3" json:"supported_chains,omitempty"`
	// Transaction commission rate with 1 unit = 0.01%
	CommissionRate int32 `protobuf:"varint,4,opt,name=commission_rate,json=commissionRate,proto3" json:"commission_rate,omitempty"`
	// Expiration Ablock
	ExpirationBlock int32 `protobuf:"varint,5,opt,name=expiration_block,json=expirationBlock,proto3" json:"expiration_block,omitempty"`
	// Maximum number of retry when a txout fails to sign.
	MaxKeysignRetry int32 `protobuf:"varint,6,opt,name=max_keysign_retry,json=maxKeysignRetry,proto3" json:"max_keysign_retry,omitempty"`
	// Maxinum number of retry when a transfer is rejected.
	MaxRejectedTransferRetry int32 `protobuf:"varint,7,opt,name=max_rejected_transfer_retry,json=maxRejectedTransferRetry,proto3" json:"max_rejected_transfer_retry,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7faea80042f6b5b, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMajorityThreshold() int32 {
	if m != nil {
		return m.MajorityThreshold
	}
	return 0
}

func (m *Params) GetTransferOutParams() []*TransferOutParams {
	if m != nil {
		return m.TransferOutParams
	}
	return nil
}

func (m *Params) GetSupportedChains() []string {
	if m != nil {
		return m.SupportedChains
	}
	return nil
}

func (m *Params) GetCommissionRate() int32 {
	if m != nil {
		return m.CommissionRate
	}
	return 0
}

func (m *Params) GetExpirationBlock() int32 {
	if m != nil {
		return m.ExpirationBlock
	}
	return 0
}

func (m *Params) GetMaxKeysignRetry() int32 {
	if m != nil {
		return m.MaxKeysignRetry
	}
	return 0
}

func (m *Params) GetMaxRejectedTransferRetry() int32 {
	if m != nil {
		return m.MaxRejectedTransferRetry
	}
	return 0
}

func init() {
	proto.RegisterType((*TransferOutParams)(nil), "types.TransferOutParams")
	proto.RegisterType((*Params)(nil), "types.Params")
}

func init() { proto.RegisterFile("sisu/params.proto", fileDescriptor_c7faea80042f6b5b) }

var fileDescriptor_c7faea80042f6b5b = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xbf, 0x8a, 0xdb, 0x40,
	0x10, 0x87, 0x2d, 0x3b, 0x72, 0xf0, 0x1a, 0xe2, 0x68, 0x93, 0x62, 0x21, 0x20, 0x84, 0x9b, 0xc8,
	0x01, 0xcb, 0x90, 0xd4, 0x69, 0xec, 0x26, 0xe0, 0x22, 0xc7, 0xe2, 0xea, 0x1a, 0xb1, 0x96, 0xf7,
	0xac, 0xb5, 0xbd, 0x5a, 0xb1, 0x3b, 0xe2, 0xa4, 0xb7, 0xb8, 0xc7, 0xba, 0xd2, 0xe5, 0x95, 0x87,
	0x5d, 0xde, 0x4b, 0x1c, 0x5a, 0xf9, 0xcf, 0x81, 0x2b, 0x31, 0xdf, 0xfc, 0x66, 0xf8, 0x18, 0x2d,
	0xf2, 0x8c, 0x30, 0xc5, 0x24, 0x67, 0x9a, 0x49, 0x13, 0xe5, 0x5a, 0x81, 0xc2, 0x2e, 0x54, 0x39,
	0x37, 0xc3, 0x39, 0xf2, 0x16, 0x9a, 0x65, 0xe6, 0x81, 0xeb, 0xff, 0x05, 0xdc, 0xd9, 0x04, 0xfe,
	0x8e, 0xdc, 0x24, 0x65, 0x22, 0x23, 0x4e, 0xe0, 0x84, 0x3d, 0xda, 0x14, 0x38, 0x40, 0x7d, 0xc9,
	0xca, 0x29, 0x83, 0x24, 0x15, 0xd9, 0x9a, 0xb4, 0x03, 0x27, 0x74, 0xe9, 0x47, 0x34, 0x7c, 0x6b,
	0xa3, 0xee, 0x69, 0xc5, 0x18, 0x61, 0xc9, 0x36, 0x4a, 0x0b, 0xa8, 0x62, 0x48, 0x35, 0x37, 0xa9,
	0xda, 0xad, 0xec, 0x3e, 0x97, 0x7a, 0xe7, 0xce, 0xe2, 0xdc, 0xc0, 0xff, 0xd0, 0x37, 0x38, 0x69,
	0xc4, 0xaa, 0x80, 0xb8, 0x51, 0x25, 0xed, 0xa0, 0x13, 0xf6, 0x7f, 0x93, 0xc8, 0xba, 0x46, 0x37,
	0xa2, 0xd4, 0x83, 0x1b, 0xf7, 0x11, 0xfa, 0x6a, 0x8a, 0x3c, 0x57, 0x1a, 0xf8, 0x2a, 0xb6, 0xe2,
	0x86, 0x74, 0x82, 0x4e, 0xd8, 0xa3, 0x83, 0x0b, 0x9f, 0x59, 0x8c, 0x7f, 0xa2, 0x41, 0xa2, 0xa4,
	0x14, 0xc6, 0x08, 0x95, 0xc5, 0x9a, 0x01, 0x27, 0x9f, 0xac, 0xe0, 0x97, 0x2b, 0xa6, 0x0c, 0x78,
	0xbd, 0x93, 0x97, 0xb9, 0xd0, 0x0c, 0xea, 0xe0, 0x72, 0xa7, 0x92, 0x2d, 0x71, 0x6d, 0x72, 0x70,
	0xe5, 0xd3, 0x1a, 0xe3, 0x5f, 0xc8, 0x93, 0xac, 0x8c, 0xb7, 0xbc, 0x32, 0x62, 0x9d, 0xc5, 0x9a,
	0x83, 0xae, 0x48, 0xb7, 0xc9, 0x4a, 0x56, 0xce, 0x1b, 0x4e, 0x6b, 0x8c, 0xff, 0xa2, 0x1f, 0x75,
	0x56, 0xf3, 0x0d, 0x4f, 0x6a, 0xdb, 0xcb, 0x05, 0x9a, 0xa9, 0xcf, 0x76, 0x8a, 0x48, 0x56, 0xd2,
	0x53, 0xe2, 0x7c, 0x00, 0x3b, 0x3e, 0x9d, 0x3d, 0x1f, 0x7c, 0x67, 0x7f, 0xf0, 0x9d, 0xd7, 0x83,
	0xef, 0x3c, 0x1d, 0xfd, 0xd6, 0xfe, 0xe8, 0xb7, 0x5e, 0x8e, 0x7e, 0xeb, 0x7e, 0xb4, 0x16, 0x90,
	0x16, 0xcb, 0x28, 0x51, 0x72, 0x52, 0xff, 0xf9, 0x71, 0xc6, 0xe1, 0x51, 0xe9, 0xad, 0x2d, 0x26,
	0x65, 0xf3, 0xb1, 0x37, 0x5d, 0x76, 0xed, 0x6b, 0xf8, 0xf3, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x1e,
	0xb3, 0xeb, 0x30, 0x22, 0x02, 0x00, 0x00,
}

func (m *TransferOutParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferOutParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferOutParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxBatching != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxBatching))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxRejectedTransferRetry != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxRejectedTransferRetry))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxKeysignRetry != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxKeysignRetry))
		i--
		dAtA[i] = 0x30
	}
	if m.ExpirationBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ExpirationBlock))
		i--
		dAtA[i] = 0x28
	}
	if m.CommissionRate != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CommissionRate))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SupportedChains) > 0 {
		for iNdEx := len(m.SupportedChains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SupportedChains[iNdEx])
			copy(dAtA[i:], m.SupportedChains[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.SupportedChains[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TransferOutParams) > 0 {
		for iNdEx := len(m.TransferOutParams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TransferOutParams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.MajorityThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MajorityThreshold))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransferOutParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaxBatching != 0 {
		n += 1 + sovParams(uint64(m.MaxBatching))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MajorityThreshold != 0 {
		n += 1 + sovParams(uint64(m.MajorityThreshold))
	}
	if len(m.TransferOutParams) > 0 {
		for _, e := range m.TransferOutParams {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.SupportedChains) > 0 {
		for _, s := range m.SupportedChains {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.CommissionRate != 0 {
		n += 1 + sovParams(uint64(m.CommissionRate))
	}
	if m.ExpirationBlock != 0 {
		n += 1 + sovParams(uint64(m.ExpirationBlock))
	}
	if m.MaxKeysignRetry != 0 {
		n += 1 + sovParams(uint64(m.MaxKeysignRetry))
	}
	if m.MaxRejectedTransferRetry != 0 {
		n += 1 + sovParams(uint64(m.MaxRejectedTransferRetry))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransferOutParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: TransferOutParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferOutParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBatching", wireType)
			}
			m.MaxBatching = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBatching |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MajorityThreshold", wireType)
			}
			m.MajorityThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MajorityThreshold |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferOutParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransferOutParams = append(m.TransferOutParams, &TransferOutParams{})
			if err := m.TransferOutParams[len(m.TransferOutParams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedChains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportedChains = append(m.SupportedChains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
			}
			m.CommissionRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommissionRate |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationBlock", wireType)
			}
			m.ExpirationBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationBlock |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxKeysignRetry", wireType)
			}
			m.MaxKeysignRetry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxKeysignRetry |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRejectedTransferRetry", wireType)
			}
			m.MaxRejectedTransferRetry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRejectedTransferRetry |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
