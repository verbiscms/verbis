// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gabriel-vasile/mimetype"
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
		Id         int             `db:"id" json:"-"` //nolint
		UUID       uuid.UUID       `db:"uuid" json:"uuid"`
		Url        string          `db:"url" json:"url"`
		Name       string          `db:"name" json:"name"`
		BucketId   string          `db:"bucket_id" json:"bucket_id"`
		Mime       Mime            `db:"mime" json:"mime"`
		SourceType string          `db:"source_type" json:"source_type"`
		Provider   StorageProvider `db:"provider" json:"provider"`
		Region     string          `db:"region" json:"region"`
		Bucket     string          `db:"bucket" json:"bucket"`
		FileSize   int64           `db:"file_size" json:"file_size"`
		Private    types.BitBool   `db:"private" json:"private"`
	}
	// Upload represents a file to be uploaded to the
	// Verbis storage system.
	Upload struct {
		UUID       uuid.UUID
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
	// MediaSourceType defines the source type for a media
	// attachment within the files table.
	MediaSourceType = "media"
	// FormAttachmentSourceType defines the source type for
	// form attachment within the files table.
	FormAttachmentSourceType = "form_attachment"
)

// ID retrieves the full path for the upload. If the
// file is local, the prefix will be added.
func (f *File) ID(prefix string) string {
	if f.Provider == StorageLocal {
		return filepath.Join(prefix, f.BucketId)
	}
	return f.BucketId
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

// Validate checks to see if the upload is valid before
// uploading.
func (u *Upload) Validate() error {
	if u.Path == "" {
		return errors.New("no path attached to upload")
	}
	if u.Size < 1 {
		return errors.New("no size attached to upload")
	}
	if u.Contents == nil {
		return errors.New("upload contents is nil")
	}
	if u.SourceType == "" {
		return errors.New("no source type attached to upload")
	}
	if u.UUID.ID() == 0 {
		return errors.New("no uuid attached to upload")
	}
	return nil
}

// AbsPath returns the absolute path of the upload
// ready for uploading.
func (u *Upload) AbsPath() string {
	return strings.TrimPrefix(filepath.Join(filepath.Dir(u.Path), u.UUID.String()+filepath.Ext(u.Path)), ".")
}

// Mime returns the Mime type of the upload.
// Returns errors.INTERNAL if there was an error seeking
// bytes or the mime could not be obtained.
func (u *Upload) Mime() (Mime, error) {
	const op = "Domain.Upload.Mime"

	_, err := u.Contents.Seek(0, 0)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Error seeking file", Operation: op, Err: err}
	}

	mime, err := mimetype.DetectReader(u.Contents)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Error obtaining mime type", Operation: op, Err: err}
	}

	return Mime(mime.String()), nil
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
