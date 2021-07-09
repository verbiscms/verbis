// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// SetProvider satisfies the Provider interface by a
// domain.StorageProvider and updating the var
// used by storage
// an id and updating the stowLocation of storage. The
// options table is also updated.
func (s *Storage) SetProvider(provider domain.StorageProvider) error {
	const op = "Storage.SetProvider"

	location, err := s.service.Provider(provider)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error setting provider", Operation: op, Err: err}
	}

	s.stowLocation = location
	s.ProviderName = provider

	return nil
}
