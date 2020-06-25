// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/proto (interfaces: ContractRegistryServer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/proto"
	reflect "reflect"
)

// MockContractRegistryServer is a mock of ContractRegistryServer interface
type MockContractRegistryServer struct {
	ctrl     *gomock.Controller
	recorder *MockContractRegistryServerMockRecorder
}

// MockContractRegistryServerMockRecorder is the mock recorder for MockContractRegistryServer
type MockContractRegistryServerMockRecorder struct {
	mock *MockContractRegistryServer
}

// NewMockContractRegistryServer creates a new mock instance
func NewMockContractRegistryServer(ctrl *gomock.Controller) *MockContractRegistryServer {
	mock := &MockContractRegistryServer{ctrl: ctrl}
	mock.recorder = &MockContractRegistryServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContractRegistryServer) EXPECT() *MockContractRegistryServerMockRecorder {
	return m.recorder
}

// DeleteArtifact mocks base method
func (m *MockContractRegistryServer) DeleteArtifact(arg0 context.Context, arg1 *proto.DeleteArtifactRequest) (*proto.DeleteArtifactResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArtifact", arg0, arg1)
	ret0, _ := ret[0].(*proto.DeleteArtifactResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteArtifact indicates an expected call of DeleteArtifact
func (mr *MockContractRegistryServerMockRecorder) DeleteArtifact(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArtifact", reflect.TypeOf((*MockContractRegistryServer)(nil).DeleteArtifact), arg0, arg1)
}

// DeregisterContract mocks base method
func (m *MockContractRegistryServer) DeregisterContract(arg0 context.Context, arg1 *proto.DeregisterContractRequest) (*proto.DeregisterContractResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterContract", arg0, arg1)
	ret0, _ := ret[0].(*proto.DeregisterContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterContract indicates an expected call of DeregisterContract
func (mr *MockContractRegistryServerMockRecorder) DeregisterContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterContract", reflect.TypeOf((*MockContractRegistryServer)(nil).DeregisterContract), arg0, arg1)
}

// GetCatalog mocks base method
func (m *MockContractRegistryServer) GetCatalog(arg0 context.Context, arg1 *proto.GetCatalogRequest) (*proto.GetCatalogResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCatalog", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetCatalogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCatalog indicates an expected call of GetCatalog
func (mr *MockContractRegistryServerMockRecorder) GetCatalog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCatalog", reflect.TypeOf((*MockContractRegistryServer)(nil).GetCatalog), arg0, arg1)
}

// GetContract mocks base method
func (m *MockContractRegistryServer) GetContract(arg0 context.Context, arg1 *proto.GetContractRequest) (*proto.GetContractResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract
func (mr *MockContractRegistryServerMockRecorder) GetContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockContractRegistryServer)(nil).GetContract), arg0, arg1)
}

// GetContractABI mocks base method
func (m *MockContractRegistryServer) GetContractABI(arg0 context.Context, arg1 *proto.GetContractRequest) (*proto.GetContractABIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractABI", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetContractABIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractABI indicates an expected call of GetContractABI
func (mr *MockContractRegistryServerMockRecorder) GetContractABI(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractABI", reflect.TypeOf((*MockContractRegistryServer)(nil).GetContractABI), arg0, arg1)
}

// GetContractBytecode mocks base method
func (m *MockContractRegistryServer) GetContractBytecode(arg0 context.Context, arg1 *proto.GetContractRequest) (*proto.GetContractBytecodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractBytecode", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetContractBytecodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractBytecode indicates an expected call of GetContractBytecode
func (mr *MockContractRegistryServerMockRecorder) GetContractBytecode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractBytecode", reflect.TypeOf((*MockContractRegistryServer)(nil).GetContractBytecode), arg0, arg1)
}

// GetContractDeployedBytecode mocks base method
func (m *MockContractRegistryServer) GetContractDeployedBytecode(arg0 context.Context, arg1 *proto.GetContractRequest) (*proto.GetContractDeployedBytecodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractDeployedBytecode", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetContractDeployedBytecodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractDeployedBytecode indicates an expected call of GetContractDeployedBytecode
func (mr *MockContractRegistryServerMockRecorder) GetContractDeployedBytecode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractDeployedBytecode", reflect.TypeOf((*MockContractRegistryServer)(nil).GetContractDeployedBytecode), arg0, arg1)
}

// GetEventsBySigHash mocks base method
func (m *MockContractRegistryServer) GetEventsBySigHash(arg0 context.Context, arg1 *proto.GetEventsBySigHashRequest) (*proto.GetEventsBySigHashResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsBySigHash", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetEventsBySigHashResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsBySigHash indicates an expected call of GetEventsBySigHash
func (mr *MockContractRegistryServerMockRecorder) GetEventsBySigHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsBySigHash", reflect.TypeOf((*MockContractRegistryServer)(nil).GetEventsBySigHash), arg0, arg1)
}

// GetMethodSignatures mocks base method
func (m *MockContractRegistryServer) GetMethodSignatures(arg0 context.Context, arg1 *proto.GetMethodSignaturesRequest) (*proto.GetMethodSignaturesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMethodSignatures", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetMethodSignaturesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMethodSignatures indicates an expected call of GetMethodSignatures
func (mr *MockContractRegistryServerMockRecorder) GetMethodSignatures(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMethodSignatures", reflect.TypeOf((*MockContractRegistryServer)(nil).GetMethodSignatures), arg0, arg1)
}

// GetMethodsBySelector mocks base method
func (m *MockContractRegistryServer) GetMethodsBySelector(arg0 context.Context, arg1 *proto.GetMethodsBySelectorRequest) (*proto.GetMethodsBySelectorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMethodsBySelector", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetMethodsBySelectorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMethodsBySelector indicates an expected call of GetMethodsBySelector
func (mr *MockContractRegistryServerMockRecorder) GetMethodsBySelector(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMethodsBySelector", reflect.TypeOf((*MockContractRegistryServer)(nil).GetMethodsBySelector), arg0, arg1)
}

// GetTags mocks base method
func (m *MockContractRegistryServer) GetTags(arg0 context.Context, arg1 *proto.GetTagsRequest) (*proto.GetTagsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetTagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags
func (mr *MockContractRegistryServerMockRecorder) GetTags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockContractRegistryServer)(nil).GetTags), arg0, arg1)
}

// RegisterContract mocks base method
func (m *MockContractRegistryServer) RegisterContract(arg0 context.Context, arg1 *proto.RegisterContractRequest) (*proto.RegisterContractResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterContract", arg0, arg1)
	ret0, _ := ret[0].(*proto.RegisterContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterContract indicates an expected call of RegisterContract
func (mr *MockContractRegistryServerMockRecorder) RegisterContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterContract", reflect.TypeOf((*MockContractRegistryServer)(nil).RegisterContract), arg0, arg1)
}

// SetAccountCodeHash mocks base method
func (m *MockContractRegistryServer) SetAccountCodeHash(arg0 context.Context, arg1 *proto.SetAccountCodeHashRequest) (*proto.SetAccountCodeHashResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccountCodeHash", arg0, arg1)
	ret0, _ := ret[0].(*proto.SetAccountCodeHashResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetAccountCodeHash indicates an expected call of SetAccountCodeHash
func (mr *MockContractRegistryServerMockRecorder) SetAccountCodeHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccountCodeHash", reflect.TypeOf((*MockContractRegistryServer)(nil).SetAccountCodeHash), arg0, arg1)
}
