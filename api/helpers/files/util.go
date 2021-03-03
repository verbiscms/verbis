// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Delete file based on file path
// Returns errors.INTERNAL if the file failed to delete.
func Delete(path string) error {
	const op = "files.Delete"
	err := os.Remove(path)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the file with the path: %v", path), Operation: op, Err: err}
	}
	return nil
}

// Exists checks if a file exists using os.Stat
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirectoryExists checks if directory exists using os.Stat
func DirectoryExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// Check the file exists and delete
// Returns errors.NOTFOUND if the file was not found.
func CheckAndDelete(path string) {
	const op = "files.Delete"
	if Exists(path) {
		if err := Delete(path); err != nil {
			logger.WithError(err).Error()
			return
		}
	}
	logger.WithError(&errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Failed to delete file with the path: %s", path), Operation: op, Err: fmt.Errorf("filepath %v not found", path)}).Error()
}

// Save File
// Returns errors.INTERNAL if the file could not be opened or be created.
func Save(file *multipart.FileHeader, dst string) error {
	const op = "files.Save"

	src, err := file.Open()
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to open file with the name: %s", file.Filename), Operation: op, Err: err}
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to create a file with the name: %s", file.Filename), Operation: op, Err: err}
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// Remove the file extension from a file
func RemoveFileExtension(file string) string {
	return file[0 : len(file)-len(GetFileExtension(file))]
}

// Get the file extension from a file
func GetFileExtension(file string) string {
	return filepath.Ext(file)
}

// Get the filesize of a file by path
func GetFileSize(path string) int {
	fi, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return int(fi.Size() / 1024)
}

// GetFileContents of given path
// Returns errors.INTERNAL if the path was invalid
func GetFileContents(path string) (string, error) {
	const op = "files.GetFileContents"
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the file contents with the path: %s", path), Operation: op, Err: err}
	}
	return string(contents), nil
}
