package internal

import (
	"fmt"
	"github.com/graymeta/stow"
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/store/options"
	"strings"
)

// StorageServices define the methods needed to obtain
// providers, buckets and the configuration for the
// storage layer.
type StorageServices interface {
	Provider(provider domain.StorageProvider) (stow.Location, error)
	Bucket(provider domain.StorageProvider, bucket string) (stow.Container, error)
	BucketByFile(file domain.File) (stow.Container, error)
	Config() (domain.StorageProvider, string, error)
}

// Service represents the implementation of
// StorageServices.
type Service struct {
	Env     *environment.Env
	Options options.Repository
}

const (
	// ErrMessageInvalidBucket is an error message returned by
	// Bucket and BucketByFile when an invalid bucket string
	// is passed.
	ErrMessageInvalidBucket = "Error retrieving bucket"
)

// Provider returns a stow.Location from the ProviderMap by
// the given string.
// Returns errors.INVALID if the Provider does not exist
// or there was an error connecting to it.
func (s *Service) Provider(provider domain.StorageProvider) (stow.Location, error) {
	const op = "Storage.Provider"
	if !Providers.Exists(provider) {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error connecting to storage provider: " + provider.String(), Operation: op, Err: fmt.Errorf("nil provider")}
	}
	loc, err := Providers[provider].Dial(s.Env)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error connecting to storage provider: " + provider.String(), Operation: op, Err: err}
	}
	return loc, nil
}

// Bucket returns a stow.Container by the given strings.
// Returns errors.INVALID if there was an error
// obtaining the provider or bucket.
func (s *Service) Bucket(provider domain.StorageProvider, bucket string) (stow.Container, error) {
	const op = "Storage.Bucket"

	p, err := s.Provider(provider)
	if err != nil {
		return nil, err
	}

	c, err := p.Container(bucket)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: ErrMessageInvalidBucket, Operation: op, Err: err}
	}

	return c, nil
}

// BucketByFile returns a stow.Container by the given
// domain.File.
// Returns errors.INVALID if there was an error
// obtaining the provider or bucket.
func (s *Service) BucketByFile(file domain.File) (stow.Container, error) {
	const op = "Storage.BucketByFile"

	p, err := s.Provider(file.Provider)
	if err != nil {
		return nil, err
	}

	c, err := p.Container(file.Bucket)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: ErrMessageInvalidBucket, Operation: op, Err: err}
	}

	return c, nil
}

// Config returns a domain.StorageProvider, a bucket or an
// error if there was a problem obtaining the currently
// set storage providers from the options table.
func (s *Service) Config() (domain.StorageProvider, string, error) {
	// TODO - cache here?
	p, err := s.Options.Find("storage_provider")
	if err != nil {
		return "", "", err
	}

	bucket, err := s.Options.Find("storage_bucket")
	if err != nil {
		return "", "", err
	}

	provider := domain.StorageProvider(strings.ReplaceAll(cast.ToString(p), "\"", ""))
	if provider == "" {
		provider = domain.StorageLocal
	}

	return provider, strings.ReplaceAll(cast.ToString(bucket), "\"", ""), nil
}
