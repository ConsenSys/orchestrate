// Code generated by MockGen. DO NOT EDIT.
// Source: metrics.go

// Package mock is a generated GoMock package.
package mock

import (
	metrics "github.com/go-kit/kit/metrics"
	gomock "github.com/golang/mock/gomock"
	dynamic "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/http/config/dynamic"
	metrics0 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/metrics"
	reflect "reflect"
)

// MockTCP is a mock of TCP interface.
type MockTCP struct {
	ctrl     *gomock.Controller
	recorder *MockTCPMockRecorder
}

// MockTCPMockRecorder is the mock recorder for MockTCP.
type MockTCPMockRecorder struct {
	mock *MockTCP
}

// NewMockTCP creates a new mock instance.
func NewMockTCP(ctrl *gomock.Controller) *MockTCP {
	mock := &MockTCP{ctrl: ctrl}
	mock.recorder = &MockTCPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTCP) EXPECT() *MockTCPMockRecorder {
	return m.recorder
}

// AcceptedConnsCounter mocks base method.
func (m *MockTCP) AcceptedConnsCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedConnsCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// AcceptedConnsCounter indicates an expected call of AcceptedConnsCounter.
func (mr *MockTCPMockRecorder) AcceptedConnsCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedConnsCounter", reflect.TypeOf((*MockTCP)(nil).AcceptedConnsCounter))
}

// ClosedConnsCounter mocks base method.
func (m *MockTCP) ClosedConnsCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClosedConnsCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// ClosedConnsCounter indicates an expected call of ClosedConnsCounter.
func (mr *MockTCPMockRecorder) ClosedConnsCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosedConnsCounter", reflect.TypeOf((*MockTCP)(nil).ClosedConnsCounter))
}

// ConnsLatencyHistogram mocks base method.
func (m *MockTCP) ConnsLatencyHistogram() metrics.Histogram {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnsLatencyHistogram")
	ret0, _ := ret[0].(metrics.Histogram)
	return ret0
}

// ConnsLatencyHistogram indicates an expected call of ConnsLatencyHistogram.
func (mr *MockTCPMockRecorder) ConnsLatencyHistogram() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnsLatencyHistogram", reflect.TypeOf((*MockTCP)(nil).ConnsLatencyHistogram))
}

// OpenConnsGauge mocks base method.
func (m *MockTCP) OpenConnsGauge() metrics.Gauge {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenConnsGauge")
	ret0, _ := ret[0].(metrics.Gauge)
	return ret0
}

// OpenConnsGauge indicates an expected call of OpenConnsGauge.
func (mr *MockTCPMockRecorder) OpenConnsGauge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenConnsGauge", reflect.TypeOf((*MockTCP)(nil).OpenConnsGauge))
}

// MockHTTP is a mock of HTTP interface.
type MockHTTP struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPMockRecorder
}

// MockHTTPMockRecorder is the mock recorder for MockHTTP.
type MockHTTPMockRecorder struct {
	mock *MockHTTP
}

// NewMockHTTP creates a new mock instance.
func NewMockHTTP(ctrl *gomock.Controller) *MockHTTP {
	mock := &MockHTTP{ctrl: ctrl}
	mock.recorder = &MockHTTPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTP) EXPECT() *MockHTTPMockRecorder {
	return m.recorder
}

// RequestsCounter mocks base method.
func (m *MockHTTP) RequestsCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// RequestsCounter indicates an expected call of RequestsCounter.
func (mr *MockHTTPMockRecorder) RequestsCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsCounter", reflect.TypeOf((*MockHTTP)(nil).RequestsCounter))
}

// TLSRequestsCounter mocks base method.
func (m *MockHTTP) TLSRequestsCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TLSRequestsCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// TLSRequestsCounter indicates an expected call of TLSRequestsCounter.
func (mr *MockHTTPMockRecorder) TLSRequestsCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLSRequestsCounter", reflect.TypeOf((*MockHTTP)(nil).TLSRequestsCounter))
}

// RequestsLatencyHistogram mocks base method.
func (m *MockHTTP) RequestsLatencyHistogram() metrics.Histogram {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestsLatencyHistogram")
	ret0, _ := ret[0].(metrics.Histogram)
	return ret0
}

// RequestsLatencyHistogram indicates an expected call of RequestsLatencyHistogram.
func (mr *MockHTTPMockRecorder) RequestsLatencyHistogram() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestsLatencyHistogram", reflect.TypeOf((*MockHTTP)(nil).RequestsLatencyHistogram))
}

// OpenConnsGauge mocks base method.
func (m *MockHTTP) OpenConnsGauge() metrics.Gauge {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenConnsGauge")
	ret0, _ := ret[0].(metrics.Gauge)
	return ret0
}

// OpenConnsGauge indicates an expected call of OpenConnsGauge.
func (mr *MockHTTPMockRecorder) OpenConnsGauge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenConnsGauge", reflect.TypeOf((*MockHTTP)(nil).OpenConnsGauge))
}

