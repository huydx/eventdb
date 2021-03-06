// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xichen2020/eventdb/index/field (interfaces: BaseFieldIterator,CloseableNullField,CloseableBoolField,CloseableIntField,CloseableDoubleField,CloseableBytesField,CloseableTimeField,DocsField)

// Package field is a generated GoMock package.
package field

import (
	gomock "github.com/golang/mock/gomock"
	field "github.com/xichen2020/eventdb/document/field"
	filter "github.com/xichen2020/eventdb/filter"
	index "github.com/xichen2020/eventdb/index"
	values "github.com/xichen2020/eventdb/values"
	reflect "reflect"
)

// MockBaseFieldIterator is a mock of BaseFieldIterator interface
type MockBaseFieldIterator struct {
	ctrl     *gomock.Controller
	recorder *MockBaseFieldIteratorMockRecorder
}

// MockBaseFieldIteratorMockRecorder is the mock recorder for MockBaseFieldIterator
type MockBaseFieldIteratorMockRecorder struct {
	mock *MockBaseFieldIterator
}

// NewMockBaseFieldIterator creates a new mock instance
func NewMockBaseFieldIterator(ctrl *gomock.Controller) *MockBaseFieldIterator {
	mock := &MockBaseFieldIterator{ctrl: ctrl}
	mock.recorder = &MockBaseFieldIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBaseFieldIterator) EXPECT() *MockBaseFieldIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockBaseFieldIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockBaseFieldIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBaseFieldIterator)(nil).Close))
}

// DocID mocks base method
func (m *MockBaseFieldIterator) DocID() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID")
	ret0, _ := ret[0].(int32)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockBaseFieldIteratorMockRecorder) DocID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockBaseFieldIterator)(nil).DocID))
}

// Err mocks base method
func (m *MockBaseFieldIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockBaseFieldIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockBaseFieldIterator)(nil).Err))
}

// Next mocks base method
func (m *MockBaseFieldIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockBaseFieldIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockBaseFieldIterator)(nil).Next))
}

// ValueUnion mocks base method
func (m *MockBaseFieldIterator) ValueUnion() field.ValueUnion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueUnion")
	ret0, _ := ret[0].(field.ValueUnion)
	return ret0
}

// ValueUnion indicates an expected call of ValueUnion
func (mr *MockBaseFieldIteratorMockRecorder) ValueUnion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueUnion", reflect.TypeOf((*MockBaseFieldIterator)(nil).ValueUnion))
}

// MockCloseableNullField is a mock of CloseableNullField interface
type MockCloseableNullField struct {
	ctrl     *gomock.Controller
	recorder *MockCloseableNullFieldMockRecorder
}

// MockCloseableNullFieldMockRecorder is the mock recorder for MockCloseableNullField
type MockCloseableNullFieldMockRecorder struct {
	mock *MockCloseableNullField
}

// NewMockCloseableNullField creates a new mock instance
func NewMockCloseableNullField(ctrl *gomock.Controller) *MockCloseableNullField {
	mock := &MockCloseableNullField{ctrl: ctrl}
	mock.recorder = &MockCloseableNullFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloseableNullField) EXPECT() *MockCloseableNullFieldMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCloseableNullField) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockCloseableNullFieldMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloseableNullField)(nil).Close))
}

// DocIDSet mocks base method
func (m *MockCloseableNullField) DocIDSet() index.DocIDSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocIDSet")
	ret0, _ := ret[0].(index.DocIDSet)
	return ret0
}

// DocIDSet indicates an expected call of DocIDSet
func (mr *MockCloseableNullFieldMockRecorder) DocIDSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocIDSet", reflect.TypeOf((*MockCloseableNullField)(nil).DocIDSet))
}

// Fetch mocks base method
func (m *MockCloseableNullField) Fetch(arg0 index.DocIDSetIterator) MaskingNullFieldIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(MaskingNullFieldIterator)
	return ret0
}

