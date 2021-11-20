// Code generated by MockGen. DO NOT EDIT.
// Source: ../iao/user/client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	user "tplgo/pkg/iao/user"
	result "tplgo/pkg/result"

	gomock "github.com/golang/mock/gomock"
)

// MockUserIao is a mock of UserIao interface.
type MockUserIao struct {
	ctrl     *gomock.Controller
	recorder *MockUserIaoMockRecorder
}

// MockUserIaoMockRecorder is the mock recorder for MockUserIao.
type MockUserIaoMockRecorder struct {
	mock *MockUserIao
}

// NewMockUserIao creates a new mock instance.
func NewMockUserIao(ctrl *gomock.Controller) *MockUserIao {
	mock := &MockUserIao{ctrl: ctrl}
	mock.recorder = &MockUserIaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserIao) EXPECT() *MockUserIaoMockRecorder {
	return m.recorder
}

// UserInfo mocks base method.
func (m *MockUserIao) UserInfo(ctx context.Context, params *user.ParamsUserInfo) (*user.UserInfo, result.Result) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserInfo", ctx, params)
	ret0, _ := ret[0].(*user.UserInfo)
	ret1, _ := ret[1].(result.Result)
	return ret0, ret1
}

// UserInfo indicates an expected call of UserInfo.
func (mr *MockUserIaoMockRecorder) UserInfo(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserInfo", reflect.TypeOf((*MockUserIao)(nil).UserInfo), ctx, params)
}
