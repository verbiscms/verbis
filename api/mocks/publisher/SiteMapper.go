// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SiteMapper is an autogenerated mock type for the SiteMapper type
type SiteMapper struct {
	mock.Mock
}

// Index provides a mock function with given fields:
func (_m *SiteMapper) Index() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
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

// Pages provides a mock function with given fields: resource
func (_m *SiteMapper) Pages(resource string) ([]byte, error) {
	ret := _m.Called(resource)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(resource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// XSL provides a mock function with given fields: index
func (_m *SiteMapper) XSL(index bool) ([]byte, error) {
	ret := _m.Called(index)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(bool) []byte); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
