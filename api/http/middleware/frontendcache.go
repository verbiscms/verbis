package middleware

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func FrontEndCache(o models.OptionsRepository) gin.HandlerFunc {
	return func(g *gin.Context) {

		options, err := o.GetStruct()
		if err != nil {
			log.Fatal(options)
		}

		// Check
		if !options.CacheFrontend {
			g.Next()
		}

		path := g.Request.URL.Path

		// Don't cache any admin assets
		if strings.Contains(path, "admin") {
			g.Next()
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

		g.Next()
	}
}

