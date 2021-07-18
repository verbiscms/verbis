// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"

	params "github.com/ainsleyclark/verbis/api/common/params"
)

// Library is an autogenerated mock type for the Library type
type Library struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Library) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: id
func (_m *Library) Find(id int) (domain.Media, error) {
	ret := _m.Called(id)

	var r0 domain.Media
	if rf, ok := ret.Get(0).(func(int) domain.Media); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Media)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: meta
func (_m *Library) List(meta params.Params) (domain.MediaItems, int, error) {
	ret := _m.Called(meta)

	var r0 domain.MediaItems
	if rf, ok := ret.Get(0).(func(params.Params) domain.MediaItems); ok {
		r0 = rf(meta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.MediaItems)
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

// Update provides a mock function with given fields: m
func (_m *Library) Update(m domain.Media) (domain.Media, error) {
	ret := _m.Called(m)

	var r0 domain.Media
	if rf, ok := ret.Get(0).(func(domain.Media) domain.Media); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Get(0).(domain.Media)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Media) error); ok {
		r1 = rf(m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upload provides a mock function with given fields: file, userID
func (_m *Library) Upload(file *multipart.FileHeader, userID int) (domain.Media, error) {
	ret := _m.Called(file, userID)

	var r0 domain.Media
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader, int) domain.Media); ok {
		r0 = rf(file, userID)
	} else {
		r0 = ret.Get(0).(domain.Media)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*multipart.FileHeader, int) error); ok {
		r1 = rf(file, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: file
func (_m *Library) Validate(file *multipart.FileHeader) error {
	ret := _m.Called(file)

	var r0 error
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader) error); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
