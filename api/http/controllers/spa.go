package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/frontend"
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
type SpaController struct{
	config  config.Configuration
	frontend.ErrorHandler
}

// newSpa - Construct
func newSpa(config config.Configuration) *SpaController {
	return &SpaController{
		config: config,
		ErrorHandler: &frontend.Errors{},
	}
}

var (
	// Base path of the app
	basePath = paths.Base()
	// SPA path (Vue)
	adminPath = paths.Admin()
)

// Serve all of the administrator & operator assets and serve the
// file extension based on the content type.
func (c *SpaController) Serve(g *gin.Context) {

	path := g.Request.URL.Path

	// If the path is a file
	if strings.Contains(path, ".") {

		path = strings.Replace(path, "/admin", "", -1)
		extensionArr := strings.Split(path, ".")
		extension := extensionArr[len(extensionArr)-1]
		data, err := ioutil.ReadFile(adminPath + path)

		if err != nil {
			// TODO, log here! Error getting admin file
			c.ErrorHandler.NotFound(g, c.config)
			return
		}

		contentType := mime.TypeByExtension(extension)
		g.Data(200, contentType, data)

	// Page catching
	} else {
		data, err := ioutil.ReadFile(adminPath + "/index.html")

		if err != nil {
			// TODO, log here! Error getting admin file
			c.ErrorHandler.NotFound(g, c.config)
		}

		g.Data(200, "text/html; charset=utf-8", data)
	}
}
