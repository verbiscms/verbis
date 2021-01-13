package middleware

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRedirects - Test redirects from options are working correctly
// with correct code & locationjson.
func Test_Redirects(t *testing.T) {

	tt := map[string]struct {
		status      int
		url         string
		redirectUrl string
		mock        func(u *mocks.OptionsRepository)
	}{
		"No Redirects": {
			status: 200,
			url:    "/page",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetStruct").Return(domain.Options{
					SeoRedirects: nil,
				})
			},
		},
		"300": {
			status:      300,
			url:         "/page/test",
			redirectUrl: "/page",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetStruct").Return(domain.Options{
					SeoRedirects: []domain.Redirect{
						{To: "/page", From: "http://localhost:8080/page/test", Code: 300},
					},
				})
			},
		},
		"301": {
			status:      301,
			url:         "/page/test",
			redirectUrl: "/page",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetStruct").Return(domain.Options{
					SeoRedirects: []domain.Redirect{
						{To: "/page", From: "http://localhost:8080/page/test", Code: 301},
					},
				})
			},
		},
		"302": {
			status:      302,
			url:         "/page/test",
			redirectUrl: "/page",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetStruct").Return(domain.Options{
					SeoRedirects: []domain.Redirect{
						{To: "/page", From: "http://localhost:8080/page/test", Code: 302},
					},
				})
			},
		},
		"303": {
			status:      303,
			url:         "/page/test",
			redirectUrl: "/page",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetStruct").Return(domain.Options{
					SeoRedirects: []domain.Redirect{
						{To: "/page", From: "http://localhost:8080/page/test", Code: 303},
					},
				})
			},
		},
		"304": {
			status:      304,
			url:         "/page/test",
			redirectUrl: "/page",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetStruct").Return(domain.Options{
					SeoRedirects: []domain.Redirect{
						{To: "/page", From: "http://localhost:8080/page/test", Code: 304},
					},
				})
			},
		},
		"307": {
			status:      307,
			url:         "/page/test",
			redirectUrl: "/page",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetStruct").Return(domain.Options{
					SeoRedirects: []domain.Redirect{
						{To: "/page", From: "http://localhost:8080/page/test", Code: 307},
					},
				})
			},
		},
		"308": {
			status:      308,
			url:         "/page/test",
			redirectUrl: "/page",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetStruct").Return(domain.Options{
					SeoRedirects: []domain.Redirect{
						{To: "/page", From: "http://localhost:8080/page/test", Code: 308},
					},
				})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			rr := httptest.NewRecorder()
			g, engine := gin.CreateTestContext(rr)
			engine.Use(location.Default())

			mock := &mocks.OptionsRepository{}
			test.mock(mock)
			engine.Use(Redirects(mock))

			engine.GET("/page", func(context *gin.Context) {
				g.String(200, "verbis")
				return
			})

			req, err := http.NewRequest("GET", test.url, nil)
			assert.NoError(t, err)

			engine.ServeHTTP(rr, req)

			assert.Equal(t, test.status, rr.Result().StatusCode)

			if test.redirectUrl != "" {
				loc, err := rr.Result().Location()
				assert.NoError(t, err)
				assert.Equal(t, test.redirectUrl, loc.Path)
			}
		})
	}
}
