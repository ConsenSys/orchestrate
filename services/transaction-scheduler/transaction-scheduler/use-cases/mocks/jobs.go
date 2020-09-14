// Code generated by MockGen. DO NOT EDIT.
// Source: jobs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/entities"
	store "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/use-cases"
	reflect "reflect"
)

// MockJobUseCases is a mock of JobUseCases interface
type MockJobUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockJobUseCasesMockRecorder
}

// MockJobUseCasesMockRecorder is the mock recorder for MockJobUseCases
type MockJobUseCasesMockRecorder struct {
	mock *MockJobUseCases
}

// NewMockJobUseCases creates a new mock instance
func NewMockJobUseCases(ctrl *gomock.Controller) *MockJobUseCases {
	mock := &MockJobUseCases{ctrl: ctrl}
	mock.recorder = &MockJobUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJobUseCases) EXPECT() *MockJobUseCasesMockRecorder {
	return m.recorder
}

// CreateJob mocks base method
func (m *MockJobUseCases) CreateJob() usecases.CreateJobUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob")
	ret0, _ := ret[0].(usecases.CreateJobUseCase)
	return ret0
}

// CreateJob indicates an expected call of CreateJob
func (mr *MockJobUseCasesMockRecorder) CreateJob() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockJobUseCases)(nil).CreateJob))
}

// GetJob mocks base method
func (m *MockJobUseCases) GetJob() usecases.GetJobUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJob")
	ret0, _ := ret[0].(usecases.GetJobUseCase)
	return ret0
}

// GetJob indicates an expected call of GetJob
func (mr *MockJobUseCasesMockRecorder) GetJob() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockJobUseCases)(nil).GetJob))
}

// StartJob mocks base method
func (m *MockJobUseCases) StartJob() usecases.StartJobUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartJob")
	ret0, _ := ret[0].(usecases.StartJobUseCase)
	return ret0
}

// StartJob indicates an expected call of StartJob
func (mr *MockJobUseCasesMockRecorder) StartJob() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartJob", reflect.TypeOf((*MockJobUseCases)(nil).StartJob))
}

// UpdateJob mocks base method
func (m *MockJobUseCases) UpdateJob() usecases.UpdateJobUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJob")
	ret0, _ := ret[0].(usecases.UpdateJobUseCase)
	return ret0
}

// UpdateJob indicates an expected call of UpdateJob
func (mr *MockJobUseCasesMockRecorder) UpdateJob() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJob", reflect.TypeOf((*MockJobUseCases)(nil).UpdateJob))
}

// SearchJobs mocks base method
func (m *MockJobUseCases) SearchJobs() usecases.SearchJobsUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchJobs")
	ret0, _ := ret[0].(usecases.SearchJobsUseCase)
	return ret0
}

// SearchJobs indicates an expected call of SearchJobs
func (mr *MockJobUseCasesMockRecorder) SearchJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchJobs", reflect.TypeOf((*MockJobUseCases)(nil).SearchJobs))
}

// MockCreateJobUseCase is a mock of CreateJobUseCase interface
type MockCreateJobUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateJobUseCaseMockRecorder
}

// MockCreateJobUseCaseMockRecorder is the mock recorder for MockCreateJobUseCase
type MockCreateJobUseCaseMockRecorder struct {
	mock *MockCreateJobUseCase
}

// NewMockCreateJobUseCase creates a new mock instance
func NewMockCreateJobUseCase(ctrl *gomock.Controller) *MockCreateJobUseCase {
	mock := &MockCreateJobUseCase{ctrl: ctrl}
	mock.recorder = &MockCreateJobUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreateJobUseCase) EXPECT() *MockCreateJobUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockCreateJobUseCase) Execute(ctx context.Context, job *entities.Job, tenants []string) (*entities.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, job, tenants)
	ret0, _ := ret[0].(*entities.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockCreateJobUseCaseMockRecorder) Execute(ctx, job, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateJobUseCase)(nil).Execute), ctx, job, tenants)
}

