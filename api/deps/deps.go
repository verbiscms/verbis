// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deps

import (
	"fmt"
	"github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/services/site"
	"github.com/verbiscms/verbis/api/services/storage"
	"github.com/verbiscms/verbis/api/services/theme"
	"github.com/verbiscms/verbis/api/services/webp"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/sys"
	"github.com/verbiscms/verbis/api/tpl"
	"github.com/verbiscms/verbis/api/verbisfs"
	"github.com/verbiscms/verbis/api/watcher"
	"os"
)

// Deps holds dependencies used by many.
// There will be normally only one instance of deps in play
// at a given time, i.e. one per Site built.
type Deps struct {
	Env   *environment.Env
	Cache cache.Store
	// The database layer
	Store *store.Repository
	// Configuration file of the site
	Config *domain.ThemeConfig
	// Site
	Site site.Service
	// Theme
	Theme   theme.Service
	Watcher watcher.FileWatcher
	// Options
	Options *domain.Options
	// Paths
	Paths paths.Paths
	// File System (Web and SPA)
	FS *verbisfs.FileSystem
	// Webp
	WebP webp.Execer
	// template
	tmpl      tpl.TemplateHandler
	System    sys.System
	Storage   storage.Provider
	Installed bool
	Running   bool
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
	if d.Options == nil {
		d.Options = options
		return
	}
	*d.Options = *options
}

type Config struct {

	// The database layer
	Store *store.Repository

	// Env
	Env   *environment.Env
	Paths paths.Paths

	Installed bool

	System sys.System

	Running bool
}

func New(cfg Config) (*Deps, error) {
	if cfg.Store == nil && cfg.Running {
		return nil, fmt.Errorf("must have a store")
	}

	cs, err := cache.Load(cfg.Env)
	if err != nil {
		return nil, err
	}

	var opts domain.Options
	if cfg.Running {
		opts = cfg.Store.Options.Struct()
	}

	st, err := storage.New(storage.Config{
		Environment: cfg.Env,
		Options:     cfg.Store.Options,
		Files:       cfg.Store.Files,
		Cache: cs,
	})
	if err != nil {
		return nil, err
	}

	themeService := theme.New(cs, cfg.Store.Options)
	config, err := themeService.Config()
	if err != nil {
		return nil, err
	}

	d := &Deps{
		Env:   cfg.Env,
		Cache: cs,
		Store: cfg.Store,
		Config:  &config,
		Options: &opts,
		Paths:   cfg.Paths,
		tmpl:    nil,
		Running: cfg.Running,
		Site:    site.New(cfg.Store.Options, cfg.System),
		Theme:   themeService,
		FS:      verbisfs.New(api.Production, cfg.Paths),
		WebP:    webp.New(cfg.Paths.Bin + webp.Path),
		Storage: st,
		System:  cfg.System,
	}

	return d, nil
}
