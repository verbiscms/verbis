package render

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
)

type ErrorHandler interface {
	NotFound(g *gin.Context)
}

type Errors struct {
	ThemeConfig domain.ThemeConfig
	Store       *models.Store
}

func (e *Errors) NotFound(g *gin.Context) {

	//tm := tpl.NewManager(g, e.Store, &domain.PostData{}, config.Configuration{})

	//gvError := goview.New(goview.Config{
	//	Root:         paths.Theme(),
	//	Extension:    e.ThemeConfig.FileExtension,
	//	Partials:     []string{},
	//	Funcs:        tm.GetFunctions(),
	//	DisableCache: true,
	//})

	// TODO: need to log here?!
	//if err := gvError.Render(g.Writer, 404, "404", nil); err != nil {
	//	panic(err)
	//}
	//return
}
