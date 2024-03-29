// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	store "github.com/eko/gocache/v2/store"
)

// Provider is an autogenerated mock type for the provider type
type Provider struct {
	mock.Mock
}

// Driver provides a mock function with given fields:
func (_m *Provider) Driver() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Ping provides a mock function with given fields:
func (_m *Provider) Ping() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields:
func (_m *Provider) Store() store.StoreInterface {
	ret := _m.Called()

	var r0 store.StoreInterface
	if rf, ok := ret.Get(0).(func() store.StoreInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreInterface)
		}
	}

	return r0
}

// Validate provides a mock function with given fields:
func (_m *Provider) Validate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
