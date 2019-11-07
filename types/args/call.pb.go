// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/args/call.proto

package args

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	abi "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/types/abi"
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

// Call contains information for a contract call
type Call struct {
	// Contract to call method on
	Contract *abi.Contract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// Method to call
	Method *abi.Method `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// Arguments to feed on transaction call
	Args                 []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}
func (*Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9df5d94cbe4727e, []int{0}
}

func (m *Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call.Unmarshal(m, b)
}
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call.Marshal(b, m, deterministic)
}
func (m *Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call.Merge(m, src)
}
func (m *Call) XXX_Size() int {
	return xxx_messageInfo_Call.Size(m)
}
func (m *Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Call proto.InternalMessageInfo

func (m *Call) GetContract() *abi.Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *Call) GetMethod() *abi.Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *Call) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Call)(nil), "args.Call")
}

func init() { proto.RegisterFile("types/args/call.proto", fileDescriptor_a9df5d94cbe4727e) }

var fileDescriptor_a9df5d94cbe4727e = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x31, 0x4b, 0xc5, 0x30,
	0x10, 0x80, 0xa9, 0x2d, 0x45, 0x53, 0x5c, 0x22, 0x42, 0x71, 0x2a, 0xba, 0xd4, 0xc1, 0x1c, 0xe8,
	0x3f, 0xb0, 0xe2, 0xe6, 0x52, 0x37, 0xb7, 0xf4, 0x8c, 0x6d, 0x7c, 0x69, 0xae, 0x24, 0xb7, 0xf4,
	0xdf, 0x3f, 0x9a, 0x57, 0xde, 0x1b, 0x02, 0xe1, 0xfb, 0x3e, 0xb8, 0x3b, 0x71, 0xcf, 0xeb, 0x62,
	0x22, 0xe8, 0x30, 0x46, 0x40, 0xed, 0x9c, 0x5a, 0x02, 0x31, 0xc9, 0x62, 0x03, 0x0f, 0x77, 0xbb,
	0x1c, 0xec, 0xf6, 0x4e, 0xea, 0xf1, 0x5f, 0x14, 0x9d, 0x76, 0x4e, 0x3e, 0x8b, 0x6b, 0x24, 0xcf,
	0x41, 0x23, 0xd7, 0x59, 0x93, 0xb5, 0xd5, 0xeb, 0xad, 0xda, 0xaa, 0x6e, 0x87, 0xfd, 0x59, 0xcb,
	0x27, 0x51, 0xce, 0x86, 0x27, 0xfa, 0xad, 0xaf, 0x52, 0x58, 0xa5, 0xf0, 0x2b, 0xa1, 0x7e, 0x57,
	0x52, 0x8a, 0x34, 0xb4, 0xce, 0x9b, 0xbc, 0xbd, 0xe9, 0xd3, 0xff, 0xfd, 0xf3, 0xe7, 0x63, 0xb4,
	0xec, 0xf4, 0xa0, 0x90, 0x66, 0xe8, 0xc8, 0x47, 0xe3, 0xbf, 0xd7, 0x08, 0xe8, 0xac, 0xf1, 0x0c,
	0x7f, 0x01, 0x90, 0x82, 0x79, 0x89, 0xac, 0xf1, 0x00, 0x14, 0x70, 0x32, 0x91, 0x83, 0x66, 0xa3,
	0x46, 0xcb, 0x70, 0xb9, 0x6c, 0x28, 0xd3, 0xea, 0x6f, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21,
	0x20, 0x70, 0x77, 0xee, 0x00, 0x00, 0x00,
}
