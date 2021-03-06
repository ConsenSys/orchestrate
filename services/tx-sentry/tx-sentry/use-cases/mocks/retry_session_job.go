// Code generated by MockGen. DO NOT EDIT.
// Source: retry_session_job.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRetrySessionJobUseCase is a mock of RetrySessionJobUseCase interface
type MockRetrySessionJobUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockRetrySessionJobUseCaseMockRecorder
}

// MockRetrySessionJobUseCaseMockRecorder is the mock recorder for MockRetrySessionJobUseCase
type MockRetrySessionJobUseCaseMockRecorder struct {
	mock *MockRetrySessionJobUseCase
}

// NewMockRetrySessionJobUseCase creates a new mock instance
func NewMockRetrySessionJobUseCase(ctrl *gomock.Controller) *MockRetrySessionJobUseCase {
	mock := &MockRetrySessionJobUseCase{ctrl: ctrl}
	mock.recorder = &MockRetrySessionJobUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRetrySessionJobUseCase) EXPECT() *MockRetrySessionJobUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockRetrySessionJobUseCase) Execute(ctx context.Context, parentJobUUID, lastChildUUID string, nChildren int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, parentJobUUID, lastChildUUID, nChildren)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockRetrySessionJobUseCaseMockRecorder) Execute(ctx, parentJobUUID, lastChildUUID, nChildren interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockRetrySessionJobUseCase)(nil).Execute), ctx, parentJobUUID, lastChildUUID, nChildren)
}
