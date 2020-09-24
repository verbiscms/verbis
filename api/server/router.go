package server

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"strconv"
)

type Server struct {
	*gin.Engine
}

func New() (*Server, error) {

	// Force log's color
	gin.ForceConsoleColor()

	// Set mode depending on
	gin.SetMode(gin.ReleaseMode)

	// Remove from console if not super admin
	if !api.SuperAdmin {
		gin.DefaultWriter = ioutil.Discard
	}

	// New router
	r := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	//r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Set template engine
	r.HTMLRender = ginview.New(goview.Config{
		Root:      paths.Theme(),
		Extension: config.Template.FileExtension,
		Master:    "/layouts/main",
		Partials:  []string{},
		DisableCache: true,
		Funcs: template.FuncMap{},
	})

	// Instantiate the server.
	s := &Server{
		r,
	}

	return s, nil
}

// Serve the app
func (s *Server) ListenAndServe(port int) error {
	passedPort := strconv.Itoa(port)

	err := s.Run(":" + passedPort)
	if err != nil {
		return err
	}

	return nil
}
