// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Stringer is an autogenerated mock type for the Stringer type
type Stringer struct {
	mock.Mock
}

// Param provides a mock function with given fields: _a0
func (_m *Stringer) Param(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
