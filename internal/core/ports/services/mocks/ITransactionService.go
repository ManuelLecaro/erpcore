// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ManuelLecaro/erpcore/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// ITransactionService is an autogenerated mock type for the ITransactionService type
type ITransactionService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, category
func (_m *ITransactionService) Create(ctx context.Context, category domain.Transaction) (*domain.Transaction, error) {
	ret := _m.Called(ctx, category)

	var r0 *domain.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, domain.Transaction) *domain.Transaction); ok {
		r0 = rf(ctx, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Transaction) error); ok {
		r1 = rf(ctx, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewITransactionService interface {
	mock.TestingT
	Cleanup(func())
}

// NewITransactionService creates a new instance of ITransactionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewITransactionService(t mockConstructorTestingTNewITransactionService) *ITransactionService {
	mock := &ITransactionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
