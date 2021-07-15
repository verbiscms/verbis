// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

type (
	// Bucket represents a named group of Files, it could
	// be remote or local.
	Bucket struct {
		Id   string `json:"id" binding:"required"` //nolint
		Name string `json:"name"`                  //nolint
	}
	// Buckets represents the slice of Bucket's.
	Buckets []Bucket
	// StorageProvider represents a the string of a provider
	// for writing storage files.
	StorageProvider string

	StorageInfo struct {
		Provider StorageProvider `json:"provider" binding:"required"`
		Bucket   string          `json:"bucket" binding:"required"`
	}

	StorageConfiguration struct {
		ActiveProvider StorageProvider  `json:"active_provider"`
		ActiveBucket   string           `json:"active_bucket"`
		Providers      StorageProviders `json:"providers"`
	}
	// StorageProviders represents the map of providers that
	// are available within Verbis.
	StorageProviders    map[StorageProvider]StorageProviderInfo
	StorageProviderInfo struct {
		DialMessage string `json:"dial_message"`
		EnvSet      bool   `json:"env_set"`
	}
)

// String is the stringer on the StorageProvider.
func (s StorageProvider) String() string {
	return string(s)
}

// IsLocal determines if the current storage provider is
// local. Returns false if it's remote.
func (s StorageProvider) IsLocal() bool {
	return s == StorageLocal
}

const (
	// StorageLocal represents the string for the local
	// storage disk.
	StorageLocal = StorageProvider("local")
	// StorageAWS represents the string for the AWS storage
	// disk.
	StorageAWS = StorageProvider("aws")
	// StorageGCP represents the string for the GCP storage
	// disk.
	StorageGCP = StorageProvider("google")
	// StorageAzure represents the string for the Azure
	// storage disk.
	StorageAzure = StorageProvider("azure")
)
