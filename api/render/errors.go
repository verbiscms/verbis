package render

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type ErrorHandler interface {
	NotFound(g *gin.Context)
}

type Errors struct{
	ThemeConfig domain.ThemeConfig
}

func (e *Errors) NotFound(g *gin.Context) {

	gvError := goview.New(goview.Config{
		Root:         paths.Theme(),
		Extension:    e.ThemeConfig.FileExtension,
		Partials:     []string{},
		DisableCache: true,
	})

	// TODO: need to log here?!
	if err := gvError.Render(g.Writer, 404, "404", nil); err != nil {
		panic(err)
	}
	return
}
