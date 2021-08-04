// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/verbiscms/verbis/api/domain"
)

// Migrator is an autogenerated mock type for the Migrator type
type Migrator struct {
	mock.Mock
}

// Migrate provides a mock function with given fields: ctx, from, to, delete
func (_m *Migrator) Migrate(ctx context.Context, from domain.StorageChange, to domain.StorageChange, delete bool) (int, error) {
	ret := _m.Called(ctx, from, to, delete)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, domain.StorageChange, domain.StorageChange, bool) int); ok {
		r0 = rf(ctx, from, to, delete)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.StorageChange, domain.StorageChange, bool) error); ok {
		r1 = rf(ctx, from, to, delete)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
