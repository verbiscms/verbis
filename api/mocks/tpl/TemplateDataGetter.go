// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// TemplateDataGetter is an autogenerated mock type for the TemplateDataGetter type
type TemplateDataGetter struct {
	mock.Mock
}

// Data provides a mock function with given fields: ctx, post
func (_m *TemplateDataGetter) Data(ctx *gin.Context, post *domain.PostDatum) interface{} {
	ret := _m.Called(ctx, post)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*gin.Context, *domain.PostDatum) interface{}); ok {
		r0 = rf(ctx, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}
