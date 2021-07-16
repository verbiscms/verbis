package internal

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/store/options"
	"github.com/graymeta/stow"
	_ "github.com/graymeta/stow/azure"
	_ "github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/s3"
	"github.com/spf13/cast"
	"strings"
)

type StorageServices interface {
	Provider(provider domain.StorageProvider) (stow.Location, error)
	Bucket(provider domain.StorageProvider, bucket string) (stow.Container, error)
	BucketByFile(file domain.File) (stow.Container, error)
	Config() (domain.StorageProvider, string, error)
}

type Service struct {
	Env     *environment.Env
	Options options.Repository
}

const (
	ErrMessageConfigNotSet  = "Configuration not set for: "
	ErrMessageDial          = "Error dialling storage provider: "
	ErrMessageInvalidBucket = "Error retrieving bucket"
)

func NewService(env *environment.Env, options options.Repository) *Service {
	return &Service{
		Env:     env,
		Options: options,
	}
}

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

func (s *Service) Bucket(provider domain.StorageProvider, bucket string) (stow.Container, error) {
	const op = "Storage.Bucket"

	p, err := s.Provider(provider)
	if err != nil {
		return nil, err
	}

	c, err := p.Container(bucket)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: ErrMessageInvalidBucket, Operation: op, Err: err}
	}

	return c, nil
}

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

func (s *Service) Config() (domain.StorageProvider, string, error) {
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
