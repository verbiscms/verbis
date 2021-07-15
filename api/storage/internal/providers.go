package internal

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/graymeta/stow"
)

type providerMap map[domain.StorageProvider]Provider

var Providers = providerMap{}

func (p providerMap) RegisterProvider(name domain.StorageProvider, provider Provider) {
	const op = "Storage.RegisterProvider"
	_, exists := p[name]
	if exists {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error, Duplicate provider: " + name.String(), Operation: op, Err: fmt.Errorf("duplicate storage provider")}).Panic()
		return
	}
	p[name] = provider
}

func (p providerMap) Exists(name domain.StorageProvider) bool {
	_, exists := p[name]
	return exists
}

type Provider interface {
	Dial(env *environment.Env) (stow.Location, error)
	ConfigValid(env *environment.Env) bool
}

func init() {
	Providers.RegisterProvider(domain.StorageLocal, &local{})
	Providers.RegisterProvider(domain.StorageGCP, &gcp{json: nil})
	Providers.RegisterProvider(domain.StorageAWS, &amazon{})
}
