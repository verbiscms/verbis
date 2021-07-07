// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uploader

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/services/media/internal/resizer"
	"github.com/ainsleyclark/verbis/api/services/webp"
	"github.com/ainsleyclark/verbis/api/storage"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// uploader defines the helper for uploading media
// items.
type Uploader struct {
	Config
	open multipart.File
	// .png
	extension string
	// png change
	mime domain.Mime
	// gopher
	bare    string
	resizer resizer.Resizer
}

type Config struct {
	File        *multipart.FileHeader
	Options     *domain.Options
	Config      *domain.ThemeConfig
	Exists      func(fileName string) bool
	WebP        webp.Execer
	StoragePath string
	Storage     storage.Client
}

func New(cfg Config) (*Uploader, error) {

	file, err := cfg.File.Open()
	if err != nil {
		return nil, err
	}

	mimeType, err := mimetype.DetectReader(file)
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return &Uploader{
		Config:    cfg,
		open:      file,
		extension: filepath.Ext(cfg.File.Filename),
		mime:      domain.Mime(mimeType.String()),
		resizer: &resizer.Resize{
			Storage:     cfg.Storage,
			Compression: cfg.Options.MediaCompression,
		},
		bare: files.RemoveFileExtension(cfg.File.Filename),
	}, nil
}

func (u *Uploader) Close() error {
	return u.open.Close()
}

// Save obtains the directory of where the file should be
// saved cleans the file name and saves the original and
// resized (images) media item. ToWebp is
// called concurrently to convert
// images to .webp.
func (u *Uploader) Save() (domain.Media, error) {
	//var (
	//	// E.G: uploads/2021/1
	//	dir = u.dir()
	//	// E.G: image.png
	//	name = u.cleanFileName()
	//)
	//
	//// Save the original file as is.
	//key, item, err := u.saveOriginal(dir)
	//if err != nil {
	//	return domain.Media{}, err
	//}

	//logger.Debug("Saved file the name: " + name + u.extension)

	//sizes, err := u.resize(name, dir)
	//if err != nil {
	//	return domain.Media{}, err
	//}

	//p := path.Dir(item.URI.Path)

	// https:/s3-eu-west-2.amazonaws.com/reddicotest/uploads/2021/07
	// /Users/ainsley/Desktop/Reddico/apis/verbis/storage/uploads/2021/07

	m := domain.Media{
		Id:          0,
		Title:       "",
		Alt:         "",
		Description: "",
		Sizes:       nil,
		UserId:      0,
		FileId:      0,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		File:        domain.File{},
	}
	//// Convert images to WebP.
	//go u.toWebP(m)

	return m, nil
}

// Dir returns the directory of where the media file should
// be uploaded. If the options allow for organising media
// by date, a date path will be created if it does
// not exist, for example '2020/01', otherwise
// it returns an empty string.
func (u *Uploader) dir() string {
	const prefix = paths.Uploads

	if !u.Options.MediaOrganiseDate {
		return prefix
	}

	t := time.Now()
	datePath := filepath.Join(t.Format("2006"), t.Format("01")) // 2020/01

	return filepath.Join(prefix, datePath)
}

// cleanFileName returns a cleaned version of the filename
// by removing any unnecessary characters. If the filename
// already exists, a version number will be added.
func (u *Uploader) cleanFileName() string {
	cleanedFile := strings.ReplaceAll(u.bare, " ", "-")
	reg := regexp.MustCompile("[^A-Za-z0-9 -]+")
	cleanedFile = strings.ToLower(reg.ReplaceAllString(cleanedFile, ""))

	// Check if the file exists and add a version number, continue if not.
	version := 0
	for {
		if version == 0 {
			exists := u.Exists(cleanedFile + u.extension)
			if !exists {
				break
			}
		} else {
			exists := u.Exists(cleanedFile + "-" + strconv.Itoa(version) + u.extension)
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
func (u *Uploader) saveOriginal(path string) (uuid.UUID, domain.File, error) {
	//key := uuid.New()
	//path = filepath.Join(path, key.String()+u.extension)

	//upload, err := u.Storage.Upload(path, u.File.Size, u.open)
	//if err != nil {
	//	return uuid.UUID{}, domain.File{}, err
	//}

	return [16]byte{}, domain.File{}, nil

	//return key, upload, nil
}

// Resize ranges over the media sizes stored in the
// options and decodes, resizes and saves the
// media size.
// Returns nil, (with no error) if the media item can not be resized.
func (u *Uploader) resize(name, path string) (domain.MediaSizes, error) {
	//if !u.mime.CanResize() {
	//	return nil, nil
	//}
	//
	//savedSizes := make(domain.MediaSizes)
	//for key, size := range u.Options.MediaSizes {
	//	uniq := uuid.New()
	//
	//	// gopher-100x100.png
	//	urlName := name + "-" + strconv.Itoa(size.Width) + "x" + strconv.Itoa(size.Height) + u.extension
	//
	//	// /Users/admin/cms/storage/uploads/2021/1/{uuid}.png
	//	localPath := path + string(os.PathSeparator) + uniq.String() + u.extension
	//
	//	var (
	//		upload domain.File
	//		err    error
	//	)
	//
	//	// Resize and save if the file is a JPG.
	//	if u.mime.IsJPG() {
	//		j := image.JPG{File: u.open}
	//		upload, err = u.resizer.Resize(&j, localPath, size)
	//	}
	//
	//	// Resize and save if the file is a PNG.
	//	if u.mime.IsPNG() {
	//		p := image.PNG{File: u.open}
	//		upload, err = u.resizer.Resize(&p, localPath, size)
	//	}
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	fmt.Println(upload, "TODO")
	//
	//	logger.Debug("Saved resized image with the name: " + urlName)
	//}
	//
	//return savedSizes, nil
	return nil, nil
}

// toWebP Checks to see if the media is a PNG or JPG and
// then ranges over the possible files of the media item
// and converts the images to webp. If the file
// exists, and an error occurred, it will be
// logged.
func (u *Uploader) toWebP(media domain.Media) {
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
