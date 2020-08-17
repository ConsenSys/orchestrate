// Code generated by MockGen. DO NOT EDIT.
// Source: send_deploy_tx.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/entities"
	reflect "reflect"
)

// MockSendDeployTxUseCase is a mock of SendDeployTxUseCase interface
type MockSendDeployTxUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSendDeployTxUseCaseMockRecorder
}

// MockSendDeployTxUseCaseMockRecorder is the mock recorder for MockSendDeployTxUseCase
type MockSendDeployTxUseCaseMockRecorder struct {
	mock *MockSendDeployTxUseCase
}

// NewMockSendDeployTxUseCase creates a new mock instance
func NewMockSendDeployTxUseCase(ctrl *gomock.Controller) *MockSendDeployTxUseCase {
	mock := &MockSendDeployTxUseCase{ctrl: ctrl}
	mock.recorder = &MockSendDeployTxUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendDeployTxUseCase) EXPECT() *MockSendDeployTxUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSendDeployTxUseCase) Execute(ctx context.Context, txRequest *entities.TxRequest, tenantID string) (*entities.TxRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, txRequest, tenantID)
	ret0, _ := ret[0].(*entities.TxRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockSendDeployTxUseCaseMockRecorder) Execute(ctx, txRequest, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendDeployTxUseCase)(nil).Execute), ctx, txRequest, tenantID)
}
