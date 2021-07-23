// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// Save satisfies the Provider interface by changing the
// current storage providers by updating the options
// table.
func (s *Storage) Save(info domain.StorageChange) error {
	const op = "Storage.Save"

	err := s.validate(info)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Validation failed", Operation: op, Err: err}
	}

	// Ensure no bucket is set if it's local
	if info.Provider.IsLocal() {
		info.Bucket = ""
	}

	err = s.optionsRepo.Update("storage_provider", info.Provider)
	if err != nil {
		return err
	}

	err = s.optionsRepo.Update("storage_bucket", info.Bucket)
	if err != nil {
		return err
	}

	return nil
}
