package frontend

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	storeMocks "github.com/ainsleyclark/verbis/api/mocks/models"
	mocks "github.com/ainsleyclark/verbis/api/mocks/render"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/render"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

// Test_NewSEO - Test construct
func Test_NewSEO(t *testing.T) {

	optsMock := storeMocks.OptionsRepository{}
	optsMock.On("GetStruct").Return(domain.Options{})

	siteMock := storeMocks.SiteRepository{}
	siteMock.On("GetThemeConfig").Return(domain.ThemeConfig{}, nil)

	store := models.Store{
		Options: &optsMock,
		Site: &siteMock,
	}
	config := config.Configuration{}
	want := &SEO{
		store:  &store,
		config: config,
		sitemap: render.NewSitemap(&store),
		ErrorHandler: &render.Errors{},
	}

	got := NewSEO(&store, config)
	assert.ObjectsAreEqual(got, want)
}

// TestSEOController_Robots - Test robots.txt route
func TestSEOController_Robots(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)

		seoSuccess := SEO{
			options: domain.Options{
				SeoRobotsServe: true,
				SeoRobots:      "test",
			},
		}

		seoSuccess.Robots(g)

		assert.Equal(t, 200, r.Code)
		assert.Equal(t, r.Body.String(), "test")
		assert.Equal(t, r.Header().Get("Content-Type"), "text/plain")
	})

	t.Run("Error", func(t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)
		errorMock := &mocks.ErrorHandler{}
		errorMock.On("NotFound", g, config.Configuration{}).Return("error")

		seoError := SEO{
			options: domain.Options{
				SeoRobotsServe: false,
			},
			ErrorHandler: errorMock,
		}

		seoError.Robots(g)

		assert.Equal(t, r.Body.String(), "")
	})
}

// TestSEOController_SiteMapIndex - Test /sitemap.xml route
func TestSEOController_SiteMapIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)
		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetIndex").Return([]byte("test"), nil)

		seoSuccess := SEO{
			sitemap: sitemapMock,
		}

		seoSuccess.SiteMapIndex(g)

		assert.Equal(t, 200, r.Code)
		assert.Equal(t, r.Body.String(), "test")
		assert.Equal(t, r.Header().Get("Content-Type"), "application/xml; charset=utf-8")
	})

	t.Run("Error", func(t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)

		errorMock := &mocks.ErrorHandler{}
		errorMock.On("NotFound", g, config.Configuration{}).Return("error")

		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetIndex").Return(nil, fmt.Errorf("error"))

		seoError := SEO{
			sitemap:      sitemapMock,
			ErrorHandler: errorMock,
		}

		seoError.SiteMapIndex(g)

		assert.Equal(t, r.Body.String(), "")
		errorMock.AssertExpectations(t)
	})
}

// TestSEOController_SiteMapResource - Test /sitemaps/:resource route
func TestSEOController_SiteMapResource(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)
		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetPages", mock.Anything).Return([]byte("test"), nil)

		seoSuccess := SEO{
			sitemap: sitemapMock,
		}

		seoSuccess.SiteMapResource(g)

		assert.Equal(t, 200, r.Code)
		assert.Equal(t, r.Body.String(), "test")
		assert.Equal(t, r.Header().Get("Content-Type"), "application/xml; charset=utf-8")
	})

	t.Run("Error", func(t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)

		errorMock := &mocks.ErrorHandler{}
		errorMock.On("NotFound", g, config.Configuration{}).Return("error")

		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetPages", mock.Anything).Return(nil, fmt.Errorf("error"))

		seoError := SEO{
			sitemap:      sitemapMock,
			ErrorHandler: errorMock,
		}

		seoError.SiteMapResource(g)

		assert.Equal(t, r.Body.String(), "")
		errorMock.AssertExpectations(t)
	})
}

// TestSEOController_SiteMapXSL - Test .xsl files for styling XML.
func TestSEOController_SiteMapXSL(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)
		sitemapMock := &mocks.SiteMapper{}

		sitemapMock.On("GetXSL", mock.Anything).Return([]byte("test"), nil)

		seoSuccess := SEO{
			sitemap: sitemapMock,
		}

		seoSuccess.SiteMapXSL(g, true)

		assert.Equal(t, 200, r.Code)
		assert.Equal(t, r.Body.String(), "test")
		assert.Equal(t, r.Header().Get("Content-Type"), "application/xml; charset=utf-8")
	})

	t.Run("Error", func(t *testing.T) {

		r := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(r)

		errorMock := &mocks.ErrorHandler{}
		errorMock.On("NotFound", g, config.Configuration{}).Return("error")

		sitemapMock := &mocks.SiteMapper{}
		sitemapMock.On("GetXSL", mock.Anything).Return(nil, fmt.Errorf("error"))

		seoError := SEO{
			sitemap:      sitemapMock,
			ErrorHandler: errorMock,
		}

		seoError.SiteMapXSL(g, true)

		assert.Equal(t, r.Body.String(), "")
	})
}
