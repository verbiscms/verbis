// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"
)

// RoleRepository is an autogenerated mock type for the RoleRepository type
type RoleRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: r
func (_m *RoleRepository) Create(r *domain.Role) (domain.Role, error) {
	ret := _m.Called(r)

	var r0 domain.Role
	if rf, ok := ret.Get(0).(func(*domain.Role) domain.Role); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(domain.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Role) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exists provides a mock function with given fields: name
func (_m *RoleRepository) Exists(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields:
func (_m *RoleRepository) Get() ([]domain.Role, error) {
	ret := _m.Called()

	var r0 []domain.Role
	if rf, ok := ret.Get(0).(func() []domain.Role); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *RoleRepository) GetByID(id int) (domain.Role, error) {
	ret := _m.Called(id)

	var r0 domain.Role
	if rf, ok := ret.Get(0).(func(int) domain.Role); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: r
func (_m *RoleRepository) Update(r *domain.Role) (domain.Role, error) {
	ret := _m.Called(r)

	var r0 domain.Role
	if rf, ok := ret.Get(0).(func(*domain.Role) domain.Role); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(domain.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Role) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
