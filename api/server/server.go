// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/middleware"
	"github.com/verbiscms/verbis/api/logger"
	"io/ioutil"
	"net/http"
	"net/http/pprof"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Server struct {
	*gin.Engine
}

func New(d *deps.Deps) *Server {
	// Force log's color
	gin.ForceConsoleColor()

	// Set mode depending on
	if api.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	// Remove default gin write
	gin.DefaultWriter = ioutil.Discard

	// New router
	r := gin.Default()

	server := &Server{r}

	// Global middleware
	r.Use(logger.Middleware())
	r.Use(location.Default())
	r.Use(middleware.Installed(d))
	r.Use(middleware.Redirects(d))
	r.Use(middleware.Proxy(d))

	if d.Installed {
		r.Use(middleware.Setters(d))
		r.Use(middleware.Redirects(d))
	}

	if !api.Production {
		debug := r.Group("/debug/pprof")
		debug.GET("/", pprofHandler(pprof.Index))
		debug.GET("/cmdline", pprofHandler(pprof.Cmdline))
		debug.GET("/profile", pprofHandler(pprof.Profile))
		debug.POST("/symbol", pprofHandler(pprof.Symbol))
		debug.GET("/symbol", pprofHandler(pprof.Symbol))
		debug.GET("/trace", pprofHandler(pprof.Trace))
		debug.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		debug.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		debug.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		debug.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		debug.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		debug.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}

	r.Use(gin.Recovery())

	// Set up Gzip compression
	server.setupGzip(d)

	// Instantiate the server.
	return server
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// ListenAndServe runs Verbis on a given port
// Returns errors.INVALID if the server could not start
func (s *Server) ListenAndServe(port int) error {
	const op = "Router.ListenAndServe"

	passedPort := strconv.Itoa(port)

	err := s.Run(":" + passedPort)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Could not start Verbis on the port %d", port), Operation: op, Err: err}
	}

	return nil
}

// setupGzip - inits the default gzip compression for the server bt
// looking for the options & setting.
func (s *Server) setupGzip(d *deps.Deps) {
	const op = "router.ListenAndServe"

	options := d.Options

	// Bail if there is no
	if !options.Gzip {
		return
	}

	// Set the default compression & check options
	compression := gzip.DefaultCompression
	switch options.GzipCompression {
	case "best-compression":
		compression = gzip.BestCompression
	case "best-speed":
		compression = gzip.BestSpeed
	}

	// If the use paths is not set, use the excluded extensions
	// or use the excluded paths.
	if !options.GzipUsePaths {
		if len(options.GzipExcludedExtensions) > 0 {
			s.Use(gzip.Gzip(compression, gzip.WithExcludedExtensions(options.GzipExcludedExtensions)))
			return
		}
	} else {
		if len(options.GzipExcludedPaths) > 0 {
			s.Use(gzip.Gzip(compression, gzip.WithExcludedPaths(options.GzipExcludedPaths)))
			return
		}
	}

	s.Use(gzip.Gzip(compression))
}
