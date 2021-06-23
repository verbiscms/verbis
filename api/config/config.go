// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

const (
	// FileName is the default configuration file name within
	// the theme.
	FileName = "config.yml"
)

var (
	// The main theme configuration.
	cfg = &domain.ThemeConfig{}
	// Ensure the theme configuration is set only once upon
	// initialisation.
	once = sync.Once{}
	// Required to avoid concurrent map writes when writing
	// via a websocket.
	mutex = &sync.Mutex{}
	// ErrNoThemes is returned by All when no themes
	// have been found by looping over the theme's
	// directory.
	ErrNoThemes = errors.New("no page templates found")
)

// Init fetches and returns the theme configuration once.
func Init(path string) *domain.ThemeConfig {
	once.Do(func() {
		Fetch(path)
	})
	return cfg
}

// Get returns the pointer to the theme configuration.
func Get() *domain.ThemeConfig {
	mutex.Lock()
	defer mutex.Unlock()
	return cfg
}

// Set the cfg variable to a new theme configuration when
// a theme has been set by the user.
func Set(config domain.ThemeConfig) {
	mutex.Lock()
	defer mutex.Unlock()
	cfg = &config
}

// Fetch gets the themes configuration from the themes
// path.
// Logs errors.INTERNAL if the unmarshalling was
// unsuccessful and returns the DefaultTheme
// variable.
func Fetch(path string) *domain.ThemeConfig {
	mutex.Lock()
	theme, err := getThemeConfig(path, FileName)
	mutex.Unlock()
	if err != nil {
		logger.WithError(err).Error()
	}
	Set(*theme)
	return theme
}

// Find looks up for theme configuration file by the given
// path and default file name.
// Returns errors.INTERNAL if the unmarshalling was
// unsuccessful and returns the DefaultTheme
// variable.
func Find(path string) (*domain.ThemeConfig, error) {
	mutex.Lock()
	theme, err := getThemeConfig(path, FileName)
	mutex.Unlock()
	if err != nil {
		return nil, err
	}
	return theme, nil
}

// All Returns a slice of domain.ThemeConfig's by reading
// the path. If the file is a directory, it will be
// skipped until a config file has been found.
// Returns errors.NOTFOUND if no themes were found.
// Returns errors.INTERNAL if the theme path is invalid.
func All(path, activeTheme string) ([]*domain.ThemeConfig, error) {
	const op = "Theme.All"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error finding themes", Err: err, Operation: op}
	}

	var themes []*domain.ThemeConfig
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		cfg, err := Find(path + string(os.PathSeparator) + f.Name())
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
// theme configuration.
func getThemeConfig(path, filename string) (*domain.ThemeConfig, error) {
	const op = "Theme.Fetch"

	cfg := DefaultTheme

	file, err := ioutil.ReadFile(path + string(os.PathSeparator) + filename)
	if err != nil {
		return &DefaultTheme, &errors.Error{Code: errors.INTERNAL, Message: "Error retrieving theme config file", Operation: op, Err: err}
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return &DefaultTheme, &errors.Error{Code: errors.INTERNAL, Message: "Syntax error in theme config file", Operation: op, Err: err}
	}

	screenshot, err := FindScreenshot(path)
	if err == nil {
		cfg.Theme.Screenshot = screenshot
	}

	cfg.Theme.Name = filepath.Base(path)
	cfg.Theme.Active = true
	cfg.Resources = cfg.Resources.Clean()

	return &cfg, nil
}
