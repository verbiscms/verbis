package internal

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/graymeta/stow"
	_ "github.com/graymeta/stow/azure"
	"github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/google"
	"github.com/graymeta/stow/local"
	"github.com/graymeta/stow/s3"
	_ "github.com/graymeta/stow/s3"
	"io/ioutil"
)

type StorageServices interface {
	Provider(provider domain.StorageProvider) (stow.Location, error)
	Bucket(file domain.File) (stow.Container, error)
}

type Service struct {
	Env     *environment.Env
	paths   paths.Paths
	gcpJson *string
}

func NewService(env *environment.Env) *Service {
	return &Service{
		Env:     env,
		paths:   paths.Get(),
		gcpJson: nil,
	}
}

func (s *Service) Provider(provider domain.StorageProvider) (stow.Location, error) {
	const op = "Storage.Provider"

	var (
		cont stow.Location
		err  error
	)

	switch provider {
	case domain.StorageLocal:
		cont, err = stow.Dial(local.Kind, stow.ConfigMap{
			local.ConfigKeyPath: s.paths.Storage,
		})
	case domain.StorageAWS:
		cont, err = stow.Dial(s3.Kind, stow.ConfigMap{
			s3.ConfigAccessKeyID: s.Env.AWSAccessKey,
			s3.ConfigSecretKey:   s.Env.AWSSecret,
		})
	case domain.StorageGCP:
		json, err := s.getGCPJson()
		if err != nil {
			return nil, err
		}
		cont, err = stow.Dial(google.Kind, stow.ConfigMap{
			google.ConfigJSON:      json,
			google.ConfigProjectId: s.Env.GCPProjectId,
		})
	}

	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error connecting to storage provider", Operation: op, Err: err}
	}

	return cont, nil
}

func (s *Service) Bucket(file domain.File) (stow.Container, error) {
	const op = "Storage.Bucket"

	provider, err := s.Provider(file.Provider)
	if err != nil {
		// TODO varf( err etc
		return nil, err
	}

	bucket, err := provider.Container(file.Bucket)
	if err != nil {
		// TODO varf( err etc
		return nil, err
	}

	return bucket, nil
}

func (s *Service) getGCPJson() (string, error) {
	if s.gcpJson != nil {
		return *s.gcpJson, nil
	}

	bytes, err := ioutil.ReadFile(s.Env.GCPJson)
	if err != nil {
		return "", err
	}

	json := string(bytes)

	s.gcpJson = &json

	return json, nil
}
