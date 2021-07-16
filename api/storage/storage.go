// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/common/paths"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/storage/internal"
	"github.com/ainsleyclark/verbis/api/store/files"
	"github.com/ainsleyclark/verbis/api/store/options"
)

// Provider describes the main storage system for Verbis.
// A provider can be either remote (GCP, AWS or Azure),
// or it can be local, dependant on what is set on
// the environment.
type Provider interface {
	Info() (domain.StorageConfiguration, error)
	Migrate(from, to domain.StorageChange) (int, error)
	Container
	Bucket
}

// Container the methods for creating storage folders either
// remotely via GCP, AWS, Azure or Locally. Buckets can
// be set listed, created and deleted.
type Container interface {
	// ListBuckets retrieves all buckets that are currently in
	// the provider and returns a slice of domain.Buckets.
	// Returns errors.INVALID if there was an error obtaining the buckets.
	ListBuckets(provider domain.StorageProvider) (domain.Buckets, error)
	// CreateBucket creates a folder or bucket on the provider
	// by name.
	// Returns errors.INVALID if there was an error creating the bucket.
	CreateBucket(provider domain.StorageProvider, name string) error
	// DeleteBucket removes a folder or bucket from the
	// provider by name.
	// Returns errors.INVALID if there was an error deleting the bucket.
	DeleteBucket(provider domain.StorageProvider, name string) error
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

	ErrAlreadyMigrating = errors.New("migration is already in progress")
)

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

type Storage struct {
	env         *environment.Env
	optionsRepo options.Repository
	filesRepo   files.Repository
	paths       paths.Paths
	service     internal.StorageServices
	isMigrating bool
	migration   MigrationInfo
}

// Conf
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

	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	service := internal.NewService(cfg.Environment, cfg.Options)

	provider, _, err := service.Config()
	if err != nil {
		return nil, err
	}

	if !internal.Providers.Exists(provider) {
		return nil, &errors.Error{Code: errors.INVALID, Message: string("Error setting up storage with provider: " + provider), Operation: op, Err: ErrNoProvider}
	}

	s := &Storage{
		env:         cfg.Environment,
		optionsRepo: cfg.Options,
		filesRepo:   cfg.Files,
		paths:       paths.Get(),
		service:     service,
	}

	return s, nil
}
