// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	txschedulertypes "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/tx-scheduler"
	reflect "reflect"
)

// MockTransactionClient is a mock of TransactionClient interface
type MockTransactionClient struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionClientMockRecorder
}

// MockTransactionClientMockRecorder is the mock recorder for MockTransactionClient
type MockTransactionClientMockRecorder struct {
	mock *MockTransactionClient
}

// NewMockTransactionClient creates a new mock instance
func NewMockTransactionClient(ctrl *gomock.Controller) *MockTransactionClient {
	mock := &MockTransactionClient{ctrl: ctrl}
	mock.recorder = &MockTransactionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionClient) EXPECT() *MockTransactionClientMockRecorder {
	return m.recorder
}

// SendContractTransaction mocks base method
func (m *MockTransactionClient) SendContractTransaction(ctx context.Context, request *txschedulertypes.SendTransactionRequest) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendContractTransaction", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendContractTransaction indicates an expected call of SendContractTransaction
func (mr *MockTransactionClientMockRecorder) SendContractTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendContractTransaction", reflect.TypeOf((*MockTransactionClient)(nil).SendContractTransaction), ctx, request)
}

// SendDeployTransaction mocks base method
func (m *MockTransactionClient) SendDeployTransaction(ctx context.Context, request *txschedulertypes.DeployContractRequest) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDeployTransaction", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDeployTransaction indicates an expected call of SendDeployTransaction
func (mr *MockTransactionClientMockRecorder) SendDeployTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDeployTransaction", reflect.TypeOf((*MockTransactionClient)(nil).SendDeployTransaction), ctx, request)
}

// SendRawTransaction mocks base method
func (m *MockTransactionClient) SendRawTransaction(ctx context.Context, request *txschedulertypes.RawTransactionRequest) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRawTransaction", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawTransaction indicates an expected call of SendRawTransaction
func (mr *MockTransactionClientMockRecorder) SendRawTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawTransaction", reflect.TypeOf((*MockTransactionClient)(nil).SendRawTransaction), ctx, request)
}

// SendTransferTransaction mocks base method
func (m *MockTransactionClient) SendTransferTransaction(ctx context.Context, request *txschedulertypes.TransferRequest) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransferTransaction", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransferTransaction indicates an expected call of SendTransferTransaction
func (mr *MockTransactionClientMockRecorder) SendTransferTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransferTransaction", reflect.TypeOf((*MockTransactionClient)(nil).SendTransferTransaction), ctx, request)
}

// GetTxRequest mocks base method
func (m *MockTransactionClient) GetTxRequest(ctx context.Context, txRequestUUID string) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxRequest", ctx, txRequestUUID)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxRequest indicates an expected call of GetTxRequest
func (mr *MockTransactionClientMockRecorder) GetTxRequest(ctx, txRequestUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxRequest", reflect.TypeOf((*MockTransactionClient)(nil).GetTxRequest), ctx, txRequestUUID)
}

// MockScheduleClient is a mock of ScheduleClient interface
type MockScheduleClient struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleClientMockRecorder
}

// MockScheduleClientMockRecorder is the mock recorder for MockScheduleClient
type MockScheduleClientMockRecorder struct {
	mock *MockScheduleClient
}

// NewMockScheduleClient creates a new mock instance
func NewMockScheduleClient(ctrl *gomock.Controller) *MockScheduleClient {
	mock := &MockScheduleClient{ctrl: ctrl}
	mock.recorder = &MockScheduleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScheduleClient) EXPECT() *MockScheduleClientMockRecorder {
	return m.recorder
}

// GetSchedule mocks base method
func (m *MockScheduleClient) GetSchedule(ctx context.Context, scheduleUUID string) (*txschedulertypes.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", ctx, scheduleUUID)
	ret0, _ := ret[0].(*txschedulertypes.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule
func (mr *MockScheduleClientMockRecorder) GetSchedule(ctx, scheduleUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockScheduleClient)(nil).GetSchedule), ctx, scheduleUUID)
}