// RetriesCounter mocks base method.
func (m *MockHTTP) RetriesCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetriesCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// RetriesCounter indicates an expected call of RetriesCounter.
func (mr *MockHTTPMockRecorder) RetriesCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetriesCounter", reflect.TypeOf((*MockHTTP)(nil).RetriesCounter))
}

// ServerUpGauge mocks base method.
func (m *MockHTTP) ServerUpGauge() metrics.Gauge {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerUpGauge")
	ret0, _ := ret[0].(metrics.Gauge)
	return ret0
}

// ServerUpGauge indicates an expected call of ServerUpGauge.
func (mr *MockHTTPMockRecorder) ServerUpGauge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerUpGauge", reflect.TypeOf((*MockHTTP)(nil).ServerUpGauge))
}

// Switch mocks base method.
func (m *MockHTTP) Switch(cfg *dynamic.Configuration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Switch", cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Switch indicates an expected call of Switch.
func (mr *MockHTTPMockRecorder) Switch(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Switch", reflect.TypeOf((*MockHTTP)(nil).Switch), cfg)
}

// MockGRPCServer is a mock of GRPCServer interface.
type MockGRPCServer struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCServerMockRecorder
}

// MockGRPCServerMockRecorder is the mock recorder for MockGRPCServer.
type MockGRPCServerMockRecorder struct {
	mock *MockGRPCServer
}

// NewMockGRPCServer creates a new mock instance.
func NewMockGRPCServer(ctrl *gomock.Controller) *MockGRPCServer {
	mock := &MockGRPCServer{ctrl: ctrl}
	mock.recorder = &MockGRPCServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGRPCServer) EXPECT() *MockGRPCServerMockRecorder {
	return m.recorder
}

// StartedCounter mocks base method.
func (m *MockGRPCServer) StartedCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartedCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// StartedCounter indicates an expected call of StartedCounter.
func (mr *MockGRPCServerMockRecorder) StartedCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartedCounter", reflect.TypeOf((*MockGRPCServer)(nil).StartedCounter))
}

// HandledCounter mocks base method.
func (m *MockGRPCServer) HandledCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandledCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// HandledCounter indicates an expected call of HandledCounter.
func (mr *MockGRPCServerMockRecorder) HandledCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandledCounter", reflect.TypeOf((*MockGRPCServer)(nil).HandledCounter))
}

// StreamMsgReceivedCounter mocks base method.
func (m *MockGRPCServer) StreamMsgReceivedCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamMsgReceivedCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// StreamMsgReceivedCounter indicates an expected call of StreamMsgReceivedCounter.
func (mr *MockGRPCServerMockRecorder) StreamMsgReceivedCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamMsgReceivedCounter", reflect.TypeOf((*MockGRPCServer)(nil).StreamMsgReceivedCounter))
}

// StreamMsgSentCounter mocks base method.
func (m *MockGRPCServer) StreamMsgSentCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamMsgSentCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// StreamMsgSentCounter indicates an expected call of StreamMsgSentCounter.
func (mr *MockGRPCServerMockRecorder) StreamMsgSentCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamMsgSentCounter", reflect.TypeOf((*MockGRPCServer)(nil).StreamMsgSentCounter))
}

// HandledDurationHistogram mocks base method.
func (m *MockGRPCServer) HandledDurationHistogram() metrics.Histogram {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandledDurationHistogram")
	ret0, _ := ret[0].(metrics.Histogram)
	return ret0
}

// HandledDurationHistogram indicates an expected call of HandledDurationHistogram.
func (mr *MockGRPCServerMockRecorder) HandledDurationHistogram() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandledDurationHistogram", reflect.TypeOf((*MockGRPCServer)(nil).HandledDurationHistogram))
}

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// TCP mocks base method.
func (m *MockRegistry) TCP() metrics0.TCP {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TCP")
	ret0, _ := ret[0].(metrics0.TCP)
	return ret0
}

// TCP indicates an expected call of TCP.
func (mr *MockRegistryMockRecorder) TCP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TCP", reflect.TypeOf((*MockRegistry)(nil).TCP))
}

// HTTP mocks base method.
func (m *MockRegistry) HTTP() metrics0.HTTP {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTP")
	ret0, _ := ret[0].(metrics0.HTTP)
	return ret0
}

// HTTP indicates an expected call of HTTP.
func (mr *MockRegistryMockRecorder) HTTP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTP", reflect.TypeOf((*MockRegistry)(nil).HTTP))
}

// GRPCServer mocks base method.
func (m *MockRegistry) GRPCServer() metrics0.GRPCServer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GRPCServer")
	ret0, _ := ret[0].(metrics0.GRPCServer)
	return ret0
}

// GRPCServer indicates an expected call of GRPCServer.
func (mr *MockRegistryMockRecorder) GRPCServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GRPCServer", reflect.TypeOf((*MockRegistry)(nil).GRPCServer))
}
