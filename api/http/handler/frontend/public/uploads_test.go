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

func (t *PublicTestSuite) TestPublic_Uploads() {
	tt := map[string]struct {
		want    interface{}
		status  int
		content string
		mock    func(m *mocks.Publisher, ctx *gin.Context)
	}{
		"Success": {
			testString,
			http.StatusOK,
			"image/png",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Upload", ctx).Return("image/png", t.bytes, nil)
			},
		},
		"Fail": {
			testString,
			404,
			"text/html",
			func(m *mocks.Publisher, ctx *gin.Context) {
				m.On("Upload", ctx).Return("", nil, fmt.Errorf("error"))
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(404, "text/html", []byte(testString))
				})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/uploads/test.jpg", "/uploads/:file", nil, func(ctx *gin.Context) {
				t.Setup(test.mock, ctx).Uploads(ctx)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
