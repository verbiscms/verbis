// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	image "image"

	mock "github.com/stretchr/testify/mock"
)

// Imager is an autogenerated mock type for the Imager type
type Imager struct {
	mock.Mock
}

// Decode provides a mock function with given fields:
func (_m *Imager) Decode() (image.Image, error) {
	ret := _m.Called()

	var r0 image.Image
	if rf, ok := ret.Get(0).(func() image.Image); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(image.Image)
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

// Save provides a mock function with given fields: _a0, path, comp
func (_m *Imager) Save(_a0 image.Image, path string, comp int) error {
	ret := _m.Called(_a0, path, comp)

	var r0 error
	if rf, ok := ret.Get(0).(func(image.Image, string, int) error); ok {
		r0 = rf(_a0, path, comp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
