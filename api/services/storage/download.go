// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"archive/zip"
	"github.com/verbiscms/verbis/api/common/params"
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
	ff, _, err := s.filesRepo.List(params.Params{
		LimitAll: true,
	})

	if err != nil {
		return err
	}

	ar := zip.NewWriter(w)

	for _, f := range ff {
		buf, err := s.getFileBytes(f)
		if err != nil {
			continue
		}

		fz, err := ar.Create(filepath.Join(ZipEnclosingFile, f.BucketID))
		if err != nil {
			continue
		}

		_, err = fz.Write(buf)
		if err != nil {
			continue
		}
	}

	ar.Close()

	return nil
}
