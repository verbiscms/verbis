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

func (t *SEOTestSuite) TestSEO_SitemapResource() {
	tt := map[string]struct {
		want    interface{}
		status  int
		content string
		mock    func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context)
	}{
		"Success": {
			testString,
			200,
			"application/xml; charset=utf-8",
			func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context) {
				ms.On("Pages", mock.Anything).Return(*t.bytes, nil)
			},
		},
		"Fail": {
			testString,
			404,
			"text/html",
			func(m *mocks.Publisher, ms *mocks.SiteMapper, ctx *gin.Context) {
				ms.On("Pages", mock.Anything).Return(nil, fmt.Errorf("error"))
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(404, "text/html", []byte(testString))
				})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/sitemaps/posts/map.xml", "/sitemaps/:resource/:map", nil, func(ctx *gin.Context) {
				t.SetupSitemap(test.mock, ctx).SiteMapResource(ctx)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
