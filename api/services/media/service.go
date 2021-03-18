// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"mime/multipart"
)

// Library defines methods for media items to
// save, validate and delete from the
// local file system.
type Library interface {
	Upload(file *multipart.FileHeader) (domain.Media, error)
	Validate(file *multipart.FileHeader) error
	Delete(item domain.Media)
}

// Service
//
// Defines the service for uploading, validating, deleting
// and serving rich media from the Verbis media library.
type Service struct {
	Options *domain.Options
	Config  *domain.ThemeConfig
	paths   paths.Paths
	Exists  func(fileName string) bool
}

// ExistsFunc is used by the uploader to determine if a
// media item exists in the library.
type ExistsFunc func(fileName string) bool

// New
//
// Creates a new Service.
func (c Service) New(opts *domain.Options, fn ExistsFunc) Library {
	return &Service{
		Options: opts,
		Config:  config.Get(),
		paths:   paths.Get(),
		Exists:  fn,
	}
}

// Upload
//
// Satisfies the Library to upload a media item to the
// library. Media items will be opened and saved to
// the local file system. Images are resized and
// saved in correspondence to the options.
// This function expects that validate
// has been called before it is run.
//
// Returns errors.INTERNAL on any eventuality the file could not be opened.
// Returns errors.INVALID if the mimetype could not be found.
func (c *Service) Upload(file *multipart.FileHeader) (domain.Media, error) {
	return upload(file, c.paths.Uploads, c.Options, c.Config, c.Exists)
}

// Delete
//
// Satisfies the Library to remove possible media item
// combinations from the file system, if the file
// does not exist (user moved) it will be
// skipped.
//
// Logs errors.INTERNAL if the file could not be deleted.
func (c *Service) Delete(item domain.Media) {
	deleteItem(item, c.paths.Uploads)
}

// Validate
//
// Satisfies the Library to see if the media item passed
// is valid. It will check if the file is a valid
// mime type, if the file size is less than the
// size specified in the options and finally
// checks the image boundaries.
//
// Returns errors.INVALID any of the conditions fail.
func (c *Service) Validate(file *multipart.FileHeader) error {
	return validate(file, c.Options, c.Config)
}
