package mocks

import (
	core "art-local/entity/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockArtworkRepoInterface is a mock of ArtworkRepoInterface interface.
type MockArtworkRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockArtworkRepoInterfaceMockRecorder
}

// MockArtworkRepoInterfaceMockRecorder is the mock recorder for MockArtworkRepoInterface.
type MockArtworkRepoInterfaceMockRecorder struct {
	mock *MockArtworkRepoInterface
}

// NewMockArtworkRepoInterface creates a new mock instance.
func NewMockArtworkRepoInterface(ctrl *gomock.Controller) *MockArtworkRepoInterface {
	mock := &MockArtworkRepoInterface{ctrl: ctrl}
	mock.recorder = &MockArtworkRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArtworkRepoInterface) EXPECT() *MockArtworkRepoInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArtworkRepoInterface) Create(arg0 core.ArtworkCore) (core.ArtworkCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(core.ArtworkCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockArtworkRepoInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArtworkRepoInterface)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockArtworkRepoInterface) Delete(arg0 uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockArtworkRepoInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArtworkRepoInterface)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockArtworkRepoInterface) GetAll() ([]core.ArtworkCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]core.ArtworkCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockArtworkRepoInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockArtworkRepoInterface)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockArtworkRepoInterface) GetById(arg0 uint) (core.ArtworkCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(core.ArtworkCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockArtworkRepoInterfaceMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockArtworkRepoInterface)(nil).GetById), arg0)
}

// Update mocks base method.
func (m *MockArtworkRepoInterface) Update(arg0 uint, arg1 core.ArtworkCore) (core.ArtworkCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(core.ArtworkCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockArtworkRepoInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArtworkRepoInterface)(nil).Update), arg0, arg1)
}