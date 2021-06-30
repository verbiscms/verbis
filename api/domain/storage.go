// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"net/url"
	"strings"
)

type (
	// Bucket
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	Bucket struct {
		Id   string `json:"id" binding:"required"` //nolint
		Name string `json:"name"`                  //nolint
	}
	// Buckets represents the slice of Bucket's.
	Buckets []Bucket
	// StorageProvider
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	StorageProvider string
	// StorageFile
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	StorageFile struct {
		URI           *url.URL
		BaseLocalPath string
		ID            string
	}
)

const (
	// StorageLocal
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	StorageLocal = StorageProvider("local")
	// StorageAWS
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	StorageAWS = StorageProvider("aws")
	// StorageGCP
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	StorageGCP = StorageProvider("google")
	// StorageAzure
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	StorageAzure = StorageProvider("azure")
)

// clean this up, bad name
func (s StorageFile) ToURL(prefix string) string {
	if s.Provider() == StorageLocal {
		return "/" + s.CleanPath()
	}
	return s.URI.String()
}

// Provider
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (s *StorageFile) Provider() StorageProvider {
	switch s.URI.Scheme {
	case "file":
		return StorageLocal
	case "s3":
		return StorageAWS
	case "google":
		return StorageGCP
	case "azure":
		return StorageAzure
	}
	return ""
}

// CleanPath
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (s *StorageFile) CleanPath() string {
	if s.Provider() == StorageLocal {
		return strings.TrimPrefix(strings.ReplaceAll(s.URI.Path, s.BaseLocalPath, ""), "/")
	}
	return s.URI.Path
}
