package mocks

import (
	core "art-local/entity/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEventServiceInterface is a mock of EventServiceInterface interface.
type MockEventServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEventServiceInterfaceMockRecorder
}

// MockEventServiceInterfaceMockRecorder is the mock recorder for MockEventServiceInterface.
type MockEventServiceInterfaceMockRecorder struct {
	mock *MockEventServiceInterface
}

// NewMockEventServiceInterface creates a new mock instance.
func NewMockEventServiceInterface(ctrl *gomock.Controller) *MockEventServiceInterface {
	mock := &MockEventServiceInterface{ctrl: ctrl}
	mock.recorder = &MockEventServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventServiceInterface) EXPECT() *MockEventServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEventServiceInterface) Create(arg0 core.EventCore, arg1 string) (core.EventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(core.EventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockEventServiceInterfaceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEventServiceInterface)(nil).Create), arg0, arg1)
}

// CreateFollow mocks base method.
func (m *MockEventServiceInterface) CreateFollow(arg0 core.FollowEventCore, arg1 uint) (core.FollowEventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFollow", arg0, arg1)
	ret0, _ := ret[0].(core.FollowEventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFollow indicates an expected call of CreateFollow.
func (mr *MockEventServiceInterfaceMockRecorder) CreateFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFollow", reflect.TypeOf((*MockEventServiceInterface)(nil).CreateFollow), arg0, arg1)       
}

// Delete mocks base method.
func (m *MockEventServiceInterface) Delete(arg0 uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockEventServiceInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEventServiceInterface)(nil).Delete), arg0)
}

// FindEventsFollow mocks base method.
func (m *MockEventServiceInterface) FindEventsFollow(arg0 uint) []core.EventCore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventsFollow", arg0)
	ret0, _ := ret[0].([]core.EventCore)
	return ret0
}

// FindEventsFollow indicates an expected call of FindEventsFollow.
func (mr *MockEventServiceInterfaceMockRecorder) FindEventsFollow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventsFollow", reflect.TypeOf((*MockEventServiceInterface)(nil).FindEventsFollow), arg0)     
}

// GetAll mocks base method.
func (m *MockEventServiceInterface) GetAll() ([]core.EventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]core.EventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEventServiceInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEventServiceInterface)(nil).GetAll))
}

// GetAllFollowEvent mocks base method.
func (m *MockEventServiceInterface) GetAllFollowEvent(arg0 uint) ([]core.EventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFollowEvent", arg0)
	ret0, _ := ret[0].([]core.EventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFollowEvent indicates an expected call of GetAllFollowEvent.
func (mr *MockEventServiceInterfaceMockRecorder) GetAllFollowEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFollowEvent", reflect.TypeOf((*MockEventServiceInterface)(nil).GetAllFollowEvent), arg0)   
}

// GetById mocks base method.
func (m *MockEventServiceInterface) GetById(arg0 uint) (core.EventCore, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(core.EventCore)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetById indicates an expected call of GetById.
func (mr *MockEventServiceInterfaceMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockEventServiceInterface)(nil).GetById), arg0)
}

// GetByIdFollowEvent mocks base method.
func (m *MockEventServiceInterface) GetByIdFollowEvent(arg0 uint) (core.FollowEventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdFollowEvent", arg0)
	ret0, _ := ret[0].(core.FollowEventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIdFollowEvent indicates an expected call of GetByIdFollowEvent.
func (mr *MockEventServiceInterfaceMockRecorder) GetByIdFollowEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdFollowEvent", reflect.TypeOf((*MockEventServiceInterface)(nil).GetByIdFollowEvent), arg0) 
}

// Update mocks base method.
func (m *MockEventServiceInterface) Update(arg0 uint, arg1 core.EventCore) (core.EventCore, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(core.EventCore)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockEventServiceInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEventServiceInterface)(nil).Update), arg0, arg1)
}