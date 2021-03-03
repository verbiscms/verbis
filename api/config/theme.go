// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ghodss/yaml"
	"io/ioutil"
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
// Logs errors.INTERNAL if the unmarshalling was
// unsuccessful and returns the DefaultTheme
// variable.
func Fetch(path string) *domain.ThemeConfig {
	theme, err := getThemeConfig(path, FileName)
	if err != nil {
		fmt.Println(err)
		logger.WithError(err).Error()
	}
	Set(*theme)
	return theme
}

// Config
//
// TODO: Need Config E, REWORD THIS FUNCTION ITS NO GOOD
func Config(path string) (*domain.ThemeConfig, error) {
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

	theme := DefaultTheme

	cfg, err := ioutil.ReadFile(path + filename)
	if err != nil {
		return &DefaultTheme, &errors.Error{Code: errors.INTERNAL, Message: "Unable to get retrieve theme config file", Operation: op, Err: err}
	}

	err = yaml.Unmarshal(cfg, &theme)
	if err != nil {
		return &DefaultTheme, &errors.Error{Code: errors.INTERNAL, Message: "Syntax error in theme config file", Operation: op, Err: err}
	}

	return &theme, nil
}
