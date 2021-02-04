// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// MediaHandler is an autogenerated mock type for the MediaHandler type
type MediaHandler struct {
	mock.Mock
}

// Delete provides a mock function with given fields: g
func (_m *MediaHandler) Delete(g *gin.Context) {
	_m.Called(g)
}

// Get provides a mock function with given fields: g
func (_m *MediaHandler) Get(g *gin.Context) {
	_m.Called(g)
}

// GetById provides a mock function with given fields: g
func (_m *MediaHandler) GetById(g *gin.Context) {
	_m.Called(g)
}

// Update provides a mock function with given fields: g
func (_m *MediaHandler) Update(g *gin.Context) {
	_m.Called(g)
}

// Upload provides a mock function with given fields: g
func (_m *MediaHandler) Upload(g *gin.Context) {
	_m.Called(g)
}
