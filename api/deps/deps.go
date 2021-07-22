// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deps

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/common/paths"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/services/site"
	"github.com/ainsleyclark/verbis/api/services/storage"
	"github.com/ainsleyclark/verbis/api/services/theme"
	"github.com/ainsleyclark/verbis/api/services/webp"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/sys"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/ainsleyclark/verbis/api/verbisfs"
	"github.com/ainsleyclark/verbis/api/watchers"
	"os"
	"path/filepath"
)

// Deps holds dependencies used by many.
// There will be normally only one instance of deps in play
// at a given time, i.e. one per Site built.
type Deps struct {
	Env *environment.Env
	// The database layer
	Store *store.Repository
	// Configuration file of the site
	Config *domain.ThemeConfig
	// Site
	Site site.Repository
	// Theme
	Theme   theme.Repository
	Watcher *watchers.Batch
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
	*d.Options = *options
}

func (d *Deps) SetTheme(name string) error {
	err := d.Store.Options.SetTheme(name)
	if err != nil {
		return err
	}
	d.Options.ActiveTheme = name
	//d.Watcher.SetTheme(d.Paths.Themes + string(os.PathSeparator) + name)
	return nil
}

type Config struct {

	// The database layer
	Store *store.Repository

	// Env
	Env *environment.Env
	Paths paths.Paths

	Installed bool

	System sys.System

	Running bool
}

func New(cfg Config) *Deps {
	if cfg.Store == nil && cfg.Running {
		panic("Must have a store")
	}


	var opts domain.Options
	if cfg.Running {
		opts = cfg.Store.Options.Struct()
	}

	st, err := storage.New(storage.Config{
		Environment: cfg.Env,
		Options:     cfg.Store.Options,
		Files:       cfg.Store.Files,
	})
	if err != nil {
		logger.WithError(err).Panic()
	}

	activeTheme, err := cfg.Store.Options.GetTheme()
	if err != nil {
		logger.WithError(err).Panic()
	}

	config.Init(filepath.Join(cfg.Paths.Themes, activeTheme))

	d := &Deps{
		Env:     cfg.Env,
		Store:   cfg.Store,
		Config:   config.Get(),
		Options: &opts,
		Paths:   cfg.Paths,
		tmpl:    nil,
		Running: cfg.Running,
		Site:    site.New(&opts, cfg.System),
		Theme:   theme.New(),
		FS:      verbisfs.New(api.Production, cfg.Paths),
		WebP:    webp.New(cfg.Paths.Bin + webp.Path),
		Storage: st,
		System:  cfg.System,
	}

	d.Watcher = watchers.New(d.ThemePath())
	d.Watcher.Start()

	return d
}
