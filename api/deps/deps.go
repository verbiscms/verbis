// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deps

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
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
	Web     string
}

// Deps holds dependencies used by many.
// There will be normally only one instance of deps in play
// at a given time, i.e. one per Site built.
type Deps struct {

	// The database layer
	Store *models.Store

	// Configuration file of the site
	Config *config.Configuration

	// Cache

	Site domain.Site

	// Logger

	// Options
	Options *domain.Options

	// Paths
	Paths Paths

	// Theme
	Theme *domain.ThemeConfig

	tmpl tpl.TemplateHandler

	Running bool
}

func (d *Deps) Tmpl() tpl.TemplateHandler {
	return d.tmpl
}

func (d *Deps) SetTmpl(tmpl tpl.TemplateHandler) {
	d.tmpl = tmpl
}

func (d *Deps) SetOptions(options *domain.Options) {
	d.Options = options
}

type DepsConfig struct {

	// The database layer
	Store *models.Store

	// Config
	Config *config.Configuration

	Running bool
}

func New(cfg DepsConfig) *Deps {

	if cfg.Store == nil {
		panic("Must have a store")
	}

	if cfg.Config == nil {
		panic("Must have a configuration")
	}

	opts := cfg.Store.Options.GetStruct()

	theme := cfg.Store.Site.GetThemeConfig()

	d := &Deps{
		Store:   cfg.Store,
		Config:  cfg.Config,
		Site:    cfg.Store.Site.GetGlobalConfig(),
		Options: &opts,
		Paths: Paths{
			Base:    paths.Base(),
			Admin:   paths.Admin(),
			API:     paths.Api(),
			Theme:   paths.Theme(),
			Uploads: paths.Uploads(),
			Storage: paths.Storage(),
			Web:     paths.Web(),
		},
		Theme:   &theme,
		tmpl:    nil,
		Running: cfg.Running,
	}

	return d
}
