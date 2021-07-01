// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/graymeta/stow"
	_ "github.com/graymeta/stow/azure"
	_ "github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/s3"
	"io"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

type Client interface {
	Find(path string) ([]byte, domain.StorageFile, error)
	Delete(path string) error
	Upload(path string, size int64, contents io.Reader) (domain.StorageFile, error)
	SetProvider(location domain.StorageProvider) error
	SetBucket(id string) error
	ListBuckets() (domain.Buckets, error)
}

type Storage struct {
	provider stow.Location
	bucket   stow.Container
	opts     *domain.Options
	env      *environment.Env
	paths    paths.Paths
	local    bool
}

// New parse config
func New(env *environment.Env, opts *domain.Options) (Client, error) {
	s := &Storage{
		env:   env,
		opts:  opts,
		paths: paths.Get(),
	}

	provider := opts.StorageProvider
	if provider == "" {
		provider = domain.StorageLocal
	}

	err := s.SetProvider(provider)
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
func (s *Storage) Upload(path string, size int64, contents io.Reader) (domain.StorageFile, error) {
	const op = "Storage.Upload"

	item, err := s.bucket.Put(path, contents, size, nil)
	if err != nil {
		return domain.StorageFile{}, err
	}

	sp := domain.StorageFile{
		URI:           item.URL(),
		BaseLocalPath: s.paths.Storage,
		ID:            item.ID(),
	}

	return sp, nil
}

func (s *Storage) Find(path string) ([]byte, domain.StorageFile, error) {
	const op = "Storage.Find"

	if !s.local {
		path = s.bucket.Name() + "/" + path
	} else {
		path = filepath.Join(s.paths.Storage, path)
	}

	item, err := s.provider.ItemByURL(&url.URL{Path: path})
	if err != nil {
		return nil, domain.StorageFile{}, &errors.Error{Code: errors.NOTFOUND, Message: "Error obtaining file with the path: " + path, Operation: op, Err: err}
	}

	file, err := item.Open()
	if err != nil {
		return nil, domain.StorageFile{}, &errors.Error{Code: errors.INTERNAL, Message: "Error opening file", Operation: op, Err: err}
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, domain.StorageFile{}, &errors.Error{Code: errors.INTERNAL, Message: "Error reading file", Operation: op, Err: err}
	}

	sp := domain.StorageFile{
		URI:           item.URL(),
		BaseLocalPath: s.paths.Storage,
		ID:            item.ID(),
	}

	return buf, sp, nil
}

func (s *Storage) Delete(path string) error {
	const op = "Storage.Delete"

	if s.local {
		path = filepath.Join(s.paths.Storage, path)
	}

	err := s.bucket.RemoveItem(path)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting file", Operation: op, Err: err}
	}

	return nil
}