// Fetch indicates an expected call of Fetch
func (mr *MockCloseableNullFieldMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockCloseableNullField)(nil).Fetch), arg0)
}

// Filter mocks base method
func (m *MockCloseableNullField) Filter(arg0 filter.Op, arg1 *field.ValueUnion, arg2 int32) (index.DocIDSetIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1, arg2)
	ret0, _ := ret[0].(index.DocIDSetIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *MockCloseableNullFieldMockRecorder) Filter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockCloseableNullField)(nil).Filter), arg0, arg1, arg2)
}

// Iter mocks base method
func (m *MockCloseableNullField) Iter() NullFieldIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(NullFieldIterator)
	return ret0
}

// Iter indicates an expected call of Iter
func (mr *MockCloseableNullFieldMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockCloseableNullField)(nil).Iter))
}

// ShallowCopy mocks base method
func (m *MockCloseableNullField) ShallowCopy() CloseableNullField {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShallowCopy")
	ret0, _ := ret[0].(CloseableNullField)
	return ret0
}

// ShallowCopy indicates an expected call of ShallowCopy
func (mr *MockCloseableNullFieldMockRecorder) ShallowCopy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShallowCopy", reflect.TypeOf((*MockCloseableNullField)(nil).ShallowCopy))
}

// MockCloseableBoolField is a mock of CloseableBoolField interface
type MockCloseableBoolField struct {
	ctrl     *gomock.Controller
	recorder *MockCloseableBoolFieldMockRecorder
}

// MockCloseableBoolFieldMockRecorder is the mock recorder for MockCloseableBoolField
type MockCloseableBoolFieldMockRecorder struct {
	mock *MockCloseableBoolField
}

// NewMockCloseableBoolField creates a new mock instance
func NewMockCloseableBoolField(ctrl *gomock.Controller) *MockCloseableBoolField {
	mock := &MockCloseableBoolField{ctrl: ctrl}
	mock.recorder = &MockCloseableBoolFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloseableBoolField) EXPECT() *MockCloseableBoolFieldMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCloseableBoolField) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockCloseableBoolFieldMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloseableBoolField)(nil).Close))
}

// DocIDSet mocks base method
func (m *MockCloseableBoolField) DocIDSet() index.DocIDSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocIDSet")
	ret0, _ := ret[0].(index.DocIDSet)
	return ret0
}

// DocIDSet indicates an expected call of DocIDSet
func (mr *MockCloseableBoolFieldMockRecorder) DocIDSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocIDSet", reflect.TypeOf((*MockCloseableBoolField)(nil).DocIDSet))
}

// Fetch mocks base method
func (m *MockCloseableBoolField) Fetch(arg0 index.DocIDSetIterator) (MaskingBoolFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(MaskingBoolFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockCloseableBoolFieldMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockCloseableBoolField)(nil).Fetch), arg0)
}

// Filter mocks base method
func (m *MockCloseableBoolField) Filter(arg0 filter.Op, arg1 *field.ValueUnion, arg2 int32) (index.DocIDSetIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1, arg2)
	ret0, _ := ret[0].(index.DocIDSetIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *MockCloseableBoolFieldMockRecorder) Filter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockCloseableBoolField)(nil).Filter), arg0, arg1, arg2)
}

// Iter mocks base method
func (m *MockCloseableBoolField) Iter() (BoolFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(BoolFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iter indicates an expected call of Iter
func (mr *MockCloseableBoolFieldMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockCloseableBoolField)(nil).Iter))
}

// ShallowCopy mocks base method
func (m *MockCloseableBoolField) ShallowCopy() CloseableBoolField {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShallowCopy")
	ret0, _ := ret[0].(CloseableBoolField)
	return ret0
}

// ShallowCopy indicates an expected call of ShallowCopy
func (mr *MockCloseableBoolFieldMockRecorder) ShallowCopy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShallowCopy", reflect.TypeOf((*MockCloseableBoolField)(nil).ShallowCopy))
}

// Values mocks base method
func (m *MockCloseableBoolField) Values() values.BoolValues {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].(values.BoolValues)
	return ret0
}

