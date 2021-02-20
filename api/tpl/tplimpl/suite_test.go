// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tplimpl

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

type TplTestSuite struct {
	suite.Suite
	apiPath   string
	logWriter bytes.Buffer
}

func (t *TplTestSuite) BeforeTest(suiteName, testName string) {
	b := bytes.Buffer{}
	t.logWriter = b
	log.SetOutput(&t.logWriter)
	t.SetApiPath()
}
func TestTpl(t *testing.T) {
	suite.Run(t, new(TplTestSuite))
}

func (t *TplTestSuite) SetApiPath() {
	wd, err := os.Getwd()
	t.NoError(err)
	t.apiPath = filepath.Join(filepath.Dir(wd), "..")
}

func (t *TplTestSuite) Reset() {
	t.logWriter.Reset()
}

func (t *TplTestSuite) Setup() (*TemplateManager, *gin.Context, *domain.PostData) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	ctx.Request, _ = http.NewRequest("GET", "/page", nil)
	engine.Use(location.Default())

	engine.GET("/page", func(g *gin.Context) {
		ctx = g
		return
	})

	req, err := http.NewRequest("GET", "http://verbiscms.com/page?page=2&foo=bar", nil)
	t.NoError(err)
	engine.ServeHTTP(rr, req)

	os.Setenv("foo", "bar")

	post := &domain.PostData{
		Post: domain.Post{
			Id:           1,
			Slug:         "/page",
			Title:        "My Verbis Page",
			Status:       "published",
			Resource:     nil,
			PageTemplate: "single",
			PageLayout:   "main",
			UserId:       1,
		},
		Fields: []domain.PostField{
			{PostId: 1, Type: "text", Name: "text", Key: "", OriginalValue: "Hello World!"},
		},
	}

	d := &deps.Deps{
		Store:  nil,
		Config: nil,
		Site:   domain.Site{},
		Options: &domain.Options{
			GeneralLocale: "en-gb",
		},
		Paths:   deps.Paths{},
		Theme:   &domain.ThemeConfig{},
		Running: false,
	}

	return &TemplateManager{deps: d}, ctx, post
}
