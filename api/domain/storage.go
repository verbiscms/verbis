// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"strings"
)

type (
	// Bucket represents a named group of Files, it could
	// be remote or local.
	Bucket struct {
		ID   string `json:"id" binding:"required"`
		Name string `json:"name"`
	}
	// Buckets represents the slice of Bucket's.
	Buckets []Bucket
	// StorageProvider represents a the string of a provider
	// for writing storage files.
	StorageProvider string
	// StorageProviders represents the map of providers that
	// are available within Verbis.
	StorageProviders map[StorageProvider]StorageProviderInfo
	// StorageProviderInfo represents the data of a storage
	// provider, including if the provider is connected,
	// environment keys and relevant names etc.
	StorageProviderInfo struct {
		Name            string      `json:"name"`
		Order           int         `json:"-"`
		Connected       bool        `json:"connected"`
		Error           interface{} `json:"error"`
		Instructions    string      `json:"instructions"`
		EnvironmentKeys []string    `json:"environment_keys"`
		EnvironmentSet  bool        `json:"environment_set"`
	}
	// StorageConfig represents the data needed to change or
	// modify and existing storage provider or bucket.
	StorageConfig struct {
		Provider     StorageProvider `json:"provider" binding:"required"`
		Bucket       string          `json:"bucket"`
		Region       string          `json:"region"`
		UploadRemote bool            `json:"upload_remote"`
		LocalBackup  bool            `json:"local_backup"`
		RemoteBackup bool            `json:"remote_backup"`
	}
)

// IsValid checks to see if a bucket passed is represented
// in the buckets slice. Returns false if there is no
// match.
func (b Buckets) IsValid(bucket string) bool {
	for _, v := range b {
		if v.ID == bucket {
			return true
		}
	}
	return false
}

// String is the stringer on the StorageProvider.
func (s StorageProvider) String() string {
	return string(s)
}

// IsLocal determines if the current storage provider is
// local. Returns false if it's remote.
func (s StorageProvider) IsLocal() bool {
	return s == StorageLocal
}

// TitleCase returns a StorageProvider with title case.
func (s StorageProvider) TitleCase() StorageProvider {
	return StorageProvider(strings.Title(s.String()))
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
