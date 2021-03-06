// Code generated by MockGen. DO NOT EDIT.
// Source: ./persistence/user.go

// Package mock_persistence is a generated GoMock package.
package mock_persistence

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	model "sampleRestApp/model"
)

// MockUserPersistence is a mock of UserPersistence interface
type MockUserPersistence struct {
	ctrl     *gomock.Controller
	recorder *MockUserPersistenceMockRecorder
}

// MockUserPersistenceMockRecorder is the mock recorder for MockUserPersistence
type MockUserPersistenceMockRecorder struct {
	mock *MockUserPersistence
}

// NewMockUserPersistence creates a new mock instance
func NewMockUserPersistence(ctrl *gomock.Controller) *MockUserPersistence {
	mock := &MockUserPersistence{ctrl: ctrl}
	mock.recorder = &MockUserPersistenceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserPersistence) EXPECT() *MockUserPersistenceMockRecorder {
	return m.recorder
}

// GetAllUser mocks base method
func (m *MockUserPersistence) GetAllUser() ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUser")
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUser indicates an expected call of GetAllUser
func (mr *MockUserPersistenceMockRecorder) GetAllUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUser", reflect.TypeOf((*MockUserPersistence)(nil).GetAllUser))
}

// CreateUser mocks base method
func (m *MockUserPersistence) CreateUser(arg0 model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserPersistenceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserPersistence)(nil).CreateUser), arg0)
}

// UpdateUser mocks base method
func (m *MockUserPersistence) UpdateUser(arg0 uint64, arg1 model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserPersistenceMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserPersistence)(nil).UpdateUser), arg0, arg1)
}

// DeleteUser mocks base method
func (m *MockUserPersistence) DeleteUser(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockUserPersistenceMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserPersistence)(nil).DeleteUser), arg0)
}
