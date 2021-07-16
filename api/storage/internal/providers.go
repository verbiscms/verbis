package internal

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/graymeta/stow"
)

type Provider interface {
	Dial(env *environment.Env) (stow.Location, error)
	Info(env *environment.Env) domain.StorageProviderInfo
}

type ProviderMap map[domain.StorageProvider]Provider

var Providers = ProviderMap{}

func (p ProviderMap) RegisterProvider(name domain.StorageProvider, provider Provider) {
	const op = "Storage.RegisterProvider"
	_, exists := p[name]
	if exists {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error, Duplicate provider: " + name.String(), Operation: op, Err: fmt.Errorf("duplicate storage provider")}).Panic()
		return
	}
	p[name] = provider
}

func (p ProviderMap) Exists(name domain.StorageProvider) bool {
	_, exists := p[name]
	return exists
}

func init() {
	Providers.RegisterProvider(domain.StorageLocal, &local{})
	Providers.RegisterProvider(domain.StorageGCP, &gcp{json: nil})
	Providers.RegisterProvider(domain.StorageAWS, &amazon{})
	Providers.RegisterProvider(domain.StorageAzure, &azure{})
}
