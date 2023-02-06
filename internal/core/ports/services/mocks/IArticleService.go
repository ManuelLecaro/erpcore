// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ManuelLecaro/erpcore/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// IArticleService is an autogenerated mock type for the IArticleService type
type IArticleService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, article
func (_m *IArticleService) Create(ctx context.Context, article domain.Article) (*domain.Article, error) {
	ret := _m.Called(ctx, article)

	var r0 *domain.Article
	if rf, ok := ret.Get(0).(func(context.Context, domain.Article) *domain.Article); ok {
		r0 = rf(ctx, article)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Article) error); ok {
		r1 = rf(ctx, article)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IArticleService) GetByID(ctx context.Context, id string) (*domain.Article, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Article
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Article); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Article)
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

// Search provides a mock function with given fields: ctx, fields
func (_m *IArticleService) Search(ctx context.Context, fields ...string) ([]*domain.Article, error) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*domain.Article
	if rf, ok := ret.Get(0).(func(context.Context, ...string) []*domain.Article); ok {
		r0 = rf(ctx, fields...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, fields...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIArticleService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIArticleService creates a new instance of IArticleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIArticleService(t mockConstructorTestingTNewIArticleService) *IArticleService {
	mock := &IArticleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}