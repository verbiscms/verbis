// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"
)

// Finder is an autogenerated mock type for the Finder type
type Finder struct {
	mock.Mock
}

// GetLayout provides a mock function with given fields: post, cacheable
func (_m *Finder) GetLayout(post domain.PostData, cacheable bool) []domain.FieldGroup {
	ret := _m.Called(post, cacheable)

	var r0 []domain.FieldGroup
	if rf, ok := ret.Get(0).(func(domain.PostData, bool) []domain.FieldGroup); ok {
		r0 = rf(post, cacheable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.FieldGroup)
		}
	}

	return r0
}
