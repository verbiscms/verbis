// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	bytes "bytes"

	domain "github.com/ainsleyclark/verbis/api/domain"
	image "github.com/ainsleyclark/verbis/api/services/media/image"

	mock "github.com/stretchr/testify/mock"
)

// Resizer is an autogenerated mock type for the Resizer type
type Resizer struct {
	mock.Mock
}

// Resize provides a mock function with given fields: imager, compression, media
func (_m *Resizer) Resize(imager image.Imager, compression int, media domain.MediaSize) (*bytes.Reader, error) {
	ret := _m.Called(imager, compression, media)

	var r0 *bytes.Reader
	if rf, ok := ret.Get(0).(func(image.Imager, int, domain.MediaSize) *bytes.Reader); ok {
		r0 = rf(imager, compression, media)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bytes.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(image.Imager, int, domain.MediaSize) error); ok {
		r1 = rf(imager, compression, media)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
