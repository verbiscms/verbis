// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/services/webp"
	"github.com/gabriel-vasile/mimetype"
	"mime/multipart"
	"path/filepath"
)

// Library defines methods for media items to
// save, validate and delete from the
// local file system.
type Library interface {
	Upload(file *multipart.FileHeader) (domain.Media, error)
	Serve(media domain.Media, path string, acceptWebP bool) ([]byte, domain.Mime, error)
	Validate(file *multipart.FileHeader) error
	Delete(item domain.Media)
}

// Service
//
// Defines the service for uploading, validating, deleting
// and serving rich media from the Verbis media library.
type Service struct {
	options *domain.Options
	config  *domain.ThemeConfig
	paths   paths.Paths
	exists  func(fileName string) bool
	webp    webp.Execer
}

// ExistsFunc is used by the uploader to determine if a
// media item exists in the library.
type ExistsFunc func(fileName string) bool

// New
//
// Creates a new Service.
func New(opts *domain.Options, fn ExistsFunc) *Service {
	p := paths.Get()
	return &Service{
		options: opts,
		config:  config.Get(),
		paths:   p,
		exists:  fn,
		webp:    webp.New(p.Bin + webp.Path),
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
func (s *Service) Upload(file *multipart.FileHeader) (domain.Media, error) {
	const op = "Media.Uploader.Upload"

	h, err := file.Open()
	defer func() {
		err := h.Close()
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error closing file with the name: " + file.Filename, Operation: op, Err: err})
		}
	}()
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INVALID, Message: "Error opening file with the name: " + file.Filename, Operation: op, Err: err}
	}

	m, err := mimetype.DetectReader(h)
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INVALID, Message: "Mime type not found", Operation: op, Err: err}
	}

	_, err = h.Seek(0, 0)
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error seeking file", Operation: op, Err: err}
	}

	u := uploader{
		File:       h,
		Options:    s.options,
		Config:     s.config,
		Exists:     s.exists,
		UploadPath: s.paths.Uploads,
		FileName:   file.Filename,
		Extension:  filepath.Ext(file.Filename),
		Size:       file.Size,
		Mime:       domain.Mime(m.String()),
		Resizer:    &Resize{},
		WebP:       s.webp,
	}

	return u.Save()
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
func (s *Service) Validate(file *multipart.FileHeader) error {
	return validate(file, s.options, s.config)
}

// Delete
//
// Satisfies the Library to remove possible media item
// combinations from the file system, if the file
// does not exist (user moved) it will be
// skipped.
//
// Logs errors.INTERNAL if the file could not be deleted.
func (s *Service) Delete(item domain.Media) {
	deleteItem(item, s.paths.Uploads)
}
