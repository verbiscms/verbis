// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	tpl "github.com/verbiscms/verbis/api/tpl"
)

// Resolver is an autogenerated mock type for the Resolver type
type Resolver struct {
	mock.Mock
}

// Execute provides a mock function with given fields: custom
func (_m *Resolver) Execute(custom bool) (string, tpl.TemplateExecutor, bool) {
	ret := _m.Called(custom)

	var r0 string
	if rf, ok := ret.Get(0).(func(bool) string); ok {
		r0 = rf(custom)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 tpl.TemplateExecutor
	if rf, ok := ret.Get(1).(func(bool) tpl.TemplateExecutor); ok {
		r1 = rf(custom)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(tpl.TemplateExecutor)
		}
	}

	var r2 bool
	if rf, ok := ret.Get(2).(func(bool) bool); ok {
		r2 = rf(custom)
	} else {
		r2 = ret.Get(2).(bool)
	}

	return r0, r1, r2
}
