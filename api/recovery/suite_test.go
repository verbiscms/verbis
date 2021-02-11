// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

type RecoverTestSuite struct {
	suite.Suite
	deps      *deps.Deps
	logWriter bytes.Buffer
	apiPath   string
}

type noStringer struct{}

func TestFields(t *testing.T) {
	suite.Run(t, new(RecoverTestSuite))
}

func (t *RecoverTestSuite) BeforeTest(suiteName, testName string) {
	b := bytes.Buffer{}
	t.logWriter = b
	log.SetOutput(&t.logWriter)
	t.SetPath()
	t.SetDeps()
}

func (t *RecoverTestSuite) SetPath() {
	wd, err := os.Getwd()
	t.NoError(err)
	t.apiPath = filepath.Join(filepath.Dir(wd))
}

func (t *RecoverTestSuite) SetDeps() {
	t.deps = &deps.Deps{
		Paths: deps.Paths{
			Theme: "theme",
		},
		Theme: &domain.ThemeConfig{
			FileExtension: "cms",
			TemplateDir:   "template",
		},
	}
}

func (t *RecoverTestSuite) Reset() {
	t.logWriter.Reset()
}

func (t *RecoverTestSuite) RequestSetup(body io.Reader, cookie *http.Cookie, fn func(ctx *gin.Context)) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(rr)
	engine.Use(location.Default())

	engine.GET("/test", func(g *gin.Context) {
		fn(g)
		return
	})

	request, err := http.NewRequest("GET", "/test?page=test", body)
	t.NoError(err)

	request.Header.Set("header", "test")

	if cookie != nil {
		request.AddCookie(cookie)
	}

	engine.ServeHTTP(rr, request)
}