// Values indicates an expected call of Values
func (mr *MockCloseableBoolFieldMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockCloseableBoolField)(nil).Values))
}

// MockCloseableIntField is a mock of CloseableIntField interface
type MockCloseableIntField struct {
	ctrl     *gomock.Controller
	recorder *MockCloseableIntFieldMockRecorder
}

// MockCloseableIntFieldMockRecorder is the mock recorder for MockCloseableIntField
type MockCloseableIntFieldMockRecorder struct {
	mock *MockCloseableIntField
}

// NewMockCloseableIntField creates a new mock instance
func NewMockCloseableIntField(ctrl *gomock.Controller) *MockCloseableIntField {
	mock := &MockCloseableIntField{ctrl: ctrl}
	mock.recorder = &MockCloseableIntFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloseableIntField) EXPECT() *MockCloseableIntFieldMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCloseableIntField) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockCloseableIntFieldMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloseableIntField)(nil).Close))
}

// DocIDSet mocks base method
func (m *MockCloseableIntField) DocIDSet() index.DocIDSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocIDSet")
	ret0, _ := ret[0].(index.DocIDSet)
	return ret0
}

// DocIDSet indicates an expected call of DocIDSet
func (mr *MockCloseableIntFieldMockRecorder) DocIDSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocIDSet", reflect.TypeOf((*MockCloseableIntField)(nil).DocIDSet))
}

// Fetch mocks base method
func (m *MockCloseableIntField) Fetch(arg0 index.DocIDSetIterator) (MaskingIntFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(MaskingIntFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockCloseableIntFieldMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockCloseableIntField)(nil).Fetch), arg0)
}

// Filter mocks base method
func (m *MockCloseableIntField) Filter(arg0 filter.Op, arg1 *field.ValueUnion, arg2 int32) (index.DocIDSetIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1, arg2)
	ret0, _ := ret[0].(index.DocIDSetIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *MockCloseableIntFieldMockRecorder) Filter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockCloseableIntField)(nil).Filter), arg0, arg1, arg2)
}

// Iter mocks base method
func (m *MockCloseableIntField) Iter() (IntFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(IntFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iter indicates an expected call of Iter
func (mr *MockCloseableIntFieldMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockCloseableIntField)(nil).Iter))
}

// ShallowCopy mocks base method
func (m *MockCloseableIntField) ShallowCopy() CloseableIntField {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShallowCopy")
	ret0, _ := ret[0].(CloseableIntField)
	return ret0
}

// ShallowCopy indicates an expected call of ShallowCopy
func (mr *MockCloseableIntFieldMockRecorder) ShallowCopy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShallowCopy", reflect.TypeOf((*MockCloseableIntField)(nil).ShallowCopy))
}

// Values mocks base method
func (m *MockCloseableIntField) Values() values.IntValues {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].(values.IntValues)
	return ret0
}

// Values indicates an expected call of Values
func (mr *MockCloseableIntFieldMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockCloseableIntField)(nil).Values))
}

// MockCloseableDoubleField is a mock of CloseableDoubleField interface
type MockCloseableDoubleField struct {
	ctrl     *gomock.Controller
	recorder *MockCloseableDoubleFieldMockRecorder
}

// MockCloseableDoubleFieldMockRecorder is the mock recorder for MockCloseableDoubleField
type MockCloseableDoubleFieldMockRecorder struct {
	mock *MockCloseableDoubleField
}

// NewMockCloseableDoubleField creates a new mock instance
func NewMockCloseableDoubleField(ctrl *gomock.Controller) *MockCloseableDoubleField {
	mock := &MockCloseableDoubleField{ctrl: ctrl}
	mock.recorder = &MockCloseableDoubleFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloseableDoubleField) EXPECT() *MockCloseableDoubleFieldMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCloseableDoubleField) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockCloseableDoubleFieldMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloseableDoubleField)(nil).Close))
}

