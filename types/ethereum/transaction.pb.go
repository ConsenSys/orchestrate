// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/ethereum/transaction.proto

package ethereum

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// Transaction data
type TxData struct {
	// QUANTITY - Integer of a nonce.
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// DATA (20 Bytes) - The address of the receiver. null when it’s a contract creation transaction.
	// e.g. 0xAf84242d70aE9D268E2bE3616ED497BA28A7b62C
	To *Account `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// QUANTITY - Integer of the value sent with this transaction.
	// e.g 0xaf23
	Value *Quantity `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// QUANTITY - Integer of the gas provided for the transaction execution.
	Gas uint64 `protobuf:"varint,4,opt,name=gas,proto3" json:"gas,omitempty"`
	// QUANTITY - Integer of the gas price used for each paid gas.
	// e.g 0xaf23b
	GasPrice *Quantity `protobuf:"bytes,5,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// DATA - Hash of the method signature (4 bytes) followed by encoded parameters.
	// e.g 0xa9059cbb000000000000000000000000ff778b716fc07d98839f48ddb88d8be583beb684000000000000000000000000000000000000000000000000002386f26fc10000
	Data                 *Data    `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxData) Reset()         { *m = TxData{} }
func (m *TxData) String() string { return proto.CompactTextString(m) }
func (*TxData) ProtoMessage()    {}
func (*TxData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c15aa666d7a06d, []int{0}
}

func (m *TxData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxData.Unmarshal(m, b)
}
func (m *TxData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxData.Marshal(b, m, deterministic)
}
func (m *TxData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxData.Merge(m, src)
}
func (m *TxData) XXX_Size() int {
	return xxx_messageInfo_TxData.Size(m)
}
func (m *TxData) XXX_DiscardUnknown() {
	xxx_messageInfo_TxData.DiscardUnknown(m)
}

var xxx_messageInfo_TxData proto.InternalMessageInfo

func (m *TxData) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TxData) GetTo() *Account {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TxData) GetValue() *Quantity {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TxData) GetGas() uint64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *TxData) GetGasPrice() *Quantity {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *TxData) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

// Transaction
type Transaction struct {
	TxData *TxData `protobuf:"bytes,1,opt,name=tx_data,json=txData,proto3" json:"tx_data,omitempty"`
	// DATA - The signed, RLP encoded transaction
	Raw *Data `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	// DATA (32 Bytes) - Hash of the transaction.
	// e.g. 0x0a0cafa26ca3f411e6629e9e02c53f23713b0033d7a72e534136104b5447a210
	Hash                 *Hash    `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c15aa666d7a06d, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTxData() *TxData {
	if m != nil {
		return m.TxData
	}
	return nil
}

func (m *Transaction) GetRaw() *Data {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Transaction) GetHash() *Hash {
	if m != nil {
		return m.Hash
	}
	return nil
}

func init() {
	proto.RegisterType((*TxData)(nil), "ethereum.TxData")
	proto.RegisterType((*Transaction)(nil), "ethereum.Transaction")
}

func init() { proto.RegisterFile("types/ethereum/transaction.proto", fileDescriptor_e8c15aa666d7a06d) }

var fileDescriptor_e8c15aa666d7a06d = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x95, 0xfe, 0x09, 0xc5, 0x95, 0x50, 0xb1, 0x18, 0x02, 0x53, 0xe8, 0x54, 0x06, 0x62,
	0x09, 0x3e, 0x01, 0x7f, 0x86, 0x8a, 0x09, 0x42, 0x27, 0x96, 0xea, 0x6a, 0x8e, 0x24, 0xa2, 0xb5,
	0x2b, 0xdf, 0x05, 0x9a, 0x85, 0x0f, 0xc9, 0x27, 0x42, 0x71, 0x5a, 0x05, 0x2a, 0xb1, 0xd9, 0x7a,
	0xbf, 0x77, 0xa7, 0x77, 0x4f, 0xc4, 0x5c, 0xad, 0x91, 0x14, 0x72, 0x8e, 0x0e, 0xcb, 0x95, 0x62,
	0x07, 0x86, 0x40, 0x73, 0x61, 0x4d, 0xb2, 0x76, 0x96, 0xad, 0x1c, 0xec, 0xb4, 0xb3, 0xd3, 0x3d,
	0x76, 0x01, 0x84, 0x0d, 0x34, 0xfe, 0x0e, 0x44, 0x38, 0xdb, 0xdc, 0x03, 0x83, 0x3c, 0x11, 0x7d,
	0x63, 0x8d, 0xc6, 0x28, 0x88, 0x83, 0x49, 0x2f, 0x6d, 0x3e, 0xf2, 0x5c, 0x74, 0xd8, 0x46, 0x9d,
	0x38, 0x98, 0x0c, 0xaf, 0x8e, 0x93, 0xdd, 0x88, 0xe4, 0x46, 0x6b, 0x5b, 0x1a, 0x4e, 0x3b, 0x6c,
	0xe5, 0x44, 0xf4, 0x3f, 0x60, 0x59, 0x62, 0xd4, 0xf5, 0x94, 0x6c, 0xa9, 0xa7, 0x12, 0x0c, 0x17,
	0x5c, 0xa5, 0x0d, 0x20, 0x47, 0xa2, 0x9b, 0x01, 0x45, 0x3d, 0xbf, 0xa0, 0x7e, 0x4a, 0x25, 0x0e,
	0x33, 0xa0, 0xf9, 0xda, 0x15, 0x1a, 0xa3, 0xfe, 0xbf, 0xfe, 0x41, 0x06, 0xf4, 0x58, 0x33, 0x72,
	0x2c, 0x7a, 0xaf, 0xc0, 0x10, 0x85, 0x9e, 0x3d, 0x6a, 0xd9, 0x3a, 0x43, 0xea, 0xb5, 0xf1, 0x97,
	0x18, 0xce, 0xda, 0x73, 0xc8, 0x0b, 0x71, 0xc0, 0x9b, 0xb9, 0x77, 0x05, 0xde, 0x35, 0x6a, 0x5d,
	0x4d, 0xf6, 0x34, 0xe4, 0xe6, 0x06, 0xb1, 0xe8, 0x3a, 0xf8, 0xdc, 0xc6, 0xdd, 0x1f, 0x5e, 0x4b,
	0xf5, 0xfe, 0x1c, 0x28, 0xdf, 0x66, 0xfd, 0x85, 0x4c, 0x81, 0xf2, 0xd4, 0x6b, 0xb7, 0x0f, 0x2f,
	0xd3, 0xac, 0xe0, 0x25, 0x2c, 0x12, 0x6d, 0x57, 0xea, 0xce, 0x1a, 0x42, 0xf3, 0x5c, 0x91, 0xd2,
	0xcb, 0x02, 0x0d, 0xab, 0x37, 0xa7, 0xb4, 0x75, 0x78, 0x49, 0x0c, 0xfa, 0x5d, 0x59, 0xa7, 0x73,
	0x24, 0x76, 0xc0, 0x98, 0x64, 0x05, 0xab, 0xbf, 0x5d, 0x2d, 0x42, 0xdf, 0xd3, 0xf5, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xad, 0xa9, 0x21, 0x78, 0xf0, 0x01, 0x00, 0x00,
}
