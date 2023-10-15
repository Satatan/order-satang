// Code generated by MockGen. DO NOT EDIT.
// Source: ./rank.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	models "order_satang/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRankUsecaseInterface is a mock of RankUsecaseInterface interface.
type MockRankUsecaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRankUsecaseInterfaceMockRecorder
}

// MockRankUsecaseInterfaceMockRecorder is the mock recorder for MockRankUsecaseInterface.
type MockRankUsecaseInterfaceMockRecorder struct {
	mock *MockRankUsecaseInterface
}

// NewMockRankUsecaseInterface creates a new mock instance.
func NewMockRankUsecaseInterface(ctrl *gomock.Controller) *MockRankUsecaseInterface {
	mock := &MockRankUsecaseInterface{ctrl: ctrl}
	mock.recorder = &MockRankUsecaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRankUsecaseInterface) EXPECT() *MockRankUsecaseInterfaceMockRecorder {
	return m.recorder
}

// GetUserRank mocks base method.
func (m *MockRankUsecaseInterface) GetUserRank() ([]models.UserRank, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRank")
	ret0, _ := ret[0].([]models.UserRank)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRank indicates an expected call of GetUserRank.
func (mr *MockRankUsecaseInterfaceMockRecorder) GetUserRank() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRank", reflect.TypeOf((*MockRankUsecaseInterface)(nil).GetUserRank))
}

// MockRankRepositoryInterface is a mock of RankRepositoryInterface interface.
type MockRankRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRankRepositoryInterfaceMockRecorder
}

// MockRankRepositoryInterfaceMockRecorder is the mock recorder for MockRankRepositoryInterface.
type MockRankRepositoryInterfaceMockRecorder struct {
	mock *MockRankRepositoryInterface
}

// NewMockRankRepositoryInterface creates a new mock instance.
func NewMockRankRepositoryInterface(ctrl *gomock.Controller) *MockRankRepositoryInterface {
	mock := &MockRankRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRankRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRankRepositoryInterface) EXPECT() *MockRankRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetUserRank mocks base method.
func (m *MockRankRepositoryInterface) GetUserRank() ([]models.UserRank, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRank")
	ret0, _ := ret[0].([]models.UserRank)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRank indicates an expected call of GetUserRank.
func (mr *MockRankRepositoryInterfaceMockRecorder) GetUserRank() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRank", reflect.TypeOf((*MockRankRepositoryInterface)(nil).GetUserRank))
}