// DocIDSet mocks base method
func (m *MockCloseableDoubleField) DocIDSet() index.DocIDSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocIDSet")
	ret0, _ := ret[0].(index.DocIDSet)
	return ret0
}

// DocIDSet indicates an expected call of DocIDSet
func (mr *MockCloseableDoubleFieldMockRecorder) DocIDSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocIDSet", reflect.TypeOf((*MockCloseableDoubleField)(nil).DocIDSet))
}

// Fetch mocks base method
func (m *MockCloseableDoubleField) Fetch(arg0 index.DocIDSetIterator) (MaskingDoubleFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(MaskingDoubleFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockCloseableDoubleFieldMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockCloseableDoubleField)(nil).Fetch), arg0)
}

// Filter mocks base method
func (m *MockCloseableDoubleField) Filter(arg0 filter.Op, arg1 *field.ValueUnion, arg2 int32) (index.DocIDSetIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1, arg2)
	ret0, _ := ret[0].(index.DocIDSetIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *MockCloseableDoubleFieldMockRecorder) Filter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockCloseableDoubleField)(nil).Filter), arg0, arg1, arg2)
}

// Iter mocks base method
func (m *MockCloseableDoubleField) Iter() (DoubleFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(DoubleFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iter indicates an expected call of Iter
func (mr *MockCloseableDoubleFieldMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockCloseableDoubleField)(nil).Iter))
}

// ShallowCopy mocks base method
func (m *MockCloseableDoubleField) ShallowCopy() CloseableDoubleField {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShallowCopy")
	ret0, _ := ret[0].(CloseableDoubleField)
	return ret0
}

// ShallowCopy indicates an expected call of ShallowCopy
func (mr *MockCloseableDoubleFieldMockRecorder) ShallowCopy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShallowCopy", reflect.TypeOf((*MockCloseableDoubleField)(nil).ShallowCopy))
}

// Values mocks base method
func (m *MockCloseableDoubleField) Values() values.DoubleValues {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].(values.DoubleValues)
	return ret0
}

// Values indicates an expected call of Values
func (mr *MockCloseableDoubleFieldMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockCloseableDoubleField)(nil).Values))
}

// MockCloseableBytesField is a mock of CloseableBytesField interface
type MockCloseableBytesField struct {
	ctrl     *gomock.Controller
	recorder *MockCloseableBytesFieldMockRecorder
}

// MockCloseableBytesFieldMockRecorder is the mock recorder for MockCloseableBytesField
type MockCloseableBytesFieldMockRecorder struct {
	mock *MockCloseableBytesField
}

// NewMockCloseableBytesField creates a new mock instance
func NewMockCloseableBytesField(ctrl *gomock.Controller) *MockCloseableBytesField {
	mock := &MockCloseableBytesField{ctrl: ctrl}
	mock.recorder = &MockCloseableBytesFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloseableBytesField) EXPECT() *MockCloseableBytesFieldMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCloseableBytesField) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockCloseableBytesFieldMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloseableBytesField)(nil).Close))
}

// DocIDSet mocks base method
func (m *MockCloseableBytesField) DocIDSet() index.DocIDSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocIDSet")
	ret0, _ := ret[0].(index.DocIDSet)
	return ret0
}

// DocIDSet indicates an expected call of DocIDSet
func (mr *MockCloseableBytesFieldMockRecorder) DocIDSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocIDSet", reflect.TypeOf((*MockCloseableBytesField)(nil).DocIDSet))
}

// Fetch mocks base method
func (m *MockCloseableBytesField) Fetch(arg0 index.DocIDSetIterator) (MaskingBytesFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(MaskingBytesFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockCloseableBytesFieldMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockCloseableBytesField)(nil).Fetch), arg0)
}

// Filter mocks base method
func (m *MockCloseableBytesField) Filter(arg0 filter.Op, arg1 *field.ValueUnion, arg2 int32) (index.DocIDSetIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1, arg2)
	ret0, _ := ret[0].(index.DocIDSetIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *MockCloseableBytesFieldMockRecorder) Filter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockCloseableBytesField)(nil).Filter), arg0, arg1, arg2)
}

