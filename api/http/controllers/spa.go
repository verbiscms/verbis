package controllers

import (
	"cms/api/helpers/mime"
	"cms/api/helpers/paths"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

type SpaController struct {
}

type SpaHandler interface {
	Serve(g *gin.Context)
}

// Construct
func newSpa() *SpaController {
	return &SpaController{}
}

// Serve all of the administrator & operator assets and serve the
// file extension based on the content type.
func (c *SpaController) Serve(g *gin.Context) {

	path := g.Request.URL.Path

	// If the path is a file
	if strings.Contains(path, ".") {
		path = strings.Replace(path, "/admin", "", -1)

		extensionArr := strings.Split(path, ".")
		extension := extensionArr[len(extensionArr)-1]

		data, _ := ioutil.ReadFile(paths.Admin() + "/dist" + path)

		contentType := mime.TypeByExtension(extension)

		g.Data(200, contentType, data)

	// Page catching
	} else {
		data, _ := ioutil.ReadFile(paths.Admin() + "/dist/index.html")
		g.Data(200, "text/html; charset=utf-8", data)
	}
}
