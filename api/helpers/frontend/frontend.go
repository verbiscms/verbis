package frontend

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strconv"
	"strings"
)

type Cacher interface {
	Cache(g *gin.Context)
}

type Cache struct {
	options models.OptionsRepository
}

func NewCache(o models.OptionsRepository) *Cache {
	return &Cache{
		options: o,
	}
}

func (t *Cache) Cache(g *gin.Context) {
	const op = "Cacheer.Cache"

	options, err := t.options.GetStruct()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: fmt.Errorf("could not get the options struct")},
		}).Fatal()
	}

	// Bail if the cache frontend is disabled
	if !options.CacheFrontend {
		return
	}

	path := g.Request.URL.Path

	// Don't cache any admin assets
	if strings.Contains(path, "admin") {
		return
	}

	// Get the expiration
	expiration := options.CacheFrontendSeconds

	// Get the request type
	request := options.CacheFrontendRequest
	allowedRequest := []string{"max-age", "max-stale", "min-fresh", "no-cache", "no-store", "no-transform", "only-if-cached"}
	if request == "" || !helpers.StringInSlice(request, allowedRequest) {
		request = "max-age"
	}

	// Get the extensions to be cached
	extensionsAllowed := options.CacheFrontendExtension
	extension := filepath.Ext(path)

	// Check if the extensions
	if len(extensionsAllowed) > 0 {
		for _, v := range extensionsAllowed {
			if extension == "." + v {
				cache := ""
				if request == "max-age" || request == "min-fresh" || request == "max-stale" {
					cache = fmt.Sprintf("%s=%s, %s", request, strconv.FormatInt(expiration, 10), "public")
				} else {
					cache = request
				}
				g.Header("Cache-Control", cache)
			}
		}
	}
}