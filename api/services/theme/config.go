// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"context"
	"fmt"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	configCacheKey = "theme_config"
)

// Config satisfies the Theme interface by retrieving a
// theme configuration. The config will be cached
// forever until it is wiped by Set().
func (t *Theme) Config() (domain.ThemeConfig, error) {
	c := domain.ThemeConfig{}
	err := t.cache.Get(context.Background(), configCacheKey, &c)
	if err == nil {
		return c, nil
	}

	theme, err := t.options.GetTheme()
	if err != nil {
		return domain.ThemeConfig{}, err
	}

	cfg, err := t.config.Get(theme)
	if err != nil {
		return domain.ThemeConfig{}, err
	}

	t.cache.Set(context.Background(), configCacheKey, cfg, cache.Options{
		Expiration: cache.RememberForever,
	})

	return cfg, nil
}

// Activate satisfies the Theme interface by activating a
// theme by name.
func (t *Theme) Activate(theme string) (domain.ThemeConfig, error) {
	const op = "Theme.Activate"

	ok := t.Exists(theme)
	if !ok {
		return domain.ThemeConfig{}, &errors.Error{Code: errors.INVALID, Message: "Error finding theme: " + theme, Operation: op, Err: fmt.Errorf("no theme found")}
	}

	err := t.options.SetTheme(theme)
	if err != nil {
		return domain.ThemeConfig{}, err
	}

	t.cache.Delete(context.Background(), configCacheKey)

	cfg, err := t.Config()
	if err != nil {
		return domain.ThemeConfig{}, err
	}

	return cfg, nil
}

// Exists satisfies the Theme interface by checking to see
// if a theme exists by name.
func (t *Theme) Exists(theme string) bool {
	_, err := os.Stat(filepath.Join(t.themesPath, theme))
	return !os.IsNotExist(err)
}

// Find satisfies the Theme interface by retrieving a config
// by name.
func (t *Theme) Find(theme string) (domain.ThemeConfig, error) {
	return t.config.Get(theme)
}

// List satisfies the Theme interface by listing all theme
// configurations in the theme path.
func (t *Theme) List() ([]domain.ThemeConfig, error) {
	const op = "Theme.List"

	files, err := ioutil.ReadDir(t.themesPath)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error finding themes", Err: err, Operation: op}
	}

	theme, err := t.options.GetTheme()
	if err != nil {
		return nil, err
	}

	var themes []domain.ThemeConfig
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		cfg, err := t.config.Get(f.Name())
		if err != nil {
			logger.WithError(err).Error()
			continue
		}
		if theme == cfg.Theme.Name {
			cfg.Theme.Active = true
		}
		themes = append(themes, cfg)
	}

	if len(themes) == 0 {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No themes available", Operation: op, Err: ErrNoThemes}
	}

	return themes, nil
}
