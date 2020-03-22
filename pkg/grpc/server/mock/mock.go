// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockBuilder is a mock of Builder interface
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method
func (m *MockBuilder) Build(ctx context.Context, name string, configuration interface{}) (*grpc.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx, name, configuration)
	ret0, _ := ret[0].(*grpc.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build
func (mr *MockBuilderMockRecorder) Build(ctx, name, configuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuilder)(nil).Build), ctx, name, configuration)
}

// MockOptionsBuilder is a mock of OptionsBuilder interface
type MockOptionsBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsBuilderMockRecorder
}

// MockOptionsBuilderMockRecorder is the mock recorder for MockOptionsBuilder
type MockOptionsBuilderMockRecorder struct {
	mock *MockOptionsBuilder
}

// NewMockOptionsBuilder creates a new mock instance
func NewMockOptionsBuilder(ctrl *gomock.Controller) *MockOptionsBuilder {
	mock := &MockOptionsBuilder{ctrl: ctrl}
	mock.recorder = &MockOptionsBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOptionsBuilder) EXPECT() *MockOptionsBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method
func (m *MockOptionsBuilder) Build(ctx context.Context, name string, configuration interface{}) ([]grpc.ServerOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx, name, configuration)
	ret0, _ := ret[0].([]grpc.ServerOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build
func (mr *MockOptionsBuilderMockRecorder) Build(ctx, name, configuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockOptionsBuilder)(nil).Build), ctx, name, configuration)
}
