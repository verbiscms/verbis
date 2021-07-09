// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/storage/internal"
	"github.com/ainsleyclark/verbis/api/store/files"
	"github.com/ainsleyclark/verbis/api/store/options"
	"github.com/graymeta/stow"
)

// Provider describes the main storage system for Verbis.
// A provider can be either remote (GCP, AWS or Azure),
// or it can be local, dependant on what is set on
// the environment.
type Provider interface {
	// SetProvider sets the storage system to the given
	// provider.
	// Returns errors.INVALID if there was an error setting the provider.
	SetProvider(provider domain.StorageProvider) error
	Container
	Bucket
}

// Container the methods for creating storage folders either
// remotely via GCP, AWS, Azure or Locally. Buckets can
// be set listed, created and deleted.
type Container interface {
	// SetBucket updates the storage system with a new bucket.
	// The options table will be updated from the database.
	// Returns errors.INVALID if the bucket could not be set.
	// Returns errors.INTERNAL if there was an error updating the options DB table.
	SetBucket(id string) error
	// ListBuckets retrieves all buckets that are currently in
	// the provider and returns a slice of domain.Buckets.
	// Returns errors.INVALID if there was an error obtaining the buckets.
	ListBuckets() (domain.Buckets, error)
	// CreateBucket creates a folder or bucket on the provider
	// by name.
	// Returns errors.INVALID if there was an error creating the bucket.
	CreateBucket(name string) error
	// DeleteBucket removes a folder or bucket from the
	// provider by name.
	// Returns errors.INVALID if there was an error deleting the bucket.
	DeleteBucket(name string) error
}

// Bucket describes the methods used for interacting with
// the Verbis storage system. The client can be remote
// or work with the local file system dependant on
// what is set in the options table.
type Bucket interface {
	// Find looks up the file with the given URL and retrieves
	// the appropriate bucket to obtain the file contents.
	// It returns the byte value of the file as well as
	// the domain.File.
	// Returns errors.INTERNAL if the file could not be opened or read.
	// Returns errors.NOTFOUND if the file could not be retrieved from the bucket.
	Find(url string) ([]byte, domain.File, error)
	// Upload adds a domain.Upload to the database as well as
	// the bucket that is currently set in the env. The
	// file is seeked, the mime type is obtained and it
	// is inserted into the database and uploaded to
	// the bucket.
	// Returns errors.INVALID if the bucket could not be obtained.
	// Returns errors.INTERNAL if the contents couldn't be seeked or the mime
	// type could not be obtained.
	Upload(upload domain.Upload) (domain.File, error)
	// Delete removes an item from the the bucket. It first retrieves
	// the file by a lookup from the database, obtains the correct
	// bucket, then removes the file from the storage provider.
	// The file data will also be deleted from
	// the database.
	// Returns errors.INVALID if the file could not be deleted from the bucket.
	Delete(id int) error
	// Exists queries the database by the given name and
	// returns true if there was a match.
	Exists(name string) bool
}

var (
	// ErrNoProvider is returned by New and SetProvider when
	// there is no match from the options table.
	ErrNoProvider = errors.New("invalid provider")
)

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

type Storage struct {
	ProviderName  domain.StorageProvider
	env           *environment.Env
	optionsRepo   options.Repository
	filesRepo     files.Repository
	options       *domain.Options
	paths         paths.Paths
	stowLocation  stow.Location
	stowContainer stow.Container
	service       internal.StorageServices
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

//
type Config struct {
	Environment *environment.Env
	Options     options.Repository
	Files       files.Repository
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (c Config) Validate() error {
	const op = "Storage.Config.Validate"
	if c.Environment == nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error, no Environment set", Operation: op, Err: fmt.Errorf("nil environment")}
	}
	if c.Options == nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error, no options repository set", Operation: op, Err: fmt.Errorf("nil options store")}
	}
	if c.Files == nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error, no files repository set", Operation: op, Err: fmt.Errorf("nil files store")}
	}
	return nil
}

// New parse config
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

func New(cfg Config) (*Storage, error) {
	const op = "Storage.New"

	// Validate configuration
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	var (
		service = internal.NewService(cfg.Environment)
		opts    = cfg.Options.Struct()
	)

	provider := opts.StorageProvider
	if provider == "" {
		provider = domain.StorageLocal
	}

	if !provider.Validate() {
		return nil, &errors.Error{Code: errors.INVALID, Message: string("Error setting up storage with provider: " + provider), Operation: op, Err: ErrNoProvider}
	}

	s := &Storage{
		ProviderName: provider,
		env:          cfg.Environment,
		optionsRepo:  cfg.Options,
		filesRepo:    cfg.Files,
		options:      opts,
		paths:        paths.Get(),
		service:      service,
	}

	// Set Provider
	err = s.SetProvider(provider)
	if err != nil {
		return nil, err
	}

	// Set Bucket
	// TODO, do we want to be srtting options on initialisation?
	err = s.SetBucket(opts.StorageBucket)
	if err != nil {
		return nil, err
	}

	return s, nil
}
