// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resizer

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/services/media/internal/img"
	"github.com/ainsleyclark/verbis/api/storage"
	"github.com/disintegration/imaging"
	"image"
)

// Resizer describes the method for resizing images for
// the library.
type Resizer interface {
	Resize(imager img.Imager, dest string, media domain.MediaSize) (domain.StorageFile, error)
}

// Resize implements the Resizer interface.
type Resize struct {
	Storage     storage.Client
	Compression int
}

// Resize satisfies the Resizer by decoding, cropping and
// resizing and finally saving the resized image.
func (r *Resize) Resize(imager img.Imager, dest string, media domain.MediaSize) (domain.StorageFile, error) {
	i, err := imager.Decode()
	if err != nil {
		return err
	}

	var resized *image.NRGBA
	if media.Crop {
		resized = imaging.Fill(i, media.Width, media.Height, imaging.Center, imaging.Lanczos)
	} else {
		resized = imaging.Resize(i, media.Width, media.Height, imaging.Lanczos)
	}

	enc, err := imager.Encode(resized, r.Compression)
	if err != nil {
		return err
	}

	upload, err := r.Storage.Upload(dest, int64(enc.Len()), enc)
	if err != nil {
		return err
	}

	return upload, nil
}
