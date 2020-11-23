package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
	mocks "github.com/ainsleyclark/verbis/api/mocks/frontend"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"testing"
)

// getSPAMock is a helper to obtain a mock SPA controller
// for testing.
func getSPAMock() *SpaController {
	mockError := mocks.ErrorHandler{}
	mockError.On("NotFound", mock.Anything, mock.Anything).Return(mock.Anything)
	return &SpaController{
		config:       config.Configuration{},
		ErrorHandler: &mockError,
	}
}

// TestSpaController_Serve - Test serving of files for SPA handler.
func TestSpaController_Serve(t *testing.T) {

	// Save current function and restore at the end:
	oldBasePath := basePath
	oldAdminPath := adminPath
	defer func() {
		basePath = oldBasePath
		adminPath = oldAdminPath
	}()

	// Reassign paths
	adminPath = "/Users/ainsley/Desktop/Reddico/apis/verbis/api/test/testdata"
	basePath = "/Users/ainsley/Desktop/Reddico/apis/verbis/api/test/testdata"
	imagePath := "/images/gopher.svg"

	// Test success getting svg file
	t.Run("Invalid", func(t *testing.T) {
		test := newResponseRecorder(t)

		test.RequestAndServe("GET", "/admin"+imagePath, "/admin"+imagePath, nil, func(g *gin.Context) {
			getSPAMock().Serve(g)
		})

		data, err := ioutil.ReadFile(basePath + imagePath)
		if err != nil {
			t.Errorf("could not open file with the path %s", basePath+adminPath)
		}

		assert.Equal(t, test.recorder.Body.Bytes(), data)
		assert.Equal(t, test.recorder.Code, 200)
		assert.Equal(t, test.recorder.Header().Get("Content-Type"), "image/svg+xml")
	})

	// Test 404 of file
	t.Run("404", func(t *testing.T) {
		test := newResponseRecorder(t)

		test.RequestAndServe("GET", "/admin"+imagePath, "/admin/wrongimage.svg", nil, func(g *gin.Context) {
			getSPAMock().Serve(g)
		})

		assert.Equal(t, test.recorder.Code, 404)
	})

	// Test success getting html file
	t.Run("Invalid", func(t *testing.T) {
		test := newResponseRecorder(t)

		test.RequestAndServe("GET", "/admin/html/index.html", "/admin/html/index.html", nil, func(g *gin.Context) {
			getSPAMock().Serve(g)
		})

		data, err := ioutil.ReadFile(basePath + "/html/index.html")
		if err != nil {
			t.Errorf("could not open file with the path %s", basePath+adminPath)
		}

		assert.Equal(t, test.recorder.Body.String(), string(data))
		assert.Equal(t, test.recorder.Code, 200)
		assert.Equal(t, test.recorder.Header().Get("Content-Type"), "text/html")
	})
}
