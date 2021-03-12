// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Client interface {
	Upload(file *multipart.FileHeader) error
	Validate(file *multipart.FileHeader) error
	Delete(path string, sizes domain.MediaSizes) error
}

type Library struct {
	Options  *domain.Options
	Config   *domain.ThemeConfig
	paths    paths.Paths
	datePath string
	ext      string
	Exists   func(fileName string) bool
}

func (c *Library) Upload(file *multipart.FileHeader) error {
	const op = "Client.Upload"

	// E.G: Image20@.png
	name := file.Filename

	// E.G: .png
	//extension := files.GetFileExtension(name)

	src, err := file.Open()
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error opening file with the name: " + name, Operation: op, Err: err}
	}
	defer src.Close()

	return nil
}

func (c *Library) openFile(file *multipart.FileHeader) (multipart.File, func() error, error) {
	const op = "MediaClient.openFile"

	src, err := file.Open()
	if err != nil {
		return nil, nil, &errors.Error{Code: errors.INTERNAL, Message: "Error opening file with the name: " + file.Filename, Operation: op, Err: err}
	}
	return src, src.Close, nil
}

// Dir
//
//
func (c *Library) Dir() string {
	const op = "MediaClient.Dir"

	if !c.Options.MediaOrganiseDate {
		return c.paths.Uploads
	}

	t := time.Now()

	// 2020/01
	datePath := t.Format("2006") + string(os.PathSeparator) + t.Format("01")
	path := c.paths.Uploads + "/" + datePath

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			// Handle
		}
	}

	c.datePath = datePath
	return path
}

// fileName
//
//
func (c *Library) fileName(file, extension string) {
	name := files.RemoveFileExtension(file)

	cleanedFile := strings.ReplaceAll(name, " ", "-")
	reg := regexp.MustCompile("[^A-Za-z0-9 -]+")
	cleanedFile = strings.ToLower(reg.ReplaceAllString(cleanedFile, ""))

	// Check if the file exists and add a version number, continue if not
	version := 0
	for {
		if version == 0 {
			exists := c.Exists(cleanedFile + extension)
			if !exists {
				break
			}
		} else {
			exists := c.Exists(cleanedFile + "-" + strconv.Itoa(version) + extension)
			if !exists {
				cleanedFile = cleanedFile + "-" + strconv.Itoa(version)
				break
			}
		}
		version++
	}
}
