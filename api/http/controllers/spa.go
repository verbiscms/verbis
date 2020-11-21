package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

// SpaHandler defines methods for the SPA (Vue) to interact with the server
type SpaHandler interface {
	Serve(g *gin.Context)
}

// SpaController defines the handler for the SPA
type SpaController struct{}

// newSpa - Construct
func newSpa() *SpaController {
	return &SpaController{}
}

// Serve all of the administrator & operator assets and serve the
// file extension based on the content type.
func (c *SpaController) Serve(g *gin.Context) {

	path := g.Request.URL.Path

	fmt.Println(path)

	// If the path is a file
	if strings.Contains(path, ".") {

		path = strings.Replace(path, "/admin", "", -1)
		extensionArr := strings.Split(path, ".")
		extension := extensionArr[len(extensionArr)-1]
		data, _ := ioutil.ReadFile(paths.Admin() + path)
		contentType := mime.TypeByExtension(extension)
		g.Data(200, contentType, data)

		// Page catching
	} else {
		data, _ := ioutil.ReadFile(paths.Admin() + "/index.html")
		g.Data(200, "text/html; charset=utf-8", data)
	}
}
