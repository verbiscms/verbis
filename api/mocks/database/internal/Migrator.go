// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	version "github.com/hashicorp/go-version"
	mock "github.com/stretchr/testify/mock"
)

// Migrator is an autogenerated mock type for the Migrator type
type Migrator struct {
	mock.Mock
}

// Migrate provides a mock function with given fields: _a0
func (_m *Migrator) Migrate(_a0 *version.Version) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*version.Version) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
