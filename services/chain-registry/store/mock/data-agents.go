// Code generated by MockGen. DO NOT EDIT.
// Source: data-agents.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/models"
	reflect "reflect"
)

// MockChainAgent is a mock of ChainAgent interface
type MockChainAgent struct {
	ctrl     *gomock.Controller
	recorder *MockChainAgentMockRecorder
}

// MockChainAgentMockRecorder is the mock recorder for MockChainAgent
type MockChainAgentMockRecorder struct {
	mock *MockChainAgent
}

// NewMockChainAgent creates a new mock instance
func NewMockChainAgent(ctrl *gomock.Controller) *MockChainAgent {
	mock := &MockChainAgent{ctrl: ctrl}
	mock.recorder = &MockChainAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainAgent) EXPECT() *MockChainAgentMockRecorder {
	return m.recorder
}

// RegisterChain mocks base method
func (m *MockChainAgent) RegisterChain(ctx context.Context, chain *models.Chain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterChain", ctx, chain)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterChain indicates an expected call of RegisterChain
func (mr *MockChainAgentMockRecorder) RegisterChain(ctx, chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterChain", reflect.TypeOf((*MockChainAgent)(nil).RegisterChain), ctx, chain)
}

// GetChains mocks base method
func (m *MockChainAgent) GetChains(ctx context.Context, tenants []string, filters map[string]string) ([]*models.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChains", ctx, tenants, filters)
	ret0, _ := ret[0].([]*models.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChains indicates an expected call of GetChains
func (mr *MockChainAgentMockRecorder) GetChains(ctx, tenants, filters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChains", reflect.TypeOf((*MockChainAgent)(nil).GetChains), ctx, tenants, filters)
}

// GetChain mocks base method
func (m *MockChainAgent) GetChain(ctx context.Context, uuid string, tenants []string) (*models.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChain", ctx, uuid, tenants)
	ret0, _ := ret[0].(*models.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChain indicates an expected call of GetChain
func (mr *MockChainAgentMockRecorder) GetChain(ctx, uuid, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChain", reflect.TypeOf((*MockChainAgent)(nil).GetChain), ctx, uuid, tenants)
}

// UpdateChain mocks base method
func (m *MockChainAgent) UpdateChain(ctx context.Context, uuid string, tenants []string, chain *models.Chain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChain", ctx, uuid, tenants, chain)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateChain indicates an expected call of UpdateChain
func (mr *MockChainAgentMockRecorder) UpdateChain(ctx, uuid, tenants, chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChain", reflect.TypeOf((*MockChainAgent)(nil).UpdateChain), ctx, uuid, tenants, chain)
}

// UpdateChainByName mocks base method
func (m *MockChainAgent) UpdateChainByName(ctx context.Context, name string, tenants []string, chain *models.Chain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChainByName", ctx, name, tenants, chain)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateChainByName indicates an expected call of UpdateChainByName
func (mr *MockChainAgentMockRecorder) UpdateChainByName(ctx, name, tenants, chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChainByName", reflect.TypeOf((*MockChainAgent)(nil).UpdateChainByName), ctx, name, tenants, chain)
}

// DeleteChain mocks base method
func (m *MockChainAgent) DeleteChain(ctx context.Context, uuid string, tenants []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChain", ctx, uuid, tenants)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChain indicates an expected call of DeleteChain
func (mr *MockChainAgentMockRecorder) DeleteChain(ctx, uuid, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChain", reflect.TypeOf((*MockChainAgent)(nil).DeleteChain), ctx, uuid, tenants)
}

// MockFaucetAgent is a mock of FaucetAgent interface
type MockFaucetAgent struct {
	ctrl     *gomock.Controller
	recorder *MockFaucetAgentMockRecorder
}

// MockFaucetAgentMockRecorder is the mock recorder for MockFaucetAgent
type MockFaucetAgentMockRecorder struct {
	mock *MockFaucetAgent
}

// NewMockFaucetAgent creates a new mock instance
func NewMockFaucetAgent(ctrl *gomock.Controller) *MockFaucetAgent {
	mock := &MockFaucetAgent{ctrl: ctrl}
	mock.recorder = &MockFaucetAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFaucetAgent) EXPECT() *MockFaucetAgentMockRecorder {
	return m.recorder
}

// RegisterFaucet mocks base method
func (m *MockFaucetAgent) RegisterFaucet(ctx context.Context, faucet *models.Faucet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterFaucet", ctx, faucet)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterFaucet indicates an expected call of RegisterFaucet
func (mr *MockFaucetAgentMockRecorder) RegisterFaucet(ctx, faucet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFaucet", reflect.TypeOf((*MockFaucetAgent)(nil).RegisterFaucet), ctx, faucet)
}

// GetFaucets mocks base method
func (m *MockFaucetAgent) GetFaucets(ctx context.Context, tenants []string, filters map[string]string) ([]*models.Faucet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFaucets", ctx, tenants, filters)
	ret0, _ := ret[0].([]*models.Faucet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFaucets indicates an expected call of GetFaucets
func (mr *MockFaucetAgentMockRecorder) GetFaucets(ctx, tenants, filters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFaucets", reflect.TypeOf((*MockFaucetAgent)(nil).GetFaucets), ctx, tenants, filters)
}

// GetFaucet mocks base method
func (m *MockFaucetAgent) GetFaucet(ctx context.Context, uuid string, tenants []string) (*models.Faucet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFaucet", ctx, uuid, tenants)
	ret0, _ := ret[0].(*models.Faucet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFaucet indicates an expected call of GetFaucet
func (mr *MockFaucetAgentMockRecorder) GetFaucet(ctx, uuid, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFaucet", reflect.TypeOf((*MockFaucetAgent)(nil).GetFaucet), ctx, uuid, tenants)
}

// UpdateFaucet mocks base method
func (m *MockFaucetAgent) UpdateFaucet(ctx context.Context, uuid string, tenants []string, faucet *models.Faucet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFaucet", ctx, uuid, tenants, faucet)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFaucet indicates an expected call of UpdateFaucet
func (mr *MockFaucetAgentMockRecorder) UpdateFaucet(ctx, uuid, tenants, faucet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFaucet", reflect.TypeOf((*MockFaucetAgent)(nil).UpdateFaucet), ctx, uuid, tenants, faucet)
}

// DeleteFaucet mocks base method
func (m *MockFaucetAgent) DeleteFaucet(ctx context.Context, uuid string, tenants []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFaucet", ctx, uuid, tenants)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFaucet indicates an expected call of DeleteFaucet
func (mr *MockFaucetAgentMockRecorder) DeleteFaucet(ctx, uuid, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFaucet", reflect.TypeOf((*MockFaucetAgent)(nil).DeleteFaucet), ctx, uuid, tenants)
}
