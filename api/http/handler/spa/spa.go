// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spa

import (
	"github.com/gin-gonic/gin"
	app "github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/common/mime"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/publisher"
	"net/http"
	"strings"
)

// Handler defines methods for the SPA (Vue) to interact with the server.
type Handler interface {
	Serve(ctx *gin.Context)
}

// SPA defines the handler for all SPA routes.
type SPA struct {
	*deps.Deps
	publisher publisher.Publisher
}

// New
//
// Creates a new spa handler.
func New(d *deps.Deps) *SPA {
	s := &SPA{
		Deps: d,
	}
	if d.Installed {
		s.publisher = publisher.NewRender(d)
	}
	return s
}

// Operation for Errors
const op = "SPA.Serve"

// Serve
//
// Serve all of the administrator & operator assets and serve the
// file extension based on the content type.
func (s *SPA) Serve(ctx *gin.Context) {
	urlPath := ctx.Request.URL.Path

	// Check if the urlPath is the installed urlPath and
	// the app is installed, if it is the user
	// should not be there, abort.
	if urlPath == app.AdminInstallPath && s.Installed {
		s.publisher.NotFound(ctx)
		return
	}

	// If the urlPath is a file
	if strings.Contains(urlPath, ".") {
		s.file(urlPath, ctx)
		return
	}

	// If the urlPath is index.html
	s.page(ctx)
}

// file
//
// Returns any files for the SPA.
func (s *SPA) file(path string, ctx *gin.Context) {
	file := strings.ReplaceAll(path, app.AdminPath, "")
	extensionArr := strings.Split(file, ".")
	extension := extensionArr[len(extensionArr)-1]

	data, err := s.FS.SPA.ReadFile(file)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error reading admin admin file with the path: " + path, Operation: op, Err: err}).Error()
		if s.Installed {
			s.publisher.NotFound(ctx)
		}
		return
	}

	contentType := mime.TypeByExtension(extension)
	ctx.Data(http.StatusOK, contentType, data)
}

// page
//
// Returns the index.html in bytes
func (s *SPA) page(ctx *gin.Context) {
	data, err := s.FS.SPA.ReadFile("/index.html")

	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error reading admin admin file with the path: " + "index.html", Operation: op, Err: err}).Error()
		if s.Installed {
			s.publisher.NotFound(ctx)
		}
		return
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", data)
}
