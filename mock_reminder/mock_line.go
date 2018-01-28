// Code generated by MockGen. DO NOT EDIT.
// Source: reminder/line.go

// Package mock_reminder is a generated GoMock package.
package mock_reminder

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLineService is a mock of LineService interface
type MockLineService struct {
	ctrl     *gomock.Controller
	recorder *MockLineServiceMockRecorder
}

// MockLineServiceMockRecorder is the mock recorder for MockLineService
type MockLineServiceMockRecorder struct {
	mock *MockLineService
}

// NewMockLineService creates a new mock instance
func NewMockLineService(ctrl *gomock.Controller) *MockLineService {
	mock := &MockLineService{ctrl: ctrl}
	mock.recorder = &MockLineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLineService) EXPECT() *MockLineServiceMockRecorder {
	return m.recorder
}

// PostMessage mocks base method
func (m *MockLineService) PostMessage(message string) error {
	ret := m.ctrl.Call(m, "PostMessage", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostMessage indicates an expected call of PostMessage
func (mr *MockLineServiceMockRecorder) PostMessage(message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessage", reflect.TypeOf((*MockLineService)(nil).PostMessage), message)
}

// ReplyMessage mocks base method
func (m *MockLineService) ReplyMessage(token, message string) error {
	ret := m.ctrl.Call(m, "ReplyMessage", token, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplyMessage indicates an expected call of ReplyMessage
func (mr *MockLineServiceMockRecorder) ReplyMessage(token, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyMessage", reflect.TypeOf((*MockLineService)(nil).ReplyMessage), token, message)
}

// GetProfile mocks base method
func (m *MockLineService) GetProfile(id string) (string, error) {
	ret := m.ctrl.Call(m, "GetProfile", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile
func (mr *MockLineServiceMockRecorder) GetProfile(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockLineService)(nil).GetProfile), id)
}

// PostReport mocks base method
func (m *MockLineService) PostReport(id string) (string, error) {
	ret := m.ctrl.Call(m, "PostReport", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostReport indicates an expected call of PostReport
func (mr *MockLineServiceMockRecorder) PostReport(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostReport", reflect.TypeOf((*MockLineService)(nil).PostReport), id)
}

// PostReminder mocks base method
func (m *MockLineService) PostReminder(id string) (string, error) {
	ret := m.ctrl.Call(m, "PostReminder", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostReminder indicates an expected call of PostReminder
func (mr *MockLineServiceMockRecorder) PostReminder(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostReminder", reflect.TypeOf((*MockLineService)(nil).PostReminder), id)
}
