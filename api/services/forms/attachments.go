// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"encoding/base64"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/gabriel-vasile/mimetype"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

const (
	// The total upload limit allowed to upload file
	// attachments.
	UploadLimit = 5
)

var (
	// AllowedMimes represents the mime types permitted
	// to be attached to forms.
	AllowedMimes = []string{
		"text/plain",
		"image/jpeg",
		"image/png",
		"image/svg+xml",
		"application/pdf",
		"application/msword",
		"application/vnd.ms-word.document.macroenabled.12",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/msword",
	}
)

// Attachments defines the slice of mail attachments
type Attachments []Attachment

// Attachment defines the mail file that has been
// uploaded via the forms endpoint. It contains
// useful information for sending files over
// the mail driver.
type Attachment struct {
	MIMEType string
	Filename string
	MD5name  string
	B64Data  *string
	Size     int64
}

// SizeMB
//
// Returns the attachment file size in megabytes.
func (a *Attachment) SizeMB() int {
	return int(a.Size / 1024) //nolint
}

// getAttachment
//
//
func getAttachment(i interface{}, uploadsPath string) (*Attachment, error) {
	const op = "Forms.getAttachement"

	m, ok := i.(*multipart.FileHeader)
	if !ok {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "", Operation: op, Err: fmt.Errorf("")}
	}

	path, file, teardown, err := createTempFile(m)
	if err != nil {
		return nil, err
	}
	defer teardown()

	mt, err := validateFile(file, m.Size)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "", Operation: op, Err: err}
	}

	md5Name, err := dumpFile(bytes, m.Filename, uploadsPath)
	if err != nil {
		return nil, err
	}

	data := b64(bytes)

	return &Attachment{
		MIMEType: mt,
		Filename: m.Filename,
		MD5name:  md5Name,
		B64Data:  &data,
		Size:     m.Size,
	}, nil
}

// TODO: ptrToRefParam: consider to make non-pointer type for `*multipart.File`
func createTempFile(m *multipart.FileHeader) (string, *multipart.File, func(), error) { //nolint
	const op = "Forms.createTempFile"

	path := os.TempDir() + "/verbis-" + encryption.MD5Hash(time.Now().String()) + filepath.Ext(m.Filename)

	file, err := m.Open()
	if err != nil {
		return "", nil, nil, &errors.Error{Code: errors.INTERNAL, Message: "", Operation: op, Err: err}
	}

	out, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}

	t := func() {
		_ = file.Close()
		_ = out.Close()
		_ = os.Remove(path)
	}

	_, err = io.Copy(out, file)
	if err != nil {
		return "", nil, t, &errors.Error{Code: errors.INTERNAL, Message: "", Operation: op, Err: err}
	}

	return path, &file, t, nil
}

// validateFile
//
// Validates the file attachment for mime types and file sizes.
//
// Returns errors.INVALID if the mime type could not to be detected,
// the mime type is not in the list of permitted types or the
// file is above the UploadLimit.
//
// TODO: ptrToRefParam: consider `file' to be of non-pointer type
func validateFile(file *multipart.File, size int64) (string, error) { //nolint
	const op = "Forms.validateFile"

	typ, err := mimetype.DetectReader(*file)
	if err != nil {
		return "", &errors.Error{Code: errors.INVALID, Message: "Unable to detect filetype", Operation: op, Err: err}
	}

	if !mime.IsValidMime(AllowedMimes, typ.String()) {
		return "", &errors.Error{Code: errors.INVALID, Message: "Mime type not permitted", Operation: op, Err: fmt.Errorf("mime for the uploaded file is not permitted")}
	}

	fileSize := int((size / 1024) / 1024) //nolint

	if fileSize > UploadLimit {
		return "", &errors.Error{Code: errors.INVALID, Message: "File is too large to upload", Operation: op, Err: fmt.Errorf("the file exceeds the upload limit for uploading")}
	}

	return typ.String(), nil
}

// b64
//
// Base64 encodes the attachment to be sent via the
// mailer.
func b64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// dumpFile
//
// Saves the mail attachment to the system by taking
// in the data and name of the file, a new MD5
// hash of the filename will be created and
// saved to the forms storage folder.
//
// Returns errors.INTERNAL if the file could not be created or saved.
func dumpFile(b []byte, name, path string) (string, error) {
	const op = "Forms.dumpFile"

	ext := filepath.Ext(name)
	file := encryption.MD5Hash(name+time.Now().String()) + ext
	dst := path + "/forms/" + file

	err := ioutil.WriteFile(dst, b, os.ModePerm)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Unable to create file to save mail attachment to the system.", Operation: op, Err: err}
	}

	return file, nil
}
