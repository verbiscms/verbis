// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/services/media/image"
	"github.com/ainsleyclark/verbis/api/services/webp"
	"github.com/google/uuid"
	"github.com/gookit/color"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// uploader defines the helper for uploading media
// items.
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
	Resizer    image.Resizer
	WebP       webp.Execer
}

// Save obtains the directory of where the file should be
// saved cleans the file name and saves the original and
// resized (images) media item. ToWebp is
// called concurrently to convert
// images to .webp.
func (u *uploader) Save() (domain.Media, error) {
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

	logger.Debug("Saved file the name: " + name + u.Extension)

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
		Mime:     u.Mime,
	}

	// Convert images to WebP.
	go u.ToWebP(m)

	return m, nil
}

// Dir returns the directory of where the media file should
// be uploaded. If the options allow for organising media
// by date, a date path will be created if it does
// not exist, for example '2020/01', otherwise
// it returns an empty string.
func (u *uploader) Dir() (string, error) {
	const op = "Media.Uploader.Dir"

	if !u.Options.MediaOrganiseDate {
		return "", nil
	}

	t := time.Now()
	datePath := filepath.Join(t.Format("2006"), t.Format("01")) // 2020/01
	path := filepath.Join(u.UploadPath, datePath)

	_, err := os.Stat(path)
	color.Green.Println(err, path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		fmt.Println(path)
		if err != nil {
			return "", &errors.Error{Code: errors.INVALID, Message: "Error creating media uploads folder with the path: " + path, Operation: op, Err: err}
		}
	}

	return datePath, nil
}

// CleanFileName returns a cleaned version of the filename
// by removing any unnecessary characters. If the filename
// already exists, a version number will be added.
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

// SaveOriginal saves the original file to disk and returns
// a new UUID when it has been saved successfully.
// Returns errors.INTERNAL if the file could not be copied
// or created.
func (u *uploader) SaveOriginal(path string) (uuid.UUID, error) {
	const op = "Media.Uploader.Save"

	key := uuid.New()
	dest := path + string(os.PathSeparator) + key.String() + u.Extension

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

// Resize ranges over the media sizes stored in the
// options and decodes, resizes and saves the
// media size.
// Returns nil, (with no error) if the media item can not be resized.
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
			j := image.JPG{File: u.File}
			err := u.Resizer.Resize(&j, localPath, size, comp)
			if err != nil {
				return nil, err
			}
		}

		// Resize and save if the file is a PNG.
		if u.Mime.IsPNG() {
			p := image.PNG{File: u.File}
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

// URL gets the public url of the file according to date
// and month if the organise year variable in the config
// is set to true. If false the function will
// return the public uploads folder
// by default.
func (u *uploader) URL() string {
	path := "/" + strings.ReplaceAll(u.Config.Media.UploadPath, "/", "")
	if !u.Options.MediaOrganiseDate {
		return path
	} else {
		t := time.Now()
		return path + "/" + t.Format("2006") + "/" + t.Format("01")
	}
}

// FileSize obtains the filesize of the path given, if the
// path does not exist, 0 will be returned.
func (u *uploader) FileSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fi.Size()
}

// ToWebP Checks to see if the media is a PNG or JPG and
// then ranges over the possible files of the media item
// and converts the images to webp. If the file
// exists, and an error occurred, it will be
// logged.
func (u *uploader) ToWebP(media domain.Media) {
	if !media.Mime.CanResize() {
		return
	}

	comp := u.Options.MediaCompression

	logger.Debug("Attempting to convert original image to WebP: " + media.FileName)
	u.WebP.Convert(u.UploadPath+string(os.PathSeparator)+media.UploadPath(), comp)

	for _, v := range media.Sizes {
		logger.Debug("Attempting to convert media size image to WebP: " + v.Name)
		path := u.UploadPath + string(os.PathSeparator) + media.FilePath + string(os.PathSeparator) + v.UUID.String() + media.Extension()
		u.WebP.Convert(path, comp)
	}
}
