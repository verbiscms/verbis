// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	fields "github.com/ainsleyclark/verbis/api/fields"

	mock "github.com/stretchr/testify/mock"
)

// FieldService is an autogenerated mock type for the FieldService type
type FieldService struct {
	mock.Mock
}

// GetField provides a mock function with given fields: name, args
func (_m *FieldService) GetField(name string, args ...interface{}) (interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(name, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(name, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFields provides a mock function with given fields: args
func (_m *FieldService) GetFields(args ...interface{}) (fields.Fields, error) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 fields.Fields
	if rf, ok := ret.Get(0).(func(...interface{}) fields.Fields); ok {
		r0 = rf(args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fields.Fields)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlexible provides a mock function with given fields: name, args
func (_m *FieldService) GetFlexible(name string, args ...interface{}) (fields.Flexible, error) {
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 fields.Flexible
	if rf, ok := ret.Get(0).(func(string, ...interface{}) fields.Flexible); ok {
		r0 = rf(name, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fields.Flexible)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(name, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLayout provides a mock function with given fields: name, args
func (_m *FieldService) GetLayout(name string, args ...interface{}) (domain.Field, error) {
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 domain.Field
	if rf, ok := ret.Get(0).(func(string, ...interface{}) domain.Field); ok {
		r0 = rf(name, args...)
	} else {
		r0 = ret.Get(0).(domain.Field)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(name, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLayouts provides a mock function with given fields: args
func (_m *FieldService) GetLayouts(args ...interface{}) []domain.FieldGroup {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 []domain.FieldGroup
	if rf, ok := ret.Get(0).(func(...interface{}) []domain.FieldGroup); ok {
		r0 = rf(args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.FieldGroup)
		}
	}

	return r0
}

// GetRepeater provides a mock function with given fields: name, args
func (_m *FieldService) GetRepeater(name string, args ...interface{}) (fields.Repeater, error) {
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 fields.Repeater
	if rf, ok := ret.Get(0).(func(string, ...interface{}) fields.Repeater); ok {
		r0 = rf(name, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fields.Repeater)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(name, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
