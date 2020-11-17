package controllers

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/gin-gonic/gin"
)

// CacheHandler defines methods for fields to interact with the server
type CacheHandler interface {
	Clear(g *gin.Context)
}

// CacheController defines the handler for Cache
type CacheController struct{}

// newCache - Construct
func newCache() *CacheController {
	return &CacheController{}
}

// Clear server cache
func (c *CacheController) Clear(g *gin.Context) {
	const op = "CacheHandler.Clear"
	cache.Store.Flush()
	Respond(g, 200, "Successfully cleared server cache", nil)
}
