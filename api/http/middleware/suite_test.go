package middleware

import (
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"testing"
)

// MiddlewareTestSuite defines the helper used for middleware
// testing.
type MiddlewareTestSuite struct {
	api.HandlerSuite
}

// TestMiddleware
//
// Assert testing has begun.
func TestMiddleware(t *testing.T) {
	suite.Run(t, &MiddlewareTestSuite{
		HandlerSuite: api.TestSuite(),
	})
}

// DefaultHandler
//
// Is a helper func for returning data for testing.
func (t *MiddlewareTestSuite) DefaultHandler(g *gin.Context) {
	g.String(200, "verbis")
	return
}
