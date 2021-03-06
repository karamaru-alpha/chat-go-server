// Code generated by MockGen. DO NOT EDIT.
// Source: domain/model/room/repository.go

// Package mock_room is a generated GoMock package.
package mock_room

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	room "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockIRepository) Find(arg0 room.ID) (room.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(room.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIRepository)(nil).Find), arg0)
}

// FindAll mocks base method.
func (m *MockIRepository) FindAll() ([]room.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]room.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockIRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIRepository)(nil).FindAll))
}

// FindByTitle mocks base method.
func (m *MockIRepository) FindByTitle(arg0 room.Title) (room.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTitle", arg0)
	ret0, _ := ret[0].(room.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTitle indicates an expected call of FindByTitle.
func (mr *MockIRepositoryMockRecorder) FindByTitle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTitle", reflect.TypeOf((*MockIRepository)(nil).FindByTitle), arg0)
}

// Save mocks base method.
func (m *MockIRepository) Save(arg0 room.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIRepository)(nil).Save), arg0)
}
