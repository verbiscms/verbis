// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"sort"
	"strings"
)

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

	StorageConfiguration struct {
		ActiveProvider StorageProvider  `json:"active_provider"`
		ActiveBucket   string           `json:"active_bucket"`
		Providers      StorageProviders `json:"providers"`
	}
	// StorageProviders represents the map of providers that
	// are available within Verbis.
	StorageProviders map[StorageProvider]StorageProviderInfo

	StorageProviderInfo struct {
		key             StorageProvider
		Name            string      `json:"name"`
		Order           int         `json:"-"`
		Connected       bool        `json:"connected"`
		Error           interface{} `json:"error"`
		Instructions    string      `json:"instructions"`
		EnvironmentKeys []string    `json:"environment_keys"`
		EnvironmentSet  bool        `json:"environment_set"`
	}
	StorageChange struct {
		Provider StorageProvider `json:"provider" binding:"required"`
		Bucket   string          `json:"bucket"`
		Region   string          `json:"region"`
	}
)

func (b Buckets) IsValid(bucket string) bool {
	for _, v := range b {
		if v.Id == bucket {
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

type providers []StorageProviderInfo

// Len is part of sort.Interface.
func (p providers) Len() int {
	return len(p)
}

// Less is part of sort.Interface. We use count as the
// value to sort by.
func (p providers) Less(i, j int) bool {
	return p[i].Order < p[j].Order
}

// Swap is part of sort.Interface.
func (p providers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (s StorageProviders) Sort() StorageProviders {
	pp := make(providers, 0, len(s))
	for k, v := range s {
		v.key = k
		pp = append(pp, v)
	}
	sort.Sort(pp)
	m := make(StorageProviders, len(pp))
	for _, v := range pp {
		m[v.key] = v
	}
	return m
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
