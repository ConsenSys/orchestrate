// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/envelope-store/proto (interfaces: EnvelopeStoreClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/envelope-store/proto"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockEnvelopeStoreClient is a mock of EnvelopeStoreClient interface
type MockEnvelopeStoreClient struct {
	ctrl     *gomock.Controller
	recorder *MockEnvelopeStoreClientMockRecorder
}

// MockEnvelopeStoreClientMockRecorder is the mock recorder for MockEnvelopeStoreClient
type MockEnvelopeStoreClientMockRecorder struct {
	mock *MockEnvelopeStoreClient
}

// NewMockEnvelopeStoreClient creates a new mock instance
func NewMockEnvelopeStoreClient(ctrl *gomock.Controller) *MockEnvelopeStoreClient {
	mock := &MockEnvelopeStoreClient{ctrl: ctrl}
	mock.recorder = &MockEnvelopeStoreClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvelopeStoreClient) EXPECT() *MockEnvelopeStoreClientMockRecorder {
	return m.recorder
}

// LoadByID mocks base method
func (m *MockEnvelopeStoreClient) LoadByID(arg0 context.Context, arg1 *proto.LoadByIDRequest, arg2 ...grpc.CallOption) (*proto.StoreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoadByID", varargs...)
	ret0, _ := ret[0].(*proto.StoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadByID indicates an expected call of LoadByID
func (mr *MockEnvelopeStoreClientMockRecorder) LoadByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadByID", reflect.TypeOf((*MockEnvelopeStoreClient)(nil).LoadByID), varargs...)
}

// LoadByTxHash mocks base method
func (m *MockEnvelopeStoreClient) LoadByTxHash(arg0 context.Context, arg1 *proto.LoadByTxHashRequest, arg2 ...grpc.CallOption) (*proto.StoreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoadByTxHash", varargs...)
	ret0, _ := ret[0].(*proto.StoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadByTxHash indicates an expected call of LoadByTxHash
func (mr *MockEnvelopeStoreClientMockRecorder) LoadByTxHash(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadByTxHash", reflect.TypeOf((*MockEnvelopeStoreClient)(nil).LoadByTxHash), varargs...)
}

// LoadPending mocks base method
func (m *MockEnvelopeStoreClient) LoadPending(arg0 context.Context, arg1 *proto.LoadPendingRequest, arg2 ...grpc.CallOption) (*proto.LoadPendingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoadPending", varargs...)
	ret0, _ := ret[0].(*proto.LoadPendingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadPending indicates an expected call of LoadPending
func (mr *MockEnvelopeStoreClientMockRecorder) LoadPending(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPending", reflect.TypeOf((*MockEnvelopeStoreClient)(nil).LoadPending), varargs...)
}

// SetStatus mocks base method
func (m *MockEnvelopeStoreClient) SetStatus(arg0 context.Context, arg1 *proto.SetStatusRequest, arg2 ...grpc.CallOption) (*proto.StatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetStatus", varargs...)
	ret0, _ := ret[0].(*proto.StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetStatus indicates an expected call of SetStatus
func (mr *MockEnvelopeStoreClientMockRecorder) SetStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockEnvelopeStoreClient)(nil).SetStatus), varargs...)
}

// Store mocks base method
func (m *MockEnvelopeStoreClient) Store(arg0 context.Context, arg1 *proto.StoreRequest, arg2 ...grpc.CallOption) (*proto.StoreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Store", varargs...)
	ret0, _ := ret[0].(*proto.StoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockEnvelopeStoreClientMockRecorder) Store(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockEnvelopeStoreClient)(nil).Store), varargs...)
}
