package redirects

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gin-gonic/gin"
)

// RedirectHandler defines methods for Redirect routes to interact with the server.
type RedirectHandler interface {
	List(g *gin.Context)
	Find(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// Redirects defines the handler for all Redirect Routes.
type Redirects struct {
	*deps.Deps
}