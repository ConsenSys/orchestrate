// Code generated by MockGen. DO NOT EDIT.
// Source: message.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	engine "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/engine"
	reflect "reflect"
)

// MockMsg is a mock of Msg interface
type MockMsg struct {
	ctrl     *gomock.Controller
	recorder *MockMsgMockRecorder
}

// MockMsgMockRecorder is the mock recorder for MockMsg
type MockMsgMockRecorder struct {
	mock *MockMsg
}

// NewMockMsg creates a new mock instance
func NewMockMsg(ctrl *gomock.Controller) *MockMsg {
	mock := &MockMsg{ctrl: ctrl}
	mock.recorder = &MockMsgMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMsg) EXPECT() *MockMsgMockRecorder {
	return m.recorder
}

// Entrypoint mocks base method
func (m *MockMsg) Entrypoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entrypoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// Entrypoint indicates an expected call of Entrypoint
func (mr *MockMsgMockRecorder) Entrypoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entrypoint", reflect.TypeOf((*MockMsg)(nil).Entrypoint))
}

// Value mocks base method
func (m *MockMsg) Value() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockMsgMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockMsg)(nil).Value))
}

// Key mocks base method
func (m *MockMsg) Key() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Key indicates an expected call of Key
func (mr *MockMsgMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockMsg)(nil).Key))
}

// Header mocks base method
func (m *MockMsg) Header() engine.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(engine.Header)
	return ret0
}

// Header indicates an expected call of Header
func (mr *MockMsgMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockMsg)(nil).Header))
}

// MockHeader is a mock of Header interface
type MockHeader struct {
	ctrl     *gomock.Controller
	recorder *MockHeaderMockRecorder
}

// MockHeaderMockRecorder is the mock recorder for MockHeader
type MockHeaderMockRecorder struct {
	mock *MockHeader
}

// NewMockHeader creates a new mock instance
func NewMockHeader(ctrl *gomock.Controller) *MockHeader {
	mock := &MockHeader{ctrl: ctrl}
	mock.recorder = &MockHeaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHeader) EXPECT() *MockHeaderMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockHeader) Add(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", key, value)
}

// Add indicates an expected call of Add
func (mr *MockHeaderMockRecorder) Add(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockHeader)(nil).Add), key, value)
}

// Del mocks base method
func (m *MockHeader) Del(key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Del", key)
}

// Del indicates an expected call of Del
func (mr *MockHeaderMockRecorder) Del(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockHeader)(nil).Del), key)
}

// Get mocks base method
func (m *MockHeader) Get(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockHeaderMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHeader)(nil).Get), key)
}

// Set mocks base method
func (m *MockHeader) Set(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set
func (mr *MockHeaderMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockHeader)(nil).Set), key, value)
}
