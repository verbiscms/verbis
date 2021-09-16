// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"archive/zip"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"io"
	"path/filepath"
	"sync"
)

// DownloadFileName is the recommended name for
// downloading the zip file from the library.
const DownloadFileName = "verbis-library.zip"

// ZipEnclosingFile is the directory the downloaded
// zip file will be enclosed in.
var ZipEnclosingFile = "storage"

const (
	// downloadConcurrentAllowance is the amount of files that
	// are allowed to be download concurrently.
	downloadConcurrentAllowance = 10
)

type zipWriter interface {
	Create(name string) (io.Writer, error)
}

// Download satisfies the Provider interface by accepting an
// io.Writer and writing a zip file to the writer.
func (s *Storage) Download(w io.Writer) error {
	ff, _, err := s.filesRepo.List(params.Params{
		LimitAll: true,
	})

	if err != nil {
		return err
	}

	var (
		ar = zip.NewWriter(w)
		wg sync.WaitGroup
		// c is the channel used for sending and processing downloaded
		// files using the downloadConcurrentAllowance
		c = make(chan bool, downloadConcurrentAllowance)
	)

	// Range over the files and increment the
	// wait group. Process the file and
	// add it to the zipWriter.
	for _, f := range ff {
		wg.Add(1)
		c <- true
		go s.addDownloadToZip(ar, f, c, &wg)
	}

	wg.Wait()

	ar.Close()

	return nil
}

func (s *Storage) addDownloadToZip(w zipWriter, file domain.File, channel chan bool, wg *sync.WaitGroup) {
	const op = "Storage.Download"

	// Remove a bool from the channel and release
	// the wait group.
	<-channel
	defer wg.Done()

	// Obtain the value of the storage file in bytes.
	// Downloads the file from the provider.
	buf, err := s.getFileBytes(file)
	if err != nil {
		logger.WithError(err).Warning()
		return
	}

	// Create s file within the zipWriter with wrapping
	// the enclosing file name (ZipEnclosingFile) and
	// the BucketID (file path).
	fz, err := w.Create(filepath.Join(ZipEnclosingFile, file.BucketID))
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error creating zip file", Operation: op, Err: err}).Warning()
		return
	}

	// Write the buffer to the zip file.
	_, err = fz.Write(buf)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error writing zip file", Operation: op, Err: err}).Warning()
		return
	}
}