// WithDBTransaction mocks base method
func (m *MockCreateJobUseCase) WithDBTransaction(dbtx store.Tx) usecases.CreateJobUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDBTransaction", dbtx)
	ret0, _ := ret[0].(usecases.CreateJobUseCase)
	return ret0
}

// WithDBTransaction indicates an expected call of WithDBTransaction
func (mr *MockCreateJobUseCaseMockRecorder) WithDBTransaction(dbtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDBTransaction", reflect.TypeOf((*MockCreateJobUseCase)(nil).WithDBTransaction), dbtx)
}

// MockGetJobUseCase is a mock of GetJobUseCase interface
type MockGetJobUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetJobUseCaseMockRecorder
}

// MockGetJobUseCaseMockRecorder is the mock recorder for MockGetJobUseCase
type MockGetJobUseCaseMockRecorder struct {
	mock *MockGetJobUseCase
}

// NewMockGetJobUseCase creates a new mock instance
func NewMockGetJobUseCase(ctrl *gomock.Controller) *MockGetJobUseCase {
	mock := &MockGetJobUseCase{ctrl: ctrl}
	mock.recorder = &MockGetJobUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetJobUseCase) EXPECT() *MockGetJobUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetJobUseCase) Execute(ctx context.Context, jobUUID string, tenants []string) (*entities.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, jobUUID, tenants)
	ret0, _ := ret[0].(*entities.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetJobUseCaseMockRecorder) Execute(ctx, jobUUID, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetJobUseCase)(nil).Execute), ctx, jobUUID, tenants)
}

// MockSearchJobsUseCase is a mock of SearchJobsUseCase interface
type MockSearchJobsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSearchJobsUseCaseMockRecorder
}

// MockSearchJobsUseCaseMockRecorder is the mock recorder for MockSearchJobsUseCase
type MockSearchJobsUseCaseMockRecorder struct {
	mock *MockSearchJobsUseCase
}

// NewMockSearchJobsUseCase creates a new mock instance
func NewMockSearchJobsUseCase(ctrl *gomock.Controller) *MockSearchJobsUseCase {
	mock := &MockSearchJobsUseCase{ctrl: ctrl}
	mock.recorder = &MockSearchJobsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchJobsUseCase) EXPECT() *MockSearchJobsUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSearchJobsUseCase) Execute(ctx context.Context, filters *entities.JobFilters, tenants []string) ([]*entities.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, filters, tenants)
	ret0, _ := ret[0].([]*entities.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockSearchJobsUseCaseMockRecorder) Execute(ctx, filters, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSearchJobsUseCase)(nil).Execute), ctx, filters, tenants)
}

// MockStartJobUseCase is a mock of StartJobUseCase interface
type MockStartJobUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockStartJobUseCaseMockRecorder
}

// MockStartJobUseCaseMockRecorder is the mock recorder for MockStartJobUseCase
type MockStartJobUseCaseMockRecorder struct {
	mock *MockStartJobUseCase
}

// NewMockStartJobUseCase creates a new mock instance
func NewMockStartJobUseCase(ctrl *gomock.Controller) *MockStartJobUseCase {
	mock := &MockStartJobUseCase{ctrl: ctrl}
	mock.recorder = &MockStartJobUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStartJobUseCase) EXPECT() *MockStartJobUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockStartJobUseCase) Execute(ctx context.Context, jobUUID string, tenants []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, jobUUID, tenants)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockStartJobUseCaseMockRecorder) Execute(ctx, jobUUID, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockStartJobUseCase)(nil).Execute), ctx, jobUUID, tenants)
}

// MockStartNextJobUseCase is a mock of StartNextJobUseCase interface
type MockStartNextJobUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockStartNextJobUseCaseMockRecorder
}

