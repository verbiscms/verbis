// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resizer

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/services/media/image"
	stdimage "image"
)

// Resizer describes the method for resizing images for
// the library.
type Resizer interface {
	Resize(imager image.Imager, compression int, media domain.MediaSize) (*bytes.Reader, error)
}

// Resize defines the data needed for resizing images.
type Resize struct{}

var (
	// ErrNilImager is returned by Resize when a nil Imager
	// has been passed to the function.
	ErrNilImager = fmt.Errorf("nil imager passed to resize")
)

// Resize satisfies the Resizer by decoding, cropping and
// resizing and finally saving the resized image.
func (r *Resize) Resize(imager image.Imager, compression int, media domain.MediaSize) (*bytes.Reader, error) {
	const op = "Resizer.Resize"

	if imager == nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error resizing, nil Image", Operation: op, Err: ErrNilImager}
	}

	img, err := imager.Decode()
	if err != nil {
		return nil, err
	}

	var resized *stdimage.NRGBA
	if media.Crop {
		resized = imaging.Fill(img, media.Width, media.Height, imaging.Center, imaging.Lanczos)
	} else {
		resized = imaging.Resize(img, media.Width, media.Height, imaging.Lanczos)
	}

	enc, err := imager.Encode(resized, compression)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(enc.Bytes()), nil
}
