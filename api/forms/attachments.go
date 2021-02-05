// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/gabriel-vasile/mimetype"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

const (
	// The total upload limit allowed to upload file
	// attachments.
	UploadLimit = 5.0
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
type Attachments []*Attachment

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

// getAttachment
//
//
func getAttachment(i interface{}) (*Attachment, error) {
	const op = "Forms.getAttachement"

	mFile, ok := i.(*multipart.FileHeader)
	if !ok {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not cast to multipart FileHeader", Operation: op, Err: fmt.Errorf("unable to cast to multipart.FileHeader")}
	}

	// TODO: If nil

	// TODO: This needs to be dynamic based in the options.
	name, err := dumpFile(mFile)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not dump file to file system", Operation: op, Err: err}
	}

	file, err := mFile.Open()
	if err != nil {

	}
	defer file.Close()

	mimeType, err := validateFile(file, mFile.Size)
	if err != nil {

	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	data := b64(buf.Bytes())

	return &Attachment{
		MIMEType: mimeType,
		Filename: mFile.Filename,
		MD5name:  name,
		B64Data:  &data,
		Size:     mFile.Size,
	}, nil
}

// validateFile
//
// Validates the file attachment for mime types and file sizes.
//
// Returns errors.INVALID if the mime type could not to be detected,
// the mime type is not in the list of permitted types or the
// file is above the UploadLimit.
func validateFile(file multipart.File, size int64) (string, error) {
	const op = "Forms.validateFile"

	typ, err := mimetype.DetectReader(file)
	if err != nil {
		return "", &errors.Error{Code: errors.INVALID, Message: "Unable to detect filetype", Operation: op, Err: err}
	}

	if !mime.IsValidMime(AllowedMimes, typ.String()) {
		return "", &errors.Error{Code: errors.INVALID, Message: "Mime type not permitted", Operation: op, Err: fmt.Errorf("mime for the uploaded file is not permitted")}
	}

	if float64(size / 1024) / 1024  > UploadLimit {
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
func dumpFile(mp *multipart.FileHeader) (string, error) {
	const op = "Forms.dumpFile"

	name := mp.Filename
	ext := filepath.Ext(name)
	file := encryption.MD5Hash(name+time.Now().String()) + ext
	dst := paths.Forms() + string(os.PathSeparator) + file

	return file, files.Save(mp, dst)
}