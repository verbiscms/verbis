// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"fmt"
	mocks "github.com/ainsleyclark/verbis/api/mocks/publisher"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
)

func (t *PublicTestSuite) TestPublic_Assets() {
	tt := map[string]struct {
		want    interface{}
		status  int
		content string
		mock    func(m *mocks.Publisher, ctx *gin.Context)
	}{
		"Success": {
			testString,
			200,
			"image/png",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Asset", ctx).Return("image/png", t.bytes, nil)
			},
		},
		"Fail": {
			testString,
			404,
			"text/html",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Asset", ctx).Return("", nil, fmt.Errorf("error"))
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(404, "text/html", []byte(testString))
				})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/assets/test.jpg", "/assets/:file", nil, func(ctx *gin.Context) {
				t.Setup(test.mock, ctx).Assets(ctx)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
