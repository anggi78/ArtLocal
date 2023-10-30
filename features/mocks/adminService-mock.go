package mocks

import (
	core "art-local/entity/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAdminServiceInterface is a mock of AdminServiceInterface interface.
type MockAdminServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAdminServiceInterfaceMockRecorder
}

// MockAdminServiceInterfaceMockRecorder is the mock recorder for MockAdminServiceInterface.
type MockAdminServiceInterfaceMockRecorder struct {
	mock *MockAdminServiceInterface
}

// NewMockAdminServiceInterface creates a new mock instance.
func NewMockAdminServiceInterface(ctrl *gomock.Controller) *MockAdminServiceInterface {
	mock := &MockAdminServiceInterface{ctrl: ctrl}
	mock.recorder = &MockAdminServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminServiceInterface) EXPECT() *MockAdminServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateAdmin mocks base method.
func (m *MockAdminServiceInterface) CreateAdmin(arg0 core.Admin) (core.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", arg0)
	ret0, _ := ret[0].(core.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockAdminServiceInterfaceMockRecorder) CreateAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockAdminServiceInterface)(nil).CreateAdmin), arg0)
}

// CreateEvent mocks base method.
func (m *MockAdminServiceInterface) CreateEvent(arg0 core.EventCore) (core.EventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0)
	ret0, _ := ret[0].(core.EventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockAdminServiceInterfaceMockRecorder) CreateEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockAdminServiceInterface)(nil).CreateEvent), arg0)
}

// LoginAdmin mocks base method.
func (m *MockAdminServiceInterface) LoginAdmin(arg0, arg1 string) (core.Admin, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginAdmin", arg0, arg1)
	ret0, _ := ret[0].(core.Admin)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoginAdmin indicates an expected call of LoginAdmin.
func (mr *MockAdminServiceInterfaceMockRecorder) LoginAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginAdmin", reflect.TypeOf((*MockAdminServiceInterface)(nil).LoginAdmin), arg0, arg1)
}

// Update mocks base method.
func (m *MockAdminServiceInterface) Update(arg0 int, arg1 *core.Admin) (*core.Admin, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*core.Admin)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockAdminServiceInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdminServiceInterface)(nil).Update), arg0, arg1)
}