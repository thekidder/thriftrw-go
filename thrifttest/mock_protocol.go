// Code generated by MockGen. DO NOT EDIT.

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
// Source: go.uber.org/thriftrw/protocol (interfaces: Protocol,EnvelopeAgnosticProtocol,Responder)

// Package thrifttest is a generated GoMock package.
package thrifttest

import (
	gomock "github.com/golang/mock/gomock"
	protocol "go.uber.org/thriftrw/protocol"
	wire "go.uber.org/thriftrw/wire"
	io "io"
	reflect "reflect"
)

// MockProtocol is a mock of Protocol interface
type MockProtocol struct {
	ctrl     *gomock.Controller
	recorder *MockProtocolMockRecorder
}

// MockProtocolMockRecorder is the mock recorder for MockProtocol
type MockProtocolMockRecorder struct {
	mock *MockProtocol
}

// NewMockProtocol creates a new mock instance
func NewMockProtocol(ctrl *gomock.Controller) *MockProtocol {
	mock := &MockProtocol{ctrl: ctrl}
	mock.recorder = &MockProtocolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProtocol) EXPECT() *MockProtocolMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockProtocol) Decode(arg0 io.ReaderAt, arg1 wire.Type) (wire.Value, error) {
	ret := m.ctrl.Call(m, "Decode", arg0, arg1)
	ret0, _ := ret[0].(wire.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockProtocolMockRecorder) Decode(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockProtocol)(nil).Decode), arg0, arg1)
}

// DecodeEnveloped mocks base method
func (m *MockProtocol) DecodeEnveloped(arg0 io.ReaderAt) (wire.Envelope, error) {
	ret := m.ctrl.Call(m, "DecodeEnveloped", arg0)
	ret0, _ := ret[0].(wire.Envelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeEnveloped indicates an expected call of DecodeEnveloped
func (mr *MockProtocolMockRecorder) DecodeEnveloped(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeEnveloped", reflect.TypeOf((*MockProtocol)(nil).DecodeEnveloped), arg0)
}

// Encode mocks base method
func (m *MockProtocol) Encode(arg0 wire.Value, arg1 io.Writer) error {
	ret := m.ctrl.Call(m, "Encode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Encode indicates an expected call of Encode
func (mr *MockProtocolMockRecorder) Encode(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockProtocol)(nil).Encode), arg0, arg1)
}

// EncodeEnveloped mocks base method
func (m *MockProtocol) EncodeEnveloped(arg0 wire.Envelope, arg1 io.Writer) error {
	ret := m.ctrl.Call(m, "EncodeEnveloped", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EncodeEnveloped indicates an expected call of EncodeEnveloped
func (mr *MockProtocolMockRecorder) EncodeEnveloped(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncodeEnveloped", reflect.TypeOf((*MockProtocol)(nil).EncodeEnveloped), arg0, arg1)
}

// MockEnvelopeAgnosticProtocol is a mock of EnvelopeAgnosticProtocol interface
type MockEnvelopeAgnosticProtocol struct {
	ctrl     *gomock.Controller
	recorder *MockEnvelopeAgnosticProtocolMockRecorder
}

// MockEnvelopeAgnosticProtocolMockRecorder is the mock recorder for MockEnvelopeAgnosticProtocol
type MockEnvelopeAgnosticProtocolMockRecorder struct {
	mock *MockEnvelopeAgnosticProtocol
}

// NewMockEnvelopeAgnosticProtocol creates a new mock instance
func NewMockEnvelopeAgnosticProtocol(ctrl *gomock.Controller) *MockEnvelopeAgnosticProtocol {
	mock := &MockEnvelopeAgnosticProtocol{ctrl: ctrl}
	mock.recorder = &MockEnvelopeAgnosticProtocolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvelopeAgnosticProtocol) EXPECT() *MockEnvelopeAgnosticProtocolMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockEnvelopeAgnosticProtocol) Decode(arg0 io.ReaderAt, arg1 wire.Type) (wire.Value, error) {
	ret := m.ctrl.Call(m, "Decode", arg0, arg1)
	ret0, _ := ret[0].(wire.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockEnvelopeAgnosticProtocolMockRecorder) Decode(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockEnvelopeAgnosticProtocol)(nil).Decode), arg0, arg1)
}

// DecodeEnveloped mocks base method
func (m *MockEnvelopeAgnosticProtocol) DecodeEnveloped(arg0 io.ReaderAt) (wire.Envelope, error) {
	ret := m.ctrl.Call(m, "DecodeEnveloped", arg0)
	ret0, _ := ret[0].(wire.Envelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeEnveloped indicates an expected call of DecodeEnveloped
func (mr *MockEnvelopeAgnosticProtocolMockRecorder) DecodeEnveloped(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeEnveloped", reflect.TypeOf((*MockEnvelopeAgnosticProtocol)(nil).DecodeEnveloped), arg0)
}

// DecodeRequest mocks base method
func (m *MockEnvelopeAgnosticProtocol) DecodeRequest(arg0 wire.EnvelopeType, arg1 io.ReaderAt) (wire.Value, protocol.Responder, error) {
	ret := m.ctrl.Call(m, "DecodeRequest", arg0, arg1)
	ret0, _ := ret[0].(wire.Value)
	ret1, _ := ret[1].(protocol.Responder)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DecodeRequest indicates an expected call of DecodeRequest
func (mr *MockEnvelopeAgnosticProtocolMockRecorder) DecodeRequest(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeRequest", reflect.TypeOf((*MockEnvelopeAgnosticProtocol)(nil).DecodeRequest), arg0, arg1)
}

// Encode mocks base method
func (m *MockEnvelopeAgnosticProtocol) Encode(arg0 wire.Value, arg1 io.Writer) error {
	ret := m.ctrl.Call(m, "Encode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Encode indicates an expected call of Encode
func (mr *MockEnvelopeAgnosticProtocolMockRecorder) Encode(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEnvelopeAgnosticProtocol)(nil).Encode), arg0, arg1)
}

// EncodeEnveloped mocks base method
func (m *MockEnvelopeAgnosticProtocol) EncodeEnveloped(arg0 wire.Envelope, arg1 io.Writer) error {
	ret := m.ctrl.Call(m, "EncodeEnveloped", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EncodeEnveloped indicates an expected call of EncodeEnveloped
func (mr *MockEnvelopeAgnosticProtocolMockRecorder) EncodeEnveloped(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncodeEnveloped", reflect.TypeOf((*MockEnvelopeAgnosticProtocol)(nil).EncodeEnveloped), arg0, arg1)
}

// MockResponder is a mock of Responder interface
type MockResponder struct {
	ctrl     *gomock.Controller
	recorder *MockResponderMockRecorder
}

// MockResponderMockRecorder is the mock recorder for MockResponder
type MockResponderMockRecorder struct {
	mock *MockResponder
}

// NewMockResponder creates a new mock instance
func NewMockResponder(ctrl *gomock.Controller) *MockResponder {
	mock := &MockResponder{ctrl: ctrl}
	mock.recorder = &MockResponderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResponder) EXPECT() *MockResponderMockRecorder {
	return m.recorder
}

// EncodeResponse mocks base method
func (m *MockResponder) EncodeResponse(arg0 wire.Value, arg1 wire.EnvelopeType, arg2 io.Writer) error {
	ret := m.ctrl.Call(m, "EncodeResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EncodeResponse indicates an expected call of EncodeResponse
func (mr *MockResponderMockRecorder) EncodeResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncodeResponse", reflect.TypeOf((*MockResponder)(nil).EncodeResponse), arg0, arg1, arg2)
}