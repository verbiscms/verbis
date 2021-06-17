// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deps

import (
	"fmt"
	"github.com/ainsleyclark/updater"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/services/site"
	"github.com/ainsleyclark/verbis/api/services/theme"
	"github.com/ainsleyclark/verbis/api/services/webp"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/ainsleyclark/verbis/api/verbisfs"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/ainsleyclark/verbis/api/watchers"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
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
	Theme theme.Repository

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
	tmpl tpl.TemplateHandler

	Updater updater.Patcher

	Installed bool

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

	// Config
	Config *domain.ThemeConfig

	Paths paths.Paths

	Installed bool

	Running bool
}

func New(cfg Config) *Deps {
	if cfg.Store == nil && cfg.Running {
		panic("Must have a store")
	}

	if cfg.Config == nil && cfg.Running {
		panic("Must have a configuration")
	}

	var opts domain.Options
	if cfg.Running {
		opts = cfg.Store.Options.Struct()
	}

	p := paths.Get()

	d := &Deps{
		Env:     cfg.Env,
		Store:   cfg.Store,
		Config:  cfg.Config,
		Options: &opts,
		Paths:   p,
		tmpl:    nil,
		Running: cfg.Running,
		Site:    site.New(&opts),
		Theme:   theme.New(),
		FS:      verbisfs.New(api.Production, p),
		WebP:    webp.New(p.Bin + webp.Path),
		Updater: getUpdater(),
	}

	d.Watcher = watchers.New(d.ThemePath())
	d.Watcher.Start()

	return d
}

func getUpdater() updater.Patcher {
	const op = "Deps.GetUpdater"
	opts := &updater.Options{
		RepositoryURL: api.Repo,
		ArchiveName:   fmt.Sprintf("verbis_%s_%s_%s.zip", "0.0.1", runtime.GOOS, runtime.GOARCH),
		Version:       version.Version,
	}
	u, err := updater.New(opts)
	if err != nil {
		log.Fatal(&errors.Error{Code: errors.INTERNAL, Message: "Error creating new Verbis updater", Operation: op, Err: err})
	}
	return u
}
