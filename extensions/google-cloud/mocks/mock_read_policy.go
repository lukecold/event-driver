// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/honestbank/event-driver/extensions/google-cloud/storage/gcs_event_store (interfaces: ObjectIterator)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_read_policy.go -package=mocks github.com/honestbank/event-driver/extensions/google-cloud/storage/gcs_event_store ObjectIterator
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	storage "cloud.google.com/go/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockObjectIterator is a mock of ObjectIterator interface.
type MockObjectIterator struct {
	ctrl     *gomock.Controller
	recorder *MockObjectIteratorMockRecorder
}

// MockObjectIteratorMockRecorder is the mock recorder for MockObjectIterator.
type MockObjectIteratorMockRecorder struct {
	mock *MockObjectIterator
}

// NewMockObjectIterator creates a new mock instance.
func NewMockObjectIterator(ctrl *gomock.Controller) *MockObjectIterator {
	mock := &MockObjectIterator{ctrl: ctrl}
	mock.recorder = &MockObjectIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectIterator) EXPECT() *MockObjectIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method.
func (m *MockObjectIterator) Next() (*storage.ObjectAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*storage.ObjectAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockObjectIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockObjectIterator)(nil).Next))
}
