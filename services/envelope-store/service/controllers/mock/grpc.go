// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/envelope-store/proto (interfaces: EnvelopeStoreServer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/envelope-store/proto"
	reflect "reflect"
)

// MockEnvelopeStoreServer is a mock of EnvelopeStoreServer interface
type MockEnvelopeStoreServer struct {
	ctrl     *gomock.Controller
	recorder *MockEnvelopeStoreServerMockRecorder
}

// MockEnvelopeStoreServerMockRecorder is the mock recorder for MockEnvelopeStoreServer
type MockEnvelopeStoreServerMockRecorder struct {
	mock *MockEnvelopeStoreServer
}

// NewMockEnvelopeStoreServer creates a new mock instance
func NewMockEnvelopeStoreServer(ctrl *gomock.Controller) *MockEnvelopeStoreServer {
	mock := &MockEnvelopeStoreServer{ctrl: ctrl}
	mock.recorder = &MockEnvelopeStoreServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvelopeStoreServer) EXPECT() *MockEnvelopeStoreServerMockRecorder {
	return m.recorder
}

// LoadByID mocks base method
func (m *MockEnvelopeStoreServer) LoadByID(arg0 context.Context, arg1 *proto.LoadByIDRequest) (*proto.StoreResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadByID", arg0, arg1)
	ret0, _ := ret[0].(*proto.StoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadByID indicates an expected call of LoadByID
func (mr *MockEnvelopeStoreServerMockRecorder) LoadByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadByID", reflect.TypeOf((*MockEnvelopeStoreServer)(nil).LoadByID), arg0, arg1)
}

// LoadByTxHash mocks base method
func (m *MockEnvelopeStoreServer) LoadByTxHash(arg0 context.Context, arg1 *proto.LoadByTxHashRequest) (*proto.StoreResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadByTxHash", arg0, arg1)
	ret0, _ := ret[0].(*proto.StoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadByTxHash indicates an expected call of LoadByTxHash
func (mr *MockEnvelopeStoreServerMockRecorder) LoadByTxHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadByTxHash", reflect.TypeOf((*MockEnvelopeStoreServer)(nil).LoadByTxHash), arg0, arg1)
}

// LoadByTxHashes mocks base method
func (m *MockEnvelopeStoreServer) LoadByTxHashes(arg0 context.Context, arg1 *proto.LoadByTxHashesRequest) (*proto.LoadByTxHashesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadByTxHashes", arg0, arg1)
	ret0, _ := ret[0].(*proto.LoadByTxHashesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadByTxHashes indicates an expected call of LoadByTxHashes
func (mr *MockEnvelopeStoreServerMockRecorder) LoadByTxHashes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadByTxHashes", reflect.TypeOf((*MockEnvelopeStoreServer)(nil).LoadByTxHashes), arg0, arg1)
}

// LoadPending mocks base method
func (m *MockEnvelopeStoreServer) LoadPending(arg0 context.Context, arg1 *proto.LoadPendingRequest) (*proto.LoadPendingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPending", arg0, arg1)
	ret0, _ := ret[0].(*proto.LoadPendingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadPending indicates an expected call of LoadPending
func (mr *MockEnvelopeStoreServerMockRecorder) LoadPending(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPending", reflect.TypeOf((*MockEnvelopeStoreServer)(nil).LoadPending), arg0, arg1)
}

// SetStatus mocks base method
func (m *MockEnvelopeStoreServer) SetStatus(arg0 context.Context, arg1 *proto.SetStatusRequest) (*proto.StatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", arg0, arg1)
	ret0, _ := ret[0].(*proto.StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetStatus indicates an expected call of SetStatus
func (mr *MockEnvelopeStoreServerMockRecorder) SetStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockEnvelopeStoreServer)(nil).SetStatus), arg0, arg1)
}

// Store mocks base method
func (m *MockEnvelopeStoreServer) Store(arg0 context.Context, arg1 *proto.StoreRequest) (*proto.StoreResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(*proto.StoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockEnvelopeStoreServerMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockEnvelopeStoreServer)(nil).Store), arg0, arg1)
}