// MockStartNextJobUseCaseMockRecorder is the mock recorder for MockStartNextJobUseCase
type MockStartNextJobUseCaseMockRecorder struct {
	mock *MockStartNextJobUseCase
}

// NewMockStartNextJobUseCase creates a new mock instance
func NewMockStartNextJobUseCase(ctrl *gomock.Controller) *MockStartNextJobUseCase {
	mock := &MockStartNextJobUseCase{ctrl: ctrl}
	mock.recorder = &MockStartNextJobUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStartNextJobUseCase) EXPECT() *MockStartNextJobUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockStartNextJobUseCase) Execute(ctx context.Context, prevJobUUID string, tenants []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, prevJobUUID, tenants)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockStartNextJobUseCaseMockRecorder) Execute(ctx, prevJobUUID, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockStartNextJobUseCase)(nil).Execute), ctx, prevJobUUID, tenants)
}

// MockUpdateJobUseCase is a mock of UpdateJobUseCase interface
type MockUpdateJobUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateJobUseCaseMockRecorder
}

// MockUpdateJobUseCaseMockRecorder is the mock recorder for MockUpdateJobUseCase
type MockUpdateJobUseCaseMockRecorder struct {
	mock *MockUpdateJobUseCase
}

// NewMockUpdateJobUseCase creates a new mock instance
func NewMockUpdateJobUseCase(ctrl *gomock.Controller) *MockUpdateJobUseCase {
	mock := &MockUpdateJobUseCase{ctrl: ctrl}
	mock.recorder = &MockUpdateJobUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpdateJobUseCase) EXPECT() *MockUpdateJobUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockUpdateJobUseCase) Execute(ctx context.Context, jobEntity *entities.Job, nextStatus, logMessage string, tenants []string) (*entities.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, jobEntity, nextStatus, logMessage, tenants)
	ret0, _ := ret[0].(*entities.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockUpdateJobUseCaseMockRecorder) Execute(ctx, jobEntity, nextStatus, logMessage, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateJobUseCase)(nil).Execute), ctx, jobEntity, nextStatus, logMessage, tenants)
}

// MockUpdateChildrenUseCase is a mock of UpdateChildrenUseCase interface
type MockUpdateChildrenUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateChildrenUseCaseMockRecorder
}

// MockUpdateChildrenUseCaseMockRecorder is the mock recorder for MockUpdateChildrenUseCase
type MockUpdateChildrenUseCaseMockRecorder struct {
	mock *MockUpdateChildrenUseCase
}

// NewMockUpdateChildrenUseCase creates a new mock instance
func NewMockUpdateChildrenUseCase(ctrl *gomock.Controller) *MockUpdateChildrenUseCase {
	mock := &MockUpdateChildrenUseCase{ctrl: ctrl}
	mock.recorder = &MockUpdateChildrenUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpdateChildrenUseCase) EXPECT() *MockUpdateChildrenUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockUpdateChildrenUseCase) Execute(ctx context.Context, jobUUID, parentJobUUID, nextStatus string, tenants []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, jobUUID, parentJobUUID, nextStatus, tenants)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockUpdateChildrenUseCaseMockRecorder) Execute(ctx, jobUUID, parentJobUUID, nextStatus, tenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateChildrenUseCase)(nil).Execute), ctx, jobUUID, parentJobUUID, nextStatus, tenants)
}

// WithDBTransaction mocks base method
func (m *MockUpdateChildrenUseCase) WithDBTransaction(dbtx store.Tx) usecases.UpdateChildrenUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDBTransaction", dbtx)
	ret0, _ := ret[0].(usecases.UpdateChildrenUseCase)
	return ret0
}

// WithDBTransaction indicates an expected call of WithDBTransaction
func (mr *MockUpdateChildrenUseCaseMockRecorder) WithDBTransaction(dbtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDBTransaction", reflect.TypeOf((*MockUpdateChildrenUseCase)(nil).WithDBTransaction), dbtx)
}
