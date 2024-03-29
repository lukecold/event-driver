// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lukecold/event-driver/handlers/cache (interfaces: ConflictResolver)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_cache.go -package=mocks github.com/lukecold/event-driver/handlers/cache ConflictResolver
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	event "github.com/lukecold/event-driver/event"
	handlers "github.com/lukecold/event-driver/handlers"
	gomock "go.uber.org/mock/gomock"
)

// MockConflictResolver is a mock of ConflictResolver interface.
type MockConflictResolver struct {
	ctrl     *gomock.Controller
	recorder *MockConflictResolverMockRecorder
}

// MockConflictResolverMockRecorder is the mock recorder for MockConflictResolver.
type MockConflictResolverMockRecorder struct {
	mock *MockConflictResolver
}

// NewMockConflictResolver creates a new mock instance.
func NewMockConflictResolver(ctrl *gomock.Controller) *MockConflictResolver {
	mock := &MockConflictResolver{ctrl: ctrl}
	mock.recorder = &MockConflictResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConflictResolver) EXPECT() *MockConflictResolverMockRecorder {
	return m.recorder
}

// Resolve mocks base method.
func (m *MockConflictResolver) Resolve(arg0 context.Context, arg1 *event.Message, arg2 handlers.CallNext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resolve indicates an expected call of Resolve.
func (mr *MockConflictResolverMockRecorder) Resolve(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockConflictResolver)(nil).Resolve), arg0, arg1, arg2)
}
