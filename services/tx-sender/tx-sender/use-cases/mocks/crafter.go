// Code generated by MockGen. DO NOT EDIT.
// Source: crafter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "github.com/consensys/orchestrate/pkg/types/entities"
	reflect "reflect"
)

// MockCraftTransactionUseCase is a mock of CraftTransactionUseCase interface
type MockCraftTransactionUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCraftTransactionUseCaseMockRecorder
}

// MockCraftTransactionUseCaseMockRecorder is the mock recorder for MockCraftTransactionUseCase
type MockCraftTransactionUseCaseMockRecorder struct {
	mock *MockCraftTransactionUseCase
}

// NewMockCraftTransactionUseCase creates a new mock instance
func NewMockCraftTransactionUseCase(ctrl *gomock.Controller) *MockCraftTransactionUseCase {
	mock := &MockCraftTransactionUseCase{ctrl: ctrl}
	mock.recorder = &MockCraftTransactionUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCraftTransactionUseCase) EXPECT() *MockCraftTransactionUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockCraftTransactionUseCase) Execute(ctx context.Context, job *entities.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockCraftTransactionUseCaseMockRecorder) Execute(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCraftTransactionUseCase)(nil).Execute), ctx, job)
}
