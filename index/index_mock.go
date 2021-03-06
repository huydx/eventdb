// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xichen2020/eventdb/index (interfaces: DocIDSet,DocIDSetIterator,SeekableDocIDSetIterator,DocIDPositionIterator)

// Package index is a generated GoMock package.
package index

import (
	"bytes"
	"io"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockDocIDSet is a mock of DocIDSet interface
type MockDocIDSet struct {
	ctrl     *gomock.Controller
	recorder *MockDocIDSetMockRecorder
}

// MockDocIDSetMockRecorder is the mock recorder for MockDocIDSet
type MockDocIDSetMockRecorder struct {
	mock *MockDocIDSet
}

// NewMockDocIDSet creates a new mock instance
func NewMockDocIDSet(ctrl *gomock.Controller) *MockDocIDSet {
	mock := &MockDocIDSet{ctrl: ctrl}
	mock.recorder = &MockDocIDSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDocIDSet) EXPECT() *MockDocIDSetMockRecorder {
	return m.recorder
}

// Intersect mocks base method
func (m *MockDocIDSet) Intersect(arg0 DocIDSetIterator) DocIDPositionIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersect", arg0)
	ret0, _ := ret[0].(DocIDPositionIterator)
	return ret0
}

// Intersect indicates an expected call of Intersect
func (mr *MockDocIDSetMockRecorder) Intersect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersect", reflect.TypeOf((*MockDocIDSet)(nil).Intersect), arg0)
}

// Iter mocks base method
func (m *MockDocIDSet) Iter() DocIDSetIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(DocIDSetIterator)
	return ret0
}

// Iter indicates an expected call of Iter
func (mr *MockDocIDSetMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockDocIDSet)(nil).Iter))
}

// WriteTo mocks base method
func (m *MockDocIDSet) WriteTo(arg0 io.Writer, arg1 *bytes.Buffer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTo indicates an expected call of WriteTo
func (mr *MockDocIDSetMockRecorder) WriteTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockDocIDSet)(nil).WriteTo), arg0, arg1)
}

// MockDocIDSetIterator is a mock of DocIDSetIterator interface
type MockDocIDSetIterator struct {
	ctrl     *gomock.Controller
	recorder *MockDocIDSetIteratorMockRecorder
}

// MockDocIDSetIteratorMockRecorder is the mock recorder for MockDocIDSetIterator
type MockDocIDSetIteratorMockRecorder struct {
	mock *MockDocIDSetIterator
}

// NewMockDocIDSetIterator creates a new mock instance
func NewMockDocIDSetIterator(ctrl *gomock.Controller) *MockDocIDSetIterator {
	mock := &MockDocIDSetIterator{ctrl: ctrl}
	mock.recorder = &MockDocIDSetIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDocIDSetIterator) EXPECT() *MockDocIDSetIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockDocIDSetIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDocIDSetIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDocIDSetIterator)(nil).Close))
}

// DocID mocks base method
func (m *MockDocIDSetIterator) DocID() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID")
	ret0, _ := ret[0].(int32)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockDocIDSetIteratorMockRecorder) DocID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockDocIDSetIterator)(nil).DocID))
}

// Err mocks base method
func (m *MockDocIDSetIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockDocIDSetIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockDocIDSetIterator)(nil).Err))
}

// Next mocks base method
func (m *MockDocIDSetIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockDocIDSetIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockDocIDSetIterator)(nil).Next))
}

// MockSeekableDocIDSetIterator is a mock of SeekableDocIDSetIterator interface
type MockSeekableDocIDSetIterator struct {
	ctrl     *gomock.Controller
	recorder *MockSeekableDocIDSetIteratorMockRecorder
}

// MockSeekableDocIDSetIteratorMockRecorder is the mock recorder for MockSeekableDocIDSetIterator
type MockSeekableDocIDSetIteratorMockRecorder struct {
	mock *MockSeekableDocIDSetIterator
}

// NewMockSeekableDocIDSetIterator creates a new mock instance
func NewMockSeekableDocIDSetIterator(ctrl *gomock.Controller) *MockSeekableDocIDSetIterator {
	mock := &MockSeekableDocIDSetIterator{ctrl: ctrl}
	mock.recorder = &MockSeekableDocIDSetIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSeekableDocIDSetIterator) EXPECT() *MockSeekableDocIDSetIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSeekableDocIDSetIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSeekableDocIDSetIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSeekableDocIDSetIterator)(nil).Close))
}

// DocID mocks base method
func (m *MockSeekableDocIDSetIterator) DocID() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID")
	ret0, _ := ret[0].(int32)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockSeekableDocIDSetIteratorMockRecorder) DocID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockSeekableDocIDSetIterator)(nil).DocID))
}

// Err mocks base method
func (m *MockSeekableDocIDSetIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockSeekableDocIDSetIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSeekableDocIDSetIterator)(nil).Err))
}

// Next mocks base method
func (m *MockSeekableDocIDSetIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockSeekableDocIDSetIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockSeekableDocIDSetIterator)(nil).Next))
}

// SeekForward mocks base method
func (m *MockSeekableDocIDSetIterator) SeekForward(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeekForward", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SeekForward indicates an expected call of SeekForward
func (mr *MockSeekableDocIDSetIteratorMockRecorder) SeekForward(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeekForward", reflect.TypeOf((*MockSeekableDocIDSetIterator)(nil).SeekForward), arg0)
}

// MockDocIDPositionIterator is a mock of DocIDPositionIterator interface
type MockDocIDPositionIterator struct {
	ctrl     *gomock.Controller
	recorder *MockDocIDPositionIteratorMockRecorder
}

// MockDocIDPositionIteratorMockRecorder is the mock recorder for MockDocIDPositionIterator
type MockDocIDPositionIteratorMockRecorder struct {
	mock *MockDocIDPositionIterator
}

// NewMockDocIDPositionIterator creates a new mock instance
func NewMockDocIDPositionIterator(ctrl *gomock.Controller) *MockDocIDPositionIterator {
	mock := &MockDocIDPositionIterator{ctrl: ctrl}
	mock.recorder = &MockDocIDPositionIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDocIDPositionIterator) EXPECT() *MockDocIDPositionIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockDocIDPositionIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDocIDPositionIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDocIDPositionIterator)(nil).Close))
}

// DocID mocks base method
func (m *MockDocIDPositionIterator) DocID() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID")
	ret0, _ := ret[0].(int32)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockDocIDPositionIteratorMockRecorder) DocID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockDocIDPositionIterator)(nil).DocID))
}

// Err mocks base method
func (m *MockDocIDPositionIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockDocIDPositionIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockDocIDPositionIterator)(nil).Err))
}

// MaskingPosition mocks base method
func (m *MockDocIDPositionIterator) MaskingPosition() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaskingPosition")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaskingPosition indicates an expected call of MaskingPosition
func (mr *MockDocIDPositionIteratorMockRecorder) MaskingPosition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaskingPosition", reflect.TypeOf((*MockDocIDPositionIterator)(nil).MaskingPosition))
}

// Next mocks base method
func (m *MockDocIDPositionIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockDocIDPositionIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockDocIDPositionIterator)(nil).Next))
}

// Position mocks base method
func (m *MockDocIDPositionIterator) Position() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Position")
	ret0, _ := ret[0].(int)
	return ret0
}

// Position indicates an expected call of Position
func (mr *MockDocIDPositionIteratorMockRecorder) Position() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Position", reflect.TypeOf((*MockDocIDPositionIterator)(nil).Position))
}
