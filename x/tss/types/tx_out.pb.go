// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tx_out.proto

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

type TxOutType int32

const (
	TxOutType_NORMAL              TxOutType = 0
	TxOutType_CONTRACT_DEPLOYMENT TxOutType = 1
)

var TxOutType_name = map[int32]string{
	0: "NORMAL",
	1: "CONTRACT_DEPLOYMENT",
}

var TxOutType_value = map[string]int32{
	"NORMAL":              0,
	"CONTRACT_DEPLOYMENT": 1,
}

func (x TxOutType) String() string {
	return proto.EnumName(TxOutType_name, int32(x))
}

func (TxOutType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f8b99e4d8e912782, []int{0}
}

type TxOutWithSigner struct {
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Data   *TxOut `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TxOutWithSigner) Reset()         { *m = TxOutWithSigner{} }
func (m *TxOutWithSigner) String() string { return proto.CompactTextString(m) }
func (*TxOutWithSigner) ProtoMessage()    {}
func (*TxOutWithSigner) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8b99e4d8e912782, []int{0}
}
func (m *TxOutWithSigner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxOutWithSigner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxOutWithSigner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxOutWithSigner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxOutWithSigner.Merge(m, src)
}
func (m *TxOutWithSigner) XXX_Size() int {
	return m.Size()
}
func (m *TxOutWithSigner) XXX_DiscardUnknown() {
	xxx_messageInfo_TxOutWithSigner.DiscardUnknown(m)
}

var xxx_messageInfo_TxOutWithSigner proto.InternalMessageInfo

func (m *TxOutWithSigner) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *TxOutWithSigner) GetData() *TxOut {
	if m != nil {
		return m.Data
	}
	return nil
}

type TxOut struct {
	TxType        TxOutType `protobuf:"varint,1,opt,name=txType,proto3,enum=types.TxOutType" json:"txType,omitempty"`
	InChain       string    `protobuf:"bytes,2,opt,name=inChain,proto3" json:"inChain,omitempty"`
	OutChain      string    `protobuf:"bytes,3,opt,name=outChain,proto3" json:"outChain,omitempty"`
	InBlockHeight int64     `protobuf:"varint,4,opt,name=inBlockHeight,proto3" json:"inBlockHeight,omitempty"`
	InHash        string    `protobuf:"bytes,5,opt,name=inHash,proto3" json:"inHash,omitempty"`
	OutBytes      []byte    `protobuf:"bytes,6,opt,name=outBytes,proto3" json:"outBytes,omitempty"`
	// optional contract hash
	ContractHash string `protobuf:"bytes,7,opt,name=contractHash,proto3" json:"contractHash,omitempty"`
}

func (m *TxOut) Reset()         { *m = TxOut{} }
func (m *TxOut) String() string { return proto.CompactTextString(m) }
func (*TxOut) ProtoMessage()    {}
func (*TxOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8b99e4d8e912782, []int{1}
}
func (m *TxOut) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxOut.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxOut.Merge(m, src)
}
func (m *TxOut) XXX_Size() int {
	return m.Size()
}
func (m *TxOut) XXX_DiscardUnknown() {
	xxx_messageInfo_TxOut.DiscardUnknown(m)
}

var xxx_messageInfo_TxOut proto.InternalMessageInfo

func (m *TxOut) GetTxType() TxOutType {
	if m != nil {
		return m.TxType
	}
	return TxOutType_NORMAL
}

func (m *TxOut) GetInChain() string {
	if m != nil {
		return m.InChain
	}
	return ""
}

func (m *TxOut) GetOutChain() string {
	if m != nil {
		return m.OutChain
	}
	return ""
}

func (m *TxOut) GetInBlockHeight() int64 {
	if m != nil {
		return m.InBlockHeight
	}
	return 0
}

func (m *TxOut) GetInHash() string {
	if m != nil {
		return m.InHash
	}
	return ""
}

func (m *TxOut) GetOutBytes() []byte {
	if m != nil {
		return m.OutBytes
	}
	return nil
}

func (m *TxOut) GetContractHash() string {
	if m != nil {
		return m.ContractHash
	}
	return ""
}

type TxOutConfirmWithSigner struct {
	Signer string        `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Data   *TxOutConfirm `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TxOutConfirmWithSigner) Reset()         { *m = TxOutConfirmWithSigner{} }
func (m *TxOutConfirmWithSigner) String() string { return proto.CompactTextString(m) }
func (*TxOutConfirmWithSigner) ProtoMessage()    {}
func (*TxOutConfirmWithSigner) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8b99e4d8e912782, []int{2}
}
func (m *TxOutConfirmWithSigner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxOutConfirmWithSigner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxOutConfirmWithSigner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxOutConfirmWithSigner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxOutConfirmWithSigner.Merge(m, src)
}
func (m *TxOutConfirmWithSigner) XXX_Size() int {
	return m.Size()
}
func (m *TxOutConfirmWithSigner) XXX_DiscardUnknown() {
	xxx_messageInfo_TxOutConfirmWithSigner.DiscardUnknown(m)
}

var xxx_messageInfo_TxOutConfirmWithSigner proto.InternalMessageInfo

func (m *TxOutConfirmWithSigner) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *TxOutConfirmWithSigner) GetData() *TxOutConfirm {
	if m != nil {
		return m.Data
	}
	return nil
}

type TxOutConfirm struct {
	TxType      TxOutType `protobuf:"varint,1,opt,name=txType,proto3,enum=types.TxOutType" json:"txType,omitempty"`
	OutChain    string    `protobuf:"bytes,2,opt,name=outChain,proto3" json:"outChain,omitempty"`
	OutHash     string    `protobuf:"bytes,3,opt,name=outHash,proto3" json:"outHash,omitempty"`
	BlockHeight int64     `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	// optional contract address
	ContractAddress string `protobuf:"bytes,5,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
}

func (m *TxOutConfirm) Reset()         { *m = TxOutConfirm{} }
func (m *TxOutConfirm) String() string { return proto.CompactTextString(m) }
func (*TxOutConfirm) ProtoMessage()    {}
func (*TxOutConfirm) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8b99e4d8e912782, []int{3}
}
func (m *TxOutConfirm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxOutConfirm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxOutConfirm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxOutConfirm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxOutConfirm.Merge(m, src)
}
func (m *TxOutConfirm) XXX_Size() int {
	return m.Size()
}
func (m *TxOutConfirm) XXX_DiscardUnknown() {
	xxx_messageInfo_TxOutConfirm.DiscardUnknown(m)
}

var xxx_messageInfo_TxOutConfirm proto.InternalMessageInfo

func (m *TxOutConfirm) GetTxType() TxOutType {
	if m != nil {
		return m.TxType
	}
	return TxOutType_NORMAL
}

func (m *TxOutConfirm) GetOutChain() string {
	if m != nil {
		return m.OutChain
	}
	return ""
}

func (m *TxOutConfirm) GetOutHash() string {
	if m != nil {
		return m.OutHash
	}
	return ""
}

func (m *TxOutConfirm) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *TxOutConfirm) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("types.TxOutType", TxOutType_name, TxOutType_value)
	proto.RegisterType((*TxOutWithSigner)(nil), "types.TxOutWithSigner")
	proto.RegisterType((*TxOut)(nil), "types.TxOut")
	proto.RegisterType((*TxOutConfirmWithSigner)(nil), "types.TxOutConfirmWithSigner")
	proto.RegisterType((*TxOutConfirm)(nil), "types.TxOutConfirm")
}

func init() { proto.RegisterFile("tx_out.proto", fileDescriptor_f8b99e4d8e912782) }

var fileDescriptor_f8b99e4d8e912782 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6a, 0xe2, 0x50,
	0x18, 0xc5, 0x73, 0xfd, 0x13, 0xc7, 0xcf, 0xcc, 0x28, 0x57, 0x70, 0x2e, 0xb3, 0x08, 0x21, 0x0c,
	0x4c, 0x98, 0x85, 0x0c, 0xce, 0x13, 0x68, 0x46, 0x10, 0x46, 0xcd, 0x70, 0x27, 0x50, 0x5c, 0x49,
	0xd4, 0xd4, 0x5c, 0xda, 0xe6, 0x4a, 0x72, 0x03, 0xfa, 0x16, 0x7d, 0x9e, 0x3e, 0x41, 0x97, 0x2e,
	0xbb, 0x2c, 0x4a, 0xdf, 0xa3, 0x78, 0x13, 0x25, 0xda, 0x4d, 0xbb, 0xcb, 0x39, 0x27, 0xdf, 0x21,
	0xe7, 0x47, 0x40, 0x13, 0xeb, 0x29, 0x4f, 0x44, 0x7b, 0x15, 0x71, 0xc1, 0x71, 0x59, 0x6c, 0x56,
	0x7e, 0x6c, 0xfe, 0x85, 0xba, 0xbb, 0x76, 0x12, 0x71, 0xc5, 0x44, 0xf0, 0x9f, 0x2d, 0x43, 0x3f,
	0xc2, 0x2d, 0x50, 0x63, 0xf9, 0x44, 0x90, 0x81, 0xac, 0x2a, 0xcd, 0x14, 0x36, 0xa0, 0xb4, 0xf0,
	0x84, 0x47, 0x0a, 0x06, 0xb2, 0x6a, 0x1d, 0xad, 0x2d, 0x0b, 0xda, 0xf2, 0x9a, 0xca, 0xc4, 0x7c,
	0x41, 0x50, 0x96, 0x1a, 0x5b, 0xa0, 0x8a, 0xb5, 0xbb, 0x59, 0xf9, 0xb2, 0xe3, 0x4b, 0xa7, 0x91,
	0x7f, 0xfb, 0xe0, 0xd3, 0x2c, 0xc7, 0x04, 0x2a, 0x2c, 0xb4, 0x03, 0x8f, 0x85, 0xb2, 0xb8, 0x4a,
	0x8f, 0x12, 0x7f, 0x83, 0x4f, 0x3c, 0x11, 0x69, 0x54, 0x94, 0xd1, 0x49, 0xe3, 0xef, 0xf0, 0x99,
	0x85, 0xbd, 0x5b, 0x3e, 0xbf, 0x19, 0xf8, 0x6c, 0x19, 0x08, 0x52, 0x32, 0x90, 0x55, 0xa4, 0xe7,
	0xe6, 0x61, 0x09, 0x0b, 0x07, 0x5e, 0x1c, 0x90, 0x72, 0xba, 0x24, 0x55, 0x59, 0x73, 0x6f, 0x23,
	0xfc, 0x98, 0xa8, 0x06, 0xb2, 0x34, 0x7a, 0xd2, 0xd8, 0x04, 0x6d, 0xce, 0x43, 0x11, 0x79, 0x73,
	0x21, 0x2f, 0x2b, 0xf2, 0xf2, 0xcc, 0x33, 0x27, 0xd0, 0x92, 0x43, 0x6c, 0x1e, 0x5e, 0xb3, 0xe8,
	0xee, 0x1d, 0xec, 0x7e, 0x9c, 0xb1, 0x6b, 0xe6, 0x69, 0x64, 0x25, 0x19, 0xc2, 0x07, 0x04, 0x5a,
	0xde, 0xfe, 0x00, 0xc9, 0x3c, 0xaf, 0xc2, 0x05, 0x2f, 0x02, 0x15, 0x9e, 0xa4, 0x83, 0x52, 0x94,
	0x47, 0x89, 0x0d, 0xa8, 0xcd, 0xde, 0x70, 0xcc, 0x5b, 0xd8, 0x82, 0xfa, 0x71, 0x7d, 0x77, 0xb1,
	0x88, 0xfc, 0x38, 0xce, 0x70, 0x5e, 0xda, 0x3f, 0x7f, 0x41, 0xf5, 0xf4, 0x59, 0x18, 0x40, 0x1d,
	0x3b, 0x74, 0xd4, 0x1d, 0x36, 0x14, 0xfc, 0x15, 0x9a, 0xb6, 0x33, 0x76, 0x69, 0xd7, 0x76, 0xa7,
	0x7f, 0xfa, 0xff, 0x86, 0xce, 0x64, 0xd4, 0x1f, 0xbb, 0x0d, 0xd4, 0x23, 0x8f, 0x3b, 0x1d, 0x6d,
	0x77, 0x3a, 0x7a, 0xde, 0xe9, 0xe8, 0x7e, 0xaf, 0x2b, 0xdb, 0xbd, 0xae, 0x3c, 0xed, 0x75, 0x65,
	0xa6, 0xca, 0xdf, 0xf4, 0xf7, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x6e, 0xbd, 0xc8, 0xb6,
	0x02, 0x00, 0x00,
}

func (m *TxOutWithSigner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxOutWithSigner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxOutWithSigner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTxOut(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxOut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxOut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxOut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractHash) > 0 {
		i -= len(m.ContractHash)
		copy(dAtA[i:], m.ContractHash)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.ContractHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.OutBytes) > 0 {
		i -= len(m.OutBytes)
		copy(dAtA[i:], m.OutBytes)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.OutBytes)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.InHash) > 0 {
		i -= len(m.InHash)
		copy(dAtA[i:], m.InHash)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.InHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.InBlockHeight != 0 {
		i = encodeVarintTxOut(dAtA, i, uint64(m.InBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.OutChain) > 0 {
		i -= len(m.OutChain)
		copy(dAtA[i:], m.OutChain)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.OutChain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InChain) > 0 {
		i -= len(m.InChain)
		copy(dAtA[i:], m.InChain)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.InChain)))
		i--
		dAtA[i] = 0x12
	}
	if m.TxType != 0 {
		i = encodeVarintTxOut(dAtA, i, uint64(m.TxType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TxOutConfirmWithSigner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxOutConfirmWithSigner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxOutConfirmWithSigner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTxOut(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxOutConfirm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxOutConfirm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxOutConfirm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTxOut(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.OutHash) > 0 {
		i -= len(m.OutHash)
		copy(dAtA[i:], m.OutHash)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.OutHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OutChain) > 0 {
		i -= len(m.OutChain)
		copy(dAtA[i:], m.OutChain)
		i = encodeVarintTxOut(dAtA, i, uint64(len(m.OutChain)))
		i--
		dAtA[i] = 0x12
	}
	if m.TxType != 0 {
		i = encodeVarintTxOut(dAtA, i, uint64(m.TxType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTxOut(dAtA []byte, offset int, v uint64) int {
	offset -= sovTxOut(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxOutWithSigner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTxOut(uint64(l))
	}
	return n
}

func (m *TxOut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxType != 0 {
		n += 1 + sovTxOut(uint64(m.TxType))
	}
	l = len(m.InChain)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	l = len(m.OutChain)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	if m.InBlockHeight != 0 {
		n += 1 + sovTxOut(uint64(m.InBlockHeight))
	}
	l = len(m.InHash)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	l = len(m.OutBytes)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	l = len(m.ContractHash)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	return n
}

func (m *TxOutConfirmWithSigner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTxOut(uint64(l))
	}
	return n
}

func (m *TxOutConfirm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxType != 0 {
		n += 1 + sovTxOut(uint64(m.TxType))
	}
	l = len(m.OutChain)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	l = len(m.OutHash)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTxOut(uint64(m.BlockHeight))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTxOut(uint64(l))
	}
	return n
}

func sovTxOut(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTxOut(x uint64) (n int) {
	return sovTxOut(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxOutWithSigner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxOut
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
			return fmt.Errorf("proto: TxOutWithSigner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxOutWithSigner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &TxOut{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxOut(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxOut
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
func (m *TxOut) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxOut
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
			return fmt.Errorf("proto: TxOut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxOut: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			m.TxType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxType |= TxOutType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InBlockHeight", wireType)
			}
			m.InBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutBytes = append(m.OutBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.OutBytes == nil {
				m.OutBytes = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxOut(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxOut
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
func (m *TxOutConfirmWithSigner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxOut
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
			return fmt.Errorf("proto: TxOutConfirmWithSigner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxOutConfirmWithSigner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &TxOutConfirm{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxOut(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxOut
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
func (m *TxOutConfirm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxOut
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
			return fmt.Errorf("proto: TxOutConfirm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxOutConfirm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			m.TxType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxType |= TxOutType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxOut
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
				return ErrInvalidLengthTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxOut(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxOut
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
func skipTxOut(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxOut
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
					return 0, ErrIntOverflowTxOut
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
					return 0, ErrIntOverflowTxOut
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
				return 0, ErrInvalidLengthTxOut
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTxOut
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTxOut
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTxOut        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxOut          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTxOut = fmt.Errorf("proto: unexpected end of group")
)