// Iter mocks base method
func (m *MockCloseableBytesField) Iter() (BytesFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(BytesFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iter indicates an expected call of Iter
func (mr *MockCloseableBytesFieldMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockCloseableBytesField)(nil).Iter))
}

// ShallowCopy mocks base method
func (m *MockCloseableBytesField) ShallowCopy() CloseableBytesField {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShallowCopy")
	ret0, _ := ret[0].(CloseableBytesField)
	return ret0
}

// ShallowCopy indicates an expected call of ShallowCopy
func (mr *MockCloseableBytesFieldMockRecorder) ShallowCopy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShallowCopy", reflect.TypeOf((*MockCloseableBytesField)(nil).ShallowCopy))
}

// Values mocks base method
func (m *MockCloseableBytesField) Values() values.BytesValues {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].(values.BytesValues)
	return ret0
}

// Values indicates an expected call of Values
func (mr *MockCloseableBytesFieldMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockCloseableBytesField)(nil).Values))
}

// MockCloseableTimeField is a mock of CloseableTimeField interface
type MockCloseableTimeField struct {
	ctrl     *gomock.Controller
	recorder *MockCloseableTimeFieldMockRecorder
}

// MockCloseableTimeFieldMockRecorder is the mock recorder for MockCloseableTimeField
type MockCloseableTimeFieldMockRecorder struct {
	mock *MockCloseableTimeField
}

// NewMockCloseableTimeField creates a new mock instance
func NewMockCloseableTimeField(ctrl *gomock.Controller) *MockCloseableTimeField {
	mock := &MockCloseableTimeField{ctrl: ctrl}
	mock.recorder = &MockCloseableTimeFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloseableTimeField) EXPECT() *MockCloseableTimeFieldMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCloseableTimeField) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockCloseableTimeFieldMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloseableTimeField)(nil).Close))
}

// DocIDSet mocks base method
func (m *MockCloseableTimeField) DocIDSet() index.DocIDSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocIDSet")
	ret0, _ := ret[0].(index.DocIDSet)
	return ret0
}

// DocIDSet indicates an expected call of DocIDSet
func (mr *MockCloseableTimeFieldMockRecorder) DocIDSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocIDSet", reflect.TypeOf((*MockCloseableTimeField)(nil).DocIDSet))
}

// Fetch mocks base method
func (m *MockCloseableTimeField) Fetch(arg0 index.DocIDSetIterator) (MaskingTimeFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(MaskingTimeFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockCloseableTimeFieldMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockCloseableTimeField)(nil).Fetch), arg0)
}

// Filter mocks base method
func (m *MockCloseableTimeField) Filter(arg0 filter.Op, arg1 *field.ValueUnion, arg2 int32) (index.DocIDSetIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1, arg2)
	ret0, _ := ret[0].(index.DocIDSetIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *MockCloseableTimeFieldMockRecorder) Filter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockCloseableTimeField)(nil).Filter), arg0, arg1, arg2)
}

// Iter mocks base method
func (m *MockCloseableTimeField) Iter() (TimeFieldIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(TimeFieldIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iter indicates an expected call of Iter
func (mr *MockCloseableTimeFieldMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockCloseableTimeField)(nil).Iter))
}

// ShallowCopy mocks base method
func (m *MockCloseableTimeField) ShallowCopy() CloseableTimeField {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShallowCopy")
	ret0, _ := ret[0].(CloseableTimeField)
	return ret0
}

// ShallowCopy indicates an expected call of ShallowCopy
func (mr *MockCloseableTimeFieldMockRecorder) ShallowCopy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShallowCopy", reflect.TypeOf((*MockCloseableTimeField)(nil).ShallowCopy))
}

// Values mocks base method
func (m *MockCloseableTimeField) Values() values.TimeValues {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].(values.TimeValues)
	return ret0
}

