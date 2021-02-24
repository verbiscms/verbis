// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spa

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/publisher"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

const (
	// The SPA test dir.
	TestPath = "/test/testdata/spa"
)

// SPATestSuite defines the helper used for SPA
// testing.
type SPATestSuite struct {
	api.HandlerSuite
}

// TestSPA
//
// Assert testing has begun.
func TestSPA(t *testing.T) {
	suite.Run(t, &SPATestSuite{
		HandlerSuite: api.TestSuite(),
	})
}

// Setup
//
// A helper to obtain a SPA handler for testing.
func (t *SPATestSuite) Setup(mf func(m *mocks.Renderer, ctx *gin.Context), admin string, ctx *gin.Context) *SPA {
	wd, err := os.Getwd()
	t.NoError(err)
	apiPath := filepath.Join(filepath.Dir(wd), "../..")

	m := &mocks.Renderer{}
	if mf != nil {
		mf(m, ctx)
	}

	return &SPA{
		Deps: &deps.Deps{
			Paths: deps.Paths{
				API:   apiPath,
				Admin: apiPath + admin,
			},
		},
		Publisher: m,
	}
}

func (t *SPATestSuite) TestSPA() {

	tt := map[string]struct {
		want      string
		status    int
		content   string
		url       string
		adminPath string
		mock      func(m *mocks.Renderer, ctx *gin.Context)
	}{
		"Success File": {
			"/images/gopher.svg",
			200,
			"image/svg+xml",
			"/images/gopher.svg",
			TestPath,
			nil,
		},
		"Not Found File": {
			"test",
			404,
			"text/html",
			"/images/wrongpath.svg",
			TestPath,
			func(m *mocks.Renderer, ctx *gin.Context) {
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(404, "text/html", []byte("test"))
				})
			},
		},
		"Success Page": {
			"/index.html",
			200,
			"text/html; charset=utf-8",
			"/",
			TestPath,
			nil,
		},
		"Not Found Page": {
			"test",
			404,
			"text/html",
			"/",
			"wrong",
			func(m *mocks.Renderer, ctx *gin.Context) {
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(404, "text/html", []byte("test"))
				})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/admin"+test.url, "*any", nil, func(ctx *gin.Context) {
				spa := t.Setup(test.mock, test.adminPath, ctx)
				spa.Serve(ctx)

				data, err := ioutil.ReadFile(spa.Paths.API + test.adminPath + test.want)
				if err == nil {
					test.want = string(data)
				}
			})

			t.Equal(test.status, t.Status())
			t.Equal(test.content, t.ContentType())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
