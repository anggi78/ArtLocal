package mocks

import (
	core "art-local/features/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockArtServiceInterface is a mock of ArtServiceInterface interface.
type MockArtServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockArtServiceInterfaceMockRecorder
}

// MockArtServiceInterfaceMockRecorder is the mock recorder for MockArtServiceInterface.
type MockArtServiceInterfaceMockRecorder struct {
	mock *MockArtServiceInterface
}

// NewMockArtServiceInterface creates a new mock instance.
func NewMockArtServiceInterface(ctrl *gomock.Controller) *MockArtServiceInterface {
	mock := &MockArtServiceInterface{ctrl: ctrl}
	mock.recorder = &MockArtServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArtServiceInterface) EXPECT() *MockArtServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArtServiceInterface) Create(arg0 core.ArtworkCore) (core.ArtworkCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(core.ArtworkCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockArtServiceInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArtServiceInterface)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockArtServiceInterface) Delete(arg0 uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockArtServiceInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArtServiceInterface)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockArtServiceInterface) GetAll() ([]core.ArtworkCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]core.ArtworkCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockArtServiceInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockArtServiceInterface)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockArtServiceInterface) GetById(arg0 uint) (core.ArtworkCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(core.ArtworkCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockArtServiceInterfaceMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockArtServiceInterface)(nil).GetById), arg0)
}

// Update mocks base method.
func (m *MockArtServiceInterface) Update(arg0 uint, arg1 core.ArtworkCore) (core.ArtworkCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(core.ArtworkCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockArtServiceInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArtServiceInterface)(nil).Update), arg0, arg1)
}