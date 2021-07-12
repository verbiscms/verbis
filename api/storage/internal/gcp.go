package internal

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/local"
)

type providerFunc func(env *environment.Env, storagePath string) (stow.Location, error)

type providerMap map[domain.StorageProvider]providerFunc

var providers = providerMap{}

func (p providerMap) RegisterProvider(provider domain.StorageProvider, fn providerFunc) {
	const op = "Storage.RegisterProvider"
	_, exists := p[provider]
	if exists {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error, Duplicate provider: " + provider.String(), Operation: op, Err: fmt.Errorf("duplicate storage provider")}).Panic()
		return
	}
	p[provider] = fn
}

func (p providerMap) Dial(provider domain.StorageProvider, storagePath string, env *environment.Env) (stow.Location, error) {
	return p[provider](env, storagePath)
}

func init() {
	providers.RegisterProvider(domain.StorageLocal, func(env *environment.Env, storagePath string) (stow.Location, error) {
		return stow.Dial(local.Kind, stow.ConfigMap{
			local.ConfigKeyPath: storagePath,
		})
	})
}
