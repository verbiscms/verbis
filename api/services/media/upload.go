// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/services/media/internal/image"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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

	out, err := file.Open()
	if err != nil {
		return domain.Media{}, err
	}

	defer func() {
		err = out.Close()
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error closing file with the name: " + file.Filename, Operation: op, Err: err})
		}
	}()

	_, err = out.Seek(0, 0)
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error seeking file", Operation: op, Err: err}
	}

	var (
		ext  = filepath.Ext(file.Filename)
		path = filepath.Join(s.dir(), s.cleanFileName(file.Filename, ext)+ext)
	)

	upload, err := s.storage.Upload(domain.Upload{
		Path:       path,
		Size:       file.Size,
		Contents:   out,
		Private:    false,
		SourceType: domain.MediaSourceType,
	})

	if err != nil {
		return domain.Media{}, err
	}

	sizes, err := s.resize(upload, out)
	if err != nil {
		return domain.Media{}, err
	}

	media, err := s.repo.Create(domain.Media{
		Sizes:  sizes,
		UserId: 0,
		FileId: upload.Id,
		File:   upload,
	})

	if err != nil {
		return domain.Media{}, err
	}

	return media, nil
}

// Dir returns the directory of where the media file should
// be uploaded. If the options allow for organising media
// by date, a date path will be created if it does
// not exist, for example '2020/01', otherwise
// it returns an empty string.
func (s *Service) dir() string {
	const prefix = paths.Uploads

	if !s.options.MediaOrganiseDate {
		return prefix
	}

	t := time.Now()
	datePath := filepath.Join(t.Format("2006"), t.Format("01")) // 2020/01

	return filepath.Join(prefix, datePath)
}

// cleanFileName returns a cleaned version of the filename
// by removing any unnecessary characters. If the filename
// already exists, a version number will be added.
func (s *Service) cleanFileName(name string, ext string) string {
	var (
		bare = files.RemoveFileExtension(name)
	)

	cleanedFile := strings.ReplaceAll(bare, " ", "-")
	reg := regexp.MustCompile("[^A-Za-z0-9 -]+")
	cleanedFile = strings.ToLower(reg.ReplaceAllString(cleanedFile, ""))

	// Check if the file exists and add a version number, continue if not.
	version := 0
	for {
		if version == 0 {
			exists := s.storage.Exists(cleanedFile + ext)
			if !exists {
				break
			}
		} else {
			exists := s.storage.Exists(cleanedFile + "-" + strconv.Itoa(version) + ext)
			if !exists {
				cleanedFile = cleanedFile + "-" + strconv.Itoa(version)
				break
			}
		}
		version++
	}

	return cleanedFile
}

// Resize ranges over the media sizes stored in the
// options and decodes, resizes and saves the
// media size.
// Returns nil, (with no error) if the media item can not be resized.
func (s *Service) resize(file domain.File, mp multipart.File) (domain.MediaSizes, error) {
	if !file.Mime.CanResize() {
		return nil, nil
	}

	var (
		ext = file.Extension()
	)

	savedSizes := make(domain.MediaSizes)
	for key, size := range s.options.MediaSizes {
		// gopher-100x100.png
		urlName := file.Name + "-" + strconv.Itoa(size.Width) + "x" + strconv.Itoa(size.Height) + ext

		// /Users/admin/cms/storage/uploads/2021/1/{uuid}.png
		localPath := filepath.Join(file.Path, urlName)

		var (
			buf *bytes.Reader
			err error
		)

		// Resize and save if the file is a JPG.
		if file.Mime.IsJPG() {
			j := image.JPG{File: mp}
			buf, err = s.resizer.Resize(&j, size)
		}

		// Resize and save if the file is a PNG.
		if file.Mime.IsPNG() {
			p := image.PNG{File: mp}
			buf, err = s.resizer.Resize(&p, size)
		}

		if err != nil {
			return nil, err
		}

		upload, err := s.storage.Upload(domain.Upload{
			Path:       localPath,
			Size:       int64(buf.Len()),
			Contents:   buf,
			Private:    false,
			SourceType: domain.MediaSourceType,
		})

		if err != nil {
			return nil, err
		}

		savedSizes[key] = domain.MediaSize{
			FileId:   upload.Id,
			SizeKey:  key,
			SizeName: urlName,
			Width:    size.Width,
			Height:   size.Height,
			Crop:     size.Crop,
			File:     upload,
		}

		logger.Debug("Saved resized image with the name: " + urlName)
	}

	return savedSizes, nil
}

// toWebP Checks to see if the media is a PNG or JPG and
// then ranges over the possible files of the media item
// and converts the images to webp. If the file
// exists, and an error occurred, it will be
// logged.
func (s *Service) toWebP(media domain.Media) {
	//if !u.Options.MediaConvertWebP {
	//	return
	//}
	//
	//if !media.Mime.CanResize() {
	//	return
	//}
	//
	//comp := u.Options.MediaCompression
	//
	//logger.Debug("Attempting to convert original image to WebP: " + media.FileName)
	//u.WebP.Convert(media.PrivatePath(u.StoragePath), comp)
	//
	//for _, v := range media.Sizes {
	//	logger.Debug("Attempting to convert media size image to WebP: " + v.SizeName)
	//	path := filepath.Join(u.StoragePath, media.FilePath, v.UUID.String()+media.Extension())
	//	u.WebP.Convert(path, comp)
	//}
}
