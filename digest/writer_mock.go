// Copyright (c) 2018 Uber Technologies, Inc.
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/xichen2020/eventdb/digest (interfaces: FdWithDigestWriter)

package digest

import (
	"hash"
	"os"

	"github.com/golang/mock/gomock"
)

// Mock of FdWithDigestWriter interface
type MockFdWithDigestWriter struct {
	ctrl     *gomock.Controller
	recorder *_MockFdWithDigestWriterRecorder
}

// Recorder for MockFdWithDigestWriter (not exported)
type _MockFdWithDigestWriterRecorder struct {
	mock *MockFdWithDigestWriter
}

func NewMockFdWithDigestWriter(ctrl *gomock.Controller) *MockFdWithDigestWriter {
	mock := &MockFdWithDigestWriter{ctrl: ctrl}
	mock.recorder = &_MockFdWithDigestWriterRecorder{mock}
	return mock
}

func (_m *MockFdWithDigestWriter) EXPECT() *_MockFdWithDigestWriterRecorder {
	return _m.recorder
}

func (_m *MockFdWithDigestWriter) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFdWithDigestWriterRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockFdWithDigestWriter) Digest() hash.Hash32 {
	ret := _m.ctrl.Call(_m, "Digest")
	ret0, _ := ret[0].(hash.Hash32)
	return ret0
}

func (_mr *_MockFdWithDigestWriterRecorder) Digest() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Digest")
}

func (_m *MockFdWithDigestWriter) Fd() *os.File {
	ret := _m.ctrl.Call(_m, "Fd")
	ret0, _ := ret[0].(*os.File)
	return ret0
}

func (_mr *_MockFdWithDigestWriterRecorder) Fd() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fd")
}

func (_m *MockFdWithDigestWriter) Flush() error {
	ret := _m.ctrl.Call(_m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFdWithDigestWriterRecorder) Flush() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush")
}

func (_m *MockFdWithDigestWriter) Reset(_param0 *os.File) {
	_m.ctrl.Call(_m, "Reset", _param0)
}

func (_mr *_MockFdWithDigestWriterRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

func (_m *MockFdWithDigestWriter) Write(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Write", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFdWithDigestWriterRecorder) Write(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0)
}