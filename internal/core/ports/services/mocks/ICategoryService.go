// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ManuelLecaro/erpcore/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// ICategoryService is an autogenerated mock type for the ICategoryService type
type ICategoryService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, article
func (_m *ICategoryService) Create(ctx context.Context, article domain.Category) (*domain.Category, error) {
	ret := _m.Called(ctx, article)

	var r0 *domain.Category
	if rf, ok := ret.Get(0).(func(context.Context, domain.Category) *domain.Category); ok {
		r0 = rf(ctx, article)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Category) error); ok {
		r1 = rf(ctx, article)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *ICategoryService) GetAll(ctx context.Context) ([]*domain.Category, error) {
	ret := _m.Called(ctx)

	var r0 []*domain.Category
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Category); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ICategoryService) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Category
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Category); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewICategoryService interface {
	mock.TestingT
	Cleanup(func())
}

// NewICategoryService creates a new instance of ICategoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICategoryService(t mockConstructorTestingTNewICategoryService) *ICategoryService {
	mock := &ICategoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
