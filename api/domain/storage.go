// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

type StorageProvider string

type Bucket struct {
	Id string  `json:"id" binding:"required"` //nolint
	Name string  `json:"name"` //nolint
}

type Buckets []Bucket

const (
	StorageLocal = StorageProvider("local")
	StorageAWS   = StorageProvider("aws")
	StorageGCP   = StorageProvider("google")
	StorageAzure = StorageProvider("azure")
)
