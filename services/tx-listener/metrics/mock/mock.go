// Code generated by MockGen. DO NOT EDIT.
// Source: exported.go

// Package mock is a generated GoMock package.
package mock

import (
	metrics "github.com/go-kit/kit/metrics"
	gomock "github.com/golang/mock/gomock"
	prometheus "github.com/prometheus/client_golang/prometheus"
	reflect "reflect"
)

// MockListenerMetrics is a mock of ListenerMetrics interface
type MockListenerMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockListenerMetricsMockRecorder
}

// MockListenerMetricsMockRecorder is the mock recorder for MockListenerMetrics
type MockListenerMetricsMockRecorder struct {
	mock *MockListenerMetrics
}

// NewMockListenerMetrics creates a new mock instance
func NewMockListenerMetrics(ctrl *gomock.Controller) *MockListenerMetrics {
	mock := &MockListenerMetrics{ctrl: ctrl}
	mock.recorder = &MockListenerMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockListenerMetrics) EXPECT() *MockListenerMetricsMockRecorder {
	return m.recorder
}

// BlockCounter mocks base method
func (m *MockListenerMetrics) BlockCounter() metrics.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockCounter")
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// BlockCounter indicates an expected call of BlockCounter
func (mr *MockListenerMetricsMockRecorder) BlockCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockCounter", reflect.TypeOf((*MockListenerMetrics)(nil).BlockCounter))
}

// Describe mocks base method
func (m *MockListenerMetrics) Describe(arg0 chan<- *prometheus.Desc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Describe", arg0)
}

// Describe indicates an expected call of Describe
func (mr *MockListenerMetricsMockRecorder) Describe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockListenerMetrics)(nil).Describe), arg0)
}

// Collect mocks base method
func (m *MockListenerMetrics) Collect(arg0 chan<- prometheus.Metric) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Collect", arg0)
}

// Collect indicates an expected call of Collect
func (mr *MockListenerMetricsMockRecorder) Collect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockListenerMetrics)(nil).Collect), arg0)
}
