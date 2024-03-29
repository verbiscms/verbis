// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	tpl "github.com/verbiscms/verbis/api/tpl"
)

// FileHandler is an autogenerated mock type for the FileHandler type
type FileHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: config, template
func (_m *FileHandler) Execute(config tpl.TemplateConfig, template string) (string, error) {
	ret := _m.Called(config, template)

	var r0 string
	if rf, ok := ret.Get(0).(func(tpl.TemplateConfig, string) string); ok {
		r0 = rf(config, template)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(tpl.TemplateConfig, string) error); ok {
		r1 = rf(config, template)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
