// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/trace/trace.proto

package trace

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/protos/common"
	ethereum "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/protos/ethereum"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Metadata attached to the trace
type Metadata struct {
	// ID of the trace  in UUID RFC 4122, ISO/IEC 9834-8:2005 format
	// e.g a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Extra information (optional)
	Extra                map[string]string `protobuf:"bytes,2,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_d68378d0a5cb15c9, []int{0}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Metadata) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

// Trace object storing contextual information about transaction execution
type Trace struct {
	// Chain to send transaction to
	Chain *common.Chain `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	// Ethereum transaction call information
	// Sender for the transaction
	Sender   *common.Account `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver *common.Account `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Call     *common.Call    `protobuf:"bytes,4,opt,name=call,proto3" json:"call,omitempty"`
	// Transaction
	Tx *ethereum.Transaction `protobuf:"bytes,5,opt,name=tx,proto3" json:"tx,omitempty"`
	// Receipt
	Receipt *ethereum.Receipt `protobuf:"bytes,6,opt,name=receipt,proto3" json:"receipt,omitempty"`
	// Errors cought during transaction execution
	Errors []*common.Error `protobuf:"bytes,7,rep,name=errors,proto3" json:"errors,omitempty"`
	// Metadata of the trace
	Metadata             *Metadata `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_d68378d0a5cb15c9, []int{1}
}

func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetChain() *common.Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Trace) GetSender() *common.Account {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Trace) GetReceiver() *common.Account {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *Trace) GetCall() *common.Call {
	if m != nil {
		return m.Call
	}
	return nil
}

func (m *Trace) GetTx() *ethereum.Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *Trace) GetReceipt() *ethereum.Receipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *Trace) GetErrors() []*common.Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Trace) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "trace.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "trace.Metadata.ExtraEntry")
	proto.RegisterType((*Trace)(nil), "trace.Trace")
}

func init() { proto.RegisterFile("protos/trace/trace.proto", fileDescriptor_d68378d0a5cb15c9) }

var fileDescriptor_d68378d0a5cb15c9 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x6a, 0x1b, 0x31,
	0x10, 0xc7, 0xb1, 0x9c, 0x75, 0xdc, 0x71, 0x3f, 0x45, 0x0b, 0xaa, 0x4b, 0xc1, 0x4d, 0x09, 0x0d,
	0x84, 0x4a, 0xc5, 0xbd, 0x84, 0x9e, 0xda, 0x04, 0x1f, 0x7b, 0x51, 0x73, 0xea, 0x4d, 0x96, 0x27,
	0x8e, 0xf0, 0x5a, 0x32, 0x5a, 0x39, 0xd8, 0x2f, 0xd0, 0xa7, 0xeb, 0x43, 0x05, 0x7d, 0xec, 0x9a,
	0x84, 0x5c, 0x8c, 0x35, 0xbf, 0x9f, 0x66, 0xf5, 0x1f, 0x09, 0xd8, 0xc6, 0xbb, 0xe0, 0x1a, 0x11,
	0xbc, 0xd2, 0x98, 0x7f, 0x79, 0x2a, 0xd1, 0x2a, 0x2d, 0xc6, 0x1f, 0x8a, 0xa0, 0xdd, 0x7a, 0xed,
	0xac, 0x50, 0x5a, 0xbb, 0xad, 0x0d, 0xd9, 0x19, 0xb3, 0x87, 0x50, 0xab, 0xba, 0x2e, 0xe4, 0xfd,
	0x23, 0x72, 0xab, 0x8c, 0x7d, 0x1a, 0xa1, 0xf7, 0xce, 0x17, 0xf4, 0xa9, 0x20, 0x0c, 0xb7, 0xe8,
	0x71, 0xbb, 0x8e, 0x07, 0xb2, 0x8d, 0xd2, 0xc1, 0xb8, 0x76, 0xf7, 0xc7, 0xc7, 0x8a, 0x47, 0x8d,
	0x66, 0x53, 0x4e, 0x74, 0xf2, 0xaf, 0x07, 0xc3, 0xdf, 0x18, 0xd4, 0x42, 0x05, 0x45, 0x5f, 0x02,
	0x31, 0x0b, 0xd6, 0x9b, 0xf4, 0xce, 0x9e, 0x49, 0x62, 0x16, 0xf4, 0x1b, 0x54, 0xb8, 0x0b, 0x5e,
	0x31, 0x32, 0xe9, 0x9f, 0x8d, 0xa6, 0x63, 0x9e, 0xf3, 0xb6, 0x3e, 0x9f, 0x45, 0x38, 0xb3, 0xc1,
	0xef, 0x65, 0x16, 0xc7, 0x17, 0x00, 0x87, 0x22, 0x7d, 0x0d, 0xfd, 0x15, 0xee, 0x4b, 0xc3, 0xf8,
	0x97, 0xbe, 0x85, 0xea, 0x4e, 0xd5, 0x5b, 0x64, 0x24, 0xd5, 0xf2, 0xe2, 0x07, 0xb9, 0xe8, 0x9d,
	0xfc, 0x27, 0x50, 0x5d, 0xc7, 0xf6, 0xf4, 0x33, 0x54, 0x29, 0x7e, 0xda, 0x37, 0x9a, 0xbe, 0xe0,
	0x39, 0x38, 0xbf, 0x8a, 0x45, 0x99, 0x19, 0xfd, 0x02, 0x83, 0x06, 0xed, 0x02, 0x7d, 0xea, 0x34,
	0x9a, 0xbe, 0x6a, 0xad, 0x5f, 0x79, 0xe0, 0xb2, 0x60, 0x7a, 0x0e, 0xc3, 0x94, 0xf8, 0x0e, 0x3d,
	0xeb, 0x3f, 0xad, 0x76, 0x02, 0x9d, 0xc0, 0x51, 0xbc, 0x13, 0x76, 0x94, 0xc4, 0xe7, 0xdd, 0x97,
	0x55, 0x5d, 0xcb, 0x44, 0xe8, 0x29, 0x90, 0xb0, 0x63, 0x55, 0xe2, 0xef, 0x78, 0x3b, 0x54, 0x7e,
	0x7d, 0x98, 0xbb, 0x24, 0x61, 0x47, 0xcf, 0xe1, 0xb8, 0xcc, 0x99, 0x0d, 0x92, 0xfb, 0xe6, 0xe0,
	0xca, 0x0c, 0x64, 0x6b, 0xd0, 0x53, 0x18, 0xa4, 0x4b, 0x6d, 0xd8, 0x71, 0x9a, 0x73, 0x97, 0x78,
	0x16, 0xab, 0xb2, 0xc0, 0x98, 0x64, 0x5d, 0x26, 0xcf, 0x86, 0x25, 0xc9, 0xc3, 0x0b, 0x91, 0x9d,
	0x70, 0x79, 0xf9, 0xf7, 0xe7, 0xd2, 0x84, 0x5a, 0xcd, 0x63, 0x2f, 0x71, 0xe5, 0x6c, 0x83, 0xf6,
	0xcf, 0xbe, 0x11, 0xba, 0x36, 0x68, 0x83, 0xb8, 0xf1, 0x42, 0x3b, 0x8f, 0x5f, 0x9b, 0xa0, 0xf4,
	0x4a, 0x6c, 0x56, 0x4b, 0xbe, 0x34, 0x41, 0xa4, 0x27, 0x31, 0xdf, 0xde, 0xe4, 0x77, 0x3d, 0x1f,
	0xa4, 0xf5, 0xf7, 0xfb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x85, 0x87, 0x2b, 0xf4, 0x02, 0x00,
	0x00,
}
