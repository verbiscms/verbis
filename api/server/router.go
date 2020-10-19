package server

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

type Server struct {
	*gin.Engine
}

func New() *Server {

	// Force log's color
	gin.ForceConsoleColor()

	// Set mode depending on
	gin.SetMode(gin.ReleaseMode)

	// Remove default gin write
	gin.DefaultWriter = ioutil.Discard

	// New router
	r := gin.Default()

	//r.LoadHTMLGlob(paths.Theme() + "/**/*")

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	//r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Instantiate the server.
	return &Server{
		r,
	}
}

// ListenAndServe runs Verbis on a given port
// Returns errors.INVALID if the server could not start
func (s *Server) ListenAndServe(port int) error {
	const op = "router.ListenAndServe"
	passedPort := strconv.Itoa(port)
	err := s.Run(":" + passedPort)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Could not start Verbis on the port %d", port), Operation: op, Err: err}
	}
	return nil
}
