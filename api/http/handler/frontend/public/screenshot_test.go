// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	publisher "github.com/verbiscms/verbis/api/mocks/publisher"
	theme "github.com/verbiscms/verbis/api/mocks/services/theme"
	"net/http"
)

func (t *PublicTestSuite) TestPublic_Screenshot() {
	tt := map[string]struct {
		want    interface{}
		status  int
		content string
		mock    func(m *publisher.Publisher, mt *theme.Repository, ctx *gin.Context)
		url     string
	}{
		"Success": {
			testString,
			http.StatusOK,
			"image/png",
			func(m *publisher.Publisher, mt *theme.Repository, ctx *gin.Context) {
				mt.On("Screenshot", "theme", "screenshot.jpg").Return(*t.bytes, domain.Mime("image/png"), nil)
			},
			"/theme/screenshot.jpg",
		},
		"Error": {
			testString,
			http.StatusNotFound,
			"text/html",
			func(m *publisher.Publisher, mt *theme.Repository, ctx *gin.Context) {
				mt.On("Screenshot", "theme", "screenshot.jpg").Return(*t.bytes, domain.Mime("image/png"), &errors.Error{Code: errors.NOTFOUND})
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(http.StatusNotFound, "text/html", []byte(testString))
				})
			},
			"/theme/screenshot.jpg",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, test.url, "/:theme/:file", nil, func(ctx *gin.Context) {
				t.SetupTheme(test.mock, ctx).Screenshot(ctx)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
