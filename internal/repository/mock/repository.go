// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/asetriza/golang-web-app/internal/repository (interfaces: Authorization,Todo)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	common "github.com/asetriza/golang-web-app/internal/common"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(arg0 context.Context, arg1 common.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), arg0, arg1)
}

// CreateUserSession mocks base method.
func (m *MockAuthorization) CreateUserSession(arg0 context.Context, arg1 int, arg2, arg3 string, arg4 int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserSession", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserSession indicates an expected call of CreateUserSession.
func (mr *MockAuthorizationMockRecorder) CreateUserSession(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserSession", reflect.TypeOf((*MockAuthorization)(nil).CreateUserSession), arg0, arg1, arg2, arg3, arg4)
}

// GetUser mocks base method.
func (m *MockAuthorization) GetUser(arg0 context.Context, arg1, arg2 string) (common.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(common.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockAuthorizationMockRecorder) GetUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthorization)(nil).GetUser), arg0, arg1, arg2)
}

// GetUserSession mocks base method.
func (m *MockAuthorization) GetUserSession(arg0 context.Context, arg1 int, arg2 string) (common.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(common.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSession indicates an expected call of GetUserSession.
func (mr *MockAuthorizationMockRecorder) GetUserSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSession", reflect.TypeOf((*MockAuthorization)(nil).GetUserSession), arg0, arg1, arg2)
}

// UpdateUserSession mocks base method.
func (m *MockAuthorization) UpdateUserSession(arg0 context.Context, arg1 int, arg2 string, arg3 int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserSession", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserSession indicates an expected call of UpdateUserSession.
func (mr *MockAuthorizationMockRecorder) UpdateUserSession(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserSession", reflect.TypeOf((*MockAuthorization)(nil).UpdateUserSession), arg0, arg1, arg2, arg3)
}

// MockTodo is a mock of Todo interface.
type MockTodo struct {
	ctrl     *gomock.Controller
	recorder *MockTodoMockRecorder
}

// MockTodoMockRecorder is the mock recorder for MockTodo.
type MockTodoMockRecorder struct {
	mock *MockTodo
}

// NewMockTodo creates a new mock instance.
func NewMockTodo(ctrl *gomock.Controller) *MockTodo {
	mock := &MockTodo{ctrl: ctrl}
	mock.recorder = &MockTodoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodo) EXPECT() *MockTodoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodo) Create(arg0 context.Context, arg1 common.Todo) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockTodo) Delete(arg0 context.Context, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodo)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockTodo) Get(arg0 context.Context, arg1 int) (common.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(common.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTodoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTodo)(nil).Get), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockTodo) GetAll(arg0 context.Context, arg1 int) ([]common.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].([]common.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTodoMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodo)(nil).GetAll), arg0, arg1)
}

// Update mocks base method.
func (m *MockTodo) Update(arg0 context.Context, arg1 common.Todo) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTodoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodo)(nil).Update), arg0, arg1)
}
