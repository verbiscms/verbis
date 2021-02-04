// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// FieldHandler is an autogenerated mock type for the FieldHandler type
type FieldHandler struct {
	mock.Mock
}

// Get provides a mock function with given fields: g
func (_m *FieldHandler) Get(g *gin.Context) {
	_m.Called(g)
}
