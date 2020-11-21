package frontend

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Renderer
type Renderer interface {
	Cache(g *gin.Context)
}

// Render
type Render struct {
	gin *gin.Context
	store *models.Store
	config  config.Configuration
	options domain.Options
	post *domain.Post
	path string
}

// NewRender - Construct
func NewRender(g *gin.Context, m *models.Store, config config.Configuration) *Render {
	const op = "Renderer.NewRender"

	options, err := m.Options.GetStruct()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: err},
		}).Fatal()
	}

	return &Render{
		store: m,
		options: options,
		config:  config,
		path: g.Request.URL.Path,
	}
}

// checkExists checks a page exists in the database or
// the controller should render a 404.
func (c *Render) checkExists() error {
	post, err := c.store.Posts.GetBySlug(c.path)
	if err != nil {
		return err
	}
	c.post = &post
	return nil
}

func (c *Render) checkStatus() error {
	_, err := c.gin.Cookie("verbis-session")
	if err != nil && c.post.Status != "published" {
		return &errors.Error{
			Code:      errors.INVALID,
			Message:   "",
			Operation: "",
			Err:       nil,
			Stack:     nil,
		}
	}
	return nil
}

// Check if the post has been cached
func (c *Render) checkCache() {
	var foundCache bool
	if c.options.CacheServerAssets {
		var cachedTemplate interface{}
		cachedTemplate, foundCache = cache.Store.Get(c.path)

		if cachedTemplate != nil && foundCache {
			c.gin.Writer.WriteHeader(200)
			c.gin.Writer.Write(cachedTemplate.([]byte))
			return
		}
	}
}