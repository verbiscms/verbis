// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"context"
	"fmt"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/services/storage/internal"
)

// Configuration represents the information returned
// by the client of the current state of storage.
type Configuration struct {
	Connected     bool                    `json:"connected"`
	Info          domain.StorageConfig    `json:"info"`
	Providers     domain.StorageProviders `json:"providers"`
	IsMigrating   bool                    `json:"is_migrating"`
	MigrationInfo *MigrationInfo          `json:"migration"`
}

// Info satisfies the Provider interface by returning a
// domain.StorageConfiguration.
func (s *Storage) Info(ctx context.Context) (Configuration, error) {
	info := s.service.Config()
	var m = make(domain.StorageProviders)
	for k, v := range internal.Providers {
		if k == domain.StorageLocal {
			continue
		}
		m[k] = v.Info(s.env)
	}

	connected := false
	if m[info.Provider].Connected && !info.Provider.IsLocal() {
		connected = true
	}

	isMigrating := s.isMigrating(ctx)

	fmt.Println(isMigrating)

	var migrationInfo *MigrationInfo
	if isMigrating {
		mi, err := s.getMigration()
		if err != nil {
			return Configuration{}, err
		}
		migrationInfo = mi
	}

	c := Configuration{
		Connected:     connected,
		Info:          info,
		Providers:     m,
		IsMigrating:   isMigrating,
		MigrationInfo: migrationInfo,
	}

	return c, nil
}
