// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/ainsleyclark/verbis/api/tpl/tplimpl"
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

	d := &deps.Deps{
		Store:   e.Store,
		Site:    e.Store.Site.GetGlobalConfig(),
		Options: e.Store.Options.GetStruct(),
		Paths: deps.Paths{
			Base:    paths.Base(),
			Admin:   paths.Admin(),
			API:     paths.Api(),
			Theme:   paths.Theme(),
			Uploads: paths.Uploads(),
			Storage: paths.Storage(),
		},
		Theme: e.ThemeConfig,
	}

	d.Tpl = tplimpl.New(d)

	err := d.Tpl.Prepare(tpl.Config{
		Root:      paths.Theme(),
		Extension: e.ThemeConfig.FileExtension,
		Master:    "",
	}).Execute(g.Writer, "404", nil)

	if err != nil {
		panic(err)
	}

	//
	//gvError := goview.New(goview.Config{
	//	Root:        ,
	//	Extension:    ,
	//	Partials:     []string{},
	//	//Funcs:        tm.Funcs(),
	//	DisableCache: true,
	//})
	//
	//if err := gvError.Render(g.Writer, 404, "404", tm.Data()); err != nil {
	//
	//	return
	//}

	return
}
