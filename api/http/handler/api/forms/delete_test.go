// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *FormsTestSuite) TestForms_Delete() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.FormRepository)
		url     string
	}{
		"Success": {
			nil,
			200,
			"Successfully deleted form with ID: 123",
			func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(nil)
			},
			"/forms/123",
		},
		"Invalid ID": {
			nil,
			400,
			"A valid ID is required to delete a form",
			func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(nil)
			},
			"/forms/wrongid",
		},
		"Not Found": {
			nil,
			400,
			"not found",
			func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/forms/123",
		},
		"Conflict": {
			nil,
			400,
			"conflict",
			func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			"/forms/123",
		},
		"Internal": {
			nil,
			500,
			"internal",
			func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/forms/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodDelete, test.url, "/forms/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Delete(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}