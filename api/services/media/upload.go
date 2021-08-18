// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/verbiscms/verbis/api/common/files"
	"github.com/verbiscms/verbis/api/common/paths"
	vstrings "github.com/verbiscms/verbis/api/common/strings"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/services/media/image"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Upload satisfies the Library to upload a media item
// to the library.
// Returns errors.INTERNAL on any eventuality the file could not be opened.
// Returns errors.INVALID if the mimetype could not be found.
func (s *Service) Upload(file *multipart.FileHeader, userID int) (domain.Media, error) {
	const op = "Media.Upload"

	out, err := file.Open()
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INVALID, Message: "Error opening file", Operation: op, Err: err}
	}

	defer func() {
		err := out.Close()
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error closing file with the name: " + file.Filename, Operation: op, Err: err})
		}
	}()

	opts := s.options.Struct()

	_, err = out.Seek(0, 0)
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error seeking file", Operation: op, Err: err}
	}

	var (
		// E.g. .jpg or .png
		ext = filepath.Ext(file.Filename)
		// E.g. uploads/2020/01/gopher.png
		path = filepath.Join(s.dir(opts.MediaOrganiseDate), s.cleanFileName(file.Filename, ext)+ext)
	)

	upload, err := s.storage.Upload(domain.Upload{
		UUID:       uuid.New(),
		Path:       path,
		Size:       file.Size,
		Contents:   out,
		Private:    false,
		SourceType: domain.MediaSourceType,
	})

	if err != nil {
		return domain.Media{}, err
	}

	sizes, err := s.resize(upload, out, opts)
	if err != nil {
		return domain.Media{}, err
	}

	media, err := s.repo.Create(domain.Media{
		Sizes:  sizes,
		UserID: userID,
		FileID: upload.ID,
		File:   upload,
	})

	if err != nil {
		return domain.Media{}, err
	}

	if opts.MediaConvertWebP {
		go s.toWebP(media, opts.MediaCompression)
	}

	return media, nil
}

// Dir returns the directory of where the testMedia file should
// be uploaded. If the options allow for organising testMedia
// by date, a date path will be created if it does
// not exist, for example '2020/01', otherwise
// it returns an empty string.
func (s *Service) dir(organiseDate bool) string {
	prefix := filepath.Base(paths.Uploads)

	if !organiseDate {
		return prefix
	}

	t := time.Now()
	datePath := filepath.Join(t.Format("2006"), t.Format("01")) // 2020/01

	return filepath.Join(prefix, datePath)
}

// cleanFileName returns a cleaned version of the filename
// by removing any unnecessary characters. If the filename
// already exists, a version number will be added.
func (s *Service) cleanFileName(name, ext string) string {
	var (
		bare          = files.RemoveFileExtension(name)
		removedDashes = strings.ReplaceAll(bare, " ", "-")
		reg           = regexp.MustCompile("[^A-Za-z0-9 -]+")
		cleanedFile   = strings.ToLower(reg.ReplaceAllString(removedDashes, ""))
	)

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

// Resize ranges over the testMedia sizes stored in the
// options and decodes, resizes and saves the
// testMedia size.
// Returns nil, (with no error) if the testMedia item can not be resized.
func (s *Service) resize(file domain.File, mp multipart.File, opts domain.Options) (domain.MediaSizes, error) {
	if !file.Mime.CanResize() {
		return nil, nil
	}

	var (
		ext        = file.Extension()
		savedSizes = make(domain.MediaSizes)
	)

	for key, size := range opts.MediaSizes {
		var (
			// E.g. gopher
			extRemoved = files.RemoveFileExtension(file.Name)
			// E.g. gopher-100x100.png
			urlName = extRemoved + "-" + strconv.Itoa(size.Width) + "x" + strconv.Itoa(size.Height) + ext
			// E.g. uploads/2020/01/gopher-100x100.png
			path = vstrings.TrimSlashes(filepath.Join(filepath.Dir(file.URL), urlName))
			// For resizing image
			buf *bytes.Reader
			// Error resizes
			err error
		)

		logger.Debug("Attempting to resize image: " + path)

		// Resize and save if the file is a JPG.
		if file.Mime.IsJPG() {
			j := image.JPG{File: mp}
			buf, err = s.resizer.Resize(&j, opts.MediaCompression, size)
		}

		// Resize and save if the file is a PNG.
		if file.Mime.IsPNG() {
			p := image.PNG{File: mp}
			buf, err = s.resizer.Resize(&p, opts.MediaCompression, size)
		}

		if err != nil {
			return nil, err
		}

		upload, err := s.storage.Upload(domain.Upload{
			UUID:       uuid.New(),
			Path:       path,
			Size:       int64(buf.Len()),
			Contents:   buf,
			Private:    false,
			SourceType: domain.MediaSourceType,
		})

		if err != nil {
			return nil, err
		}

		savedSizes[key] = domain.MediaSize{
			FileID:   upload.ID,
			SizeKey:  key,
			SizeName: size.SizeName,
			Width:    size.Width,
			Height:   size.Height,
			Crop:     size.Crop,
			File:     upload,
		}

		logger.Info("Saved resized image with the path: " + path)
	}

	if len(savedSizes) == 0 {
		savedSizes = nil
	}

	return savedSizes, nil
}

// toWebP Checks to see if the testMedia is a PNG or JPG and
// then ranges over the possible files of the testMedia item
// and converts the images to webp. If the file
// exists, and an error occurred, it will be
// logged.
func (s *Service) toWebP(media domain.Media, compression int) {
	if !media.File.Mime.CanResize() {
		return
	}

	s.fileToWebP(media.File, compression)

	for _, v := range media.Sizes {
		s.fileToWebP(v.File, compression)
	}
}

// fileToWebP converts a domain.File to a WebP image.
// Logs errors if the item failed to convert.
func (s *Service) fileToWebP(file domain.File, compression int) {
	path := vstrings.TrimSlashes(file.URL + domain.WebPExtension)

	logger.Debug("Attempting to convert image to WebP: " + path)

	b, file, err := s.storage.Find(file.URL)
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	read := bytes.NewReader(b)
	convert, err := s.webp.Convert(read, compression)
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	_, err = s.storage.Upload(domain.Upload{
		UUID:       uuid.New(),
		Path:       path,
		Size:       convert.Size(),
		Contents:   convert,
		Private:    false,
		SourceType: domain.MediaSourceType,
	})

	if err != nil {
		logger.WithError(err).Error()
		return
	}

	logger.Info("Successfully converted to WebP image with the path: " + path)
}
