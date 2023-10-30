package mocks

import (
	core "art-local/entity/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEventRepoInterface is a mock of EventRepoInterface interface.
type MockEventRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepoInterfaceMockRecorder
}

// MockEventRepoInterfaceMockRecorder is the mock recorder for MockEventRepoInterface.
type MockEventRepoInterfaceMockRecorder struct {
	mock *MockEventRepoInterface
}

// NewMockEventRepoInterface creates a new mock instance.
func NewMockEventRepoInterface(ctrl *gomock.Controller) *MockEventRepoInterface {
	mock := &MockEventRepoInterface{ctrl: ctrl}
	mock.recorder = &MockEventRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepoInterface) EXPECT() *MockEventRepoInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEventRepoInterface) Create(arg0 core.EventCore) (core.EventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(core.EventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockEventRepoInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEventRepoInterface)(nil).Create), arg0)
}

// CreateFollow mocks base method.
func (m *MockEventRepoInterface) CreateFollow(arg0 core.FollowEventCore, arg1 uint) (core.FollowEventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFollow", arg0, arg1)
	ret0, _ := ret[0].(core.FollowEventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFollow indicates an expected call of CreateFollow.
func (mr *MockEventRepoInterfaceMockRecorder) CreateFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFollow", reflect.TypeOf((*MockEventRepoInterface)(nil).CreateFollow), arg0, arg1)
}

// Delete mocks base method.
func (m *MockEventRepoInterface) Delete(arg0 uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockEventRepoInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEventRepoInterface)(nil).Delete), arg0)
}

// FindEventByID mocks base method.
func (m *MockEventRepoInterface) FindEventByID(arg0 uint) core.EventCore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventByID", arg0)
	ret0, _ := ret[0].(core.EventCore)
	return ret0
}

// FindEventByID indicates an expected call of FindEventByID.
func (mr *MockEventRepoInterfaceMockRecorder) FindEventByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventByID", reflect.TypeOf((*MockEventRepoInterface)(nil).FindEventByID), arg0)
}

// FindEventID mocks base method.
func (m *MockEventRepoInterface) FindEventID(arg0 uint) []core.FollowEventCore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventID", arg0)
	ret0, _ := ret[0].([]core.FollowEventCore)
	return ret0
}

// FindEventID indicates an expected call of FindEventID.
func (mr *MockEventRepoInterfaceMockRecorder) FindEventID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventID", reflect.TypeOf((*MockEventRepoInterface)(nil).FindEventID), arg0)
}

// FindName mocks base method.
func (m *MockEventRepoInterface) FindName(arg0 uint) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FindName indicates an expected call of FindName.
func (mr *MockEventRepoInterfaceMockRecorder) FindName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindName", reflect.TypeOf((*MockEventRepoInterface)(nil).FindName), arg0)
}

// GetAll mocks base method.
func (m *MockEventRepoInterface) GetAll() ([]core.EventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]core.EventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEventRepoInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEventRepoInterface)(nil).GetAll))
}

// GetAllFollowEvent mocks base method.
func (m *MockEventRepoInterface) GetAllFollowEvent(arg0 uint) ([]core.EventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFollowEvent", arg0)
	ret0, _ := ret[0].([]core.EventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFollowEvent indicates an expected call of GetAllFollowEvent.
func (mr *MockEventRepoInterfaceMockRecorder) GetAllFollowEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFollowEvent", reflect.TypeOf((*MockEventRepoInterface)(nil).GetAllFollowEvent), arg0)      
}

// GetById mocks base method.
func (m *MockEventRepoInterface) GetById(arg0 uint) (core.EventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(core.EventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockEventRepoInterfaceMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockEventRepoInterface)(nil).GetById), arg0)
}

// GetByIdFollowEvent mocks base method.
func (m *MockEventRepoInterface) GetByIdFollowEvent(arg0 uint) (core.FollowEventCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdFollowEvent", arg0)
	ret0, _ := ret[0].(core.FollowEventCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIdFollowEvent indicates an expected call of GetByIdFollowEvent.
func (mr *MockEventRepoInterfaceMockRecorder) GetByIdFollowEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdFollowEvent", reflect.TypeOf((*MockEventRepoInterface)(nil).GetByIdFollowEvent), arg0)    
}

// Update mocks base method.
func (m *MockEventRepoInterface) Update(arg0 uint, arg1 core.EventCore) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockEventRepoInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEventRepoInterface)(nil).Update), arg0, arg1)
}