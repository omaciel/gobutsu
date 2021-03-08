// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/omaciel/gobutsu/db/sqlc (interfaces: App)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/omaciel/gobutsu/db/sqlc"
)

// MockApp is a mock of App interface.
type MockApp struct {
	ctrl     *gomock.Controller
	recorder *MockAppMockRecorder
}

// MockAppMockRecorder is the mock recorder for MockApp.
type MockAppMockRecorder struct {
	mock *MockApp
}

// NewMockApp creates a new mock instance.
func NewMockApp(ctrl *gomock.Controller) *MockApp {
	mock := &MockApp{ctrl: ctrl}
	mock.recorder = &MockAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApp) EXPECT() *MockAppMockRecorder {
	return m.recorder
}

// CreateTestCase mocks base method.
func (m *MockApp) CreateTestCase(arg0 context.Context, arg1 db.CreateTestCaseParams) (db.Testcase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestCase", arg0, arg1)
	ret0, _ := ret[0].(db.Testcase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestCase indicates an expected call of CreateTestCase.
func (mr *MockAppMockRecorder) CreateTestCase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestCase", reflect.TypeOf((*MockApp)(nil).CreateTestCase), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockApp) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAppMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockApp)(nil).CreateUser), arg0, arg1)
}

// DeleteTestCase mocks base method.
func (m *MockApp) DeleteTestCase(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTestCase", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTestCase indicates an expected call of DeleteTestCase.
func (mr *MockAppMockRecorder) DeleteTestCase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTestCase", reflect.TypeOf((*MockApp)(nil).DeleteTestCase), arg0, arg1)
}

// GetTestCase mocks base method.
func (m *MockApp) GetTestCase(arg0 context.Context, arg1 int64) (db.Testcase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestCase", arg0, arg1)
	ret0, _ := ret[0].(db.Testcase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestCase indicates an expected call of GetTestCase.
func (mr *MockAppMockRecorder) GetTestCase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestCase", reflect.TypeOf((*MockApp)(nil).GetTestCase), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockApp) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockAppMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockApp)(nil).GetUser), arg0, arg1)
}

// ListTestCases mocks base method.
func (m *MockApp) ListTestCases(arg0 context.Context, arg1 db.ListTestCasesParams) ([]db.Testcase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTestCases", arg0, arg1)
	ret0, _ := ret[0].([]db.Testcase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTestCases indicates an expected call of ListTestCases.
func (mr *MockAppMockRecorder) ListTestCases(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTestCases", reflect.TypeOf((*MockApp)(nil).ListTestCases), arg0, arg1)
}