package frontend

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

func Error(g *gin.Context, config config.Configuration) {
	gvError := goview.New(goview.Config{
		Root:         paths.Theme(),
		Extension:    config.Template.FileExtension,
		Partials:     []string{},
		DisableCache: true,
	})


	// need to log here?!

	if err := gvError.Render(g.Writer, 404, "404", nil); err != nil {
		panic(err)
	}
	return
}
