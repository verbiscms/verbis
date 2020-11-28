package spa

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	mocks "github.com/ainsleyclark/verbis/api/mocks/frontend"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// getSPAMock is a helper to obtain a mock SPA controller
// for testing.
func getSPAMock() *SPA {
	mockError := mocks.ErrorHandler{}
	mockError.On("NotFound", mock.Anything, mock.Anything)
	return &SPA{
		config:       config.Configuration{},
		ErrorHandler: &mockError,
	}
}

// spaTest represents the suite of testing methods for SPA routes.
type spaTest struct {
	testing  *testing.T
	recorder *httptest.ResponseRecorder
	gin      *gin.Context
	engine   *gin.Engine
}

// setup helper for SPA routes.
func setup(t *testing.T) *spaTest {
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	g, engine := gin.CreateTestContext(rr)

	return &spaTest{
		testing:  t,
		recorder: rr,
		gin:      g,
		engine:   engine,
	}
}

// TestSPA_Serve - Test serving of files for SPA handler.
func TestSPA_Serve(t *testing.T) {

	// Save current function and restore at the end:
	oldBasePath := basePath
	oldAdminPath := adminPath
	defer func() {
		basePath = oldBasePath
		adminPath = oldAdminPath
	}()

	// Set api path
	wd, err := os.Getwd()
	assert.NoError(t, err)
	apiPath := filepath.Join(filepath.Dir(wd), "../..")

	// Reassign paths
	adminPath = apiPath + "/test/testdata/spa"
	basePath = apiPath + "/test/testdata/spa"
	imagePath := "/images/gopher.svg"
	htmlPath := "/index.html"

	// Test success getting svg file
	t.Run("Success File", func(t *testing.T) {
		rr := setup(t)

		req, err := http.NewRequest("GET", "/admin"+imagePath, nil)
		assert.NoError(t, err)

		rr.engine.GET("/admin"+imagePath, func(g *gin.Context) {
			getSPAMock().Serve(g)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		data, err := ioutil.ReadFile(basePath + imagePath)
		if err != nil {
			fmt.Println(err)
			t.Errorf("could not open file with the path %s", basePath+imagePath)
		}

		assert.Equal(t, data, rr.recorder.Body.Bytes())
		assert.Equal(t, 200, rr.recorder.Code)
		assert.Equal(t, "image/svg+xml", rr.recorder.Header().Get("Content-Type"))
	})

	// Test 404 of file
	t.Run("404 File", func(t *testing.T) {
		rr := setup(t)

		req, err := http.NewRequest("GET", "/admin"+imagePath, nil)
		assert.NoError(t, err)

		rr.engine.GET("/admin/wrongimage.svg", func(g *gin.Context) {
			mockError := mocks.ErrorHandler{}
			mockError.On("NotFound", g, mock.Anything).Run(func(args mock.Arguments) {
				g.AbortWithStatus(404)
			})
			spa := &SPA{
				config:       config.Configuration{},
				ErrorHandler: &mockError,
			}
			spa.Serve(g)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		assert.Equal(t, 404, rr.recorder.Code)
	})

	//Test success getting html file
	t.Run("Success HTML", func(t *testing.T) {
		rr := setup(t)

		req, err := http.NewRequest("GET", "/admin", nil)
		assert.NoError(t, err)

		rr.engine.GET("/admin", func(g *gin.Context) {
			getSPAMock().Serve(g)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		data, err := ioutil.ReadFile(basePath + htmlPath)
		if err != nil {
			t.Errorf("could not open file with the path %s", basePath+htmlPath)
		}

		assert.Equal(t, string(data), rr.recorder.Body.String())
		assert.Equal(t, 200, rr.recorder.Code)
		assert.Equal(t, "text/html; charset=utf-8", rr.recorder.Header().Get("Content-Type"))
	})

	// Test 404 of file
	t.Run("404 HTML", func(t *testing.T) {
		rr := setup(t)

		adminPath = apiPath + "/test/testdata"

		req, err := http.NewRequest("GET", "/admin", nil)
		assert.NoError(t, err)

		rr.engine.GET("/admin", func(g *gin.Context) {
			mockError := mocks.ErrorHandler{}
			mockError.On("NotFound", g, mock.Anything).Run(func(args mock.Arguments) {
				g.AbortWithStatus(404)
			})
			spa := &SPA{
				config:       config.Configuration{},
				ErrorHandler: &mockError,
			}
			spa.Serve(g)
		})
		rr.engine.ServeHTTP(rr.recorder, req)

		assert.Equal(t, 404, rr.recorder.Code)
	})
}
