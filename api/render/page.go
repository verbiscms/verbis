package render

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/templates"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (r *Render) Page(g *gin.Context) ([]byte, error) {
	const op = "Render.GetPage"

	api.ServeChan <- 1
	defer func() {
		<-api.ServeChan
	}()

	url := g.Request.URL.Path
	post, err := r.store.Posts.GetBySlug(url)

	if err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No page found with the url: %s", url), Operation: op, Err: err}
	}

	// Check if the file has been cached
	var foundCache bool
	if r.options.CacheServerAssets {
		var cachedTemplate interface{}
		cachedTemplate, foundCache = cache.Store.Get(url)

		if cachedTemplate != nil && foundCache {
			return cachedTemplate.([]byte), nil
		}
	}

	_, err = g.Cookie("verbis-session")
	if err != nil && post.Status != "published" {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Page not published, or user is not logged in", Operation: op, Err: err}
	}

	pt := "index"
	if post.PageTemplate != "default" {
		pt = r.config.Template.TemplateDir + "/" + post.PageTemplate
	}

	master := ""
	if post.Layout != "default" {
		master = r.config.Template.LayoutDir + "/" + post.Layout
	} else {
		pt = pt + r.config.Template.FileExtension
	}

	tf := templates.NewFunctions(g, r.store, &post)
	gvFrontend := goview.New(goview.Config{
		Root:         paths.Theme(),
		Extension:    r.store.Config.Template.FileExtension,
		Master:       master,
		Partials:     []string{},
		Funcs:        tf.GetFunctions(),
		DisableCache: !environment.IsProduction(),
	})

	data, err := tf.GetData()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get template data", Operation: op, Err: err},
		}).Fatal()
	}

	var b bytes.Buffer
	if err := gvFrontend.RenderWriter(&b, pt, data); err != nil {
		panic(err)
	}

	minified, err := r.minify.MinifyBytes(&b, "text/html")
	if err != nil || minified == nil {
		return b.Bytes(), nil
	}

	go func() {
		if r.options.CacheServerTemplates && !foundCache {
			cache.Store.Set(url, minified, cache.RememberForever)
		}
	}()

	return minified, nil
}
