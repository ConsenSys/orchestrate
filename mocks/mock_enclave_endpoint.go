// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/ConsenSys/client/fr/core-stack/service/ethereum.git/tessera (interfaces: EnclaveEndpoint)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEnclaveEndpoint is a mock of EnclaveEndpoint interface
type MockEnclaveEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockEnclaveEndpointMockRecorder
}

// MockEnclaveEndpointMockRecorder is the mock recorder for MockEnclaveEndpoint
type MockEnclaveEndpointMockRecorder struct {
	mock *MockEnclaveEndpoint
}

// NewMockEnclaveEndpoint creates a new mock instance
func NewMockEnclaveEndpoint(ctrl *gomock.Controller) *MockEnclaveEndpoint {
	mock := &MockEnclaveEndpoint{ctrl: ctrl}
	mock.recorder = &MockEnclaveEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnclaveEndpoint) EXPECT() *MockEnclaveEndpointMockRecorder {
	return m.recorder
}

// GetRequest mocks base method
func (m *MockEnclaveEndpoint) GetRequest(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequest indicates an expected call of GetRequest
func (mr *MockEnclaveEndpointMockRecorder) GetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockEnclaveEndpoint)(nil).GetRequest), arg0)
}

// PostRequest mocks base method
func (m *MockEnclaveEndpoint) PostRequest(arg0 string, arg1, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostRequest indicates an expected call of PostRequest
func (mr *MockEnclaveEndpointMockRecorder) PostRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRequest", reflect.TypeOf((*MockEnclaveEndpoint)(nil).PostRequest), arg0, arg1, arg2)
}