// GetSchedules mocks base method
func (m *MockScheduleClient) GetSchedules(ctx context.Context) ([]*txschedulertypes.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedules", ctx)
	ret0, _ := ret[0].([]*txschedulertypes.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedules indicates an expected call of GetSchedules
func (mr *MockScheduleClientMockRecorder) GetSchedules(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedules", reflect.TypeOf((*MockScheduleClient)(nil).GetSchedules), ctx)
}

// CreateSchedule mocks base method
func (m *MockScheduleClient) CreateSchedule(ctx context.Context, request *txschedulertypes.CreateScheduleRequest) (*txschedulertypes.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchedule", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSchedule indicates an expected call of CreateSchedule
func (mr *MockScheduleClientMockRecorder) CreateSchedule(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchedule", reflect.TypeOf((*MockScheduleClient)(nil).CreateSchedule), ctx, request)
}

// MockJobClient is a mock of JobClient interface
type MockJobClient struct {
	ctrl     *gomock.Controller
	recorder *MockJobClientMockRecorder
}

// MockJobClientMockRecorder is the mock recorder for MockJobClient
type MockJobClientMockRecorder struct {
	mock *MockJobClient
}

// NewMockJobClient creates a new mock instance
func NewMockJobClient(ctrl *gomock.Controller) *MockJobClient {
	mock := &MockJobClient{ctrl: ctrl}
	mock.recorder = &MockJobClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJobClient) EXPECT() *MockJobClientMockRecorder {
	return m.recorder
}

// GetJob mocks base method
func (m *MockJobClient) GetJob(ctx context.Context, jobUUID string) (*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJob", ctx, jobUUID)
	ret0, _ := ret[0].(*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob
func (mr *MockJobClientMockRecorder) GetJob(ctx, jobUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockJobClient)(nil).GetJob), ctx, jobUUID)
}

// GetJobs mocks base method
func (m *MockJobClient) GetJobs(ctx context.Context) ([]*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobs", ctx)
	ret0, _ := ret[0].([]*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs
func (mr *MockJobClientMockRecorder) GetJobs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockJobClient)(nil).GetJobs), ctx)
}

// CreateJob mocks base method
func (m *MockJobClient) CreateJob(ctx context.Context, request *txschedulertypes.CreateJobRequest) (*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob
func (mr *MockJobClientMockRecorder) CreateJob(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockJobClient)(nil).CreateJob), ctx, request)
}

// UpdateJob mocks base method
func (m *MockJobClient) UpdateJob(ctx context.Context, jobUUID string, request *txschedulertypes.UpdateJobRequest) (*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJob", ctx, jobUUID, request)
	ret0, _ := ret[0].(*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJob indicates an expected call of UpdateJob
func (mr *MockJobClientMockRecorder) UpdateJob(ctx, jobUUID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJob", reflect.TypeOf((*MockJobClient)(nil).UpdateJob), ctx, jobUUID, request)
}

// StartJob mocks base method
func (m *MockJobClient) StartJob(ctx context.Context, jobUUID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartJob", ctx, jobUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartJob indicates an expected call of StartJob
func (mr *MockJobClientMockRecorder) StartJob(ctx, jobUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartJob", reflect.TypeOf((*MockJobClient)(nil).StartJob), ctx, jobUUID)
}

// SearchJob mocks base method
func (m *MockJobClient) SearchJob(ctx context.Context, txHashes []string, chainUUID, status string) ([]*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchJob", ctx, txHashes, chainUUID, status)
	ret0, _ := ret[0].([]*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchJob indicates an expected call of SearchJob
func (mr *MockJobClientMockRecorder) SearchJob(ctx, txHashes, chainUUID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchJob", reflect.TypeOf((*MockJobClient)(nil).SearchJob), ctx, txHashes, chainUUID, status)
}

// MockTransactionSchedulerClient is a mock of TransactionSchedulerClient interface
type MockTransactionSchedulerClient struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionSchedulerClientMockRecorder
}

// MockTransactionSchedulerClientMockRecorder is the mock recorder for MockTransactionSchedulerClient
type MockTransactionSchedulerClientMockRecorder struct {
	mock *MockTransactionSchedulerClient
}

// NewMockTransactionSchedulerClient creates a new mock instance
func NewMockTransactionSchedulerClient(ctrl *gomock.Controller) *MockTransactionSchedulerClient {
	mock := &MockTransactionSchedulerClient{ctrl: ctrl}
	mock.recorder = &MockTransactionSchedulerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionSchedulerClient) EXPECT() *MockTransactionSchedulerClientMockRecorder {
	return m.recorder
}

// SendContractTransaction mocks base method
func (m *MockTransactionSchedulerClient) SendContractTransaction(ctx context.Context, request *txschedulertypes.SendTransactionRequest) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendContractTransaction", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendContractTransaction indicates an expected call of SendContractTransaction
func (mr *MockTransactionSchedulerClientMockRecorder) SendContractTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendContractTransaction", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).SendContractTransaction), ctx, request)
}

// SendDeployTransaction mocks base method
func (m *MockTransactionSchedulerClient) SendDeployTransaction(ctx context.Context, request *txschedulertypes.DeployContractRequest) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDeployTransaction", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDeployTransaction indicates an expected call of SendDeployTransaction
func (mr *MockTransactionSchedulerClientMockRecorder) SendDeployTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDeployTransaction", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).SendDeployTransaction), ctx, request)
}

// SendRawTransaction mocks base method
func (m *MockTransactionSchedulerClient) SendRawTransaction(ctx context.Context, request *txschedulertypes.RawTransactionRequest) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRawTransaction", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawTransaction indicates an expected call of SendRawTransaction
func (mr *MockTransactionSchedulerClientMockRecorder) SendRawTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawTransaction", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).SendRawTransaction), ctx, request)
}

// SendTransferTransaction mocks base method
func (m *MockTransactionSchedulerClient) SendTransferTransaction(ctx context.Context, request *txschedulertypes.TransferRequest) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransferTransaction", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransferTransaction indicates an expected call of SendTransferTransaction
func (mr *MockTransactionSchedulerClientMockRecorder) SendTransferTransaction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransferTransaction", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).SendTransferTransaction), ctx, request)
}

