package render

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Renderer
type Renderer interface {
	Asset(g *gin.Context) (*string, *[]byte, error)
	Upload(g *gin.Context) (*string, *[]byte, error)
	Page(g *gin.Context) ([]byte, error)
}

// Render
type Render struct {
	store   *models.Store
	config  config.Configuration
	minify  minifier
	cacher  headerWriter
	options domain.Options
	theme   domain.ThemeConfig
}

// NewRender - Construct
func NewRender(m *models.Store, config config.Configuration) *Render {
	const op = "Assets.NewAssets"

	options := m.Options.GetStruct()

	theme, err := m.Site.GetThemeConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get theme config", Operation: op, Err: err},
		}).Fatal()
	}

	return &Render{
		store:   m,
		config:  config,
		minify:  newMinify(options),
		cacher:  newHeaders(options),
		options: options,
		theme:   theme,
	}
}
