// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	recovery "github.com/ainsleyclark/verbis/api/recovery"
	mock "github.com/stretchr/testify/mock"
)

// TplData is an autogenerated mock type for the TplData type
type TplData struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *TplData) Execute() *recovery.Data {
	ret := _m.Called()

	var r0 *recovery.Data
	if rf, ok := ret.Get(0).(func() *recovery.Data); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*recovery.Data)
		}
	}

	return r0
}
