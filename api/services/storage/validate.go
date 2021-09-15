// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE downloadFile.

package storage

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/domain"
)

// validate determines if the domain.StorageConfig is
// valid before proceeding. It will check if the
// bucket is empty (if it is a remote source),
// check if the provider is connected, and
// see if the bucket is valid from the
// given provider.
func (s *Storage) validate(info domain.StorageConfig) error {
	// TODO, we don't need this, we just need a disconnect?
	if !info.Provider.IsLocal() && info.Bucket == "" {
		return fmt.Errorf("bucket cannot be empty")
	}

	cfg := s.Info(context.Background())

	// If the configuration provider is not connected
	// and the provider is not local, return the
	// provider connection error.
	if !cfg.Providers[info.Provider].Connected && !info.Provider.IsLocal() {
		return fmt.Errorf(cast.ToString(cfg.Providers[info.Provider].Error))
	}

	// Obtain the buckets from the provider for matching.
	buckets, err := s.ListBuckets(info.Provider)
	if err != nil {
		return err
	}

	// TODO, do we need local checks?
	// Compare the bucket passed with the buckets listed
	// within the provider to see if it exists.
	if !info.Provider.IsLocal() && !buckets.IsValid(info.Bucket) {
		return fmt.Errorf("invalid storage bucket: %s", info.Bucket)
	}

	return nil
}
