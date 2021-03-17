// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"io"
	"mime/multipart"
	"net/http"
	"path"
)

// DownloadFile
//
// Retrieves a file from a specific url and copies it to a
// multipart.FileHeader ready to be uploaded by the
// media repository.
//
// Returns errors.NOTFOUND if the status code is anything but 200.
// Returns errors.INTERNAL if the file could not be created, copied, closed or read.
func DownloadFile(url string) (*multipart.FileHeader, error) {
	const op = "Importer.DownloadFile"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("File with the path %s not found", path.Base(url)), Operation: op, Err: fmt.Errorf("status code of %v with the url of: %s", resp.StatusCode, url)}
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", path.Base(url))
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not create new file", Operation: op, Err: err}
	}

	_, err = io.Copy(part, resp.Body)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not copy new file", Operation: op, Err: err}
	}

	err = writer.Close()
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not close file", Operation: op, Err: err}
	}

	mr := multipart.NewReader(body, writer.Boundary())
	mt, err := mr.ReadForm(99999) //nolint
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not read new file", Operation: op, Err: err}
	}

	return mt.File["file"][0], nil
}
