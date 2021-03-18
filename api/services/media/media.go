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

type client struct {
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
// Creates a new client.
func (c client) New(opts *domain.Options, fn ExistsFunc) Library {
	return &client{
		Options: opts,
		Config:  config.Get(),
		paths:   paths.Get(),
		Exists:  fn,
	}
}

// Upload
//
// Satisfies the Library to upload a media item to the
// library.
// TODO: Carry on!
func (c *client) Upload(file *multipart.FileHeader) (domain.Media, error) {
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
func (c *client) Delete(item domain.Media) {
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
func (c *client) Validate(file *multipart.FileHeader) error {
	return validate(file, c.Options, c.Config)
}
