// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spa

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/publisher"
	"github.com/gin-gonic/gin"
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
	return &SPA{
		Deps:      d,
		publisher: publisher.NewRender(d),
	}
}

// Operation for Errors
const op = "SPA.Serve"

// Serve
//
// Serve all of the administrator & operator assets and serve the
// file extension based on the content type.
func (s *SPA) Serve(ctx *gin.Context) {
	path := ctx.Request.URL.Path

	// If the path is a file
	if strings.Contains(path, ".") {
		s.file(path, ctx)
		return
	}

	// If the path is index.html
	s.page(ctx)
}

// file
//
// Returns any files for the SPA.
func (s *SPA) file(path string, ctx *gin.Context) {
	file := strings.ReplaceAll(path, "/admin", "")
	extensionArr := strings.Split(file, ".")
	extension := extensionArr[len(extensionArr)-1]

	data, err := s.FS.SPA.ReadFile(file)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error reading admin admin file with the path: " + path, Operation: op, Err: err}).Error()
		s.publisher.NotFound(ctx)
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
		s.publisher.NotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", data)
}
