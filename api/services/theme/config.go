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

func (t *Theme) Config() (domain.ThemeConfig, error) {
	const op = "Theme.Config"

	c, err := t.cache.Get(context.Background(), configCacheKey)
	if err == nil {
		cfg, ok := c.(domain.ThemeConfig)
		if ok {
			return cfg, nil
		}
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error casting cache item to theme config", Operation: op, Err: fmt.Errorf("bad cast")})
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

func (t *Theme) Set(theme string) (domain.ThemeConfig, error) {
	const op = "Theme.Set"

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

// Exists checks if a Theme exists by name.
func (t *Theme) Exists(theme string) bool {
	_, err := os.Stat(filepath.Join(t.themesPath, theme))
	return !os.IsNotExist(err)
}

func (t *Theme) Find(theme string) (domain.ThemeConfig, error) {
	return t.config.Get(theme)
}

// List all Theme configurations.
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
