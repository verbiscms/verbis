package deps

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/models"
)

// Deps holds dependencies used by many.
// There will be normally only one instance of deps in play
// at a given time, i.e. one per Site built.
type Deps struct {

	// The database layer
	Store *models.Store

	// Configuration file of the site
	Config config.Configuration

	// Cache

	// Logger

	//
}
