// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	//"bytes"
	"encoding/base64"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/gabriel-vasile/mimetype"
	"io"
	"io/ioutil"

	//"io/ioutil"
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

// SizeMB
//
// Returns the attachment file size in megabytes.
func (a *Attachment) SizeMB() int {
	return int(a.Size / 1024)
}

// getAttachment
//
//
func getAttachment(i interface{}) (*Attachment, error) {
	const op = "Forms.getAttachement"

	m, ok := i.(*multipart.FileHeader)
	if !ok {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "", Operation: op, Err: fmt.Errorf("")}
	}

	file, err := m.Open()
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "", Operation: op, Err: err}
	}
	defer file.Close()

	out, err := os.Create(paths.Storage() + "/forms/temp.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)

	bytes, err := ioutil.ReadFile(paths.Storage() + "/forms/temp.jpg")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", b64(bytes))

	err = ioutil.WriteFile(paths.Storage()+"/forms/final.jpg", bytes, 777)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		// delete temp
	}()

	//mt, err := validateFile(file, m.Size)
	//if err != nil {
	//	return nil, err
	//}

	//err = ioutil.WriteFile("/Users/ainsley/Desktop/test.jpg", buf, 777)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//return nil, err

	//
	////_, err = ioutil.ReadFile(m.Filename)
	//buf := bytes.NewBuffer(nil)
	//_, err = buf.ReadFrom(file)
	//if err != nil {
	//	return nil, err
	//}
	//
	////test, err := file.Read(buf)
	//
	//fmt.Printf("%+v\n", buf)
	//
	////buf := bytes.NewBuffer(nil)
	////_, err = io.Copy(buf, file)
	////if err != nil {
	////	return nil, &errors.Error{Code: errors.INTERNAL, Message: "", Operation: op, Err: err}
	////}
	//
	//// TODO: This needs to be dynamic based in the options.
	//name, err := dumpFile(buf.String(), m.Filename)
	//if err != nil {
	//	return nil, &errors.Error{Code: errors.INTERNAL, Message: "", Operation: op, Err: err}
	//}
	//
	data := b64(bytes)

	return &Attachment{
		MIMEType: "image/jpg",
		Filename: m.Filename,
		MD5name:  "fuckwank.jpg",
		B64Data:  &data,
		Size:     m.Size,
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

	fileSize := int(1024 / size)
	fmt.Println(fileSize)
	fmt.Println(size)
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
func dumpFile(s string, name string) (string, error) {
	const op = "Forms.dumpFile"

	ext := filepath.Ext(name)
	file := encryption.MD5Hash(name+time.Now().String()) + ext
	dst := paths.Forms() + "/" + file

	f, err := os.Create(dst)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Unable to create file to save mail attachment to the system.", Operation: op, Err: err}
	}

	_, err = f.WriteString(s)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Unable to save mail attachment to the system.", Operation: op, Err: err}
	}

	return file, nil
}
