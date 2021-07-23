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
const pageSize = 100

// ListBuckets satisfies the Container interface by listing
// all the available buckets in the provider.
func (s *Storage) ListBuckets(provider domain.StorageProvider) (domain.Buckets, error) {
	const op = "Storage.ListBuckets"

	prov, err := s.service.Provider(provider)
	if err != nil {
		return nil, err
	}

	var (
		cursor     = stow.CursorStart
		buckets    = make(domain.Buckets, 0)
		containers []stow.Container
		contErr    error
	)

	for {
		containers, cursor, contErr = prov.Containers("", cursor, pageSize)
		if contErr != nil {
			return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error obtaining buckets", Operation: op, Err: err}
		}
		for _, container := range containers {
			buckets = append(buckets, domain.Bucket{
				Id:   container.ID(),
				Name: container.Name(),
			})
		}
		if stow.IsCursorEnd(cursor) {
			break
		}
	}

	return buckets, nil
}

// CreateBucket satisfies the Container interface by
// accepting a name, and creating a new bucket
// from the provider.
func (s *Storage) CreateBucket(provider domain.StorageProvider, name string) (domain.Bucket, error) {
	const op = "Storage.CreateBucket"

	prov, err := s.service.Provider(provider)
	if err != nil {
		return domain.Bucket{}, err
	}

	bucket, err := prov.CreateContainer(name)
	if err != nil {
		return domain.Bucket{}, &errors.Error{Code: errors.INVALID, Message: "Error creating bucket with the name: " + name, Operation: op, Err: err}
	}

	return domain.Bucket{
		Id:   bucket.ID(),
		Name: bucket.Name(),
	}, nil
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
