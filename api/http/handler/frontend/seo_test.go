package frontend

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

// getFieldsMock is a helper to obtain a mock fields controller
// for testing.
//func getSeoMock(m models.SeoMetaRepository) *SEOController {
//	return &SEOController{
//		store: &models.Store{
//			SEOController{}: m,
//		},
//	}
//}

// controllerTest represents the suite of testing methods for controllers.
type controllerTest struct {
	testing  *testing.T
	recorder *httptest.ResponseRecorder
	gin      *gin.Context
	engine   *gin.Engine
}

func TestSEOController_Robots(t *testing.T) {

	gin.SetMode(gin.TestMode)

	tt := map[string]struct {
		want    string
		status  int
		options  domain.Options
	}{
		"Success": {
			want:    `test`,
			status:  200,
			options: domain.Options{
				SeoRobotsServe: true,
				SeoRobots:      "test",
			},
		},
		"Error": {
			want:    `test`,
			status:  200,
			options: domain.Options{
				SeoRobotsServe: false,
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			//rr := newTestSuite(t)
			//mock := &mocks.MediaRepository{}
			//test.mock(mock)
			//
			//rr.RequestAndServe("GET", "/users", "/users", nil, func(g *gin.Context) {
			//	getMediaMock(mock).Get(g)
			//})
			//
			//rr.Run(test.want, test.status, test.message)
		})
	}

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
		errorMock.AssertExpectations(t)
	})
}

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
		sitemapMock.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {

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
		sitemapMock.AssertExpectations(t)
	})
}

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
		sitemapMock.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {

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
		sitemapMock.AssertExpectations(t)
	})
}

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
		sitemapMock.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {

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
		errorMock.AssertExpectations(t)
		sitemapMock.AssertExpectations(t)
	})
}
