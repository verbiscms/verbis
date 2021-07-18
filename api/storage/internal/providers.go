package internal

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/graymeta/stow"
)

// Provider defines the methods used for dialling and
// obtaining information about the registered
// providers.
type Provider interface {
	// Dial returns a new stow.Location or an error if there
	// was a problem connecting to the provider.
	Dial(env *environment.Env) (stow.Location, error)
	// Info returns the current state of the provider, if it
	// is connected, if environment variables are set and
	 // name and order etc.
	Info(env *environment.Env) domain.StorageProviderInfo
}

// ProviderMap map is a map containing storage providers
// by name.
type ProviderMap map[domain.StorageProvider]Provider

const (
	// ErrMessageConfigNotSet is an error message returned by
	// the Provider info method when there is no
	// environment variables set.
	ErrMessageConfigNotSet  = "Configuration not set for: "
	// ErrMessageDial is an error message returned by the
	// Provider when there was a problem dialling.
	ErrMessageDial          = "Error dialling storage provider: "
)

var (
	// Providers is the in memory store of the ProviderMap
	// at runtime.
	Providers = ProviderMap{}
	// dialler is an alias for stow.Dial used for testing.
	dialler = stow.Dial
)

// RegisterProvider takes in a domain.StorageProvider an
// a provider implementation. If the provider already
// exists, the function will panic with
// errors.INTERNAL. Otherwise it will
// be added to the ProviderMap.
func (p ProviderMap) RegisterProvider(name domain.StorageProvider, provider Provider) {
	const op = "Storage.RegisterProvider"
	_, exists := p[name]
	if exists {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error, Duplicate provider: " + name.String(), Operation: op, Err: fmt.Errorf("duplicate storage provider")}).Panic()
		return
	}
	p[name] = provider
}

// Exists determines if the Provider exists in the map
// with the given name.
func (p ProviderMap) Exists(name domain.StorageProvider) bool {
	_, exists := p[name]
	return exists
}

// init Register's the necessary providers at runtime.
func init() {
	Providers.RegisterProvider(domain.StorageLocal, &local{})
	Providers.RegisterProvider(domain.StorageGCP, &gcp{json: nil})
	Providers.RegisterProvider(domain.StorageAWS, &amazon{})
	Providers.RegisterProvider(domain.StorageAzure, &azure{})
}
