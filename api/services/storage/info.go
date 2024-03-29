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
	Connected     bool                    `json:"connected"`
	Info          domain.StorageConfig    `json:"info"`
	Providers     domain.StorageProviders `json:"providers"`
	IsMigrating   bool                    `json:"is_migrating"`
	MigrationInfo *MigrationInfo          `json:"migration"`
}

// Info satisfies the Provider interface by returning a
// domain.StorageConfiguration.
func (s *Storage) Info(ctx context.Context) Configuration {
	info := s.service.Config()
	var m = make(domain.StorageProviders)
	for k, v := range internal.Providers {
		if k == domain.StorageLocal {
			continue
		}
		m[k] = v.Info(s.env)
	}

	if info.Provider.IsLocal() {
		info.Provider = ""
	}

	connected := false
	if m[info.Provider].Connected && !info.Provider.IsLocal() {
		connected = true
	}

	c := Configuration{
		Connected:     connected,
		Info:          info,
		Providers:     m,
		IsMigrating:   s.isMigrating(ctx),
		MigrationInfo: s.getMigration(ctx),
	}

	return c
}
