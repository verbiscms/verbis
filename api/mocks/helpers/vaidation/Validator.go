// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	validator "github.com/go-playground/validator/v10"
	mock "github.com/stretchr/testify/mock"
)

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// CmdCheck provides a mock function with given fields: key, data
func (_m *Validator) CmdCheck(key string, data interface{}) error {
	ret := _m.Called(key, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Process provides a mock function with given fields: errors
func (_m *Validator) Process(errors validator.ValidationErrors) []validation.ValidationError {
	ret := _m.Called(errors)

	var r0 []validation.ValidationError
	if rf, ok := ret.Get(0).(func(validator.ValidationErrors) []validation.ValidationError); ok {
		r0 = rf(errors)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]validation.ValidationError)
		}
	}

	return r0
}

// message provides a mock function with given fields: kind, field, param
func (_m *Validator) message(kind string, field string, param string) string {
	ret := _m.Called(kind, field, param)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(kind, field, param)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
