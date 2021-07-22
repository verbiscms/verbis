// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"

	params "github.com/ainsleyclark/verbis/api/common/params"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: f
func (_m *Repository) Create(f domain.File) (domain.File, error) {
	ret := _m.Called(f)

	var r0 domain.File
	if rf, ok := ret.Get(0).(func(domain.File) domain.File); ok {
		r0 = rf(f)
	} else {
		r0 = ret.Get(0).(domain.File)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.File) error); ok {
		r1 = rf(f)
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

// Exists provides a mock function with given fields: fileName
func (_m *Repository) Exists(fileName string) bool {
	ret := _m.Called(fileName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(fileName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Find provides a mock function with given fields: id
func (_m *Repository) Find(id int) (domain.File, error) {
	ret := _m.Called(id)

	var r0 domain.File
	if rf, ok := ret.Get(0).(func(int) domain.File); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.File)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByURL provides a mock function with given fields: url
func (_m *Repository) FindByURL(url string) (domain.File, error) {
	ret := _m.Called(url)

	var r0 domain.File
	if rf, ok := ret.Get(0).(func(string) domain.File); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(domain.File)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: meta
func (_m *Repository) List(meta params.Params) (domain.Files, int, error) {
	ret := _m.Called(meta)

	var r0 domain.Files
	if rf, ok := ret.Get(0).(func(params.Params) domain.Files); ok {
		r0 = rf(meta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Files)
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

// Update provides a mock function with given fields: f
func (_m *Repository) Update(f domain.File) (domain.File, error) {
	ret := _m.Called(f)

	var r0 domain.File
	if rf, ok := ret.Get(0).(func(domain.File) domain.File); ok {
		r0 = rf(f)
	} else {
		r0 = ret.Get(0).(domain.File)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.File) error); ok {
		r1 = rf(f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