// GetTxRequest mocks base method
func (m *MockTransactionSchedulerClient) GetTxRequest(ctx context.Context, txRequestUUID string) (*txschedulertypes.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxRequest", ctx, txRequestUUID)
	ret0, _ := ret[0].(*txschedulertypes.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxRequest indicates an expected call of GetTxRequest
func (mr *MockTransactionSchedulerClientMockRecorder) GetTxRequest(ctx, txRequestUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxRequest", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).GetTxRequest), ctx, txRequestUUID)
}

// GetSchedule mocks base method
func (m *MockTransactionSchedulerClient) GetSchedule(ctx context.Context, scheduleUUID string) (*txschedulertypes.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", ctx, scheduleUUID)
	ret0, _ := ret[0].(*txschedulertypes.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule
func (mr *MockTransactionSchedulerClientMockRecorder) GetSchedule(ctx, scheduleUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).GetSchedule), ctx, scheduleUUID)
}

// GetSchedules mocks base method
func (m *MockTransactionSchedulerClient) GetSchedules(ctx context.Context) ([]*txschedulertypes.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedules", ctx)
	ret0, _ := ret[0].([]*txschedulertypes.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedules indicates an expected call of GetSchedules
func (mr *MockTransactionSchedulerClientMockRecorder) GetSchedules(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedules", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).GetSchedules), ctx)
}

// CreateSchedule mocks base method
func (m *MockTransactionSchedulerClient) CreateSchedule(ctx context.Context, request *txschedulertypes.CreateScheduleRequest) (*txschedulertypes.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchedule", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSchedule indicates an expected call of CreateSchedule
func (mr *MockTransactionSchedulerClientMockRecorder) CreateSchedule(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchedule", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).CreateSchedule), ctx, request)
}

// GetJob mocks base method
func (m *MockTransactionSchedulerClient) GetJob(ctx context.Context, jobUUID string) (*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJob", ctx, jobUUID)
	ret0, _ := ret[0].(*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob
func (mr *MockTransactionSchedulerClientMockRecorder) GetJob(ctx, jobUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).GetJob), ctx, jobUUID)
}

// GetJobs mocks base method
func (m *MockTransactionSchedulerClient) GetJobs(ctx context.Context) ([]*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobs", ctx)
	ret0, _ := ret[0].([]*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs
func (mr *MockTransactionSchedulerClientMockRecorder) GetJobs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).GetJobs), ctx)
}

// CreateJob mocks base method
func (m *MockTransactionSchedulerClient) CreateJob(ctx context.Context, request *txschedulertypes.CreateJobRequest) (*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", ctx, request)
	ret0, _ := ret[0].(*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob
func (mr *MockTransactionSchedulerClientMockRecorder) CreateJob(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).CreateJob), ctx, request)
}

// UpdateJob mocks base method
func (m *MockTransactionSchedulerClient) UpdateJob(ctx context.Context, jobUUID string, request *txschedulertypes.UpdateJobRequest) (*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJob", ctx, jobUUID, request)
	ret0, _ := ret[0].(*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJob indicates an expected call of UpdateJob
func (mr *MockTransactionSchedulerClientMockRecorder) UpdateJob(ctx, jobUUID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJob", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).UpdateJob), ctx, jobUUID, request)
}

// StartJob mocks base method
func (m *MockTransactionSchedulerClient) StartJob(ctx context.Context, jobUUID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartJob", ctx, jobUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartJob indicates an expected call of StartJob
func (mr *MockTransactionSchedulerClientMockRecorder) StartJob(ctx, jobUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartJob", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).StartJob), ctx, jobUUID)
}

// SearchJob mocks base method
func (m *MockTransactionSchedulerClient) SearchJob(ctx context.Context, txHashes []string, chainUUID, status string) ([]*txschedulertypes.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchJob", ctx, txHashes, chainUUID, status)
	ret0, _ := ret[0].([]*txschedulertypes.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchJob indicates an expected call of SearchJob
func (mr *MockTransactionSchedulerClientMockRecorder) SearchJob(ctx, txHashes, chainUUID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchJob", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).SearchJob), ctx, txHashes, chainUUID, status)
}
