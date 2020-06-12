// Code generated by MockGen. DO NOT EDIT.
// Source: send_contract_tx.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/entities"
	reflect "reflect"
)

// MockSendContractTxUseCase is a mock of SendContractTxUseCase interface
type MockSendContractTxUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSendContractTxUseCaseMockRecorder
}

// MockSendContractTxUseCaseMockRecorder is the mock recorder for MockSendContractTxUseCase
type MockSendContractTxUseCaseMockRecorder struct {
	mock *MockSendContractTxUseCase
}

// NewMockSendContractTxUseCase creates a new mock instance
func NewMockSendContractTxUseCase(ctrl *gomock.Controller) *MockSendContractTxUseCase {
	mock := &MockSendContractTxUseCase{ctrl: ctrl}
	mock.recorder = &MockSendContractTxUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendContractTxUseCase) EXPECT() *MockSendContractTxUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSendContractTxUseCase) Execute(ctx context.Context, txRequest *entities.TxRequest, chainUUID, tenantID string) (*entities.TxRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, txRequest, chainUUID, tenantID)
	ret0, _ := ret[0].(*entities.TxRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockSendContractTxUseCaseMockRecorder) Execute(ctx, txRequest, chainUUID, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendContractTxUseCase)(nil).Execute), ctx, txRequest, chainUUID, tenantID)
}
