// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"archive/zip"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"io"
	"path/filepath"
)

// DownloadFileName is the recommended name for
// downloading the zip file from the library.
const DownloadFileName = "verbis-library.zip"

// ZipEnclosingFile is the directory the downloaded
// zip file will be enclosed in.
var ZipEnclosingFile = "storage"

// Download satisfies the Provider interface by accepting an
// io.Writer and writing a zip file to the writer.
func (s *Storage) Download(w io.Writer) error {
	const op = "Storage.Download"

	ff, _, err := s.filesRepo.List(params.Params{
		LimitAll: true,
	})

	if err != nil {
		return err
	}

	ar := zip.NewWriter(w)

	for _, f := range ff {
		// Obtain the value of the storage file in bytes.
		buf, err := s.getFileBytes(f)
		if err != nil {
			logger.WithError(err).Warning()
			continue
		}

		// Create s file within the zipWriter with wrapping
		// the enclosing file name (ZipEnclosingFile) and
		// the BucketID (file path).
		fz, err := ar.Create(filepath.Join(ZipEnclosingFile, f.BucketID))
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error creating zip file", Operation: op, Err: err}).Warning()
			continue
		}

		// Write the buffer to the zip file.
		_, err = fz.Write(buf)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error writing zip file", Operation: op, Err: err}).Warning()
			continue
		}
	}

	ar.Close()

	return nil
}
