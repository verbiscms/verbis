// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/verbiscms/verbis/api/domain"
)

// Bucket is an autogenerated mock type for the Bucket type
type Bucket struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Bucket) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: name
func (_m *Bucket) Exists(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Find provides a mock function with given fields: url
func (_m *Bucket) Find(url string) ([]byte, domain.File, error) {
	ret := _m.Called(url)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 domain.File
	if rf, ok := ret.Get(1).(func(string) domain.File); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Get(1).(domain.File)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(url)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Upload provides a mock function with given fields: upload
func (_m *Bucket) Upload(upload domain.Upload) (domain.File, error) {
	ret := _m.Called(upload)

	var r0 domain.File
	if rf, ok := ret.Get(0).(func(domain.Upload) domain.File); ok {
		r0 = rf(upload)
	} else {
		r0 = ret.Get(0).(domain.File)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Upload) error); ok {
		r1 = rf(upload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
