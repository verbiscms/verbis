// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"
)

// SeoMetaRepository is an autogenerated mock type for the SeoMetaRepository type
type SeoMetaRepository struct {
	mock.Mock
}

// UpdateCreate provides a mock function with given fields: p
func (_m *SeoMetaRepository) UpdateCreate(p *domain.PostData) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.PostData) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
