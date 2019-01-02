// Copyright (c) 2019 Uber Technologies, Inc.
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
// Source: github.com/xichen2020/eventdb/encoding (interfaces: RewindableIntIterator,RewindableStringIterator,ForwardDoubleIterator,ForwardBoolIterator,RewindableTimeIterator)

package encoding

import (
	"github.com/golang/mock/gomock"
)

// Mock of RewindableIntIterator interface
type MockRewindableIntIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockRewindableIntIteratorRecorder
}

// Recorder for MockRewindableIntIterator (not exported)
type _MockRewindableIntIteratorRecorder struct {
	mock *MockRewindableIntIterator
}

func NewMockRewindableIntIterator(ctrl *gomock.Controller) *MockRewindableIntIterator {
	mock := &MockRewindableIntIterator{ctrl: ctrl}
	mock.recorder = &_MockRewindableIntIteratorRecorder{mock}
	return mock
}

func (_m *MockRewindableIntIterator) EXPECT() *_MockRewindableIntIteratorRecorder {
	return _m.recorder
}

func (_m *MockRewindableIntIterator) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRewindableIntIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockRewindableIntIterator) Current() int {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockRewindableIntIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockRewindableIntIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRewindableIntIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockRewindableIntIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockRewindableIntIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockRewindableIntIterator) Reset(_param0 []int) {
	_m.ctrl.Call(_m, "Reset", _param0)
}

func (_mr *_MockRewindableIntIteratorRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

func (_m *MockRewindableIntIterator) Rewind() {
	_m.ctrl.Call(_m, "Rewind")
}

func (_mr *_MockRewindableIntIteratorRecorder) Rewind() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rewind")
}

// Mock of RewindableStringIterator interface
type MockRewindableStringIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockRewindableStringIteratorRecorder
}

// Recorder for MockRewindableStringIterator (not exported)
type _MockRewindableStringIteratorRecorder struct {
	mock *MockRewindableStringIterator
}

func NewMockRewindableStringIterator(ctrl *gomock.Controller) *MockRewindableStringIterator {
	mock := &MockRewindableStringIterator{ctrl: ctrl}
	mock.recorder = &_MockRewindableStringIteratorRecorder{mock}
	return mock
}

func (_m *MockRewindableStringIterator) EXPECT() *_MockRewindableStringIteratorRecorder {
	return _m.recorder
}

func (_m *MockRewindableStringIterator) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRewindableStringIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockRewindableStringIterator) Current() string {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRewindableStringIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockRewindableStringIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRewindableStringIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockRewindableStringIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockRewindableStringIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockRewindableStringIterator) Reset(_param0 []string) {
	_m.ctrl.Call(_m, "Reset", _param0)
}

func (_mr *_MockRewindableStringIteratorRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

func (_m *MockRewindableStringIterator) Rewind() {
	_m.ctrl.Call(_m, "Rewind")
}

func (_mr *_MockRewindableStringIteratorRecorder) Rewind() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rewind")
}

// Mock of ForwardDoubleIterator interface
type MockForwardDoubleIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockForwardDoubleIteratorRecorder
}

// Recorder for MockForwardDoubleIterator (not exported)
type _MockForwardDoubleIteratorRecorder struct {
	mock *MockForwardDoubleIterator
}

func NewMockForwardDoubleIterator(ctrl *gomock.Controller) *MockForwardDoubleIterator {
	mock := &MockForwardDoubleIterator{ctrl: ctrl}
	mock.recorder = &_MockForwardDoubleIteratorRecorder{mock}
	return mock
}

func (_m *MockForwardDoubleIterator) EXPECT() *_MockForwardDoubleIteratorRecorder {
	return _m.recorder
}

func (_m *MockForwardDoubleIterator) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockForwardDoubleIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockForwardDoubleIterator) Current() float64 {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(float64)
	return ret0
}

func (_mr *_MockForwardDoubleIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockForwardDoubleIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockForwardDoubleIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockForwardDoubleIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockForwardDoubleIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

// Mock of ForwardBoolIterator interface
type MockForwardBoolIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockForwardBoolIteratorRecorder
}

// Recorder for MockForwardBoolIterator (not exported)
type _MockForwardBoolIteratorRecorder struct {
	mock *MockForwardBoolIterator
}

func NewMockForwardBoolIterator(ctrl *gomock.Controller) *MockForwardBoolIterator {
	mock := &MockForwardBoolIterator{ctrl: ctrl}
	mock.recorder = &_MockForwardBoolIteratorRecorder{mock}
	return mock
}

func (_m *MockForwardBoolIterator) EXPECT() *_MockForwardBoolIteratorRecorder {
	return _m.recorder
}

func (_m *MockForwardBoolIterator) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockForwardBoolIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockForwardBoolIterator) Current() bool {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockForwardBoolIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockForwardBoolIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockForwardBoolIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockForwardBoolIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockForwardBoolIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

// Mock of RewindableTimeIterator interface
type MockRewindableTimeIterator struct {
	ctrl     *gomock.Controller
	recorder *_MockRewindableTimeIteratorRecorder
}

// Recorder for MockRewindableTimeIterator (not exported)
type _MockRewindableTimeIteratorRecorder struct {
	mock *MockRewindableTimeIterator
}

func NewMockRewindableTimeIterator(ctrl *gomock.Controller) *MockRewindableTimeIterator {
	mock := &MockRewindableTimeIterator{ctrl: ctrl}
	mock.recorder = &_MockRewindableTimeIteratorRecorder{mock}
	return mock
}

func (_m *MockRewindableTimeIterator) EXPECT() *_MockRewindableTimeIteratorRecorder {
	return _m.recorder
}

func (_m *MockRewindableTimeIterator) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRewindableTimeIteratorRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockRewindableTimeIterator) Current() int64 {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockRewindableTimeIteratorRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

func (_m *MockRewindableTimeIterator) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRewindableTimeIteratorRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockRewindableTimeIterator) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockRewindableTimeIteratorRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockRewindableTimeIterator) Reset(_param0 []int64) {
	_m.ctrl.Call(_m, "Reset", _param0)
}

func (_mr *_MockRewindableTimeIteratorRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

func (_m *MockRewindableTimeIterator) Rewind() {
	_m.ctrl.Call(_m, "Rewind")
}

func (_mr *_MockRewindableTimeIteratorRecorder) Rewind() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rewind")
}
