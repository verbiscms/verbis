// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/publisher"
	"net/http"
)

func (t *PublicTestSuite) TestPublic_Assets() {
	tt := map[string]struct {
		want    interface{}
		status  int
		content string
		mock    func(m *mocks.Publisher, ctx *gin.Context)
		url     string
	}{
		"Success": {
			testString,
			http.StatusOK,
			"image/png",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Asset", ctx, true).Return(t.bytes, domain.Mime("image/png"), nil)
			},
			"/assets/test.jpg",
		},
		"Fail": {
			testString,
			http.StatusNotFound,
			"text/html",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Asset", ctx, true).Return(nil, domain.Mime(""), fmt.Errorf("error"))
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(http.StatusNotFound, "text/html", []byte(testString))
				})
			},
			"/assets/test.jpg",
		},
		"No WebP": {
			testString,
			http.StatusOK,
			"image/png",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Asset", ctx, false).Return(t.bytes, domain.Mime("image/png"), nil)
			},
			"/assets/test.jpg?webp=false",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, test.url, "/assets/:file", nil, func(ctx *gin.Context) {
				t.Setup(test.mock, ctx).Assets(ctx)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
