// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/services/storage/internal"
	"github.com/verbiscms/verbis/api/store/files"
	"github.com/verbiscms/verbis/api/store/options"
)

// Provider describes the main storage system for Verbis.
// A provider can be either remote (GCP, AWS or Azure),
// or it can be local, dependant on what is set on
// the environment.
type Provider interface {
	// Info returns a domain.StorageConfiguration to provide
	// information about the state of the storage.
	// Which includes active provider and bucket,
	// and environment state for each provider.
	// Returns errors.INVALID if the options lookup failed.
	Info() (Configuration, error)
	// Save changes the current storage provider and bucket.
	// It will be validated before the options table is
	// updated.
	// Returns errors.INVALID if validation failed.
	// Returns errors.INTERNAL if there was a problem updating
	// the options table.
	Save(info domain.StorageChange) error
	Migrator
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
	CreateBucket(provider domain.StorageProvider, name string) (domain.Bucket, error)
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

// Migrator defines the methods used for migrating fields to
// different storage providers.
type Migrator interface {
	// Migrate migrates all files from one location to another.
	// If delete is set to true, the original files will be
	// deleted from the source destination. It returns
	// the total amount of files processing in the
	// background up on success.
	// Returns errors.INVALID if there is a migration already in progress
	// or the from and to providers are the same.
	// Returns errors.NOTFOUND if there were no files found with the from provider.
	Migrate(from, to domain.StorageChange, delete bool) (int, error)
}

// Storage represents the implementation of a Verbis
// storage provider.
type Storage struct {
	env         *environment.Env
	optionsRepo options.Repository
	filesRepo   files.Repository
	paths       paths.Paths
	service     internal.StorageServices
	cache       cache.Store
}

// Config defines the configuration passed to create a new
// storage instance.
type Config struct {
	Environment *environment.Env
	Options     options.Repository
	Files       files.Repository
	Cache       cache.Store
}

// Validate validates the configuration to ensure there are
// no nil values passed.
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
	if c.Cache == nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error, no cache set", Operation: op, Err: fmt.Errorf("nil cache store")}
	}
	return nil
}

// New creates a new Storage client. The configuration will
// be validated and a new storage client returned.
func New(cfg Config) (*Storage, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	s := &Storage{
		env:         cfg.Environment,
		optionsRepo: cfg.Options,
		filesRepo:   cfg.Files,
		paths:       paths.Get(),
		service: &internal.Service{
			Env:     cfg.Environment,
			Options: cfg.Options,
		},
	}
	return s, nil
}
