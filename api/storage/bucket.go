package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/google"
	"github.com/graymeta/stow/local"
	"github.com/graymeta/stow/s3"
	"io/ioutil"
	"net/url"
	"strings"
)

func (s *Storage) SetProvider(provider domain.StorageProvider) error {
	const op = "Storage.SetProvider"

	p, err := s.getProvider(provider)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error setting provider", Operation: op, Err: err}
	}
	s.provider = p

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
		fmt.Println(err, c.ID(), c.Name())
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

func (s *Storage) cleanLocalPath(uri *url.URL) *url.URL {
	if s.local {
		uri.Path = strings.ReplaceAll(uri.Path, s.paths.Storage, "")
	}
	return uri
}

func (s *Storage) getProvider(provider domain.StorageProvider) (stow.Location, error) {
	var (
		cont stow.Location
		err  error
	)

	s.local = false

	switch provider {
	case domain.StorageLocal:
		cont, err = stow.Dial(local.Kind, stow.ConfigMap{
			local.ConfigKeyPath: s.paths.Storage,
		})
		s.local = true
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
