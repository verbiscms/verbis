// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *OptionsTestSuite) TestOptions_Find() {
	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(m *mocks.OptionsRepository)
		url     string
	}{
		"Success": {
			`"testing"`,
			http.StatusOK,
			"Successfully obtained option with name: test",
			func(m *mocks.OptionsRepository) {
				m.On("GetByName", "test").Return("testing", nil)
			},
			"/options/test",
		},
		"Not Found": {
			`{}`,
			http.StatusOK,
			"no option found",
			func(m *mocks.OptionsRepository) {
				m.On("GetByName", "test").Return(nil, &errors.Error{Code: errors.NOTFOUND, Message: "no option found"})
			},
			"/options/test",
		},
		"Internal Error": {
			`{}`,
			http.StatusInternalServerError,
			"config",
			func(m *mocks.OptionsRepository) {
				m.On("GetByName", "test").Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "config"})
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
