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
	// The default configuration file name within the theme.
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
)

// Init
//
// Fetches and returns the theme configuration once.
func Init(path string) *domain.ThemeConfig {
	once.Do(func() {
		Fetch(path)
	})
	return cfg
}

// Get
//
// Returns the pointer to the theme configuration.
func Get() *domain.ThemeConfig {
	mutex.Lock()
	defer mutex.Unlock()
	return cfg
}

// Set
//
// Sets the cfg variable to a new theme configuration when
// a theme has been set by the user.
func Set(config domain.ThemeConfig) {
	mutex.Lock()
	defer mutex.Unlock()
	cfg = &config
}

// Fetch
//
// Get"s the themes configuration from the themes path
//
// Logs errors.INTERNAL if the unmarshalling was
// unsuccessful and returns the DefaultTheme
// variable.
func Fetch(path string) *domain.ThemeConfig {
	theme, err := getThemeConfig(path, FileName)
	if err != nil {
		logger.WithError(err).Error()
	}
	Set(*theme)
	return theme
}

// Find
//
// Looks up for theme configuration file by the given path
// and default file name.
//
// Returns errors.INTERNAL if the unmarshalling was
// unsuccessful and returns the DefaultTheme
// variable.
func Find(path string) (*domain.ThemeConfig, error) {
	theme, err := getThemeConfig(path, FileName)
	if err != nil {
		return nil, err
	}
	return theme, nil
}

// getThemeConfig
//
// Is a wrapper for Fetch taking in a path and filename
// and unmarshalling the yaml file into the theme
// configuration.
func getThemeConfig(path, filename string) (*domain.ThemeConfig, error) {
	const op = "Config.Fetch"

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

	return &cfg, nil
}
