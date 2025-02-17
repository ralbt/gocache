// Code generated by MockGen. DO NOT EDIT.
// Source: store/go_cache.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockGoCacheClientInterface is a mock of GoCacheClientInterface interface.
type MockGoCacheClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGoCacheClientInterfaceMockRecorder
}

// MockGoCacheClientInterfaceMockRecorder is the mock recorder for MockGoCacheClientInterface.
type MockGoCacheClientInterfaceMockRecorder struct {
	mock *MockGoCacheClientInterface
}

// NewMockGoCacheClientInterface creates a new mock instance.
func NewMockGoCacheClientInterface(ctrl *gomock.Controller) *MockGoCacheClientInterface {
	mock := &MockGoCacheClientInterface{ctrl: ctrl}
	mock.recorder = &MockGoCacheClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGoCacheClientInterface) EXPECT() *MockGoCacheClientInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockGoCacheClientInterface) Delete(k string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", k)
}

// Delete indicates an expected call of Delete.
func (mr *MockGoCacheClientInterfaceMockRecorder) Delete(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGoCacheClientInterface)(nil).Delete), k)
}

// Flush mocks base method.
func (m *MockGoCacheClientInterface) Flush() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush")
}

// Flush indicates an expected call of Flush.
func (mr *MockGoCacheClientInterfaceMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockGoCacheClientInterface)(nil).Flush))
}

// Get mocks base method.
func (m *MockGoCacheClientInterface) Get(k string) (any, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", k)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGoCacheClientInterfaceMockRecorder) Get(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGoCacheClientInterface)(nil).Get), k)
}

// GetWithExpiration mocks base method.
func (m *MockGoCacheClientInterface) GetWithExpiration(k string) (any, time.Time, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithExpiration", k)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// GetWithExpiration indicates an expected call of GetWithExpiration.
func (mr *MockGoCacheClientInterfaceMockRecorder) GetWithExpiration(k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithExpiration", reflect.TypeOf((*MockGoCacheClientInterface)(nil).GetWithExpiration), k)
}

// Set mocks base method.
func (m *MockGoCacheClientInterface) Set(k string, x any, d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", k, x, d)
}

// Set indicates an expected call of Set.
func (mr *MockGoCacheClientInterfaceMockRecorder) Set(k, x, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockGoCacheClientInterface)(nil).Set), k, x, d)
}
