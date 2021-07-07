// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/services/media/internal/resizer"
	"github.com/ainsleyclark/verbis/api/services/webp"
	"github.com/ainsleyclark/verbis/api/storage"
	"github.com/ainsleyclark/verbis/api/store/media"
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
	repo    media.Repository
	resizer resizer.Resizer
}

// ExistsFunc is used by the uploaderold to determine if a
// media item exists in the library.
type ExistsFunc func(fileName string) bool

// New
//
// Creates a new Service.
func New(opts *domain.Options, storage storage.Client, repo media.Repository) *Service {
	p := paths.Get()
	return &Service{
		options: opts,
		config:  config.Get(),
		paths:   p,
		webp:    webp.New(p.Bin + webp.Path),
		storage: storage,
		repo:    repo,
		resizer: &resizer.Resize{
			Compression: opts.MediaCompression,
		},
	}
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
