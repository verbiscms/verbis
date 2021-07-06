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
	Find(path string) ([]byte, domain.File, error)
	Delete(path string) error
	Upload(path string, size int64, contents io.Reader) (domain.File, error)
	IsLocal() bool
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

var (
	ErrNoProvider = errors.New("invalid provider")
)

// New parse config
func New(env *environment.Env, opts *domain.Options) (Client, error) {
	const op = "Storage.New"

	s := &Storage{
		env:   env,
		opts:  opts,
		paths: paths.Get(),
		local: false,
	}

	provider := opts.StorageProvider
	if provider == "" {
		provider = domain.StorageLocal
	}

	if !validate(provider) {
		return nil, &errors.Error{Code: errors.INVALID, Message: string("Error setting up storage with provider: " + provider), Operation: op, Err: ErrNoProvider}
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
func (s *Storage) Upload(path string, size int64, contents io.Reader) (domain.File, error) {
	const op = "Storage.Upload"

	item, err := s.bucket.Put(path, contents, size, nil)
	if err != nil {
		return domain.File{}, err
	}

	sp := domain.File{
		URL:    item.URL().Path,
		Name:   item.ID(),
		Bucket: s.bucket.Name(),
	}

	s.bucket.ID()

	return sp, nil
}

func (s *Storage) Find(path string) ([]byte, domain.File, error) {
	const op = "Storage.Find"

	if !s.local {
		path = s.bucket.Name() + "/" + path
	} else {
		path = filepath.Join(s.paths.Storage, path)
	}

	item, err := s.provider.ItemByURL(&url.URL{Path: path})
	if err != nil {
		return nil, domain.File{}, &errors.Error{Code: errors.NOTFOUND, Message: "Error obtaining file with the path: " + path, Operation: op, Err: err}
	}

	file, err := item.Open()
	if err != nil {
		return nil, domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error opening file", Operation: op, Err: err}
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error reading file", Operation: op, Err: err}
	}

	return buf, toVerbisStorage(item), nil
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

func (s *Storage) IsLocal() bool {
	return s.local
}

// InSlice checks if a string exists in a slice,
func validate(p domain.StorageProvider) bool {
	for _, sp := range domain.StorageProviders {
		if sp == p {
			return true
		}
	}
	return false
}

func toVerbisStorage(item stow.Item) domain.File {
	return domain.File{
		URL:        item.URL().Path,
		Name:       "",
		Path:       "",
		Mime:       "",
		Provider:   "",
		Region:     "",
		Bucket:     "",
		FileSize:   0,
		SourceType: "",
		Private:    false,
	}
}
