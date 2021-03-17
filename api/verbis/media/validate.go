// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gabriel-vasile/mimetype"
	"image"
	"mime/multipart"
)

var (
	ErrMimeType   = errors.New("mimetype is not permitted")
	ErrFileTooBig = errors.New("file size to big to be uploaded")
)

//
type validator struct {
	Config  *domain.ThemeConfig
	Options *domain.Options
	Size    int64
	File    multipart.File
}

// Validate
//
//
func validate(h *multipart.FileHeader, opts *domain.Options, cfg *domain.ThemeConfig) error {
	const op = "client.Validate"

	file, err := h.Open()
	defer func() {
		err := file.Close()
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error closing file with the name: " + h.Filename, Operation: op, Err: err})
		}
	}()

	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error opening file with the name: " + h.Filename, Operation: op, Err: err}
	}

	v := validator{
		Config:  cfg,
		Options: opts,
		Size:    h.Size,
		File:    file,
	}

	err = v.Mime()
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "The file is not permitted to be uploaded", Operation: op, Err: err}
	}

	err = v.FileSize()
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "The file exceeds the maximum size restriction", Operation: op, Err: err}
	}

	err = v.Image()
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "The image exceeds the width/height restrictions", Operation: op, Err: err}
	}

	return nil
}

// Mime
//
//
func (v *validator) Mime() error {
	m, err := mimetype.DetectReader(v.File)
	if err != nil {
		return ErrMimeType
	}

	valid := mime.IsValidMime(v.Config.Media.AllowedFileTypes, m.String())
	if !valid {
		return ErrMimeType
	}

	return nil
}

// FileSize
//
//
func (v *validator) FileSize() error {
	fileSize := v.Size / 1024 //nolint
	if fileSize > v.Options.MediaUploadMaxSize && v.Options.MediaUploadMaxSize != 0 {
		return ErrFileTooBig
	}
	return nil
}

// Image
//
//
func (v *validator) Image() error {
	img, _, err := image.Decode(v.File)
	if err != nil {
		return nil // Is not an image
	}

	bounds := img.Bounds().Max
	if int64(bounds.X) > v.Options.MediaUploadMaxWidth && v.Options.MediaUploadMaxWidth != 0 {
		return errors.New("image exceeds the maximum upload width")
	}

	if int64(bounds.Y) > v.Options.MediaUploadMaxHeight && v.Options.MediaUploadMaxHeight != 0 {
		return errors.New("image exceeds the maximum upload height")
	}

	return nil
}
