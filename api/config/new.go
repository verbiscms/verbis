package config

import (
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/domain"
)

type Provider interface {
	Get(path string) (domain.ThemeConfig, error)
}

type Config struct{
	cache cache.Store
}

func (c *Config) Get(path string) (domain.ThemeConfig, error) {


	return domain.ThemeConfig{}, nil
}
