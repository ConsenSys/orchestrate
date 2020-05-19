// Code generated by MockGen. DO NOT EDIT.
// Source: get_job.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/types"
	reflect "reflect"
)

// MockGetJobUseCase is a mock of GetJobUseCase interface.
type MockGetJobUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetJobUseCaseMockRecorder
}

// MockGetJobUseCaseMockRecorder is the mock recorder for MockGetJobUseCase.
type MockGetJobUseCaseMockRecorder struct {
	mock *MockGetJobUseCase
}

// NewMockGetJobUseCase creates a new mock instance.
func NewMockGetJobUseCase(ctrl *gomock.Controller) *MockGetJobUseCase {
	mock := &MockGetJobUseCase{ctrl: ctrl}
	mock.recorder = &MockGetJobUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetJobUseCase) EXPECT() *MockGetJobUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetJobUseCase) Execute(ctx context.Context, jobUUID, tenantID string) (*types.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, jobUUID, tenantID)
	ret0, _ := ret[0].(*types.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetJobUseCaseMockRecorder) Execute(ctx, jobUUID, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetJobUseCase)(nil).Execute), ctx, jobUUID, tenantID)
}
