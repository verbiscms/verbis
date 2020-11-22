package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSpaController_Serve(t *testing.T) {

	test := newResponseRecorder(t)

	// Save current function and restore at the end:
	old := adminPath
	defer func() { adminPath = old }()

	adminPath = "/Users/Hello"

	spaController := SpaController{}
	test.engine.GET("/admin/file.jpg", func(g *gin.Context) {
		spaController.Serve(g)
	})

	req, _ := http.NewRequest("GET", "/admin/file.jpg", nil)

	w := httptest.NewRecorder()
	test.engine.ServeHTTP(w, req)

}
