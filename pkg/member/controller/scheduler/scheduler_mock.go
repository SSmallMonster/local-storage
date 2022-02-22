// Code generated by MockGen. DO NOT EDIT.
// Source: scheduler.go

// Package scheduler is a generated GoMock package.
package scheduler

import (
	reflect "reflect"

	v1alpha1 "github.com/HwameiStor/local-storage/pkg/apis/uds/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockScheduler is a mock of Scheduler interface.
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler.
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance.
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// Allocate mocks base method.
func (m *MockScheduler) Allocate(vol *v1alpha1.LocalVolume) (*v1alpha1.VolumeConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Allocate", vol)
	ret0, _ := ret[0].(*v1alpha1.VolumeConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Allocate indicates an expected call of Allocate.
func (mr *MockSchedulerMockRecorder) Allocate(vol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allocate", reflect.TypeOf((*MockScheduler)(nil).Allocate), vol)
}

// GetNodeCandidates mocks base method.
func (m *MockScheduler) GetNodeCandidates(vol *v1alpha1.LocalVolume) ([]*v1alpha1.LocalStorageNode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeCandidates", vol)
	ret0, _ := ret[0].([]*v1alpha1.LocalStorageNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeCandidates indicates an expected call of GetNodeCandidates.
func (mr *MockSchedulerMockRecorder) GetNodeCandidates(vol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeCandidates", reflect.TypeOf((*MockScheduler)(nil).GetNodeCandidates), vol)
}

// Init mocks base method.
func (m *MockScheduler) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init.
func (mr *MockSchedulerMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockScheduler)(nil).Init))
}
