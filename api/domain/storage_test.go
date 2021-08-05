// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuckets_IsValid(t *testing.T) {
	tt := map[string]struct {
		input   string
		buckets Buckets
		want    bool
	}{
		"Valid": {
			"bucket",
			Buckets{Bucket{Name: "bucket-name", ID: "bucket"}},
			true,
		},
		"Not Valid": {
			"bucket",
			Buckets{},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.buckets.IsValid(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestStorageProvider_String(t *testing.T) {
	got := StorageLocal.String()
	want := string(StorageLocal)
	assert.Equal(t, got, want)
}

func TestStorageProvider_IsLocal(t *testing.T) {
	tt := map[string]struct {
		input StorageProvider
		want  bool
	}{
		"Success": {
			StorageLocal,
			true,
		},
		"Remote": {
			StorageAWS,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.IsLocal()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestStorageProvider_TitleCase(t *testing.T) {
	tt := map[string]struct {
		input StorageProvider
		want  StorageProvider
	}{
		"Local": {
			StorageLocal,
			StorageProvider("Local"),
		},
		"AWS": {
			StorageAWS,
			StorageProvider("Aws"),
		},
		"Google": {
			StorageGCP,
			StorageProvider("Google"),
		},
		"Azure": {
			StorageAzure,
			StorageProvider("Azure"),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.TitleCase()
			assert.Equal(t, test.want, got)
		})
	}
}
