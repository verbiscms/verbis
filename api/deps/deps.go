package deps

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/tpl"
)

type Paths struct {
	Base    string
	Admin   string
	API     string
	Theme   string
	Uploads string
	Storage string
}

// Deps holds dependencies used by many.
// There will be normally only one instance of deps in play
// at a given time, i.e. one per Site built.
type Deps struct {

	// The database layer
	Store *models.Store

	// Configuration file of the site
	Config config.Configuration

	// Cache

	Site domain.Site
	// Logger

	// Options
	Options domain.Options

	// Paths
	Paths Paths

	// Theme
	Theme domain.ThemeConfig

	Tpl tpl.TemplateHandler
}
