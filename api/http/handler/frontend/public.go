package frontend

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/render"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

// PublicHandler defines methods for the frontend to interact with the server
type PublicHandler interface {
	GetUploads(g *gin.Context)
	GetAssets(g *gin.Context)
	Serve(g *gin.Context)
}

// Public defines the handler for all frontend routes
type Public struct {
	store  *models.Store
	config config.Configuration
	render render.Renderer
	render.ErrorHandler
}

// NewPublic - Construct
func NewPublic(m *models.Store, config config.Configuration) *Public {
	const op = "FrontendHandler.newFrontend"

	return &Public{
		store:  m,
		config: config,
		render: render.NewRender(m, config),
		ErrorHandler: &render.Errors{
			ThemeConfig: m.Site.GetThemeConfig(),
		},
	}
}

// GetUploads retrieves images & media in the uploads folder, returns webp if accepts.
func (c *Public) GetUploads(g *gin.Context) {
	const op = "FrontendHandler.GetUploads"

	mimeType, file, err := c.render.Upload(g)
	if err != nil {
		c.NotFound(g)
		return
	}

	g.Data(200, *mimeType, *file)
}

// GetAssets retrieves assets from the theme path, returns webp if accepts.
func (c *Public) GetAssets(g *gin.Context) {
	const op = "FrontendHandler.GetAssets"

	mimeType, file, err := c.render.Asset(g)
	if err != nil {
		c.NotFound(g)
		return
	}

	g.Data(200, *mimeType, *file)
}

// Serve the front end website
func (c *Public) Serve(g *gin.Context) {
	const op = "FrontendHandler.Serve"

	page, err := c.render.Page(g)
	if err != nil {
		color.Red.Println(err)
		c.NotFound(g)
		return
	}

	g.Data(200, "text/html", page)
}
