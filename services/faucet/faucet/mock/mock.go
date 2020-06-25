// Code generated by MockGen. DO NOT EDIT.
// Source: faucet.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/faucet/types"
	big "math/big"
	reflect "reflect"
)

// MockFaucet is a mock of Faucet interface
type MockFaucet struct {
	ctrl     *gomock.Controller
	recorder *MockFaucetMockRecorder
}

// MockFaucetMockRecorder is the mock recorder for MockFaucet
type MockFaucetMockRecorder struct {
	mock *MockFaucet
}

// NewMockFaucet creates a new mock instance
func NewMockFaucet(ctrl *gomock.Controller) *MockFaucet {
	mock := &MockFaucet{ctrl: ctrl}
	mock.recorder = &MockFaucetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFaucet) EXPECT() *MockFaucetMockRecorder {
	return m.recorder
}

// Credit mocks base method
func (m *MockFaucet) Credit(ctx context.Context, r *types.Request) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credit", ctx, r)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Credit indicates an expected call of Credit
func (mr *MockFaucetMockRecorder) Credit(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credit", reflect.TypeOf((*MockFaucet)(nil).Credit), ctx, r)
}
