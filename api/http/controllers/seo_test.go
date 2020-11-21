package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/frontend"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)


func TestSEOController_Robots(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func (t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)

		seoSuccess := SEOController{
			options: domain.Options{
				SeoRobotsServe: true,
				SeoRobots:  "test",
			},
		}

		seoSuccess.Robots(g)

		assert.Equal(t, 200, r.Code)
		assert.Equal(t, r.Body.String(), "test")
		assert.Equal(t, r.Header().Get("Content-Type"), "text/plain")
	})

	t.Run("Error", func (t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)
		errorMock := &mocks.ErrorHandler{}
		errorMock.On("NotFound", g, config.Configuration{}).Return("error")

		seoError := SEOController{
			options: domain.Options{
				SeoRobotsServe: false,
			},
			ErrorHandler: errorMock,
		}

		seoError.Robots(g)

		assert.Equal(t, r.Body.String(), "")
		errorMock.AssertExpectations(t)
	})
}

func TestSEOController_SiteMapIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func (t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)
		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetIndex").Return([]byte("test"), nil)

		seoSuccess := SEOController{
			sitemap: sitemapMock,
		}

		seoSuccess.SiteMapIndex(g)

		assert.Equal(t, 200, r.Code)
		assert.Equal(t, r.Body.String(), "test")
		assert.Equal(t, r.Header().Get("Content-Type"), "application/xml; charset=utf-8")
		sitemapMock.AssertExpectations(t)
	})

	t.Run("Fail", func (t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)

		errorMock := &mocks.ErrorHandler{}
		errorMock.On("NotFound", g, config.Configuration{}).Return("error")

		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetIndex").Return(nil, fmt.Errorf("error"))

		seoError := SEOController{
			sitemap: sitemapMock,
			ErrorHandler: errorMock,
		}

		seoError.SiteMapIndex(g)

		assert.Equal(t, r.Body.String(), "")
		errorMock.AssertExpectations(t)
		sitemapMock.AssertExpectations(t)
	})
}

func TestSEOController_SiteMapResource(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func (t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)
		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetPages", mock.Anything).Return([]byte("test"), nil)

		seoSuccess := SEOController{
			sitemap: sitemapMock,
		}

		seoSuccess.SiteMapResource(g)

		assert.Equal(t, 200, r.Code)
		assert.Equal(t, r.Body.String(), "test")
		assert.Equal(t, r.Header().Get("Content-Type"), "application/xml; charset=utf-8")
		sitemapMock.AssertExpectations(t)
	})

	t.Run("Fail", func (t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)

		errorMock := &mocks.ErrorHandler{}
		errorMock.On("NotFound", g, config.Configuration{}).Return("error")

		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetPages", mock.Anything).Return(nil, fmt.Errorf("error"))

		seoError := SEOController{
			sitemap: sitemapMock,
			ErrorHandler: errorMock,
		}

		seoError.SiteMapResource(g)

		assert.Equal(t, r.Body.String(), "")
		errorMock.AssertExpectations(t)
		sitemapMock.AssertExpectations(t)
	})
}


func TestSEOController_SiteMapXSL(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func (t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)
		sitemapMock := &mocks.SiteMapper{}

		sitemapMock.On("GetXSL", mock.Anything).Return([]byte("test"), nil)

		seoSuccess := SEOController{
			sitemap: sitemapMock,
		}

		seoSuccess.SiteMapXSL(g, true)

		assert.Equal(t, 200, r.Code)
		assert.Equal(t, r.Body.String(), "test")
		assert.Equal(t, r.Header().Get("Content-Type"), "application/xml; charset=utf-8")
		sitemapMock.AssertExpectations(t)
	})

	t.Run("Fail", func (t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)

		errorMock := &mocks.ErrorHandler{}
		errorMock.On("NotFound", g, config.Configuration{}).Return("error")

		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetXSL", mock.Anything).Return(nil, fmt.Errorf("error"))

		seoError := SEOController{
			sitemap: sitemapMock,
			ErrorHandler: errorMock,
		}

		seoError.SiteMapXSL(g, true)

		assert.Equal(t, r.Body.String(), "")
		errorMock.AssertExpectations(t)
		sitemapMock.AssertExpectations(t)
	})
}