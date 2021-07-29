package config

import (
	"github.com/ghodss/yaml"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"io/ioutil"
	"path/filepath"
)

// Provider describes the methods for obtaining a theme
// configuration from the given them name.
type Provider interface {
	// Get returns a domain.ThemeConfiguration upon success.
	// Returns errors.INVALID if thee theme file could not be
	// matched or
	Get(theme string) (domain.ThemeConfig, error)
}

// Config defines the struct for obtaining theme
// configurations (yaml files).
type Config struct {
	// ThemePath is the directory in where the
	// themes reside /themes relative to the base.
	ThemePath string
}

const (
	// FileName is the default configuration file name within
	// the theme.
	FileName = "config.yml"
)

// Get retrieves is the implementation of the provider.
func (c *Config) Get(theme string) (domain.ThemeConfig, error) {
	return getThemeConfig(filepath.Join(c.ThemePath, theme), FileName)
}

// getThemeConfig is a wrapper for Fetch taking in a path
// and filename and unmarshalling the yaml file into the
// theme configuration.
func getThemeConfig(path, filename string) (domain.ThemeConfig, error) {
	const op = "Config.Get"

	cfg := DefaultTheme

	file, err := ioutil.ReadFile(filepath.Join(path, filename))
	if err != nil {
		return DefaultTheme, &errors.Error{Code: errors.INVALID, Message: "Error retrieving theme config file", Operation: op, Err: err}
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return DefaultTheme, &errors.Error{Code: errors.INVALID, Message: "Syntax error in theme config file", Operation: op, Err: err}
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
