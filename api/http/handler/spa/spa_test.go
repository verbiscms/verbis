// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spa

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	app "github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/logger"
	mocks "github.com/verbiscms/verbis/api/mocks/publisher"
	mockFS "github.com/verbiscms/verbis/api/mocks/verbisfs"
	"github.com/verbiscms/verbis/api/test"
	"github.com/verbiscms/verbis/api/verbisfs"
	"io/ioutil"
	"net/http"
	"testing"
)

// SPATestSuite defines the helper used for SPA
// testing.
type SPATestSuite struct {
	test.HandlerSuite
}

// TestSPA
//
// Assert testing has begun.
func TestSPA(t *testing.T) {
	suite.Run(t, &SPATestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a SPA handler for testing.
func (t *SPATestSuite) Setup(mf func(m *mocks.Publisher, mfs *mockFS.FS, ctx *gin.Context), ctx *gin.Context) *SPA {
	logger.SetOutput(ioutil.Discard)
	m := &mocks.Publisher{}
	mfs := &mockFS.FS{}
	if mf != nil {
		mf(m, mfs, ctx)
	}
	return &SPA{
		Deps: &deps.Deps{
			FS: &verbisfs.FileSystem{
				SPA: mfs,
			},
			Installed: true,
		},
		publisher: m,
	}
}

func (t *SPATestSuite) TestSPA() {
	tt := map[string]struct {
		want   string
		status int
		url    string
		mock   func(m *mocks.Publisher, mfs *mockFS.FS, ctx *gin.Context)
	}{
		"Success File": {
			"test",
			http.StatusOK,
			"/images/gopher.svg",
			func(m *mocks.Publisher, mfs *mockFS.FS, ctx *gin.Context) {
				mfs.On("ReadFile", "/images/gopher.svg").Return([]byte("test"), nil)
			},
		},
		"Not Found File": {
			"test",
			http.StatusNotFound,
			"/images/wrongpath.svg",
			func(m *mocks.Publisher, mfs *mockFS.FS, ctx *gin.Context) {
				mfs.On("ReadFile", "/images/wrongpath.svg").Return(nil, fmt.Errorf("error"))
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(http.StatusNotFound, "text/html", []byte("test"))
				})
			},
		},
		"Success Page": {
			"test",
			http.StatusOK,
			"/",
			func(m *mocks.Publisher, mfs *mockFS.FS, ctx *gin.Context) {
				mfs.On("ReadFile", "/index.html").Return([]byte("test"), nil)
			},
		},
		"Not Found Page": {
			"test",
			http.StatusNotFound,
			"/",
			func(m *mocks.Publisher, mfs *mockFS.FS, ctx *gin.Context) {
				mfs.On("ReadFile", "/index.html").Return(nil, fmt.Errorf("error"))
				m.On("NotFound", ctx).Run(func(args mock.Arguments) {
					ctx.Data(http.StatusNotFound, "text/html", []byte("test"))
				})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, app.AdminPath+test.url, "*any", nil, func(ctx *gin.Context) {
				spa := t.Setup(test.mock, ctx)
				spa.Serve(ctx)
			})
			t.Equal(test.status, t.Status())
			t.Equal(test.want, t.Recorder.Body.String())
			t.Reset()
		})
	}
}
