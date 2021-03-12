// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/gabriel-vasile/mimetype"
	"image"
	"mime/multipart"
)

var (
	ErrMimeType = errors.New("mimetype is not permitted")
)

// Validate
//
//
func (c *Library) Validate(file *multipart.FileHeader) error {
	const op = "Client.Validate"

	io, teardown, err := c.openFile(file)
	defer teardown()
	if err != nil {
		return err
	}

	mimeType, err := mimetype.DetectReader(io)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Mime type not permitted", Operation: op, Err: ErrMimeType}
	}

	valid := mime.IsValidMime(c.Config.Media.AllowedFileTypes, mimeType.String())
	if !valid {
		return &errors.Error{Code: errors.INVALID, Message: "Mime type not permitted", Operation: op, Err: ErrMimeType}
	}

	fileSize := int(file.Size / 1024) //nolint
	if fileSize > c.Options.MediaUploadMaxSize && c.Options.MediaUploadMaxSize != 0 {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("The file exceeds the maximum size restriction of %vkb.", c.Options.MediaUploadMaxSize), Operation: op, Err: err}
	}

	return c.validateImage(io)
}

// validateImage
//
//
func (c *Library) validateImage(file multipart.File) error {
	const op = "Client.ValidateImage"

	img, _, err := image.Decode(file)
	if err != nil {
		return nil // Is not an image
	}

	bounds := img.Bounds().Max
	if bounds.X > c.Options.MediaUploadMaxWidth && c.Options.MediaUploadMaxWidth != 0 {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("The image exceeds the upload max width of %vpx.", c.Options.MediaUploadMaxWidth), Operation: op, Err: err}
	}

	if img.Bounds().Max.Y > c.Options.MediaUploadMaxHeight && c.Options.MediaUploadMaxHeight != 0 {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("The image exceeds the upload max height of %vpx.", c.Options.MediaUploadMaxHeight), Operation: op, Err: err}
	}

	return nil
}
