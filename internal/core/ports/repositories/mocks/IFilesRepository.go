// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ManuelLecaro/erpcore/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// IFilesRepository is an autogenerated mock type for the IFilesRepository type
type IFilesRepository struct {
	mock.Mock
}

// GetByUUID provides a mock function with given fields: ctx, uuid
func (_m *IFilesRepository) GetByUUID(ctx context.Context, uuid string) (string, error) {
	ret := _m.Called(ctx, uuid)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, file
func (_m *IFilesRepository) Save(ctx context.Context, file domain.User) error {
	ret := _m.Called(ctx, file)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) error); ok {
		r0 = rf(ctx, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIFilesRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIFilesRepository creates a new instance of IFilesRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIFilesRepository(t mockConstructorTestingTNewIFilesRepository) *IFilesRepository {
	mock := &IFilesRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
