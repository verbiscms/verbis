// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/storage/internal"
)

func (s *Storage) Info() (domain.StorageConfiguration, error) {
	provider, bucket, err := s.service.Config()
	if err != nil {
		return domain.StorageConfiguration{}, err
	}

	var m = make(domain.StorageProviders, 0)
	for k, v := range internal.Providers {
		m[k] = v.Info(s.env)
	}

	c := domain.StorageConfiguration{
		ActiveProvider: provider,
		ActiveBucket:   bucket,
		Providers:      m.Sort(),
	}

	return c, nil
}
