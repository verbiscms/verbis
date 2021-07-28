// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// FileName is the default configuration file name within
	// the Theme.
	FileName       = "config.yml"
	ConfigCacheKey = "theme_config"
)

func (t *Theme) Config() (domain.ThemeConfig, error) {
	c, err := t.cache.Get(context.Background(), ConfigCacheKey)
	if err != nil {
		cfg, ok := c.(domain.ThemeConfig)
		if !ok {
			return domain.ThemeConfig{}, fmt.Errorf("error casting etc")
		}
		return cfg, nil
	}

	theme, err := t.options.GetTheme()
	if err != nil {
		return domain.ThemeConfig{}, err
	}

	cfg, err := getThemeConfig(filepath.Join(t.themesPath, theme), FileName)
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
	return getThemeConfig(filepath.Join(t.themesPath, theme), FileName)
}

// List all Theme configurations.
func (t *Theme) List(activeTheme string) ([]domain.ThemeConfig, error) {
	const op = "Theme.All"

	files, err := ioutil.ReadDir(t.themesPath)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error finding themes", Err: err, Operation: op}
	}

	var themes []domain.ThemeConfig
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		cfg, err := getThemeConfig(filepath.Join(t.themesPath, f.Name()), FileName)
		if err != nil {
			logger.WithError(err).Error()
			continue
		}
		if activeTheme != cfg.Theme.Name {
			cfg.Theme.Active = false
		}
		themes = append(themes, cfg)
	}

	if len(themes) == 0 {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No themes available", Operation: op, Err: ErrNoThemes}
	}

	return themes, nil
}

// getThemeConfig is a wrapper for Fetch taking in a path
// and filename and unmarshalling the yaml file into the
// Theme configuration.
func getThemeConfig(path, filename string) (domain.ThemeConfig, error) {
	const op = "Theme.Fetch"

	file, err := ioutil.ReadFile(filepath.Join(path, filename))
	if err != nil {
		return domain.ThemeConfig{}, &errors.Error{Code: errors.INTERNAL, Message: "Error retrieving Theme config file", Operation: op, Err: err}
	}

	var cfg domain.ThemeConfig
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return domain.ThemeConfig{}, &errors.Error{Code: errors.INTERNAL, Message: "Syntax error in Theme config file", Operation: op, Err: err}
	}

	screenshot, err := findScreenshot(path)
	if err == nil {
		cfg.Theme.Screenshot = screenshot
	}

	cfg.Theme.Name = filepath.Base(path)
	cfg.Theme.Active = true
	cfg.Resources = cfg.Resources.Clean()

	return cfg, nil
}
