// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// Library is an autogenerated mock type for the Library type
type Library struct {
	mock.Mock
}

// Delete provides a mock function with given fields: item
func (_m *Library) Delete(item domain.Media) {
	_m.Called(item)
}

// Serve provides a mock function with given fields: _a0, acceptWebP
func (_m *Library) Serve(_a0 domain.Media, acceptWebP bool) ([]byte, domain.Mime, error) {
	ret := _m.Called(_a0, acceptWebP)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(domain.Media, bool) []byte); ok {
		r0 = rf(_a0, acceptWebP)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 domain.Mime
	if rf, ok := ret.Get(1).(func(domain.Media, bool) domain.Mime); ok {
		r1 = rf(_a0, acceptWebP)
	} else {
		r1 = ret.Get(1).(domain.Mime)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(domain.Media, bool) error); ok {
		r2 = rf(_a0, acceptWebP)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Upload provides a mock function with given fields: file
func (_m *Library) Upload(file *multipart.FileHeader) (domain.Media, error) {
	ret := _m.Called(file)

	var r0 domain.Media
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader) domain.Media); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(domain.Media)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*multipart.FileHeader) error); ok {
		r1 = rf(file)
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
