// Code generated by MockGen. DO NOT EDIT.
// Source: rbac.go

// Package rbac is a generated GoMock package.
package rbac

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
)

// MockRBACCreator is a mock of RBACCreator interface.
type MockRBACCreator struct {
	ctrl     *gomock.Controller
	recorder *MockRBACCreatorMockRecorder
}

// MockRBACCreatorMockRecorder is the mock recorder for MockRBACCreator.
type MockRBACCreatorMockRecorder struct {
	mock *MockRBACCreator
}

// NewMockRBACCreator creates a new mock instance.
func NewMockRBACCreator(ctrl *gomock.Controller) *MockRBACCreator {
	mock := &MockRBACCreator{ctrl: ctrl}
	mock.recorder = &MockRBACCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRBACCreator) EXPECT() *MockRBACCreatorMockRecorder {
	return m.recorder
}

// CreateDevicePluginServiceAccount mocks base method.
func (m *MockRBACCreator) CreateDevicePluginServiceAccount(ctx context.Context, mod v1beta1.Module) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevicePluginServiceAccount", ctx, mod)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDevicePluginServiceAccount indicates an expected call of CreateDevicePluginServiceAccount.
func (mr *MockRBACCreatorMockRecorder) CreateDevicePluginServiceAccount(ctx, mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevicePluginServiceAccount", reflect.TypeOf((*MockRBACCreator)(nil).CreateDevicePluginServiceAccount), ctx, mod)
}

// CreateModuleLoaderServiceAccount mocks base method.
func (m *MockRBACCreator) CreateModuleLoaderServiceAccount(ctx context.Context, mod v1beta1.Module) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateModuleLoaderServiceAccount", ctx, mod)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateModuleLoaderServiceAccount indicates an expected call of CreateModuleLoaderServiceAccount.
func (mr *MockRBACCreatorMockRecorder) CreateModuleLoaderServiceAccount(ctx, mod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModuleLoaderServiceAccount", reflect.TypeOf((*MockRBACCreator)(nil).CreateModuleLoaderServiceAccount), ctx, mod)
}
