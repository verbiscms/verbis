// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/common/params"
	"github.com/ainsleyclark/verbis/api/common/paths"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/services/media/resizer"
	"github.com/ainsleyclark/verbis/api/services/webp"
	"github.com/ainsleyclark/verbis/api/storage"
	"github.com/ainsleyclark/verbis/api/store/media"
	"mime/multipart"
)

// Library defines methods for testMedia items to
// save, validate and delete from the
// local file system.
type Library interface {
	// List returns a slice of media items with the total amount.
	// Returns errors.INTERNAL if the SQL query was invalid.
	// Returns errors.NOTFOUND if there are no media items available.
	List(meta params.Params) (domain.MediaItems, int, error)
	// Find returns a media item by searching with the given ID.
	// Returns errors.INTERNAL if there was an error executing the query.
	// Returns errors.NOTFOUND if the media item was not found by the given ID.
	Find(id int) (domain.Media, error)
	// Update returns an updated media item by updating title, alt,
	// description and updated_at fields.
	// Returns errors.CONFLICT if the validation failed.
	// Returns errors.INTERNAL if the SQL query was invalid or the function
	// could not obtain the newly created ID.
	Update(m domain.Media) (domain.Media, error)
	// Upload uploads a testMedia item to the library. Media items
	// will be opened and saved to the local file system or
	// bucket dependant on storage. Images are resized and
	// saved in correspondence to the options. This
	// function expects that validate has been
	// called before it is run.
	// Returns errors.INTERNAL on any eventuality the file could not be opened.
	// Returns errors.INVALID if the mimetype could not be found.
	Upload(file *multipart.FileHeader, userID int) (domain.Media, error)
	// Validate accepts a multipart.FileHeader to see if the
	// testMedia item is valid before uploading. It will check
	// if the file is a valid mime type, if the file size
	// is less than the size specified in the options
	// and finally checks the image boundaries.
	// Returns errors.INVALID any of the conditions fail.
	Validate(file *multipart.FileHeader) error
	// Delete removes the testMedia item from the database and
	// storage system. Generated sizes and WebP images
	// will also be removed.
	// Returns errors.NOTFOUND if the file does not exist.
	// Returns errors.INTERNAL if the file could not be deleted from the database.
	// Logs errors.INTERNAL if the file could not be deleted from the storage bucket.
	Delete(id int) error
	// ReGenerateWebP generate's WebP deletes any WebP images
	// associated with media items and their sizes. It
	// returns a total amount of media items being
	// processed in the background.
	// Returns an errors of the media items could not be listed.
	ReGenerateWebP() (int, error)
}

var (
	// ErrMimeType is returned by validate when a mimetype is
	// not permitted.
	ErrMimeType = errors.New("mimetype is not permitted")
	// ErrFileTooBig is returned by validate when a file is to
	// big to be uploaded.
	ErrFileTooBig = errors.New("file size to big to be uploaded")
)

// Service Defines the service for uploading, validating, deleting
// and serving rich testMedia from the Verbis testMedia library.
type Service struct {
	options *domain.Options
	config  *domain.ThemeConfig
	paths   paths.Paths
	webp    webp.Execer
	storage storage.Bucket
	repo    media.Repository
	resizer resizer.Resizer
}

// New creates a new testMedia Service.
func New(opts *domain.Options, storage storage.Bucket, repo media.Repository) *Service {
	p := paths.Get()
	return &Service{
		options: opts,
		config:  config.Get(),
		paths:   p,
		webp:    webp.New(p.Bin + webp.Path),
		storage: storage,
		repo:    repo,
		resizer: &resizer.Resize{},
	}
}

// List satisfies the Library to list a collection of media
// items.
func (s *Service) List(meta params.Params) (domain.MediaItems, int, error) {
	return s.repo.List(meta)
}

// Find satisfies the Library to find a media item by
// searching by ID.
func (s *Service) Find(id int) (domain.Media, error) {
	return s.repo.Find(id)
}

// Update satisfies the Library to update a media item.
func (s *Service) Update(m domain.Media) (domain.Media, error) {
	return s.repo.Update(m)
}
