package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/google"
	"github.com/graymeta/stow/local"
	"github.com/graymeta/stow/s3"
	"io/ioutil"
)

func (s *Storage) SetProvider(provider domain.StorageProvider) error {
	const op = "Storage.SetProvider"

	p, err := s.getProvider(provider)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error setting provider", Operation: op, Err: err}
	}

	err = s.optsRepo.Update("storage_provider", provider)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating options table with new provider", Operation: op, Err: err}
	}

	s.provider = p
	s.ProviderName = provider

	return nil
}

func (s *Storage) SetBucket(id string) error {
	const op = "Storage.SetBucket"

	if s.opts.StorageProvider == domain.StorageLocal {
		id = ""
	}

	container, err := s.provider.Container(id)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error setting bucket", Operation: op, Err: err}
	}
	s.bucket = container

	err = s.optsRepo.Update("storage_bucket", id)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating options table with new bucket", Operation: op, Err: err}
	}

	return nil
}

func (s *Storage) CreateBucket(name string) error {
	const op = "Storage.CreateBucket"

	_, err := s.provider.CreateContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error creating bucket: " + name, Operation: op, Err: err}
	}

	return nil
}

func (s *Storage) ListBuckets() (domain.Buckets, error) {
	const op = "Storage.ListBuckets"

	var buckets = make(domain.Buckets, 0)
	err := stow.WalkContainers(s.provider, stow.NoPrefix, 100, func(c stow.Container, err error) error {
		if err != nil {
			return err
		}
		buckets = append(buckets, domain.Bucket{
			Id:   c.ID(),
			Name: c.Name(),
		})
		return nil
	})

	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error obtaining buckets", Operation: op, Err: err}
	}

	return nil, nil
}

func (s *Storage) getBucket(file domain.File) (stow.Container, error) {
	provider, err := s.getProvider(file.Provider)
	if err != nil {
		return nil, err
	}

	bucket, err := provider.Container(file.Bucket)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func (s *Storage) getProvider(provider domain.StorageProvider) (stow.Location, error) {
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
			s3.ConfigAccessKeyID: s.env.AWSAccessKey,
			s3.ConfigSecretKey:   s.env.AWSSecret,
		})
	case domain.StorageGCP:
		json, err := ioutil.ReadFile(s.env.GCPJson)
		if err != nil {
			return nil, err
		}
		cont, err = stow.Dial(google.Kind, stow.ConfigMap{
			google.ConfigJSON:      string(json),
			google.ConfigProjectId: s.env.GCPProjectId,
		})
		// TODO, put in ENV
		//case domain.StorageAzure:
		//	cont, err = stow.Dial(azure.Kind, stow.ConfigMap{
		//		azure.ConfigKey:
		//	})
	}

	return cont, err
}
