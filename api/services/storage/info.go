// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/services/storage/internal"
)

// Configuration represents the information returned
// by the client of the current state of storage.
type Configuration struct {
	ActiveProvider domain.StorageProvider  `json:"active_provider"`
	ActiveBucket   string                  `json:"active_bucket"`
	Providers      domain.StorageProviders `json:"providers"`
	IsMigrating    bool                    `json:"is_migrating"`
	MigrationInfo  MigrationInfo           `json:"migration"`
}

// Info satisfies the Provider interface by returning a
// domain.StorageConfiguration.
func (s *Storage) Info() (Configuration, error) {
	provider, bucket, err := s.service.Config()
	if err != nil {
		return Configuration{}, err
	}

	var m = make(domain.StorageProviders)
	for k, v := range internal.Providers {
		m[k] = v.Info(s.env)
	}

	//var migrationInfo MigrationInfo
	//mi, found := cache.Get(context.Background(), MigrationCacheKey)
	//if found == nil {
	//	migrationInfo = mi.(MigrationInfo)
	//}

	c := Configuration{
		ActiveProvider: provider,
		ActiveBucket:   bucket,
		Providers:      m,
		//IsMigrating:    s.isMigrating,
		//MigrationInfo:  MigrationInfo{},
	}

	return c, nil
}
