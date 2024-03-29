// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/publisher"
	"net/http"
)

func (t *PublicTestSuite) TestPublic_Serve() {
	tt := map[string]struct {
		want    interface{}
		status  int
		content string
		mock    func(m *mocks.Publisher, ctx *gin.Context)
	}{
		"Success": {
			testString,
			http.StatusOK,
			"text/html",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Page", ctx).Return(*t.bytes, nil)
			},
		},
		"Not Found": {
			testString,
			http.StatusNotFound,
			"text/html",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Page", ctx).Return(nil, &errors.Error{Code: errors.NOTFOUND})
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(http.StatusNotFound, "text/html", []byte(testString))
				})
			},
		},
		"Internal": {
			testString,
			http.StatusInternalServerError,
			"text/html",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Page", ctx).Return(*t.bytes, &errors.Error{Code: errors.INTERNAL})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/page", "/page", nil, func(ctx *gin.Context) {
				t.Setup(test.mock, ctx).Serve(ctx)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
