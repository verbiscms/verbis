// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"fmt"
	mocks "github.com/ainsleyclark/verbis/api/mocks/publisher"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
)

func (t *SEOTestSuite) TestSEO_SitemapXSL() {
	tt := map[string]struct {
		want    interface{}
		status  int
		content string
		mock    func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context)
	}{
		"Success": {
			testString,
			http.StatusOK,
			"application/xml; charset=utf-8",
			func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context) {
				ms.On("XSL", mock.Anything).Return(*t.bytes, nil)
			},
		},
		"Fail": {
			testString,
			http.StatusNotFound,
			"text/html",
			func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context) {
				ms.On("XSL", mock.Anything).Return(nil, fmt.Errorf("error"))
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(http.StatusNotFound, "text/html", []byte(testString))
				})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/sitemaps/posts/map.xml", "/sitemaps/:resource/:map", nil, func(ctx *gin.Context) {
				t.SetupSitemap(test.mock, ctx).SiteMapXSL(ctx, true)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
