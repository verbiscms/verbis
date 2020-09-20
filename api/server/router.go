package server

import (
	"cms/api/config"
	"cms/api/environment"
	"cms/api/helpers/paths"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

type Server struct {
	*gin.Engine
}


func New() (*Server, error) {

	// Force log's color
	gin.ForceConsoleColor()

	// Set mode depending on
	ginMode := "debug"
	if environment.IsProduction() || !environment.IsDebug() {
		ginMode = "release"
	}
	gin.SetMode(ginMode)

	// New router
	r := gin.Default()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

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
