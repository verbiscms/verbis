// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/store/files"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/graymeta/stow"
	_ "github.com/graymeta/stow/azure"
	_ "github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/s3"
	"io/ioutil"
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

type Client interface {
	Find(path string) ([]byte, domain.File, error)
	Delete(id int) error
	Upload(upload domain.Upload) (domain.File, error)
	SetProvider(location domain.StorageProvider) error
	SetBucket(id string) error
	ListBuckets() (domain.Buckets, error)
}

type Storage struct {
	ProviderName domain.StorageProvider
	Local        bool
	provider     stow.Location
	bucket       stow.Container
	opts         *domain.Options
	env          *environment.Env
	paths        paths.Paths
	repo         files.Repository
}

var (
	ErrNoProvider  = errors.New("invalid provider")
	ErrInvalidMime = errors.New("error obtaining mime type")
)

// New parse config
func New(env *environment.Env, opts *domain.Options, repo files.Repository) (Client, error) {
	const op = "Storage.New"

	s := &Storage{
		env:   env,
		opts:  opts,
		paths: paths.Get(),
		Local: false,
		repo:  repo,
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
func (s *Storage) Upload(u domain.Upload) (domain.File, error) {
	const op = "Storage.Upload"

	key := uuid.New()
	absPath := strings.TrimPrefix(filepath.Dir(u.Path)+key.String()+filepath.Ext(u.Path), ".")

	item, err := s.bucket.Put(absPath, u.Contents, u.Size, nil)
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error uploading file to storage provider", Operation: op, Err: err}
	}

	_, err = u.Contents.Seek(0, 0)
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error seeking bytes", Operation: op, Err: err}
	}

	m, err := mimetype.DetectReader(u.Contents)
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error obtaining mime type", Operation: op, Err: err}
	}

	dbPath := path.Dir(item.URL().Path)
	if s.Local {
		dbPath = strings.TrimPrefix(strings.ReplaceAll(dbPath, s.paths.Storage, ""), "/")
	}

	f := domain.File{
		UUID:       key,
		URL:        "/" + strings.TrimSuffix(u.Path, "/"),
		Name:       path.Base(item.ID()),
		Path:       dbPath,
		Mime:       domain.Mime(m.String()),
		SourceType: u.SourceType,
		Provider:   s.ProviderName,
		Region:     "TODO",
		Bucket:     s.bucket.ID(),
		FileSize:   u.Size,
		Private:    false,
	}

	create, err := s.repo.Create(f)
	if err != nil {
		return domain.File{}, err
	}

	return create, nil
}

func (s *Storage) Find(uri string) ([]byte, domain.File, error) {
	const op = "Storage.Find"

	file, err := s.repo.FindByURL(uri)
	if err != nil {
		return nil, domain.File{}, err
	}

	uploadPath := file.PrivatePath(s.paths.Storage)

	item, err := s.provider.ItemByURL(&url.URL{Path: uploadPath})
	if err != nil {
		return nil, domain.File{}, &errors.Error{Code: errors.NOTFOUND, Message: "Error obtaining file with the path: " + uploadPath, Operation: op, Err: err}
	}

	open, err := item.Open()
	if err != nil {
		return nil, domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error opening file", Operation: op, Err: err}
	}
	defer open.Close()

	buf, err := ioutil.ReadAll(open)
	if err != nil {
		return nil, domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error reading file", Operation: op, Err: err}
	}

	return buf, file, nil
}

// Delete removes an item from the bucket.
func (s *Storage) Delete(id int) error {
	const op = "Storage.Delete"

	file, err := s.repo.Find(id)
	if err != nil {
		return err
	}

	bucket, err := s.getBucket(file)
	if err != nil {
		// TODO errros
		return err
	}

	err = bucket.RemoveItem(file.PrivatePath(s.paths.Storage))
	if err != nil {
		// TODO errros
		return err
	}

	err = s.repo.Delete(file.Id)
	if err != nil {
		return err
	}

	return nil
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
