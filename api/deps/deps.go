// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deps

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/site"
	"github.com/ainsleyclark/verbis/api/tpl"
	"os"
)

// Deps holds dependencies used by many.
// There will be normally only one instance of deps in play
// at a given time, i.e. one per Site built.
type Deps struct {

	// The database layer
	Store *models.Store

	// Configuration file of the site
	Config *domain.ThemeConfig

	// Cache

	Site site.Repository

	// Logger

	// Options
	Options *domain.Options

	// Paths
	Paths paths.Paths

	tmpl tpl.TemplateHandler

	Running bool
}

func (d *Deps) ThemePath() string {
	return d.Paths.Base + string(os.PathSeparator) + "themes" + string(os.PathSeparator) + d.Options.ActiveTheme
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

	// Env
	Env *environment.Env

	// Config
	Config *domain.ThemeConfig

	Paths paths.Paths

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

	d := &Deps{
		Store:   cfg.Store,
		Config:  cfg.Config,
		Options: &opts,
		Paths:   paths.Get(),
		tmpl:    nil,
		Running: cfg.Running,
		Site:    site.New(&opts),
	}

	return d
}
