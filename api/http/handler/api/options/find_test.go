// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/store/options"
	"net/http"
)

func (t *OptionsTestSuite) TestOptions_Find() {
	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			`"testing"`,
			http.StatusOK,
			"Successfully obtained option with name: test",
			func(m *mocks.Repository) {
				m.On("Find", "test").Return("testing", nil)
			},
			"/options/test",
		},
		"Not Found": {
			`{}`,
			http.StatusOK,
			"no option found",
			func(m *mocks.Repository) {
				m.On("Find", "test").Return(nil, &errors.Error{Code: errors.NOTFOUND, Message: "no option found"})
			},
			"/options/test",
		},
		"Internal Error": {
			`{}`,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("Find", "test").Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/options/test",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/options/:name", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Find(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
