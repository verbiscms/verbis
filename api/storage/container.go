// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/graymeta/stow"
)

func (s *Storage) SetBucket(id string) error {
	const op = "Storage.SetBucket"

	if s.options.StorageProvider == domain.StorageLocal {
		id = ""
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

func (s *Storage) ListBuckets() (domain.Buckets, error) {
	const op = "Container.ListBuckets"

	var buckets = make(domain.Buckets, 0)
	err := stow.WalkContainers(s.stowLocation, stow.NoPrefix, 100, func(c stow.Container, err error) error {
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

	return nil, nil
}

func (s *Storage) CreateBucket(name string) error {
	const op = "Container.CreateBucket"

	_, err := s.stowLocation.CreateContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error creating bucket: " + name, Operation: op, Err: err}
	}

	return nil
}

func (s *Storage) DeleteBucket(name string) error {
	const op = "Container.DeleteBucket"

	err := s.stowLocation.RemoveContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting bucket: " + name, Operation: op, Err: err}
	}

	return nil
}
