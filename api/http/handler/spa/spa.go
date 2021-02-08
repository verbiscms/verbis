// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spa

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/render"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

// SPAHandler defines methods for the SPA (Vue) to interact with the server
type SPAHandler interface {
	Serve(g *gin.Context)
}

// SPA defines the handler for the SPA
type SPA struct {
	*deps.Deps
	render render.Renderer
}

// newSpa - Construct
func NewSpa(d *deps.Deps) *SPA {
	return &SPA{
		render: render.NewRender(d),
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
func (c *SPA) Serve(g *gin.Context) {

	path := g.Request.URL.Path

	// If the path is a file
	if strings.Contains(path, ".") {

		path = strings.Replace(path, "/admin", "", -1)
		extensionArr := strings.Split(path, ".")
		extension := extensionArr[len(extensionArr)-1]
		data, err := ioutil.ReadFile(adminPath + path)

		if err != nil {
			// TODO, log here! Error getting admin file
			c.render.NotFound(g)
			return
		}

		contentType := mime.TypeByExtension(extension)
		g.Data(200, contentType, data)

		// Page catching
	} else {
		data, err := ioutil.ReadFile(adminPath + "/index.html")

		if err != nil {
			// TODO, log here! Error getting admin file
			c.render.NotFound(g)
			return
		}

		g.Data(200, "text/html; charset=utf-8", data)
	}
}
