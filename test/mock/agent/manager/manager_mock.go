// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/pkg/agent/manager (interfaces: Manager)

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	go_observer "github.com/imkira/go-observer"
	cache "github.com/spiffe/spire/pkg/agent/manager/cache"
	common "github.com/spiffe/spire/proto/common"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// FetchJWTSVID mocks base method
func (m *MockManager) FetchJWTSVID(arg0 context.Context, arg1 string, arg2 []string) (string, error) {
	ret := m.ctrl.Call(m, "FetchJWTSVID", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJWTSVID indicates an expected call of FetchJWTSVID
func (mr *MockManagerMockRecorder) FetchJWTSVID(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJWTSVID", reflect.TypeOf((*MockManager)(nil).FetchJWTSVID), arg0, arg1, arg2)
}

// FetchWorkloadUpdate mocks base method
func (m *MockManager) FetchWorkloadUpdate(arg0 []*common.Selector) *cache.WorkloadUpdate {
	ret := m.ctrl.Call(m, "FetchWorkloadUpdate", arg0)
	ret0, _ := ret[0].(*cache.WorkloadUpdate)
	return ret0
}

// FetchWorkloadUpdate indicates an expected call of FetchWorkloadUpdate
func (mr *MockManagerMockRecorder) FetchWorkloadUpdate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchWorkloadUpdate", reflect.TypeOf((*MockManager)(nil).FetchWorkloadUpdate), arg0)
}

// Initialize mocks base method
func (m *MockManager) Initialize(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockManagerMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockManager)(nil).Initialize), arg0)
}

// MatchingEntries mocks base method
func (m *MockManager) MatchingEntries(arg0 []*common.Selector) []*cache.Entry {
	ret := m.ctrl.Call(m, "MatchingEntries", arg0)
	ret0, _ := ret[0].([]*cache.Entry)
	return ret0
}

// MatchingEntries indicates an expected call of MatchingEntries
func (mr *MockManagerMockRecorder) MatchingEntries(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchingEntries", reflect.TypeOf((*MockManager)(nil).MatchingEntries), arg0)
}

// Run mocks base method
func (m *MockManager) Run(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockManagerMockRecorder) Run(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockManager)(nil).Run), arg0)
}

// SubscribeToBundleChanges mocks base method
func (m *MockManager) SubscribeToBundleChanges() *cache.BundleStream {
	ret := m.ctrl.Call(m, "SubscribeToBundleChanges")
	ret0, _ := ret[0].(*cache.BundleStream)
	return ret0
}

// SubscribeToBundleChanges indicates an expected call of SubscribeToBundleChanges
func (mr *MockManagerMockRecorder) SubscribeToBundleChanges() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToBundleChanges", reflect.TypeOf((*MockManager)(nil).SubscribeToBundleChanges))
}

// SubscribeToCacheChanges mocks base method
func (m *MockManager) SubscribeToCacheChanges(arg0 cache.Selectors) cache.Subscriber {
	ret := m.ctrl.Call(m, "SubscribeToCacheChanges", arg0)
	ret0, _ := ret[0].(cache.Subscriber)
	return ret0
}

// SubscribeToCacheChanges indicates an expected call of SubscribeToCacheChanges
func (mr *MockManagerMockRecorder) SubscribeToCacheChanges(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToCacheChanges", reflect.TypeOf((*MockManager)(nil).SubscribeToCacheChanges), arg0)
}

// SubscribeToSVIDChanges mocks base method
func (m *MockManager) SubscribeToSVIDChanges() go_observer.Stream {
	ret := m.ctrl.Call(m, "SubscribeToSVIDChanges")
	ret0, _ := ret[0].(go_observer.Stream)
	return ret0
}

// SubscribeToSVIDChanges indicates an expected call of SubscribeToSVIDChanges
func (mr *MockManagerMockRecorder) SubscribeToSVIDChanges() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToSVIDChanges", reflect.TypeOf((*MockManager)(nil).SubscribeToSVIDChanges))
}