// Values indicates an expected call of Values
func (mr *MockCloseableTimeFieldMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockCloseableTimeField)(nil).Values))
}

// MockDocsField is a mock of DocsField interface
type MockDocsField struct {
	ctrl     *gomock.Controller
	recorder *MockDocsFieldMockRecorder
}

// MockDocsFieldMockRecorder is the mock recorder for MockDocsField
type MockDocsFieldMockRecorder struct {
	mock *MockDocsField
}

// NewMockDocsField creates a new mock instance
func NewMockDocsField(ctrl *gomock.Controller) *MockDocsField {
	mock := &MockDocsField{ctrl: ctrl}
	mock.recorder = &MockDocsFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDocsField) EXPECT() *MockDocsFieldMockRecorder {
	return m.recorder
}

// BoolField mocks base method
func (m *MockDocsField) BoolField() (BoolField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BoolField")
	ret0, _ := ret[0].(BoolField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// BoolField indicates an expected call of BoolField
func (mr *MockDocsFieldMockRecorder) BoolField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BoolField", reflect.TypeOf((*MockDocsField)(nil).BoolField))
}

// BytesField mocks base method
func (m *MockDocsField) BytesField() (BytesField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BytesField")
	ret0, _ := ret[0].(BytesField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// BytesField indicates an expected call of BytesField
func (mr *MockDocsFieldMockRecorder) BytesField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BytesField", reflect.TypeOf((*MockDocsField)(nil).BytesField))
}

// Close mocks base method
func (m *MockDocsField) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDocsFieldMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDocsField)(nil).Close))
}

// DoubleField mocks base method
func (m *MockDocsField) DoubleField() (DoubleField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoubleField")
	ret0, _ := ret[0].(DoubleField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// DoubleField indicates an expected call of DoubleField
func (mr *MockDocsFieldMockRecorder) DoubleField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoubleField", reflect.TypeOf((*MockDocsField)(nil).DoubleField))
}

// FieldForType mocks base method
func (m *MockDocsField) FieldForType(arg0 field.ValueType) (Union, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldForType", arg0)
	ret0, _ := ret[0].(Union)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FieldForType indicates an expected call of FieldForType
func (mr *MockDocsFieldMockRecorder) FieldForType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldForType", reflect.TypeOf((*MockDocsField)(nil).FieldForType), arg0)
}

// Filter mocks base method
func (m *MockDocsField) Filter(arg0 filter.Op, arg1 *field.ValueUnion, arg2 int32) (index.DocIDSetIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1, arg2)
	ret0, _ := ret[0].(index.DocIDSetIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *MockDocsFieldMockRecorder) Filter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockDocsField)(nil).Filter), arg0, arg1, arg2)
}

// IntField mocks base method
func (m *MockDocsField) IntField() (IntField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntField")
	ret0, _ := ret[0].(IntField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IntField indicates an expected call of IntField
func (mr *MockDocsFieldMockRecorder) IntField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntField", reflect.TypeOf((*MockDocsField)(nil).IntField))
}

// Metadata mocks base method
func (m *MockDocsField) Metadata() DocsFieldMetadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata")
	ret0, _ := ret[0].(DocsFieldMetadata)
	return ret0
}

// Metadata indicates an expected call of Metadata
func (mr *MockDocsFieldMockRecorder) Metadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockDocsField)(nil).Metadata))
}

// NewDocsFieldFor mocks base method
func (m *MockDocsField) NewDocsFieldFor(arg0 field.ValueTypeSet) (DocsField, field.ValueTypeSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDocsFieldFor", arg0)
	ret0, _ := ret[0].(DocsField)
	ret1, _ := ret[1].(field.ValueTypeSet)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewDocsFieldFor indicates an expected call of NewDocsFieldFor
func (mr *MockDocsFieldMockRecorder) NewDocsFieldFor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDocsFieldFor", reflect.TypeOf((*MockDocsField)(nil).NewDocsFieldFor), arg0)
}

