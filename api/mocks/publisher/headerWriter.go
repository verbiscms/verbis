// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// HeaderWriter is an autogenerated mock type for the headerWriter type
type HeaderWriter struct {
	mock.Mock
}

// Cache provides a mock function with given fields: g
func (_m *HeaderWriter) Cache(g *gin.Context) {
	_m.Called(g)
}
