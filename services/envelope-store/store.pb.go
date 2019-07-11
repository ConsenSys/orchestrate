// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/envelope-store/store.proto

package envelope_store

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	chain "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/chain"
	envelope "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/envelope"
	error1 "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/error"
	grpc "google.golang.org/grpc"
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

type StoreRequest struct {
	Envelope             *envelope.Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3251193be0eb6bb8, []int{0}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetEnvelope() *envelope.Envelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

type StoreResponse struct {
	// Status of Envelope
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// Last update date of envelope stored
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	// Envelope
	Envelope *envelope.Envelope `protobuf:"bytes,3,opt,name=envelope,proto3" json:"envelope,omitempty"`
	// Error
	Err                  *error1.Error `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3251193be0eb6bb8, []int{1}
}

func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (m *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(m, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StoreResponse) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func (m *StoreResponse) GetEnvelope() *envelope.Envelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

func (m *StoreResponse) GetErr() *error1.Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type TxHashRequest struct {
	// Chain ID the transaction has been sent to
	ChainId *chain.Chain `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Hash of the transaction
	TxHash               string   `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxHashRequest) Reset()         { *m = TxHashRequest{} }
func (m *TxHashRequest) String() string { return proto.CompactTextString(m) }
func (*TxHashRequest) ProtoMessage()    {}
func (*TxHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3251193be0eb6bb8, []int{2}
}

func (m *TxHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxHashRequest.Unmarshal(m, b)
}
func (m *TxHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxHashRequest.Marshal(b, m, deterministic)
}
func (m *TxHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxHashRequest.Merge(m, src)
}
func (m *TxHashRequest) XXX_Size() int {
	return xxx_messageInfo_TxHashRequest.Size(m)
}
func (m *TxHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxHashRequest proto.InternalMessageInfo

func (m *TxHashRequest) GetChainId() *chain.Chain {
	if m != nil {
		return m.ChainId
	}
	return nil
}

func (m *TxHashRequest) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

type IDRequest struct {
	// Envelope identifier
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDRequest) Reset()         { *m = IDRequest{} }
func (m *IDRequest) String() string { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()    {}
func (*IDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3251193be0eb6bb8, []int{3}
}

func (m *IDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDRequest.Unmarshal(m, b)
}
func (m *IDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDRequest.Marshal(b, m, deterministic)
}
func (m *IDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDRequest.Merge(m, src)
}
func (m *IDRequest) XXX_Size() int {
	return xxx_messageInfo_IDRequest.Size(m)
}
func (m *IDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDRequest proto.InternalMessageInfo

func (m *IDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SetStatusRequest struct {
	// Envelope identifier
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Status
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetStatusRequest) Reset()         { *m = SetStatusRequest{} }
func (m *SetStatusRequest) String() string { return proto.CompactTextString(m) }
func (*SetStatusRequest) ProtoMessage()    {}
func (*SetStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3251193be0eb6bb8, []int{4}
}

func (m *SetStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStatusRequest.Unmarshal(m, b)
}
func (m *SetStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStatusRequest.Marshal(b, m, deterministic)
}
func (m *SetStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStatusRequest.Merge(m, src)
}
func (m *SetStatusRequest) XXX_Size() int {
	return xxx_messageInfo_SetStatusRequest.Size(m)
}
func (m *SetStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetStatusRequest proto.InternalMessageInfo

func (m *SetStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SetStatusRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type LoadPendingRequest struct {
	// Pending duration in nanoseconds
	Duration             int64    `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadPendingRequest) Reset()         { *m = LoadPendingRequest{} }
func (m *LoadPendingRequest) String() string { return proto.CompactTextString(m) }
func (*LoadPendingRequest) ProtoMessage()    {}
func (*LoadPendingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3251193be0eb6bb8, []int{5}
}

func (m *LoadPendingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadPendingRequest.Unmarshal(m, b)
}
func (m *LoadPendingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadPendingRequest.Marshal(b, m, deterministic)
}
func (m *LoadPendingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadPendingRequest.Merge(m, src)
}
func (m *LoadPendingRequest) XXX_Size() int {
	return xxx_messageInfo_LoadPendingRequest.Size(m)
}
func (m *LoadPendingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadPendingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadPendingRequest proto.InternalMessageInfo

func (m *LoadPendingRequest) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type LoadPendingResponse struct {
	// Pending envelopes
	Envelopes []*envelope.Envelope `protobuf:"bytes,1,rep,name=envelopes,proto3" json:"envelopes,omitempty"`
	// Error
	Err                  *error1.Error `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LoadPendingResponse) Reset()         { *m = LoadPendingResponse{} }
func (m *LoadPendingResponse) String() string { return proto.CompactTextString(m) }
func (*LoadPendingResponse) ProtoMessage()    {}
func (*LoadPendingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3251193be0eb6bb8, []int{6}
}

func (m *LoadPendingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadPendingResponse.Unmarshal(m, b)
}
func (m *LoadPendingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadPendingResponse.Marshal(b, m, deterministic)
}
func (m *LoadPendingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadPendingResponse.Merge(m, src)
}
func (m *LoadPendingResponse) XXX_Size() int {
	return xxx_messageInfo_LoadPendingResponse.Size(m)
}
func (m *LoadPendingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadPendingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadPendingResponse proto.InternalMessageInfo

func (m *LoadPendingResponse) GetEnvelopes() []*envelope.Envelope {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

func (m *LoadPendingResponse) GetErr() *error1.Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*StoreRequest)(nil), "envelopestore.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "envelopestore.StoreResponse")
	proto.RegisterType((*TxHashRequest)(nil), "envelopestore.TxHashRequest")
	proto.RegisterType((*IDRequest)(nil), "envelopestore.IDRequest")
	proto.RegisterType((*SetStatusRequest)(nil), "envelopestore.SetStatusRequest")
	proto.RegisterType((*LoadPendingRequest)(nil), "envelopestore.LoadPendingRequest")
	proto.RegisterType((*LoadPendingResponse)(nil), "envelopestore.LoadPendingResponse")
}

func init() {
	proto.RegisterFile("services/envelope-store/store.proto", fileDescriptor_3251193be0eb6bb8)
}

var fileDescriptor_3251193be0eb6bb8 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0x5e, 0xdb, 0xdf, 0xaf, 0x6b, 0x5e, 0x5b, 0x84, 0x8c, 0xc4, 0xaa, 0x0c, 0xd8, 0x08, 0x07,
	0xb8, 0x60, 0x4f, 0xe3, 0x86, 0x80, 0xc3, 0xd6, 0xc2, 0x2a, 0xed, 0x00, 0xe9, 0xe0, 0xc0, 0xa5,
	0x72, 0x13, 0x2f, 0xb5, 0xd6, 0xc6, 0xc1, 0x76, 0xa7, 0xf5, 0x1f, 0xe3, 0xc8, 0xdf, 0x86, 0x62,
	0xc7, 0xa1, 0x09, 0x94, 0x1e, 0xb8, 0x3c, 0xd9, 0x79, 0xdf, 0xfb, 0xfc, 0xf9, 0x7b, 0x2f, 0x86,
	0x67, 0x8a, 0xc9, 0x5b, 0x1e, 0x31, 0x45, 0x58, 0x7a, 0xcb, 0x16, 0x22, 0x63, 0x2f, 0x95, 0x16,
	0x92, 0x11, 0x13, 0x71, 0x26, 0x85, 0x16, 0xa8, 0xef, 0x72, 0xe6, 0xa3, 0xff, 0x58, 0xaf, 0xb3,
	0x8d, 0x82, 0x72, 0x61, 0xd1, 0xfe, 0x81, 0x4d, 0x47, 0x73, 0xca, 0x53, 0x1b, 0xab, 0x09, 0x26,
	0xa5, 0x90, 0x36, 0x16, 0x89, 0xa3, 0x44, 0x88, 0x64, 0xc1, 0x88, 0xd9, 0xcd, 0x56, 0xd7, 0x44,
	0xf3, 0x25, 0x53, 0x9a, 0x2e, 0x33, 0x0b, 0x08, 0xde, 0x41, 0x6f, 0x92, 0x1f, 0x1d, 0xb2, 0x6f,
	0x2b, 0xa6, 0x34, 0xc2, 0xd0, 0x71, 0x87, 0x0e, 0x1a, 0xc7, 0x8d, 0x17, 0xdd, 0x53, 0x84, 0x4b,
	0x15, 0xa3, 0x62, 0x11, 0x96, 0x98, 0xe0, 0x7b, 0x03, 0xfa, 0x05, 0x81, 0xca, 0x44, 0xaa, 0x18,
	0x7a, 0x08, 0x6d, 0xa5, 0xa9, 0x5e, 0x29, 0x53, 0xef, 0x85, 0xc5, 0x0e, 0xbd, 0x85, 0xde, 0x82,
	0x2a, 0x3d, 0x5d, 0x65, 0x31, 0xd5, 0x2c, 0x1e, 0x34, 0x0d, 0xbb, 0x8f, 0xad, 0x42, 0xec, 0x14,
	0xe2, 0x2b, 0xa7, 0x30, 0xec, 0xe6, 0xf8, 0xcf, 0x16, 0x5e, 0x11, 0xd6, 0xda, 0x2d, 0x0c, 0x3d,
	0x81, 0x16, 0x93, 0x72, 0xf0, 0x9f, 0x81, 0xf6, 0xb0, 0x35, 0x65, 0x94, 0xc7, 0x30, 0x4f, 0x04,
	0x9f, 0xa0, 0x7f, 0x75, 0x77, 0x41, 0xd5, 0xdc, 0xdd, 0xfc, 0x39, 0x74, 0x8c, 0xa5, 0x53, 0x1e,
	0x17, 0x37, 0xef, 0x61, 0xeb, 0xf1, 0x79, 0x1e, 0xc3, 0x7d, 0xb3, 0x19, 0xc7, 0xe8, 0x00, 0xf6,
	0xf5, 0xdd, 0x74, 0x4e, 0xd5, 0xdc, 0xdc, 0xc1, 0x0b, 0xdb, 0xda, 0x10, 0x05, 0x87, 0xe0, 0x8d,
	0x87, 0x8e, 0xee, 0x1e, 0x34, 0x0b, 0x22, 0x2f, 0x6c, 0xf2, 0x38, 0x78, 0x0d, 0xf7, 0x27, 0x4c,
	0x4f, 0x8c, 0x17, 0x5b, 0x30, 0x1b, 0xd6, 0x35, 0x37, 0xad, 0x0b, 0x4e, 0x00, 0x5d, 0x0a, 0x1a,
	0x7f, 0x64, 0x69, 0xcc, 0xd3, 0xc4, 0x55, 0xfb, 0xd0, 0x89, 0x57, 0x92, 0x6a, 0x2e, 0x52, 0xc3,
	0xd1, 0x0a, 0xcb, 0x7d, 0x90, 0xc0, 0x83, 0x4a, 0x45, 0xd1, 0x9b, 0x13, 0xf0, 0xca, 0x81, 0x1b,
	0x34, 0x8e, 0x5b, 0x5b, 0x5c, 0xfc, 0x05, 0x72, 0x36, 0x36, 0xb7, 0xd8, 0x78, 0xfa, 0xa3, 0x05,
	0xff, 0x9b, 0xfe, 0xa3, 0xa1, 0x5b, 0x1c, 0xe2, 0xca, 0x50, 0xe3, 0xcd, 0xf9, 0xf2, 0x1f, 0xfd,
	0x39, 0x69, 0xf5, 0x05, 0x7b, 0xe8, 0x12, 0x7a, 0xb9, 0xf0, 0xb3, 0xb5, 0x6d, 0x0e, 0xaa, 0xe3,
	0x2b, 0x3d, 0xdb, 0xc9, 0x36, 0x84, 0x8e, 0x65, 0x1b, 0x0f, 0xd1, 0xa0, 0x86, 0x2d, 0x5b, 0xb5,
	0x93, 0x65, 0x04, 0xde, 0x07, 0xd7, 0xba, 0x7f, 0xa0, 0x79, 0x03, 0x5e, 0x39, 0x01, 0xe8, 0xa8,
	0x0e, 0xae, 0xcd, 0x86, 0x5f, 0xf1, 0x3a, 0xd8, 0x43, 0x5f, 0xa0, 0xbb, 0xd1, 0x51, 0xf4, 0xb4,
	0x56, 0xff, 0xfb, 0x7c, 0xf8, 0xc1, 0xdf, 0x20, 0x4e, 0xd5, 0xd9, 0xc5, 0xd7, 0xf7, 0x09, 0xd7,
	0x0b, 0x3a, 0xc3, 0x91, 0x58, 0x92, 0xf3, 0xfc, 0x6b, 0x3a, 0x59, 0x2b, 0x12, 0x2d, 0x38, 0x4b,
	0x35, 0xb9, 0x96, 0x24, 0x12, 0x32, 0x7f, 0xba, 0x68, 0x74, 0x43, 0xb2, 0x9b, 0x04, 0x27, 0x5c,
	0x93, 0xea, 0x33, 0x65, 0xdf, 0xb5, 0x59, 0xdb, 0xfc, 0xc2, 0xaf, 0x7e, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x4c, 0x89, 0xd7, 0xae, 0xf9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	// Store an envelope
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	// LoadByTxHash load an envelope by transaction hash
	LoadByTxHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	// LoadByID load an envelope by identifier
	LoadByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	// GetStatus returns envelope status
	GetStatus(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	// SetStatus set an envelope status
	SetStatus(ctx context.Context, in *SetStatusRequest, opts ...grpc.CallOption) (*error1.Error, error)
	// LoadPending load envelopes of pending transactions
	LoadPending(ctx context.Context, in *LoadPendingRequest, opts ...grpc.CallOption) (*LoadPendingResponse, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) LoadByTxHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/LoadByTxHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) LoadByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/LoadByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetStatus(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) SetStatus(ctx context.Context, in *SetStatusRequest, opts ...grpc.CallOption) (*error1.Error, error) {
	out := new(error1.Error)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/SetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) LoadPending(ctx context.Context, in *LoadPendingRequest, opts ...grpc.CallOption) (*LoadPendingResponse, error) {
	out := new(LoadPendingResponse)
	err := c.cc.Invoke(ctx, "/envelopestore.Store/LoadPending", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	// Store an envelope
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
	// LoadByTxHash load an envelope by transaction hash
	LoadByTxHash(context.Context, *TxHashRequest) (*StoreResponse, error)
	// LoadByID load an envelope by identifier
	LoadByID(context.Context, *IDRequest) (*StoreResponse, error)
	// GetStatus returns envelope status
	GetStatus(context.Context, *IDRequest) (*StoreResponse, error)
	// SetStatus set an envelope status
	SetStatus(context.Context, *SetStatusRequest) (*error1.Error, error)
	// LoadPending load envelopes of pending transactions
	LoadPending(context.Context, *LoadPendingRequest) (*LoadPendingResponse, error)
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_LoadByTxHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).LoadByTxHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/LoadByTxHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).LoadByTxHash(ctx, req.(*TxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_LoadByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).LoadByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/LoadByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).LoadByID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetStatus(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_SetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).SetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/SetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).SetStatus(ctx, req.(*SetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_LoadPending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadPendingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).LoadPending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envelopestore.Store/LoadPending",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).LoadPending(ctx, req.(*LoadPendingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envelopestore.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Store_Store_Handler,
		},
		{
			MethodName: "LoadByTxHash",
			Handler:    _Store_LoadByTxHash_Handler,
		},
		{
			MethodName: "LoadByID",
			Handler:    _Store_LoadByID_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Store_GetStatus_Handler,
		},
		{
			MethodName: "SetStatus",
			Handler:    _Store_SetStatus_Handler,
		},
		{
			MethodName: "LoadPending",
			Handler:    _Store_LoadPending_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/envelope-store/store.proto",
}
