// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/graymeta/stow"
)

// SetBucket satisfies the Container interface by accepting
// an id and updating the stowLocation of storage. The
// options table is also updated.
func (s *Storage) SetBucket(id string) error {
	const op = "Storage.SetBucket"

	if s.options.StorageProvider.IsLocal() {
		return &errors.Error{Code: errors.INVALID, Message: "Error setting bucket", Operation: op, Err: ErrLocalBucket}
	}

	container, err := s.stowLocation.Container(id)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error setting bucket", Operation: op, Err: err}
	}
	s.stowContainer = container

	err = s.optionsRepo.Update("storage_bucket", id)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating options table with new bucket", Operation: op, Err: err}
	}

	return nil
}

// pageSize is the number of containers to retrieve in the
// ListBuckets function.
const pageSize = 999999

// ListBuckets satisfies the Container interface by listing
// all the available buckets in the provider.
func (s *Storage) ListBuckets() (domain.Buckets, error) {
	const op = "Storage.ListBuckets"

	if s.options.StorageProvider.IsLocal() {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error listing buckets", Operation: op, Err: ErrLocalBucket}
	}

	var buckets = make(domain.Buckets, 0)
	err := stow.WalkContainers(s.stowLocation, stow.NoPrefix, pageSize, func(c stow.Container, err error) error {
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

	return buckets, nil
}

// CreateBucket satisfies the Container interface by
// accepting a name, and creating a new bucket
// from the provider.
func (s *Storage) CreateBucket(name string) error {
	const op = "Storage.CreateBucket"

	if s.options.StorageProvider.IsLocal() {
		return &errors.Error{Code: errors.INVALID, Message: "Error creating bucket", Operation: op, Err: ErrLocalBucket}
	}

	_, err := s.stowLocation.CreateContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error creating bucket: " + name, Operation: op, Err: err}
	}

	return nil
}

// DeleteBucket satisfies the Container interface by
// accepting a name, and deleting a bucket from
// the provider.
func (s *Storage) DeleteBucket(name string) error {
	const op = "Storage.DeleteBucket"

	if s.options.StorageProvider.IsLocal() {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting bucket", Operation: op, Err: ErrLocalBucket}
	}

	err := s.stowLocation.RemoveContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting bucket with the name: " + name, Operation: op, Err: err}
	}

	return nil
}
