// Code generated by MockGen. DO NOT EDIT.
// Source: transactions.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTransactionValidator is a mock of TransactionValidator interface.
type MockTransactionValidator struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionValidatorMockRecorder
}

// MockTransactionValidatorMockRecorder is the mock recorder for MockTransactionValidator.
type MockTransactionValidatorMockRecorder struct {
	mock *MockTransactionValidator
}

// NewMockTransactionValidator creates a new mock instance.
func NewMockTransactionValidator(ctrl *gomock.Controller) *MockTransactionValidator {
	mock := &MockTransactionValidator{ctrl: ctrl}
	mock.recorder = &MockTransactionValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionValidator) EXPECT() *MockTransactionValidatorMockRecorder {
	return m.recorder
}

// ValidateRequestHash mocks base method.
func (m *MockTransactionValidator) ValidateRequestHash(ctx context.Context, chainUUID string, params interface{}, idempotencyKey string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRequestHash", ctx, chainUUID, params, idempotencyKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateRequestHash indicates an expected call of ValidateRequestHash.
func (mr *MockTransactionValidatorMockRecorder) ValidateRequestHash(ctx, chainUUID, params, idempotencyKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRequestHash", reflect.TypeOf((*MockTransactionValidator)(nil).ValidateRequestHash), ctx, chainUUID, params, idempotencyKey)
}

// ValidateChainExists mocks base method.
func (m *MockTransactionValidator) ValidateChainExists(ctx context.Context, chainUUID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateChainExists", ctx, chainUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateChainExists indicates an expected call of ValidateChainExists.
func (mr *MockTransactionValidatorMockRecorder) ValidateChainExists(ctx, chainUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateChainExists", reflect.TypeOf((*MockTransactionValidator)(nil).ValidateChainExists), ctx, chainUUID)
}
