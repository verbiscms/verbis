package api

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/gin-gonic/gin"
)

// CacheHandler defines methods for fields to interact with the server
type CacheHandler interface {
	Clear(g *gin.Context)
}

// CacheController defines the handler for Cache
type Cache struct{}

// newCache - Construct
func NewCache() *Cache {
	return &Cache{}
}

// Clear server cache
func (c *Cache) Clear(g *gin.Context) {
	const op = "CacheHandler.Clear"
	cache.Store.Flush()
	handler.Respond(g, 200, "Successfully cleared server cache", nil)
}
