package mocks

import (
	core "art-local/features/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepoInterface is a mock of UserRepoInterface interface.
type MockUserRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoInterfaceMockRecorder
}

// MockUserRepoInterfaceMockRecorder is the mock recorder for MockUserRepoInterface.
type MockUserRepoInterfaceMockRecorder struct {
	mock *MockUserRepoInterface
}

// NewMockUserRepoInterface creates a new mock instance.
func NewMockUserRepoInterface(ctrl *gomock.Controller) *MockUserRepoInterface {
	mock := &MockUserRepoInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepoInterface) EXPECT() *MockUserRepoInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepoInterface) CreateUser(arg0 core.User) (core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoInterfaceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepoInterface)(nil).CreateUser), arg0)
}

// Delete mocks base method.
func (m *MockUserRepoInterface) Delete(arg0 uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepoInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepoInterface)(nil).Delete), arg0)
}

// FindByID mocks base method.
func (m *MockUserRepoInterface) FindByID(arg0 uint) (*core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserRepoInterfaceMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserRepoInterface)(nil).FindByID), arg0)
}

// GetAll mocks base method.
func (m *MockUserRepoInterface) GetAll() ([]core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUserRepoInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserRepoInterface)(nil).GetAll))
}

// Login mocks base method.
func (m *MockUserRepoInterface) Login(arg0, arg1 string) (core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserRepoInterfaceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserRepoInterface)(nil).Login), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserRepoInterface) Update(arg0 uint, arg1 core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepoInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepoInterface)(nil).Update), arg0, arg1)
} 