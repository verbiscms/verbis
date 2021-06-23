// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"

	params "github.com/ainsleyclark/verbis/api/helpers/params"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: redirect
func (_m *Repository) Create(redirect domain.Redirect) (domain.Redirect, error) {
	ret := _m.Called(redirect)

	var r0 domain.Redirect
	if rf, ok := ret.Get(0).(func(domain.Redirect) domain.Redirect); ok {
		r0 = rf(redirect)
	} else {
		r0 = ret.Get(0).(domain.Redirect)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Redirect) error); ok {
		r1 = rf(redirect)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: id
func (_m *Repository) Exists(id int) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ExistsByFrom provides a mock function with given fields: from
func (_m *Repository) ExistsByFrom(from string) bool {
	ret := _m.Called(from)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(from)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Find provides a mock function with given fields: id
func (_m *Repository) Find(id int) (domain.Redirect, error) {
	ret := _m.Called(id)

	var r0 domain.Redirect
	if rf, ok := ret.Get(0).(func(int) domain.Redirect); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Redirect)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByFrom provides a mock function with given fields: from
func (_m *Repository) FindByFrom(from string) (domain.Redirect, error) {
	ret := _m.Called(from)

	var r0 domain.Redirect
	if rf, ok := ret.Get(0).(func(string) domain.Redirect); ok {
		r0 = rf(from)
	} else {
		r0 = ret.Get(0).(domain.Redirect)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: meta
func (_m *Repository) List(meta params.Params) (domain.Redirects, int, error) {
	ret := _m.Called(meta)

	var r0 domain.Redirects
	if rf, ok := ret.Get(0).(func(params.Params) domain.Redirects); ok {
		r0 = rf(meta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Redirects)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(params.Params) int); ok {
		r1 = rf(meta)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(params.Params) error); ok {
		r2 = rf(meta)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: redirect
func (_m *Repository) Update(redirect domain.Redirect) (domain.Redirect, error) {
	ret := _m.Called(redirect)

	var r0 domain.Redirect
	if rf, ok := ret.Get(0).(func(domain.Redirect) domain.Redirect); ok {
		r0 = rf(redirect)
	} else {
		r0 = ret.Get(0).(domain.Redirect)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Redirect) error); ok {
		r1 = rf(redirect)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
