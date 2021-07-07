// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
	"io"
	"path/filepath"
	"strings"
)

type (
	// File represents a storage file which could be stored
	// in the cloud, or on the local file system.
	File struct {
		Id         int             `db:"id" json:"-" binding:"numeric"` //nolint
		UUID       uuid.UUID       `db:"uuid" json:"uuid"`
		URL        string          `db:"url" json:"url"`
		Name       string          `db:"name" json:"name"`
		Path       string          `db:"path" json:"path"`
		Mime       Mime            `db:"mime" json:"mime"`
		SourceType string          `db:"source_type" json:"source_type"`
		Provider   StorageProvider `db:"provider" json:"provider"`
		Region     string          `db:"region" json:"region"`
		Bucket     string          `db:"bucket" json:"bucket"`
		FileSize   int64           `db:"file_size" json:"file_size"`
		Private    types.BitBool   `db:"private" json:"private"`
	}
	Upload struct {
		Path       string
		Size       int64
		Contents   io.ReadSeeker
		Private    bool
		SourceType string
	}
	// Files represents the slice of File's.
	Files []File
)

const (
	MediaSourceType          = "media"
	MediaSizeSourceType      = "media_sizes"
	FormAttachmentSourceType = "form_attachment"
)

// PrivatePath retrieves the full path for the upload. If
// the file is local, the prefix will be added.
func (f *File) PrivatePath(prefix string) string {
	if f.Provider == StorageLocal {
		return filepath.Join(prefix, f.Path, f.UUID.String()+f.Extension())
	}
	return strings.TrimSuffix(f.Path, "/") + "/" + f.UUID.String() + f.Extension()
}

// Extension returns the extension of the file by stripping
// from the url.
func (f *File) Extension() string {
	return filepath.Ext(f.Name)
}

// IsLocal determines if the file is stored on the local
// file system.
// Returns false if it's from a cloud provider.
func (f *File) IsLocal() bool {
	return f.Provider == StorageLocal
}

func (u *Upload) Validate() error {
	if u.Path == "" {

	}
	return nil
}

// Mime is a string representation of a MIME type.
type Mime string

// CanResize Returns true if the mime type is of JPG or
// PNG, determining if the image can be resized.
func (m Mime) CanResize() bool {
	return m.IsJPG() || m.IsPNG()
}

// IsJPG returns true if the mime type is of JPG.
func (m Mime) IsJPG() bool {
	return m == "image/jpeg" || m == "image/jp2"
}

// IsPNG returns true if the mime type is of PNG.
func (m Mime) IsPNG() bool {
	return m == "image/png"
}

// String is the stringer on Mime type.
func (m Mime) String() string {
	return string(m)
}
