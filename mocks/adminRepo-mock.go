package mocks

import (
	core "art-local/features/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAdminRepoInterface is a mock of AdminRepoInterface interface.
type MockAdminRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRepoInterfaceMockRecorder
}

// MockAdminRepoInterfaceMockRecorder is the mock recorder for MockAdminRepoInterface.
type MockAdminRepoInterfaceMockRecorder struct {
	mock *MockAdminRepoInterface
}

// NewMockAdminRepoInterface creates a new mock instance.
func NewMockAdminRepoInterface(ctrl *gomock.Controller) *MockAdminRepoInterface {
	mock := &MockAdminRepoInterface{ctrl: ctrl}
	mock.recorder = &MockAdminRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminRepoInterface) EXPECT() *MockAdminRepoInterfaceMockRecorder {
	return m.recorder
}

// CreateAdmin mocks base method.
func (m *MockAdminRepoInterface) CreateAdmin(arg0 core.Admin) (core.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", arg0)
	ret0, _ := ret[0].(core.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockAdminRepoInterfaceMockRecorder) CreateAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockAdminRepoInterface)(nil).CreateAdmin), arg0)
}

// CreateEvent mocks base method.
func (m *MockAdminRepoInterface) CreateEvent(arg0 core.EventCore) (core.EventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0)
	ret0, _ := ret[0].(core.EventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockAdminRepoInterfaceMockRecorder) CreateEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockAdminRepoInterface)(nil).CreateEvent), arg0)
}

// DeleteEvent mocks base method.
func (m *MockAdminRepoInterface) DeleteEvent(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteEvent", arg0)
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockAdminRepoInterfaceMockRecorder) DeleteEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockAdminRepoInterface)(nil).DeleteEvent), arg0)
}

// FindByID mocks base method.
func (m *MockAdminRepoInterface) FindByID(arg0 int) (*core.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*core.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockAdminRepoInterfaceMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockAdminRepoInterface)(nil).FindByID), arg0)
}

// LoginAdmin mocks base method.
func (m *MockAdminRepoInterface) LoginAdmin(arg0, arg1 string) (core.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginAdmin", arg0, arg1)
	ret0, _ := ret[0].(core.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginAdmin indicates an expected call of LoginAdmin.
func (mr *MockAdminRepoInterfaceMockRecorder) LoginAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginAdmin", reflect.TypeOf((*MockAdminRepoInterface)(nil).LoginAdmin), arg0, arg1)
}

// Update mocks base method.
func (m *MockAdminRepoInterface) Update(arg0 int, arg1 *core.Admin) (*core.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*core.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAdminRepoInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdminRepoInterface)(nil).Update), arg0, arg1)
}