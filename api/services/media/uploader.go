// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/webp"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type uploader struct {
	File       multipart.File
	Options    *domain.Options
	Config     *domain.ThemeConfig
	Exists     ExistsFunc
	UploadPath string
	FileName   string
	Extension  string
	Size       int64
	Mime       domain.Mime
	Resizer    Resizer
}

// upload
//
//
func upload(h *multipart.FileHeader, path string, opts *domain.Options, cfg *domain.ThemeConfig, exists ExistsFunc) (domain.Media, error) {
	const op = "Media.Uploader.Upload"

	file, err := h.Open()
	defer func() {
		err := file.Close()
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error closing file with the name: " + h.Filename, Operation: op, Err: err})
		}
	}()
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INVALID, Message: "Error opening file with the name: " + h.Filename, Operation: op, Err: err}
	}

	m, err := mimetype.DetectReader(file)
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INVALID, Message: "Mime type not found", Operation: op, Err: err}
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error seeking file", Operation: op, Err: err}
	}

	u := uploader{
		File:       file,
		Options:    opts,
		Config:     cfg,
		Exists:     exists,
		UploadPath: path,
		FileName:   h.Filename,
		Extension:  filepath.Ext(h.Filename),
		Size:       h.Size,
		Mime:       domain.Mime(m.String()),
		Resizer:    &Resize{},
	}

	return u.Save()
}

// Save
//
// Obtains the directory of where the file should be saved
// cleans the file name and saves the original and
// resized (images) media item. ToWebp is
// called concurrently to convert images
// to .webp.
func (u *uploader) Save() (domain.Media, error) {
	const op = "Media.Uploader.Save"

	// Obtain the file path.
	filePath, err := u.Dir()
	if err != nil {
		return domain.Media{}, err
	}

	// E.G: image.png
	name := u.CleanFileName()

	// E.G /Users/admin/cms/storage/uploads/2021/1
	path := u.UploadPath + string(os.PathSeparator) + filePath

	// Save the original file as is.
	key, err := u.SaveOriginal(path)
	if err != nil {
		return domain.Media{}, err
	}

	logger.Debug("Saved resized image with the name: " + name)

	// Resize and save the image sizes
	sizes, err := u.Resize(name, path)
	if err != nil {
		logger.WithError(err).Error()
	}

	m := domain.Media{
		UUID:     key,
		Url:      u.URL() + "/" + name + u.Extension,
		FilePath: filePath,
		FileSize: u.Size,
		FileName: name + u.Extension,
		Sizes:    sizes,
		Type:     u.Mime,
	}

	// Convert images to WebP.
	go u.ToWeb(m)

	return m, nil
}

// Dir
//
// Returns the directory of where the media file should be
// uploaded. If the options allow for organising media
// by date, a date path will be created if it does
// not exist, for example '2020/01', otherwise
// it returns an empty string.
func (u *uploader) Dir() (string, error) {
	const op = "Media.Uploader.Dir"

	if !u.Options.MediaOrganiseDate {
		return "", nil
	}

	t := time.Now()
	datePath := t.Format("2006") + string(os.PathSeparator) + t.Format("01") // 2020/01
	path := u.UploadPath + string(os.PathSeparator) + datePath

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return "", &errors.Error{Code: errors.INVALID, Message: "Error creating media uploads folder with the path: " + path, Operation: op, Err: err}
		}
	}

	return string(os.PathSeparator) + datePath, nil
}

// CleanFileName
//
// Returns a cleaned version of the filename by removing
// any unnecessary characters. If the filename already
// exists, a version number will be added.
func (u *uploader) CleanFileName() string {
	name := files.RemoveFileExtension(u.FileName)

	cleanedFile := strings.ReplaceAll(name, " ", "-")
	reg := regexp.MustCompile("[^A-Za-z0-9 -]+")
	cleanedFile = strings.ToLower(reg.ReplaceAllString(cleanedFile, ""))

	// Check if the file exists and add a version number, continue if not.
	version := 0
	for {
		if version == 0 {
			exists := u.Exists(cleanedFile + u.Extension)
			if !exists {
				break
			}
		} else {
			exists := u.Exists(cleanedFile + "-" + strconv.Itoa(version) + u.Extension)
			if !exists {
				cleanedFile = cleanedFile + "-" + strconv.Itoa(version)
				break
			}
		}
		version++
	}

	return cleanedFile
}

