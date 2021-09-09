// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"context"
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
	MigrationInfo  *MigrationInfo          `json:"migration"`
	LocalBackup    bool                    `json:"local_backup"`
}

// Info satisfies the Provider interface by returning a
// domain.StorageConfiguration.
func (s *Storage) Info(ctx context.Context) (Configuration, error) {
	info := s.service.Config()
	var m = make(domain.StorageProviders)
	for k, v := range internal.Providers {
		m[k] = v.Info(s.env)
	}

	isMigrating := s.isMigrating(ctx)
	var migrationInfo *MigrationInfo
	if isMigrating {
		mi, err := s.getMigration()
		if err != nil {
			return Configuration{}, err
		}
		migrationInfo = mi
	}

	c := Configuration{
		ActiveProvider: info.Provider,
		ActiveBucket:   info.Bucket,
		Providers:      m,
		IsMigrating:    isMigrating,
		MigrationInfo:  migrationInfo,
		LocalBackup:    info.LocalBackup,
	}

	return c, nil
}
