// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tplimpl

import (
	"bytes"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	mocks "github.com/verbiscms/verbis/api/mocks/services/site"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

type TplTestSuite struct {
	suite.Suite
	TestPath string
	Logger   bytes.Buffer
}

func (t *TplTestSuite) BeforeTest(suiteName, testName string) {
	b := bytes.Buffer{}
	t.Logger = b
	logger.SetOutput(&t.Logger)
	t.SetAPIPath()
}

func TestTpl(t *testing.T) {
	suite.Run(t, new(TplTestSuite))
}

func (t *TplTestSuite) SetAPIPath() {
	wd, err := os.Getwd()
	t.NoError(err)
	t.TestPath = filepath.Join(wd, "testdata")
}

func (t *TplTestSuite) Reset() {
	t.Logger.Reset()
}

func (t *TplTestSuite) Setup() (*TemplateManager, *gin.Context, *domain.PostDatum) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	ctx.Request, _ = http.NewRequest("GET", "/page", nil)
	engine.Use(location.Default())

	engine.GET("/page", func(g *gin.Context) {
		ctx = g
	})

	req, err := http.NewRequest("GET", "http://verbiscms.com/page?page=2&foo=bar", nil)
	t.NoError(err)
	engine.ServeHTTP(rr, req)

	os.Setenv("foo", "bar")

	post := &domain.PostDatum{
		Post: domain.Post{
			ID:           1,
			Slug:         "/page",
			Title:        "My Verbis Page",
			Status:       "published",
			PageTemplate: "single",
			PageLayout:   "main",
			UserID:       1,
		},
		Fields: domain.PostFields{
			{PostID: 1, Type: "text", Name: "text", Key: "", OriginalValue: "Hello World!"},
		},
	}

	mockSite := &mocks.Service{}
	mockSite.On("Global").Return(domain.Site{})

	d := &deps.Deps{
		Store: nil,
		Site:  mockSite,
		Options: &domain.Options{
			GeneralLocale: "en-gb",
		},
		Paths:   paths.Paths{},
		Config:  &domain.ThemeConfig{},
		Running: false,
	}

	return &TemplateManager{deps: d}, ctx, post
}
