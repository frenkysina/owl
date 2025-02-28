// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/task/task.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	service "github.com/ibanyu/owl/service"
	task "github.com/ibanyu/owl/service/task"
	reflect "reflect"
)

// MockTaskDao is a mock of TaskDao interface
type MockTaskDao struct {
	ctrl     *gomock.Controller
	recorder *MockTaskDaoMockRecorder
}

// MockTaskDaoMockRecorder is the mock recorder for MockTaskDao
type MockTaskDaoMockRecorder struct {
	mock *MockTaskDao
}

// NewMockTaskDao creates a new mock instance
func NewMockTaskDao(ctrl *gomock.Controller) *MockTaskDao {
	mock := &MockTaskDao{ctrl: ctrl}
	mock.recorder = &MockTaskDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskDao) EXPECT() *MockTaskDaoMockRecorder {
	return m.recorder
}

// AddTask mocks base method
func (m *MockTaskDao) AddTask(task *task.DbInjectionTask) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", task)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTask indicates an expected call of AddTask
func (mr *MockTaskDaoMockRecorder) AddTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockTaskDao)(nil).AddTask), task)
}

// UpdateTask mocks base method
func (m *MockTaskDao) UpdateTask(task *task.DbInjectionTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask
func (mr *MockTaskDaoMockRecorder) UpdateTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskDao)(nil).UpdateTask), task)
}

// ListTask mocks base method
func (m *MockTaskDao) ListTask(pagination *service.Pagination, isDBA bool, status []task.ItemStatus) ([]task.DbInjectionTask, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTask", pagination, isDBA, status)
	ret0, _ := ret[0].([]task.DbInjectionTask)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTask indicates an expected call of ListTask
func (mr *MockTaskDaoMockRecorder) ListTask(pagination, isDBA, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTask", reflect.TypeOf((*MockTaskDao)(nil).ListTask), pagination, isDBA, status)
}

// ListHistoryTask mocks base method
func (m *MockTaskDao) ListHistoryTask(page *service.Pagination, isDBA bool) ([]task.DbInjectionTask, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHistoryTask", page, isDBA)
	ret0, _ := ret[0].([]task.DbInjectionTask)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListHistoryTask indicates an expected call of ListHistoryTask
func (mr *MockTaskDaoMockRecorder) ListHistoryTask(page, isDBA interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHistoryTask", reflect.TypeOf((*MockTaskDao)(nil).ListHistoryTask), page, isDBA)
}

// GetTask mocks base method
func (m *MockTaskDao) GetTask(id int64) (*task.DbInjectionTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", id)
	ret0, _ := ret[0].(*task.DbInjectionTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask
func (mr *MockTaskDaoMockRecorder) GetTask(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockTaskDao)(nil).GetTask), id)
}

// GetExecWaitTask mocks base method
func (m *MockTaskDao) GetExecWaitTask() ([]task.DbInjectionTask, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecWaitTask")
	ret0, _ := ret[0].([]task.DbInjectionTask)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetExecWaitTask indicates an expected call of GetExecWaitTask
func (mr *MockTaskDaoMockRecorder) GetExecWaitTask() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecWaitTask", reflect.TypeOf((*MockTaskDao)(nil).GetExecWaitTask))
}

// MockSubTaskDao is a mock of SubTaskDao interface
type MockSubTaskDao struct {
	ctrl     *gomock.Controller
	recorder *MockSubTaskDaoMockRecorder
}

// MockSubTaskDaoMockRecorder is the mock recorder for MockSubTaskDao
type MockSubTaskDaoMockRecorder struct {
	mock *MockSubTaskDao
}

// NewMockSubTaskDao creates a new mock instance
func NewMockSubTaskDao(ctrl *gomock.Controller) *MockSubTaskDao {
	mock := &MockSubTaskDao{ctrl: ctrl}
	mock.recorder = &MockSubTaskDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubTaskDao) EXPECT() *MockSubTaskDaoMockRecorder {
	return m.recorder
}

// UpdateItem mocks base method
func (m *MockSubTaskDao) UpdateItem(item *task.DbInjectionExecItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItem indicates an expected call of UpdateItem
func (mr *MockSubTaskDaoMockRecorder) UpdateItem(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockSubTaskDao)(nil).UpdateItem), item)
}

// DelItem mocks base method
func (m *MockSubTaskDao) DelItem(item *task.DbInjectionExecItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelItem", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelItem indicates an expected call of DelItem
func (mr *MockSubTaskDaoMockRecorder) DelItem(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelItem", reflect.TypeOf((*MockSubTaskDao)(nil).DelItem), item)
}

// UpdateItemByBackupId mocks base method
func (m *MockSubTaskDao) UpdateItemByBackupId(item *task.DbInjectionExecItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemByBackupId", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItemByBackupId indicates an expected call of UpdateItemByBackupId
func (mr *MockSubTaskDaoMockRecorder) UpdateItemByBackupId(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemByBackupId", reflect.TypeOf((*MockSubTaskDao)(nil).UpdateItemByBackupId), item)
}
