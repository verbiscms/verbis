// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"
)

// WalkerFunc is an autogenerated mock type for the WalkerFunc type
type WalkerFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: field
func (_m *WalkerFunc) Execute(field domain.PostField) {
	_m.Called(field)
}
