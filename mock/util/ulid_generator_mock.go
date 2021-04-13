// Code generated by MockGen. DO NOT EDIT.
// Source: util/ulid_generator.go

// Package mock_util is a generated GoMock package.
package mock_util

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ulid "github.com/oklog/ulid"
)

// MockIULIDGenerator is a mock of IULIDGenerator interface.
type MockIULIDGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIULIDGeneratorMockRecorder
}

// MockIULIDGeneratorMockRecorder is the mock recorder for MockIULIDGenerator.
type MockIULIDGeneratorMockRecorder struct {
	mock *MockIULIDGenerator
}

// NewMockIULIDGenerator creates a new mock instance.
func NewMockIULIDGenerator(ctrl *gomock.Controller) *MockIULIDGenerator {
	mock := &MockIULIDGenerator{ctrl: ctrl}
	mock.recorder = &MockIULIDGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIULIDGenerator) EXPECT() *MockIULIDGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockIULIDGenerator) Generate() ulid.ULID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(ulid.ULID)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockIULIDGeneratorMockRecorder) Generate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockIULIDGenerator)(nil).Generate))
}
