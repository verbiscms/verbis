// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Exists provides a mock function with given fields: name
func (_m *Repository) Exists(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Find provides a mock function with given fields: name
func (_m *Repository) Find(name string) (interface{}, error) {
	ret := _m.Called(name)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTheme provides a mock function with given fields:
func (_m *Repository) GetTheme() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *Repository) Insert(_a0 domain.OptionsDBMap) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.OptionsDBMap) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Map provides a mock function with given fields:
func (_m *Repository) Map() (domain.OptionsDBMap, error) {
	ret := _m.Called()

	var r0 domain.OptionsDBMap
	if rf, ok := ret.Get(0).(func() domain.OptionsDBMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.OptionsDBMap)
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

// SetTheme provides a mock function with given fields: theme
func (_m *Repository) SetTheme(theme string) error {
	ret := _m.Called(theme)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(theme)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Struct provides a mock function with given fields:
func (_m *Repository) Struct() domain.Options {
	ret := _m.Called()

	var r0 domain.Options
	if rf, ok := ret.Get(0).(func() domain.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.Options)
	}

	return r0
}
