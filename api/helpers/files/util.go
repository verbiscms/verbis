package files

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
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
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Check the file exists and delete
// Returns errors.NOTFOUND if the file was not found.
func CheckAndDelete(path string) error {
	const op = "files.Delete"
	if Exists(path) {
		if err := Delete(path); err != nil {
			return err
		}
		return nil
	}
	return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Failed to delete file with the path: %s", path), Operation: op, Err: fmt.Errorf("filepath %v not found", path)}
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

// LoadFile load's the a file based on the path and returns a []byte ready for conversion
// Returns errors.INTERNAL if the configuration file failed to load.
func LoadFile(path string) ([]byte, error) {
	const op = "files.LoadFileB"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Could not load the file with the path: %s", path), Operation: op, Err: err}
	}
	return data, nil
}
