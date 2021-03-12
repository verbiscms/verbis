// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"mime/multipart"
	"os"
)

type uploader struct {
	File multipart.File
	Path string
	Webp bool
	Mime mime
}

func (u *uploader) upload() {

	//if u.webp {
	//	go webp.Convert()
	//}
}

// save
//
//
func (u *uploader) save(dest string) error {
	const op = "Uploader.Save"

	out, err := os.Create(dest)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error creating file", Operation: op, Err: err}
	}
	defer out.Close()

	_, err = io.Copy(out, u.File)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error copying file", Operation: op, Err: err}
	}

	return err
}

// resizeImage
//
//
func (u *uploader) resizeImage(srcImage image.Image, width int, height int, crop bool) image.Image {
	if crop {
		return imaging.Fill(srcImage, width, height, imaging.Center, imaging.Lanczos)
	} else {
		return imaging.Resize(srcImage, width, height, imaging.Lanczos)
	}
}