// SaveOriginal
//
// Saves the original file to disk and returns a new UUID
// when it has been saved successfully.
// Returns errors.INTERNAL if the file could not be copied
// or created.
func (u *uploader) SaveOriginal(path string) (uuid.UUID, error) {
	const op = "Media.Uploader.Save"

	key := uuid.New()
	dest := path + key.String() + u.Extension

	out, err := os.Create(dest)
	if err != nil {
		return uuid.UUID{}, &errors.Error{Code: errors.INTERNAL, Message: "Error creating file", Operation: op, Err: err}
	}
	defer out.Close()

	_, err = io.Copy(out, u.File)
	if err != nil {
		return uuid.UUID{}, &errors.Error{Code: errors.INTERNAL, Message: "Error copying file", Operation: op, Err: err}
	}

	return key, err
}

// Resize
//
// Ranges over the media sizes stored in the options and
// decodes, resizes and saves the media size.
// Returns nil, (with no error) if the media item can not
// be resized.
func (u *uploader) Resize(name, path string) (domain.MediaSizes, error) {
	if !u.Mime.CanResize() {
		return nil, nil
	}

	comp := u.Options.MediaCompression

	savedSizes := make(domain.MediaSizes)
	for key, size := range u.Options.MediaSizes {
		uniq := uuid.New()

		// gopher-100x100.png
		urlName := name + "-" + strconv.Itoa(size.Width) + "x" + strconv.Itoa(size.Height) + u.Extension

		// /Users/admin/cms/storage/uploads/2021/1/{uuid}.png
		localPath := path + string(os.PathSeparator) + uniq.String() + u.Extension

		// Resize and save if the file is a JPG.
		if u.Mime.IsJPG() {
			j := JPG{File: u.File}
			err := u.Resizer.Resize(&j, localPath, size, comp)
			if err != nil {
				return nil, err
			}
		}

		// Resize and save if the file is a PNG.
		if u.Mime.IsPNG() {
			p := PNG{File: u.File}
			err := u.Resizer.Resize(&p, localPath, size, comp)
			if err != nil {
				return nil, err
			}
		}

		logger.Debug("Saved resized image with the name: " + urlName)

		savedSizes[key] = domain.MediaSize{
			UUID:     uniq,
			Url:      u.URL() + "/" + urlName,
			Name:     urlName,
			SizeName: size.SizeName,
			FileSize: u.FileSize(path + string(os.PathSeparator) + uniq.String() + u.Extension),
			Width:    size.Width,
			Height:   size.Height,
			Crop:     size.Crop,
		}
	}

	return savedSizes, nil
}

// URL
//
// Get the public url of the file according to date and
// month if the organise year variable in the config
// is set to true. If false the function will
// return the public uploads folder
// by default.
func (u *uploader) URL() string {
	if !u.Options.MediaOrganiseDate {
		return u.Config.Media.UploadPath
	} else {
		t := time.Now()
		return u.Config.Media.UploadPath + "/" + t.Format("2006") + "/" + t.Format("01")
	}
}

// FileSize
//
// Obtains the filesize of the path given, if the path
// does not exist, 0 will be returned.
func (u *uploader) FileSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fi.Size()
}

// ToWeb
//
// Checks to see if the media is a PNG or JPG and then
// ranges over the possible files of the media item
// and converts the images to webp. If the file
// exists, and an error occured, it will be
// logged.
func (u *uploader) ToWeb(media domain.Media) {
	if !media.Type.CanResize() {
		return
	}

	for _, v := range media.PossibleFiles() {
		path := u.UploadPath + string(os.PathSeparator) + v

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}

		logger.Debug("Attempting to convert image to webp with the path: " + path)
		webp.Convert(path, 100-u.Options.MediaCompression) //nolint
	}
}
