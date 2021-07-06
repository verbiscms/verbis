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
	"github.com/ainsleyclark/verbis/api/services/media/internal/uploader"
	"github.com/ainsleyclark/verbis/api/services/webp"
	"github.com/ainsleyclark/verbis/api/storage"
	"mime/multipart"
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
	storage storage.Client
}

// ExistsFunc is used by the uploaderold to determine if a
// media item exists in the library.
type ExistsFunc func(fileName string) bool

// New
//
// Creates a new Service.
func New(opts *domain.Options, store storage.Client, fn ExistsFunc) *Service {
	p := paths.Get()
	return &Service{
		options: opts,
		config:  config.Get(),
		paths:   p,
		exists:  fn,
		webp:    webp.New(p.Bin + webp.Path),
		storage: store,
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
	const op = "Media.Upload"

	u, err := uploader.New(uploader.Config{
		File:        file,
		Options:     s.options,
		Config:      s.config,
		Exists:      s.exists,
		WebP:        s.webp,
		StoragePath: s.paths.Storage,
		Storage:     s.storage,
	})

	defer func() {
		err = u.Close()
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error closing uploaderold", Operation: op, Err: err})
		}
	}()

	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating uploaderold for file", Operation: op, Err: err}
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
	const op = "Service.Delete"

	items := item.PossibleFiles(s.paths.Uploads)
	for _, path := range items {
		err := s.storage.Delete(path)

		// Exists func
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting file with the path: " + path, Operation: op, Err: err})
		}

		logger.Debug("Deleted file with the path: " + path)
	}
}
