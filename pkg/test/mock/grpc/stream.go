// Code generated by MockGen. DO NOT EDIT.
// Source: google.golang.org/grpc (interfaces: ServerStream,ClientStream)

// Package mock_grpc is a generated GoMock package.
package mock_grpc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
)

// MockServerStream is a mock of ServerStream interface.
type MockServerStream struct {
	ctrl     *gomock.Controller
	recorder *MockServerStreamMockRecorder
}

// MockServerStreamMockRecorder is the mock recorder for MockServerStream.
type MockServerStreamMockRecorder struct {
	mock *MockServerStream
}

// NewMockServerStream creates a new mock instance.
func NewMockServerStream(ctrl *gomock.Controller) *MockServerStream {
	mock := &MockServerStream{ctrl: ctrl}
	mock.recorder = &MockServerStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerStream) EXPECT() *MockServerStreamMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockServerStream) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockServerStreamMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockServerStream)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockServerStream) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockServerStreamMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockServerStream)(nil).RecvMsg), arg0)
}

// SendHeader mocks base method.
func (m *MockServerStream) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockServerStreamMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockServerStream)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockServerStream) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockServerStreamMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockServerStream)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockServerStream) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockServerStreamMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockServerStream)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockServerStream) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockServerStreamMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockServerStream)(nil).SetTrailer), arg0)
}

// MockClientStream is a mock of ClientStream interface.
type MockClientStream struct {
	ctrl     *gomock.Controller
	recorder *MockClientStreamMockRecorder
}

// MockClientStreamMockRecorder is the mock recorder for MockClientStream.
type MockClientStreamMockRecorder struct {
	mock *MockClientStream
}

// NewMockClientStream creates a new mock instance.
func NewMockClientStream(ctrl *gomock.Controller) *MockClientStream {
	mock := &MockClientStream{ctrl: ctrl}
	mock.recorder = &MockClientStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientStream) EXPECT() *MockClientStreamMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockClientStream) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockClientStreamMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockClientStream)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockClientStream) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockClientStreamMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockClientStream)(nil).Context))
}

// Header mocks base method.
func (m *MockClientStream) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockClientStreamMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockClientStream)(nil).Header))
}

// RecvMsg mocks base method.
func (m *MockClientStream) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockClientStreamMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockClientStream)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockClientStream) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockClientStreamMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockClientStream)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockClientStream) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockClientStreamMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockClientStream)(nil).Trailer))
}
