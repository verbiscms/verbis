// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/graymeta/stow"
)

// pageSize is the number of containers to retrieve in the
// ListBuckets function.
const pageSize = 999999

// ListBuckets satisfies the Container interface by listing
// all the available buckets in the provider.
func (s *Storage) ListBuckets(provider domain.StorageProvider) (domain.Buckets, error) {
	const op = "Storage.ListBuckets"

	prov, err := s.service.Provider(provider)
	if err != nil {
		return nil, err
	}

	var buckets = make(domain.Buckets, 0)
	err = stow.WalkContainers(prov, stow.NoPrefix, pageSize, func(c stow.Container, err error) error {
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
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error obtaining buckets", Operation: op, Err: err}
	}

	return buckets, nil
}

// CreateBucket satisfies the Container interface by
// accepting a name, and creating a new bucket
// from the provider.
func (s *Storage) CreateBucket(provider domain.StorageProvider, name string) error {
	const op = "Storage.CreateBucket"

	prov, err := s.service.Provider(provider)
	if err != nil {
		return err
	}

	_, err = prov.CreateContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error creating bucket: " + name, Operation: op, Err: err}
	}

	return nil
}

// DeleteBucket satisfies the Container interface by
// accepting a name, and deleting a bucket from
// the provider.
func (s *Storage) DeleteBucket(provider domain.StorageProvider, name string) error {
	const op = "Storage.DeleteBucket"

	prov, err := s.service.Provider(provider)
	if err != nil {
		return err
	}

	err = prov.RemoveContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting bucket with the name: " + name, Operation: op, Err: err}
	}

	return nil
}
