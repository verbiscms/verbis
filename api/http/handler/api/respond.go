// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"time"
)

// RespondJSON defines the main struct for sending back
// data for API requests, it includes a status code,
// if the handler produced an error, a message,
// meta information and the main data.
type RespondJSON struct {
	Status  int         `json:"status"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Meta    Meta        `json:"meta"`
	Data    interface{} `json:"data"`
}

// Meta defines any additional information for API
// request, including pagination for list routes.
type Meta struct {
	RequestTime  string      `json:"request_time"`
	ResponseTime string      `json:"response_time"`
	LatencyTime  string      `json:"latency_time"`
	Pagination   interface{} `json:"pagination,omitempty"`
}

// ErrorJson defines the validation errors if there any
// when processing data via handlers.
type ErrorJson struct {
	Errors validation.Errors `json:"errors"`
}

// Respond
//
// Returns RespondJSON and sends back the main data for
// use with the API. Returns status, message and data.
func Respond(ctx *gin.Context, status int, message string, data interface{}, pagination ...*pagination.Pagination) {
	ctx.Set("verbis_message", message)

	hasError := false
	if status != 200 {
		hasError = true
	}

	ctx.JSON(status, RespondJSON{
		Status:  status,
		Message: message,
		Error:   hasError,
		Meta:    GetMeta(ctx, pagination),
		Data:    checkResponseData(ctx, data),
	})

	return
}

// AbortJSON
//
// Returns RespondJSON and aborts the request with given
// status, message and data.
func AbortJSON(g *gin.Context, status int, message string, data interface{}) {
	hasError := false
	if status != 200 {
		hasError = true
	}

	g.AbortWithStatusJSON(status, RespondJSON{
		Status:  status,
		Error:   hasError,
		Message: message,
		Meta:    GetMeta(g, nil),
		Data:    checkResponseData(g, data),
	})
}

// checkResponseData
//
// Checks what type of data is passed and processes it
// accordingly. errors, empty slices & interfaces as
// well as validation. Returns the original data
// if the type passed is not of type error or
// nil
func checkResponseData(ctx *gin.Context, data interface{}) interface{} {

	switch v := data.(type) {
	case nil:
		return gin.H{}
	case *errors.Error:
		ctx.Set("verbis_error", v)

		if v.Err == nil {
			return gin.H{}
		}

		errType := reflect.TypeOf(v.Err)
		if errType.String() == "validator.ValidationErrors" && v.Code == errors.INVALID {
			validationErrors := v.Err.(validator.ValidationErrors)
			val := validation.New()
			return &ErrorJson{
				Errors: val.Process(validationErrors),
			}
		} else {
			return gin.H{}
		}
	case *json.UnmarshalTypeError:
		return ErrorJson{
			Errors: validation.Errors{
				{
					Key:     v.Field,
					Type:    "Unmarshal error",
					Message: "Invalid type passed to " + v.Struct + " struct.",
				},
			},
		}
	}

	if reflect.TypeOf(data).Kind().String() == "slice" {
		s := reflect.ValueOf(data)
		ret := make([]interface{}, s.Len())
		if len(ret) == 0 {
			return gin.H{}
		}
	}

	return data
}

// GetMeta
//
// Processes the request and response time and calculates
// latency time. Sets pagination if the length is
// greater than one.
func GetMeta(ctx *gin.Context, pagination []*pagination.Pagination) Meta {

	// Calculate start, end and latency time
	var startTime = time.Now()
	requestTime, exists := ctx.Get("request_time")
	if exists {
		startTime = requestTime.(time.Time)
	}
	latencyTime := time.Since(startTime)

	// Check if the pagination is empty
	var p interface{} = nil
	if len(pagination) == 1 {
		p = pagination[0]
	}

	return Meta{
		RequestTime:  startTime.UTC().String(),
		ResponseTime: time.Now().UTC().String(),
		LatencyTime:  latencyTime.Round(time.Microsecond).String(),
		Pagination:   p,
	}
}
