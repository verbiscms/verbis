// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *FormsTestSuite) TestForms_Find() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.FormRepository)
		url     string
	}{
		"Success": {
			form,
			http.StatusOK,
			"Successfully obtained form with ID: 123",
			func(m *mocks.FormRepository) {
				m.On("GetByID", 123).Return(form, nil)
			},
			"/forms/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"Pass a valid number to obtain the form by ID",
			func(m *mocks.FormRepository) {
				m.On("GetByID", 123).Return(domain.Form{}, fmt.Errorf("error"))
			},
			"/forms/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"no forms found",
			func(m *mocks.FormRepository) {
				m.On("GetByID", 123).Return(domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: "no forms found"})
			},
			"/forms/123",
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.FormRepository) {
				m.On("GetByID", 123).Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/forms/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/forms/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Find(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
