// Code generated by MockGen then slightly modified. DO NOT EDIT.
// Source: TODO: should be github.com/go-kit/kit/metrics/metrics.go

package mock

import (
	"reflect"

	kitmetrics "github.com/go-kit/kit/metrics"
	"github.com/golang/mock/gomock"
)

// MockCounter is a mock of Counter interface
type MockCounter struct {
	ctrl     *gomock.Controller
	recorder *MockCounterMockRecorder
}

// MockCounterMockRecorder is the mock recorder for MockCounter
type MockCounterMockRecorder struct {
	mock *MockCounter
}

// NewMockCounter creates a new mock instance
func NewMockCounter(ctrl *gomock.Controller) *MockCounter {
	mock := &MockCounter{ctrl: ctrl}
	mock.recorder = &MockCounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCounter) EXPECT() *MockCounterMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockCounter) With(labelValues ...string) kitmetrics.Counter {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range labelValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "With", varargs...)
	ret0, _ := ret[0].(kitmetrics.Counter)
	return ret0
}

// With indicates an expected call of With
func (mr *MockCounterMockRecorder) With(labelValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockCounter)(nil).With), labelValues...)
}

// Add mocks base method
func (m *MockCounter) Add(delta float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", delta)
}

// Add indicates an expected call of Add
func (mr *MockCounterMockRecorder) Add(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCounter)(nil).Add), delta)
}

// MockGauge is a mock of Gauge interface
type MockGauge struct {
	ctrl     *gomock.Controller
	recorder *MockGaugeMockRecorder
}

// MockGaugeMockRecorder is the mock recorder for MockGauge
type MockGaugeMockRecorder struct {
	mock *MockGauge
}

// NewMockGauge creates a new mock instance
func NewMockGauge(ctrl *gomock.Controller) *MockGauge {
	mock := &MockGauge{ctrl: ctrl}
	mock.recorder = &MockGaugeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGauge) EXPECT() *MockGaugeMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockGauge) With(labelValues ...string) kitmetrics.Gauge {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range labelValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "With", varargs...)
	ret0, _ := ret[0].(kitmetrics.Gauge)
	return ret0
}

// With indicates an expected call of With
func (mr *MockGaugeMockRecorder) With(labelValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockGauge)(nil).With), labelValues...)
}

// Set mocks base method
func (m *MockGauge) Set(value float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", value)
}

// Set indicates an expected call of Set
func (mr *MockGaugeMockRecorder) Set(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockGauge)(nil).Set), value)
}

// Add mocks base method
func (m *MockGauge) Add(delta float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", delta)
}

// Add indicates an expected call of Add
func (mr *MockGaugeMockRecorder) Add(delta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockGauge)(nil).Add), delta)
}

// MockHistogram is a mock of Histogram interface
type MockHistogram struct {
	ctrl     *gomock.Controller
	recorder *MockHistogramMockRecorder
}

// MockHistogramMockRecorder is the mock recorder for MockHistogram
type MockHistogramMockRecorder struct {
	mock *MockHistogram
}

// NewMockHistogram creates a new mock instance
func NewMockHistogram(ctrl *gomock.Controller) *MockHistogram {
	mock := &MockHistogram{ctrl: ctrl}
	mock.recorder = &MockHistogramMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHistogram) EXPECT() *MockHistogramMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockHistogram) With(labelValues ...string) kitmetrics.Histogram {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range labelValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "With", varargs...)
	ret0, _ := ret[0].(kitmetrics.Histogram)
	return ret0
}

// With indicates an expected call of With
func (mr *MockHistogramMockRecorder) With(labelValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockHistogram)(nil).With), labelValues...)
}

// Observe mocks base method
func (m *MockHistogram) Observe(value float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Observe", value)
}

// Observe indicates an expected call of Observe
func (mr *MockHistogramMockRecorder) Observe(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Observe", reflect.TypeOf((*MockHistogram)(nil).Observe), value)
}