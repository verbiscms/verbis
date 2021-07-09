// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import "github.com/ainsleyclark/verbis/api/domain"

func (s *Storage) Name() domain.StorageProvider {
	return s.ProviderName
}

func (s *Storage) SetProvider(provider domain.StorageProvider) error {
	location, err := s.service.Provider(provider)
	if err != nil {
		return err
	}

	s.stowLocation = location

	return nil
}
