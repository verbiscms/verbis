// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/http/handler/frontend"
	storeMocks "github.com/ainsleyclark/verbis/api/mocks/models"
	mocks "github.com/ainsleyclark/verbis/api/mocks/render"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// getPublicMock is a helper to obtain a mock Public
// handler for testing.
func getPublicMock(r Renderer, gin *gin.Context) *frontend.Public {
	mockError := mocks.ErrorHandler{}
	mockError.On("NotFound", gin).Run(func(args mock.Arguments) {
		gin.AbortWithStatus(404)
		return
	})
	return &frontend.Public{
		config:       config.Configuration{},
		ErrorHandler: &mockError,
		render:       r,
	}
}

// spaTest represents the suite of testing methods for SPA routes.
type frontendTest struct {
	testing  *testing.T
	recorder *httptest.ResponseRecorder
	gin      *gin.Context
	engine   *gin.Engine
	apiPath  string
}

// setup helper for frontend routes.
func setup(t *testing.T) *frontendTest {
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	g, engine := gin.CreateTestContext(rr)

	// Set api path
	wd, err := os.Getwd()
	assert.NoError(t, err)
	apiPath := filepath.Join(filepath.Dir(wd), "../..")

	return &frontendTest{
		testing:  t,
		recorder: rr,
		gin:      g,
		engine:   engine,
		apiPath:  apiPath,
	}
}

// Test_NewFrontend - Test construct
func Test_NewFrontend(t *testing.T) {

	optsMock := storeMocks.OptionsRepository{}
	optsMock.On("GetStruct").Return(domain.Options{})

	siteMock := storeMocks.SiteRepository{}
	siteMock.On("GetThemeConfig").Return(domain.ThemeConfig{}, nil)

	store := models.Store{
		Options: &optsMock,
		Site:    &siteMock,
	}
	config := config.Configuration{}
	want := &frontend.Public{
		store:        &store,
		config:       config,
		render:       NewRender(&store, config),
		ErrorHandler: &Errors{},
	}

	got := frontend.NewPublic(&store, config)
	assert.ObjectsAreEqual(got, want)
}

// TestPublic_GetUploads - Test serving of uploads
func TestPublic_GetUploads(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		rr := setup(t)

		image := "/test/testdata/images/gopher.svg"

		req, err := http.NewRequest("GET", "/uploads/gopher.svg", nil)
		assert.NoError(t, err)

		file, err := ioutil.ReadFile(rr.apiPath + image)
		assert.NoError(t, err)
		mimeType := mime.TypeByExtension(strings.ReplaceAll(filepath.Ext(image), ".", ""))

		renderMock := mocks.Renderer{}
		renderMock.On("Upload", rr.gin).Return(&mimeType, &file, nil)

		rr.engine.GET("/uploads/*any", func(g *gin.Context) {
			getPublicMock(&renderMock, rr.gin).GetUploads(rr.gin)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		assert.Equal(t, file, rr.recorder.Body.Bytes())
		assert.Equal(t, "image/svg+xml", rr.recorder.Header().Get("Content-Type"))
		assert.Equal(t, 200, rr.recorder.Code)
	})

	t.Run("Not Found", func(t *testing.T) {
		rr := setup(t)

		req, err := http.NewRequest("GET", "/uploads/gopher.svg", nil)
		assert.NoError(t, err)

		renderMock := mocks.Renderer{}
		renderMock.On("Upload", rr.gin).Return(nil, nil, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})

		rr.engine.GET("/uploads/*any", func(g *gin.Context) {
			getPublicMock(&renderMock, rr.gin).GetUploads(rr.gin)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		assert.Equal(t, 404, rr.recorder.Code)
	})
}

// TestPublic_GetAssets - Test serving of assets (under theme path)
func TestPublic_GetAssets(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		rr := setup(t)

		image := "/test/testdata/images/gopher.svg"

		req, err := http.NewRequest("GET", "/uploads/gopher.svg", nil)
		assert.NoError(t, err)

		file, err := ioutil.ReadFile(rr.apiPath + image)
		assert.NoError(t, err)
		mimeType := mime.TypeByExtension(strings.ReplaceAll(filepath.Ext(image), ".", ""))

		renderMock := mocks.Renderer{}
		renderMock.On("Asset", rr.gin).Return(&mimeType, &file, nil)

		rr.engine.GET("/uploads/*any", func(g *gin.Context) {
			getPublicMock(&renderMock, rr.gin).GetAssets(rr.gin)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		assert.Equal(t, file, rr.recorder.Body.Bytes())
		assert.Equal(t, "image/svg+xml", rr.recorder.Header().Get("Content-Type"))
		assert.Equal(t, 200, rr.recorder.Code)
	})

	t.Run("Not Found", func(t *testing.T) {
		rr := setup(t)

		req, err := http.NewRequest("GET", "/uploads/gopher.svg", nil)
		assert.NoError(t, err)

		renderMock := mocks.Renderer{}
		renderMock.On("Asset", rr.gin).Return(nil, nil, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})

		rr.engine.GET("/uploads/*any", func(g *gin.Context) {
			getPublicMock(&renderMock, rr.gin).GetAssets(rr.gin)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		assert.Equal(t, 404, rr.recorder.Code)
	})
}

// TestPublic_Serve - Test serving of pages
func TestPublic_Serve(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		rr := setup(t)

		html := "/test/testdata/html/index.html"

		req, err := http.NewRequest("GET", "/page", nil)
		assert.NoError(t, err)

		file, err := ioutil.ReadFile(rr.apiPath + html)
		assert.NoError(t, err)
		mimeType := mime.TypeByExtension(strings.ReplaceAll(filepath.Ext(html), ".", ""))

		renderMock := mocks.Renderer{}
		renderMock.On("Page", rr.gin).Return(file, nil)

		rr.engine.GET("/page", func(g *gin.Context) {
			getPublicMock(&renderMock, rr.gin).Serve(rr.gin)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		assert.Equal(t, file, rr.recorder.Body.Bytes())
		assert.Equal(t, mimeType, rr.recorder.Header().Get("Content-Type"))
		assert.Equal(t, 200, rr.recorder.Code)
	})

	t.Run("Not Found", func(t *testing.T) {
		rr := setup(t)

		req, err := http.NewRequest("GET", "/page", nil)
		assert.NoError(t, err)

		renderMock := mocks.Renderer{}
		renderMock.On("Page", rr.gin).Return(nil, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})

		rr.engine.GET("/page", func(g *gin.Context) {
			getPublicMock(&renderMock, rr.gin).Serve(rr.gin)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		assert.Equal(t, 404, rr.recorder.Code)
	})
}
