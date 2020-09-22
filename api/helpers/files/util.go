package files

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Delete file based on file path
func Delete(path string) error {
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("Could not delete the file with the path: %v", path)
	}
	return nil
}

// Check if a file exists
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Check the file exists and delete
func CheckAndDelete(path string) error {
	if Exists(path) {
		if err := Delete(path); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("Filepath %v not found", path)
}

// Save File
func Save(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// Remove the file extension from a file
func RemoveFileExtension(file string) string {
	return file[0:len(file)-len(GetFileExtension(file))]
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

// Get file contents of given path
func GetFileContents(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		return "", fmt.Errorf("Could not get file contents: %w", err)
	}

	return string(contents), nil
}

// Get files retrieves all files based on the file path param
func GetFile(path string) ([]byte, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}