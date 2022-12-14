// Code generated by MockGen. DO NOT EDIT.
// Source: maker.go

// Package job is a generated GoMock package.
package job

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	v1 "k8s.io/api/batch/v1"
)

// MockMaker is a mock of Maker interface.
type MockMaker struct {
	ctrl     *gomock.Controller
	recorder *MockMakerMockRecorder
}

// MockMakerMockRecorder is the mock recorder for MockMaker.
type MockMakerMockRecorder struct {
	mock *MockMaker
}

// NewMockMaker creates a new mock instance.
func NewMockMaker(ctrl *gomock.Controller) *MockMaker {
	mock := &MockMaker{ctrl: ctrl}
	mock.recorder = &MockMakerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaker) EXPECT() *MockMakerMockRecorder {
	return m.recorder
}

// MakeJobTemplate mocks base method.
func (m *MockMaker) MakeJobTemplate(ctx context.Context, mod v1beta1.Module, buildConfig *v1beta1.Build, targetKernel, containerImage string, pushImage bool, registryTLS *v1beta1.TLSOptions) (*v1.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeJobTemplate", ctx, mod, buildConfig, targetKernel, containerImage, pushImage, registryTLS)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeJobTemplate indicates an expected call of MakeJobTemplate.
func (mr *MockMakerMockRecorder) MakeJobTemplate(ctx, mod, buildConfig, targetKernel, containerImage, pushImage, registryTLS interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeJobTemplate", reflect.TypeOf((*MockMaker)(nil).MakeJobTemplate), ctx, mod, buildConfig, targetKernel, containerImage, pushImage, registryTLS)
}
