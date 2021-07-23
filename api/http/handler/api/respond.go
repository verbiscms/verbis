// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/pagination"
	"github.com/verbiscms/verbis/api/version"
	"net/http"
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

// ErrorJSON defines the validation errors if there any
// when processing data via handlers.
type ErrorJSON struct {
	Errors validation.Errors `json:"errors"`
}

// Respond returns RespondJSON and sends back the main
// data for use with the API. Returns status, message
// and data.
func Respond(ctx *gin.Context, status int, message string, data interface{}, p ...*pagination.Pagination) {
	ctx.Set("verbis_message", message)

	hasError := false
	if status != http.StatusOK {
		hasError = true
	}

	ctx.Header(version.Header, version.Version)

	ctx.JSON(status, RespondJSON{
		Status:  status,
		Message: message,
		Error:   hasError,
		Meta:    GetMeta(ctx, p),
		Data:    checkResponseData(ctx, data),
	})
}

// AbortJSON returns RespondJSON and aborts the request
// with given status, message and data.
func AbortJSON(ctx *gin.Context, status int, message string, data interface{}) {
	hasError := false
	if status != http.StatusOK {
		hasError = true
	}

	ctx.Header(version.Header, version.Version)

	ctx.AbortWithStatusJSON(status, RespondJSON{
		Status:  status,
		Error:   hasError,
		Message: message,
		Meta:    GetMeta(ctx, nil),
		Data:    checkResponseData(ctx, data),
	})
}

// checkResponseData Checks what type of data is passed
// and processes it accordingly. errors, empty slices
// & interfaces as well as validation. Returns the
// original data if the type passed is not of
// type error or nil
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
			return &ErrorJSON{
				Errors: val.Process(validationErrors),
			}
		} else {
			return gin.H{}
		}
	case *json.UnmarshalTypeError:
		return ErrorJSON{
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
func GetMeta(ctx *gin.Context, p []*pagination.Pagination) Meta {
	// Calculate start, end and latency time
	var startTime = time.Now()
	requestTime, exists := ctx.Get("request_time")
	if exists {
		startTime = requestTime.(time.Time)
	}
	latencyTime := time.Since(startTime)

	// Check if the pagination is empty
	var pag interface{} = nil
	if len(p) == 1 {
		pag = p[0]
	}

	return Meta{
		RequestTime:  startTime.UTC().String(),
		ResponseTime: time.Now().UTC().String(),
		LatencyTime:  latencyTime.Round(time.Microsecond).String(),
		Pagination:   pag,
	}
}
