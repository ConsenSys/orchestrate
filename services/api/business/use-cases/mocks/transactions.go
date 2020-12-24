// Code generated by MockGen. DO NOT EDIT.
// Source: transactions.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/use-cases"
	reflect "reflect"
)

// MockTransactionUseCases is a mock of TransactionUseCases interface
type MockTransactionUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionUseCasesMockRecorder
}

// MockTransactionUseCasesMockRecorder is the mock recorder for MockTransactionUseCases
type MockTransactionUseCasesMockRecorder struct {
	mock *MockTransactionUseCases
}

// NewMockTransactionUseCases creates a new mock instance
func NewMockTransactionUseCases(ctrl *gomock.Controller) *MockTransactionUseCases {
	mock := &MockTransactionUseCases{ctrl: ctrl}
	mock.recorder = &MockTransactionUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionUseCases) EXPECT() *MockTransactionUseCasesMockRecorder {
	return m.recorder
}

// SendContractTransaction mocks base method
func (m *MockTransactionUseCases) SendContractTransaction() usecases.SendContractTxUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendContractTransaction")
	ret0, _ := ret[0].(usecases.SendContractTxUseCase)
	return ret0
}

// SendContractTransaction indicates an expected call of SendContractTransaction
func (mr *MockTransactionUseCasesMockRecorder) SendContractTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendContractTransaction", reflect.TypeOf((*MockTransactionUseCases)(nil).SendContractTransaction))
}

// SendDeployTransaction mocks base method
func (m *MockTransactionUseCases) SendDeployTransaction() usecases.SendDeployTxUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDeployTransaction")
	ret0, _ := ret[0].(usecases.SendDeployTxUseCase)
	return ret0
}

// SendDeployTransaction indicates an expected call of SendDeployTransaction
func (mr *MockTransactionUseCasesMockRecorder) SendDeployTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDeployTransaction", reflect.TypeOf((*MockTransactionUseCases)(nil).SendDeployTransaction))
}

// SendTransaction mocks base method
func (m *MockTransactionUseCases) SendTransaction() usecases.SendTxUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction")
	ret0, _ := ret[0].(usecases.SendTxUseCase)
	return ret0
}

// SendTransaction indicates an expected call of SendTransaction
func (mr *MockTransactionUseCasesMockRecorder) SendTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockTransactionUseCases)(nil).SendTransaction))
}

// GetTransaction mocks base method
func (m *MockTransactionUseCases) GetTransaction() usecases.GetTxUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction")
	ret0, _ := ret[0].(usecases.GetTxUseCase)
	return ret0
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockTransactionUseCasesMockRecorder) GetTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockTransactionUseCases)(nil).GetTransaction))
}

// SearchTransactions mocks base method
func (m *MockTransactionUseCases) SearchTransactions() usecases.SearchTransactionsUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTransactions")
	ret0, _ := ret[0].(usecases.SearchTransactionsUseCase)
	return ret0
}

// SearchTransactions indicates an expected call of SearchTransactions
func (mr *MockTransactionUseCasesMockRecorder) SearchTransactions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTransactions", reflect.TypeOf((*MockTransactionUseCases)(nil).SearchTransactions))
}

// MockGetTxUseCase is a mock of GetTxUseCase interface
type MockGetTxUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetTxUseCaseMockRecorder
}

// MockGetTxUseCaseMockRecorder is the mock recorder for MockGetTxUseCase
type MockGetTxUseCaseMockRecorder struct {
	mock *MockGetTxUseCase
}

// NewMockGetTxUseCase creates a new mock instance
func NewMockGetTxUseCase(ctrl *gomock.Controller) *MockGetTxUseCase {
	mock := &MockGetTxUseCase{ctrl: ctrl}
	mock.recorder = &MockGetTxUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetTxUseCase) EXPECT() *MockGetTxUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetTxUseCase) Execute(ctx context.Context, scheduleUUID string, tenants []string) (*entities.TxRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, scheduleUUID, tenants)
	ret0, _ := ret[0].(*entities.TxRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetTxUseCaseMockRecorder) Execute(ctx, scheduleUUID, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetTxUseCase)(nil).Execute), ctx, scheduleUUID, tenants)
}

// MockSearchTransactionsUseCase is a mock of SearchTransactionsUseCase interface
type MockSearchTransactionsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSearchTransactionsUseCaseMockRecorder
}

// MockSearchTransactionsUseCaseMockRecorder is the mock recorder for MockSearchTransactionsUseCase
type MockSearchTransactionsUseCaseMockRecorder struct {
	mock *MockSearchTransactionsUseCase
}

// NewMockSearchTransactionsUseCase creates a new mock instance
func NewMockSearchTransactionsUseCase(ctrl *gomock.Controller) *MockSearchTransactionsUseCase {
	mock := &MockSearchTransactionsUseCase{ctrl: ctrl}
	mock.recorder = &MockSearchTransactionsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchTransactionsUseCase) EXPECT() *MockSearchTransactionsUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSearchTransactionsUseCase) Execute(ctx context.Context, filters *entities.TransactionRequestFilters, tenants []string) ([]*entities.TxRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, filters, tenants)
	ret0, _ := ret[0].([]*entities.TxRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockSearchTransactionsUseCaseMockRecorder) Execute(ctx, filters, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSearchTransactionsUseCase)(nil).Execute), ctx, filters, tenants)
}

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
func (m *MockSendContractTxUseCase) Execute(ctx context.Context, txRequest *entities.TxRequest, tenantID string) (*entities.TxRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, txRequest, tenantID)
	ret0, _ := ret[0].(*entities.TxRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockSendContractTxUseCaseMockRecorder) Execute(ctx, txRequest, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendContractTxUseCase)(nil).Execute), ctx, txRequest, tenantID)
}

// MockSendTxUseCase is a mock of SendTxUseCase interface
type MockSendTxUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSendTxUseCaseMockRecorder
}

// MockSendTxUseCaseMockRecorder is the mock recorder for MockSendTxUseCase
type MockSendTxUseCaseMockRecorder struct {
	mock *MockSendTxUseCase
}

// NewMockSendTxUseCase creates a new mock instance
func NewMockSendTxUseCase(ctrl *gomock.Controller) *MockSendTxUseCase {
	mock := &MockSendTxUseCase{ctrl: ctrl}
	mock.recorder = &MockSendTxUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendTxUseCase) EXPECT() *MockSendTxUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSendTxUseCase) Execute(ctx context.Context, txRequest *entities.TxRequest, txData, tenantID string) (*entities.TxRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, txRequest, txData, tenantID)
	ret0, _ := ret[0].(*entities.TxRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockSendTxUseCaseMockRecorder) Execute(ctx, txRequest, txData, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendTxUseCase)(nil).Execute), ctx, txRequest, txData, tenantID)
}
