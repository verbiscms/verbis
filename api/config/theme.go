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
)

// Get
//
// Fetches and returns the theme configuration once.
func Get(path string) *domain.ThemeConfig {
	once.Do(func() {
		Fetch(path)
	})
	return cfg
}

// Set
//
// Sets the cfg variable to a new theme configuration when
// a theme has been set by the user.
func Set(config domain.ThemeConfig) {
	cfg = &config
}

// Fetch
//
// Get"s the themes configuration from the themes path
// Logs errors.INTERNAL if the unmarshalling was
// unsuccessful and returns the DefaultTheme
// variable.
func Fetch(path string) *domain.ThemeConfig {
	return getThemeConfig(path, FileName)
}

// getThemeConfig
//
// Is a wrapper for Fetch taking in a path and filename
// and unmarshalling the yaml file into the theme
// configuration.
func getThemeConfig(path, filename string) *domain.ThemeConfig {
	const op = "Config.Fetch"

	t := DefaultTheme
	theme, err := ioutil.ReadFile(path + string(os.PathSeparator) + filename)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Unable to get retrieve theme config file", Operation: op, Err: err}).Error()
	}

	err = yaml.Unmarshal(theme, &t)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Syntax error in theme config file", Operation: op, Err: err}).Error()
	}

	Set(t)

	return cfg
}
