// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	gohttp "net/http"
	"time"
)

type tester struct {
	Name string `json:"name"`
}

func (t *APITestSuite) TestRespond() {
	paginate := pagination.Pagination{Page: 1, Pages: 2, Limit: 10, Total: 0, Next: false, Prev: false}
	var verr = errors.Error{Err: nil}

	tt := map[string]struct {
		status     int
		message    string
		data       interface{}
		pagination *pagination.Pagination
		want       RespondJSON
	}{
		"Nil Data": {
			gohttp.StatusOK,
			"message",
			nil,
			nil,
			RespondJSON{
				Status:  gohttp.StatusOK,
				Error:   false,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Error": {
			gohttp.StatusBadRequest,
			"message",
			nil,
			nil,
			RespondJSON{
				Status:  gohttp.StatusBadRequest,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Pagination": {
			gohttp.StatusOK,
			"message",
			nil,
			&paginate,
			RespondJSON{
				Status:  gohttp.StatusOK,
				Error:   false,
				Message: "message",
				Meta:    Meta{Pagination: map[string]interface{}{"limit": float64(10), "next": false, "page": float64(1), "pages": float64(2), "prev": false, "total": float64(0)}},
				Data:    map[string]interface{}{},
			},
		},
		"Data String": {
			gohttp.StatusOK,
			"message",
			"hello",
			nil,
			RespondJSON{
				Status:  gohttp.StatusOK,
				Error:   false,
				Message: "message",
				Data:    "hello",
			},
		},
		"Data Struct": {
			gohttp.StatusOK,
			"message",
			tester{Name: "test"},
			nil,
			RespondJSON{
				Status:  gohttp.StatusOK,
				Error:   false,
				Message: "message",
				Data:    map[string]interface{}{"name": "test"},
			},
		},
		"Data Struct Slice": {
			gohttp.StatusOK,
			"message",
			[]tester{{Name: "test1"}, {Name: "test2"}},
			nil,
			RespondJSON{
				Status:  gohttp.StatusOK,
				Error:   false,
				Message: "message",
				Data:    []interface{}{map[string]interface{}{"name": "test1"}, map[string]interface{}{"name": "test2"}},
			},
		},
		"Verbis Error": {
			gohttp.StatusBadRequest,
			"message",
			&errors.Error{Code: errors.INVALID, Err: validator.ValidationErrors{}},
			nil,
			RespondJSON{
				Status:  gohttp.StatusBadRequest,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{"errors": interface{}(nil)},
			},
		},
		"Verbis Error - Not Invalid": {
			gohttp.StatusBadRequest,
			"message",
			&errors.Error{Code: errors.CONFLICT, Err: validator.ValidationErrors{}},
			nil,
			RespondJSON{
				Status:  gohttp.StatusBadRequest,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Verbis Error Nil": {
			gohttp.StatusBadRequest,
			"message",
			&verr,
			nil,
			RespondJSON{
				Status:  gohttp.StatusBadRequest,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Verbis No Error": {
			gohttp.StatusBadRequest,
			"message",
			&errors.Error{Message: "error message"},
			nil,
			RespondJSON{
				Status:  gohttp.StatusBadRequest,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Unmarshal Error": {
			gohttp.StatusBadRequest,
			"message",
			&json.UnmarshalTypeError{Struct: "struct", Field: "field"},
			nil,
			RespondJSON{
				Status:  gohttp.StatusBadRequest,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{"errors": []interface{}{map[string]interface{}{"key": "field", "message": "Invalid type passed to struct struct.", "type": "Unmarshal error"}}},
			},
		},
		"Empty Slice": {
			gohttp.StatusOK,
			"message",
			[]tester{},
			nil,
			RespondJSON{
				Status:  gohttp.StatusOK,
				Error:   false,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(gohttp.MethodGet, "/respond", "/respond", nil, func(ctx *gin.Context) {
				Respond(ctx, test.status, test.message, test.data, test.pagination)
			})
			respond, _ := t.RespondData()

			if test.pagination != nil {
				t.Equal(test.want.Meta.Pagination, respond.Meta.Pagination)
			} else {
				t.Nil(test.want.Meta.Pagination, respond.Meta.Pagination)
			}

			t.Equal(test.want.Status, respond.Status)
			t.Equal(test.want.Error, respond.Error)
			t.Equal(test.want.Message, respond.Message)
			t.Equal(test.want.Data, respond.Data)
			t.Equal(t.Recorder.Header().Get(version.Header), version.Version)

			t.Reset()
		})
	}
}

func (t *APITestSuite) TestAbortJSON() {
	tt := map[string]struct {
		status  int
		message string
		data    interface{}
		want    RespondJSON
	}{
		"With Data": {
			gohttp.StatusOK,
			"message",
			map[string]interface{}{"test": "test"},
			RespondJSON{
				Status:  gohttp.StatusOK,
				Error:   false,
				Message: "message",
				Data:    map[string]interface{}{"test": "test"},
			},
		},
		"Nil Data": {
			gohttp.StatusOK,
			"message",
			nil,
			RespondJSON{
				Status:  gohttp.StatusOK,
				Error:   false,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Error": {
			gohttp.StatusBadRequest,
			"message",
			nil,
			RespondJSON{
				Status:  gohttp.StatusBadRequest,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(gohttp.MethodGet, "/respond", "/respond", nil, func(ctx *gin.Context) {
				AbortJSON(ctx, test.status, test.message, test.data)
			})
			respond, _ := t.RespondData()

			t.Equal(test.want.Status, respond.Status)
			t.Equal(test.want.Error, respond.Error)
			t.Equal(test.want.Message, respond.Message)
			t.Equal(test.want.Data, respond.Data)
			t.Equal(t.Recorder.Header().Get(version.Header), version.Version)

			t.Reset()
		})
	}
}

func (t *APITestSuite) Test_GetMeta() {
	var got Meta
	t.RequestAndServe(gohttp.MethodGet, "/meta", "/meta", nil, func(ctx *gin.Context) {
		ctx.Set("request_time", time.Now())
		got = GetMeta(ctx, nil)
	})

	layout := "2006-01-02 15:04:05 -0700 MST"

	request, err := time.Parse(layout, got.RequestTime)
	t.NoError(err)

	response, err := time.Parse(layout, got.RequestTime)
	t.NoError(err)

	t.WithinDuration(time.Now(), request, 10*time.Millisecond)
	t.WithinDuration(time.Now(), response, 10*time.Millisecond)
}
