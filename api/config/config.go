package config

import (
	"github.com/ghodss/yaml"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"io/ioutil"
	"path/filepath"
)

type Provider interface {
	Get(theme string) (domain.ThemeConfig, error)
}

type Config struct {
	ThemePath string
}

const (
	// FileName is the default configuration file name within
	// the theme.
	FileName = "config.yml"
)

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
		return DefaultTheme, &errors.Error{Code: errors.INTERNAL, Message: "Error retrieving theme config file", Operation: op, Err: err}
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return DefaultTheme, &errors.Error{Code: errors.INTERNAL, Message: "Syntax error in theme config file", Operation: op, Err: err}
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
