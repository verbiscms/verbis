// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// SiteHandler is an autogenerated mock type for the SiteHandler type
type SiteHandler struct {
	mock.Mock
}

// GetLayouts provides a mock function with given fields: g
func (_m *SiteHandler) GetLayouts(g *gin.Context) {
	_m.Called(g)
}

// GetSite provides a mock function with given fields: g
func (_m *SiteHandler) GetSite(g *gin.Context) {
	_m.Called(g)
}

// GetTemplates provides a mock function with given fields: g
func (_m *SiteHandler) GetTemplates(g *gin.Context) {
	_m.Called(g)
}

// GetTheme provides a mock function with given fields: g
func (_m *SiteHandler) GetTheme(g *gin.Context) {
	_m.Called(g)
}
