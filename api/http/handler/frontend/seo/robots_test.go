// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/publisher"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
)

func (t *SEOTestSuite) TestSEO_Robots() {
	tt := map[string]struct {
		want    interface{}
		status  int
		content string
		options *domain.Options
		mock    func(m *mocks.Publisher, ctx *gin.Context)
	}{
		"Success": {
			"test",
			http.StatusOK,
			"text/plain",
			&domain.Options{
				SeoRobotsServe: true,
				SeoRobots:      "test",
			},
			nil,
		},
		"Disabled Serve": {
			"test",
			http.StatusNotFound,
			"text/html",
			&domain.Options{
				SeoRobotsServe: false,
			},
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(http.StatusNotFound, "text/html", []byte("test"))
				})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/robots.txt", "robots.txt", nil, func(ctx *gin.Context) {
				seo := t.Setup(test.mock, ctx)
				seo.Deps.Options = test.options
				seo.Robots(ctx)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
