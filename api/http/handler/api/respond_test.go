// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	gohttp "net/http"
	"time"
)

type tester struct {
	Name string `json:"name"`
}

func (t *ApiTestSuite) TestRespond() {

	pagination := http.Pagination{Page: 1, Pages: 2, Limit: 10, Total: 0, Next: false, Prev: false}
	var verr = errors.Error{Err: nil}

	tt := map[string]struct {
		status     int
		message    string
		data       interface{}
		pagination *http.Pagination
		want       RespondJSON
	}{
		"Nil Data": {
			200,
			"message",
			nil,
			nil,
			RespondJSON{
				Status:  200,
				Error:   false,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Error": {
			400,
			"message",
			nil,
			nil,
			RespondJSON{
				Status:  400,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Pagination": {
			200,
			"message",
			nil,
			&pagination,
			RespondJSON{
				Status:  200,
				Error:   false,
				Message: "message",
				Meta:    Meta{Pagination: map[string]interface{}{"limit": float64(10), "next": false, "page": float64(1), "pages": float64(2), "prev": false, "total": float64(0)}},
				Data:    map[string]interface{}{},
			},
		},
		"Data String": {
			200,
			"message",
			"hello",
			nil,
			RespondJSON{
				Status:  200,
				Error:   false,
				Message: "message",
				Data:    "hello",
			},
		},
		"Data Struct": {
			200,
			"message",
			tester{Name: "test"},
			nil,
			RespondJSON{
				Status:  200,
				Error:   false,
				Message: "message",
				Data:    map[string]interface{}{"name": "test"},
			},
		},
		"Data Struct Slice": {
			200,
			"message",
			[]tester{{Name: "test1"}, {Name: "test2"}},
			nil,
			RespondJSON{
				Status:  200,
				Error:   false,
				Message: "message",
				Data:    []interface{}{map[string]interface{}{"name": "test1"}, map[string]interface{}{"name": "test2"}},
			},
		},
		"Verbis Error": {
			400,
			"message",
			&errors.Error{Code: errors.INVALID, Err: validator.ValidationErrors{}},
			nil,
			RespondJSON{
				Status:  400,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{"errors": interface{}(nil)},
			},
		},
		"Verbis Error Nil": {
			400,
			"message",
			&verr,
			nil,
			RespondJSON{
				Status:  400,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Verbis No Error": {
			400,
			"message",
			&errors.Error{Message: "error message"},
			nil,
			RespondJSON{
				Status:  400,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{},
			},
		},
		"Unmarshal Error": {
			400,
			"message",
			&json.UnmarshalTypeError{Struct: "struct", Field: "field"},
			nil,
			RespondJSON{
				Status:  400,
				Error:   true,
				Message: "message",
				Data:    map[string]interface{}{"errors": []interface{}{map[string]interface{}{"key": "field", "message": "Invalid type passed to struct struct.", "type": "Unmarshal error"}}},
			},
		},
		"Empty Slice": {
			200,
			"message",
			[]tester{},
			nil,
			RespondJSON{
				Status:  200,
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

			t.Reset()
		})
	}
}

func (t *ApiTestSuite) Test_GetMeta() {
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