// NewMergedDocsField mocks base method
func (m *MockDocsField) NewMergedDocsField(arg0 DocsField) DocsField {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMergedDocsField", arg0)
	ret0, _ := ret[0].(DocsField)
	return ret0
}

// NewMergedDocsField indicates an expected call of NewMergedDocsField
func (mr *MockDocsFieldMockRecorder) NewMergedDocsField(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMergedDocsField", reflect.TypeOf((*MockDocsField)(nil).NewMergedDocsField), arg0)
}

// NullField mocks base method
func (m *MockDocsField) NullField() (NullField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NullField")
	ret0, _ := ret[0].(NullField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// NullField indicates an expected call of NullField
func (mr *MockDocsFieldMockRecorder) NullField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NullField", reflect.TypeOf((*MockDocsField)(nil).NullField))
}

// ShallowCopy mocks base method
func (m *MockDocsField) ShallowCopy() DocsField {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShallowCopy")
	ret0, _ := ret[0].(DocsField)
	return ret0
}

// ShallowCopy indicates an expected call of ShallowCopy
func (mr *MockDocsFieldMockRecorder) ShallowCopy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShallowCopy", reflect.TypeOf((*MockDocsField)(nil).ShallowCopy))
}

// TimeField mocks base method
func (m *MockDocsField) TimeField() (TimeField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeField")
	ret0, _ := ret[0].(TimeField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// TimeField indicates an expected call of TimeField
func (mr *MockDocsFieldMockRecorder) TimeField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeField", reflect.TypeOf((*MockDocsField)(nil).TimeField))
}

// closeableBoolField mocks base method
func (m *MockDocsField) closeableBoolField() (CloseableBoolField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "closeableBoolField")
	ret0, _ := ret[0].(CloseableBoolField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// closeableBoolField indicates an expected call of closeableBoolField
func (mr *MockDocsFieldMockRecorder) closeableBoolField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeableBoolField", reflect.TypeOf((*MockDocsField)(nil).closeableBoolField))
}

// closeableBytesField mocks base method
func (m *MockDocsField) closeableBytesField() (CloseableBytesField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "closeableBytesField")
	ret0, _ := ret[0].(CloseableBytesField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// closeableBytesField indicates an expected call of closeableBytesField
func (mr *MockDocsFieldMockRecorder) closeableBytesField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeableBytesField", reflect.TypeOf((*MockDocsField)(nil).closeableBytesField))
}

// closeableDoubleField mocks base method
func (m *MockDocsField) closeableDoubleField() (CloseableDoubleField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "closeableDoubleField")
	ret0, _ := ret[0].(CloseableDoubleField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// closeableDoubleField indicates an expected call of closeableDoubleField
func (mr *MockDocsFieldMockRecorder) closeableDoubleField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeableDoubleField", reflect.TypeOf((*MockDocsField)(nil).closeableDoubleField))
}

// closeableIntField mocks base method
func (m *MockDocsField) closeableIntField() (CloseableIntField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "closeableIntField")
	ret0, _ := ret[0].(CloseableIntField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// closeableIntField indicates an expected call of closeableIntField
func (mr *MockDocsFieldMockRecorder) closeableIntField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeableIntField", reflect.TypeOf((*MockDocsField)(nil).closeableIntField))
}

// closeableNullField mocks base method
func (m *MockDocsField) closeableNullField() (CloseableNullField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "closeableNullField")
	ret0, _ := ret[0].(CloseableNullField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// closeableNullField indicates an expected call of closeableNullField
func (mr *MockDocsFieldMockRecorder) closeableNullField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeableNullField", reflect.TypeOf((*MockDocsField)(nil).closeableNullField))
}

// closeableTimeField mocks base method
func (m *MockDocsField) closeableTimeField() (CloseableTimeField, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "closeableTimeField")
	ret0, _ := ret[0].(CloseableTimeField)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// closeableTimeField indicates an expected call of closeableTimeField
func (mr *MockDocsFieldMockRecorder) closeableTimeField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeableTimeField", reflect.TypeOf((*MockDocsField)(nil).closeableTimeField))
}
