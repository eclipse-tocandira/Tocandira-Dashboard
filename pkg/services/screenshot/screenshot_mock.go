// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/grafana/grafana/pkg/services/screenshot (interfaces: ScreenshotService)

// Package screenshot is a generated GoMock package.
package screenshot

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockScreenshotService is a mock of ScreenshotService interface.
type MockScreenshotService struct {
	ctrl     *gomock.Controller
	recorder *MockScreenshotServiceMockRecorder
}

// MockScreenshotServiceMockRecorder is the mock recorder for MockScreenshotService.
type MockScreenshotServiceMockRecorder struct {
	mock *MockScreenshotService
}

// NewMockScreenshotService creates a new mock instance.
func NewMockScreenshotService(ctrl *gomock.Controller) *MockScreenshotService {
	mock := &MockScreenshotService{ctrl: ctrl}
	mock.recorder = &MockScreenshotServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScreenshotService) EXPECT() *MockScreenshotServiceMockRecorder {
	return m.recorder
}

// Take mocks base method.
func (m *MockScreenshotService) Take(arg0 context.Context, arg1 ScreenshotOptions) (*Screenshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Take", arg0, arg1)
	ret0, _ := ret[0].(*Screenshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Take indicates an expected call of Take.
func (mr *MockScreenshotServiceMockRecorder) Take(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Take", reflect.TypeOf((*MockScreenshotService)(nil).Take), arg0, arg1)
}
