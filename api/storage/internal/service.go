package internal

import (
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
	Info() (domain.StorageProvider, string, error)
}

type Service struct {
	Env         *environment.Env
	Options     options.Repository
	storagePath string
	gcpJson     *string
}

func NewService(env *environment.Env, options options.Repository, storagePath string) *Service {
	return &Service{
		Env:         env,
		Options:     options,
		storagePath: storagePath,
		gcpJson:     nil,
	}
}

func (s *Service) Provider(provider domain.StorageProvider) (stow.Location, error) {
	const op = "Storage.Provider"
	loc, err := Providers[provider].Dial(s.Env)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error connecting to storage provider: " + provider.String(), Operation: op, Err: err}
	}
	return loc, nil
}

func (s *Service) BucketByFile(file domain.File) (stow.Container, error) {
	const op = "Storage.BucketByFile"

	p, err := s.Provider(file.Provider)
	if err != nil {
		return nil, err
	}

	c, err := p.Container(file.Bucket)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) Bucket(provider domain.StorageProvider, bucket string) (stow.Container, error) {
	const op = "Storage.Bucket"

	p, err := s.Provider(provider)
	if err != nil {
		return nil, err
	}

	c, err := p.Container(bucket)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) Info() (domain.StorageProvider, string, error) {
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

//
//switch provider {
//case domain.StorageLocal:
//cont, err = stow.Dial(local.Kind, stow.ConfigMap{
//local.ConfigKeyPath: s.storagePath,
//})
//case domain.StorageAWS:
//cont, err = stow.Dial(s3.Kind, stow.ConfigMap{
//s3.ConfigAccessKeyID: s.Env.AWSAccessKey,
//s3.ConfigSecretKey:   s.Env.AWSSecret,
//})
//case domain.StorageGCP:
//
//}
