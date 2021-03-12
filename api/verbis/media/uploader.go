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
	"github.com/disintegration/imaging"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"image"
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
	Mime       mimeType
}

//
//
//
func upload(h *multipart.FileHeader, path string, opts *domain.Options, cfg *domain.ThemeConfig, exists ExistsFunc) (domain.Media, error) {
	const op = "MediaClient.Dir"

	file, err := h.Open()
	defer func() {
		err := file.Close()
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error closing file with the name: " + h.Filename, Operation: op, Err: err})
		}
	}()

	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error opening file with the name: " + h.Filename, Operation: op, Err: err}
	}

	m, err := mimetype.DetectReader(file)
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "Error opening file with the name: " + h.Filename, Operation: op, Err: err}
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
		Mime:       mimeType(m.String()),
	}

	return u.Save()
}

//
//
//
func (u *uploader) Save() (domain.Media, error) {

	// Obtain the file path.
	filePath, err := u.Dir()
	if err != nil {
		return domain.Media{}, &errors.Error{Code: "", Message: "", Operation: "", Err: err}
	}

	key := uuid.New()
	path := u.UploadPath + filePath + key.String() + u.Extension

	// Save the original file as is.
	err = u.SaveOriginal(path)
	if err != nil {
		return domain.Media{}, &errors.Error{Code: "", Message: "", Operation: "", Err: err}
	}

	// Convert to WebP if the options allow.
	// TODO CHECK MIME!
	if u.Options.MediaConvertWebP {
		go webp.Convert(path, u.Options.MediaCompression)
	}

	sizes := u.SaveResized()

	// E.G: image.png
	cleanName := u.CleanFileName()

	return domain.Media{
		UUID:     key,
		Url:      u.Config.Media.UploadPath + "/" + filePath + "/" + cleanName,
		FilePath: filePath,
		FileSize: u.Size,
		FileName: cleanName + u.Extension,
		Sizes:    sizes,
		Type:     "MIMEMEEE",
	}, nil
}

// Dir
//
//
func (u *uploader) Dir() (string, error) {
	const op = "MediaClient.Dir"

	if !u.Options.MediaOrganiseDate {
		return "", nil
	}

	t := time.Now()

	// 2020/01
	datePath := t.Format("2006") + string(os.PathSeparator) + t.Format("01")
	path := u.UploadPath + string(os.PathSeparator) + datePath

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return string(os.PathSeparator) + datePath, nil
}

// fileName
//
//
func (u *uploader) CleanFileName() string {
	name := files.RemoveFileExtension(u.FileName)

	cleanedFile := strings.ReplaceAll(name, " ", "-")
	reg := regexp.MustCompile("[^A-Za-z0-9 -]+")
	cleanedFile = strings.ToLower(reg.ReplaceAllString(cleanedFile, ""))

	// Check if the file exists and add a version number, continue if not
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

// save
//
//
func (u *uploader) SaveOriginal(dest string) error {
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

// SaveResize
func (u *uploader) SaveResized() domain.MediaSizes {
	if !u.Mime.CanResize() {
		return nil
	}
	//
	//for key, size := range u.Options.MediaSizes {
	//	//uniq := uuid.New()
	//	//fileName :=  name + "-" + strconv.Itoa(size.Width) + "x" + strconv.Itoa(size.Height) + extension
	//
	//	if u.Mime.IsPNG() {
	//		p := PNG{
	//			File: u.File,
	//		}
	//		img, err := p.Decode()
	//		if err != nil {
	//			// handle
	//		}
	//	}
	//
	//}

	return nil
}

// ResizeImage
//
//
func (u *uploader) ResizeImage(srcImage image.Image, width int, height int, crop bool) image.Image {
	if crop {
		return imaging.Fill(srcImage, width, height, imaging.Center, imaging.Lanczos)
	} else {
		return imaging.Resize(srcImage, width, height, imaging.Lanczos)
	}
}
