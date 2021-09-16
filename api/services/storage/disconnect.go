// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// ErrAlreadyDisconnected is returned by Disconnected when
// the provider is already local.
var ErrAlreadyDisconnected = errors.New("provider is already local, skipping")

// Disconnect satisfies the Provider interface by disconnecting
// the current storage provider.
func (s *Storage) Disconnect() error {
	const op = "Storage.Disconnect"

	cfg := s.service.Config()
	if cfg.Provider.IsLocal() {
		return &errors.Error{Code: errors.INVALID, Message: "Error, storage is already disconnected", Operation: op, Err: ErrAlreadyDisconnected}
	}

	err := s.optionsRepo.Insert(domain.OptionsDBMap{
		"storage_provider": domain.StorageLocal,
		"storage_bucket":   "",
	})

	if err != nil {
		return err
	}

	return nil
}
