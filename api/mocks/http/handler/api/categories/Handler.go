// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx
func (_m *Handler) Create(ctx *gin.Context) {
	_m.Called(ctx)
}

// Delete provides a mock function with given fields: ctx
func (_m *Handler) Delete(ctx *gin.Context) {
	_m.Called(ctx)
}

// Find provides a mock function with given fields: ctx
func (_m *Handler) Find(ctx *gin.Context) {
	_m.Called(ctx)
}

// List provides a mock function with given fields: ctx
func (_m *Handler) List(ctx *gin.Context) {
	_m.Called(ctx)
}

// Update provides a mock function with given fields: ctx
func (_m *Handler) Update(ctx *gin.Context) {
	_m.Called(ctx)
}
