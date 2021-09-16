// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE downloadFile.

package storage

import (
	"github.com/verbiscms/verbis/api/common/params"
	vstrings "github.com/verbiscms/verbis/api/common/strings"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"io/ioutil"
	"path"
	"strings"
	"sync"
)

// List satisfies the Bucket interface by accepting an url
// and retrieving the downloadFile and byte contents of the downloadFile.
func (s *Storage) List(meta params.Params) (domain.Files, int, error) {
	return s.filesRepo.List(meta)
}

// Find satisfies the Bucket interface by accepting an url
// and retrieving the downloadFile and byte contents of the downloadFile.
func (s *Storage) Find(url string) ([]byte, domain.File, error) {
	const op = "Storage.Find"

	file, err := s.filesRepo.FindByURL(url)
	if err != nil {
		return nil, domain.File{}, err
	}

	buf, err := s.getFileBytes(file)
	if err != nil {
		return nil, domain.File{}, err
	}

	return buf, file, nil
}

// getFileBytes retrieves the downloadFile's bytes with the given
// input, returns an error if the downloadFile was not found
// or
func (s *Storage) getFileBytes(file domain.File) ([]byte, error) {
	const op = "Storage.GetFileBytes"

	bucket, err := s.service.BucketByFile(file)
	if err != nil {
		return nil, err
	}

	id := file.FullPath(s.paths.Storage)

	item, err := bucket.Item(id)
	if err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "Error obtaining downloadFile with the ID: " + id, Operation: op, Err: err}
	}

	open, err := item.Open()
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error opening downloadFile", Operation: op, Err: err}
	}
	defer open.Close()

	buf, err := ioutil.ReadAll(open)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error reading downloadFile", Operation: op, Err: err}
	}

	return buf, nil
}

// Upload satisfies the Bucket interface by accepting a
// domain.Upload and inserting into the database and
// uploading to the bucket.
func (s *Storage) Upload(u domain.Upload) (domain.File, error) {
	const op = "Storage.Upload"

	// Check for any upload errors before processing.
	err := u.Validate()
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INVALID, Message: "Validation failed", Operation: op, Err: err}
	}

	// Obtain the configuration for the upload.
	info := s.service.Config()

	logger.Debug("Processing upload: " + u.Path)

	if !info.UploadRemote {
		info.Provider = domain.StorageLocal
		info.Bucket = ""
	}

	// TODO: Data race detected here with backups!
	file, err := s.upload(&uploadCfg{
		Provider:       info.Provider,
		Bucket:         info.Bucket,
		Upload:         u,
		CreateDatabase: true,
	})
	if err == nil {
		logger.Debug("Successfully processed upload: " + u.Path)
	}

	mtx := sync.Mutex{}

	// Spawn a routine for backing up the storage files.
	go func() {
		mtx.Lock()
		defer mtx.Unlock()

		// If the local backup is enabled and the current settings
		// are remote, backup the upload to the local provider.
		if info.LocalBackup && !info.Provider.IsLocal() {
			s.backup(domain.StorageLocal, "", u)
		}

		// If the server backup is enabled and the current settings
		// are local, backup the upload to the remote provider.
		if info.RemoteBackup && info.Provider.IsLocal() {
			// Obtain the correct connected provider.
			cfg := s.service.Config()
			s.backup(cfg.Provider, cfg.Bucket, u)
		}
	}()

	return file, err
}

func (s *Storage) backup(p domain.StorageProvider, b string, u domain.Upload) {
	_, err := s.upload(&uploadCfg{
		Provider:       p,
		Bucket:         b,
		Upload:         u,
		CreateDatabase: false,
	})
	logger.Debug("Processing backup with storage provider: " + p.String() + " and filepath: " + u.Path)
	if err != nil {
		logger.WithError(err).Error()
	}
	logger.Debug("Successfully processed backup: " + u.Path)
}

type uploadCfg struct {
	Provider       domain.StorageProvider
	Bucket         string
	Upload         domain.Upload
	CreateDatabase bool
}

func (s *Storage) upload(cfg *uploadCfg) (domain.File, error) {
	const op = "Storage.Upload"

	cont, err := s.service.Bucket(cfg.Provider, cfg.Bucket)
	if err != nil {
		return domain.File{}, err
	}

	// Seek the downloadFile back to the original bytes.
	cfg.Upload.Contents.Seek(0, 0) //nolint

	item, err := cont.Put(cfg.Upload.AbsPath(), cfg.Upload.Contents, cfg.Upload.Size, nil)
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INVALID, Message: "Error uploading downloadFile to storage provider", Operation: op, Err: err}
	}

	mime, err := cfg.Upload.Mime()
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error obtaining mime type", Operation: op, Err: err}
	}

	var (
		// E.g. uploads/2021/07/ea5101e3-9730-49cd-855b-a068524c6fd5.jpg
		id = item.ID()
		// E.g. b-name
		bucket = cont.ID()
		// TODO E.g eu-west-2
		region = ""
	)

	if cfg.Provider.IsLocal() {
		id = vstrings.TrimSlashes(strings.ReplaceAll(item.URL().Path, s.paths.Storage, ""))
		bucket = ""
		region = ""
	}

	f := domain.File{
		UUID:       cfg.Upload.UUID,
		URL:        "/" + vstrings.TrimSlashes(cfg.Upload.Path),
		Name:       path.Base(cfg.Upload.Path),
		BucketID:   id,
		Mime:       mime,
		SourceType: cfg.Upload.SourceType,
		Provider:   cfg.Provider,
		Region:     region,
		Bucket:     bucket,
		FileSize:   cfg.Upload.Size,
		Private:    false,
	}

	if !cfg.CreateDatabase {
		return f, nil
	}

	file, err := s.filesRepo.Create(f)
	if err != nil {
		return domain.File{}, err
	}

	return file, nil
}

// Exists satisfies the Bucket interface by accepting name
// and determining if a downloadFile exists by name.
func (s *Storage) Exists(name string) bool {
	return s.filesRepo.Exists(name)
}

// Delete satisfies the Bucket interface by accepting an ID
// and deleting a downloadFile from the bucket and database.
func (s *Storage) Delete(id int) error {
	file, err := s.deleteFile(true, id)
	if err != nil {
		return err
	}
	go s.deleteBackups(file)
	return nil
}

func (s *Storage) deleteFile(database bool, id int) (domain.File, error) {
	const op = "Storage.Delete"

	file, err := s.filesRepo.Find(id)
	if err != nil {
		return domain.File{}, err
	}

	bucket, err := s.service.BucketByFile(file)
	if err != nil {
		return file, err
	}

	err = bucket.RemoveItem(file.FullPath(s.paths.Storage))
	if err != nil {
		return file, &errors.Error{Code: errors.INVALID, Message: "Error deleting downloadFile from storage", Operation: op, Err: err}
	}

	if !database {
		return file, nil
	}

	err = s.filesRepo.Delete(file.ID)
	if err != nil {
		return file, err
	}

	return file, nil
}

// deleteBackups deletes possible backup files from
// the remote or local provider.
func (s *Storage) deleteBackups(file domain.File) {
	cfg := s.service.Config()

	if file.IsLocal() {
		file.Provider = cfg.Provider
		file.Bucket = cfg.Bucket
	} else {
		file.Provider = domain.StorageLocal
		file.Bucket = ""
	}

	bucket, err := s.service.BucketByFile(file)
	if err != nil {
		return
	}

	err = bucket.RemoveItem(file.FullPath(s.paths.Storage))
	if err != nil {
		return
	}
}
