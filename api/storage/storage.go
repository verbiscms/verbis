// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"bytes"
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
	"io"
	"io/ioutil"
	"net/url"
	"strings"
)

// change provider
// channge provider

type Client interface {
	Find(path string) ([]byte, error)
	Upload(path string, contents io.Reader) (stow.Item, error)
	SetProvider(location domain.StorageProvider) error
	SetBucket(id string) error
	ListBuckets() error
}

type Storage struct {
	provider stow.Location
	bucket   stow.Container
	opts     *domain.Options
	env      *environment.Env
}

// New parse config
func New(env *environment.Env, opts *domain.Options) (Client, error) {
	s := &Storage{
		env:  env,
		opts: opts,
	}

	err := s.SetProvider(opts.StorageProvider)
	if err != nil {
		return nil, err
	}

	err = s.SetBucket(opts.StorageBucket)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Upload something TODO
func (s *Storage) Upload(path string, contents io.Reader) (stow.Item, error) {
	const op = "Storage.Upload"

	r := strings.NewReader("this is a test")

	buf := &bytes.Buffer{}
	length, err := io.Copy(buf, contents)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error reading file", Operation: op, Err: err}
	}

	item, err := s.bucket.Put(path, r, length, nil)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *Storage) Find(path string) ([]byte, error) {
	item, err := s.provider.ItemByURL(&url.URL{Path: path})
	if err != nil {
		return nil, err
	}

	file, err := item.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

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

func (s *Storage) ListBuckets() error {
	return nil
}

func (s *Storage) getProvider(provider domain.StorageProvider) (stow.Location, error) {
	var (
		cont stow.Location
		err  error
	)

	switch provider {
	case domain.StorageLocal:
		cont, err = stow.Dial(local.Kind, stow.ConfigMap{
			local.ConfigKeyPath: paths.Get().Storage,
		})
	case domain.StorageAWS:
		cont, err = stow.Dial(s3.Kind, stow.ConfigMap{
			s3.ConfigAccessKeyID: s.env.AWSAccessKey,
			s3.ConfigSecretKey:   s.env.AWSSecret,
			s3.ConfigRegion:      "eu-west-1",
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
