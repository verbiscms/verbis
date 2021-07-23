// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/verbiscms/verbis/api/common/params"
	vstrings "github.com/verbiscms/verbis/api/common/strings"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"io/ioutil"
	"path"
	"strings"
)

// List satisfies the Bucket interface by accepting an url
// and retrieving the file and byte contents of the file.
func (s *Storage) List(meta params.Params) (domain.Files, int, error) {
	return s.filesRepo.List(meta)
}

// Find satisfies the Bucket interface by accepting an url
// and retrieving the file and byte contents of the file.
func (s *Storage) Find(url string) ([]byte, domain.File, error) {
	const op = "Storage.Find"

	file, err := s.filesRepo.FindByURL(url)
	if err != nil {
		return nil, domain.File{}, err
	}

	bucket, err := s.service.BucketByFile(file)
	if err != nil {
		return nil, domain.File{}, err
	}

	id := file.ID(s.paths.Storage)

	item, err := bucket.Item(id)
	if err != nil {
		return nil, domain.File{}, &errors.Error{Code: errors.NOTFOUND, Message: "Error obtaining file with the ID: " + id, Operation: op, Err: err}
	}

	open, err := item.Open()
	if err != nil {
		return nil, domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error opening file", Operation: op, Err: err}
	}
	defer open.Close()

	buf, err := ioutil.ReadAll(open)
	if err != nil {
		return nil, domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error reading file", Operation: op, Err: err}
	}

	return buf, file, nil
}

// Upload satisfies the Bucket interface by accepting a
// domain.Upload and inserting into the database and
// uploading to the bucket.
func (s *Storage) Upload(u domain.Upload) (domain.File, error) {
	const op = "Storage.Upload"

	err := u.Validate()
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INVALID, Message: "Validation failed", Operation: op, Err: err}
	}

	provider, bucket, err := s.service.Config()
	if err != nil {
		return domain.File{}, err
	}

	return s.upload(provider, bucket, u, true)
}

func (s *Storage) upload(p domain.StorageProvider, b string, u domain.Upload, createDB bool) (domain.File, error) {
	const op = "Storage.Upload"

	cont, err := s.service.Bucket(p, b)
	if err != nil {
		return domain.File{}, err
	}

	item, err := cont.Put(u.AbsPath(), u.Contents, u.Size, nil)
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INVALID, Message: "Error uploading file to storage provider", Operation: op, Err: err}
	}

	mime, err := u.Mime()
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

	if p.IsLocal() {
		id = vstrings.TrimSlashes(strings.ReplaceAll(item.URL().Path, s.paths.Storage, ""))
		bucket = ""
		region = ""
	}

	f := domain.File{
		UUID:       u.UUID,
		Url:        "/" + vstrings.TrimSlashes(u.Path),
		Name:       path.Base(u.Path),
		BucketId:   id,
		Mime:       mime,
		SourceType: u.SourceType,
		Provider:   p,
		Region:     region,
		Bucket:     bucket,
		FileSize:   u.Size,
		Private:    false,
	}

	if !createDB {
		return f, nil
	}

	file, err := s.filesRepo.Create(f)
	if err != nil {
		return domain.File{}, err
	}

	return file, nil
}

// Exists satisfies the Bucket interface by accepting name
// and determining if a file exists by name.
func (s *Storage) Exists(name string) bool {
	return s.filesRepo.Exists(name)
}

// Delete satisfies the Bucket interface by accepting an ID
// and deleting a file from the bucket and database.
func (s *Storage) Delete(id int) error {
	return s.deleteFile(true, id)
}

func (s *Storage) deleteFile(database bool, id int) error {
	const op = "Storage.Delete"

	file, err := s.filesRepo.Find(id)
	if err != nil {
		return err
	}

	bucket, err := s.service.BucketByFile(file)
	if err != nil {
		return err
	}

	err = bucket.RemoveItem(file.ID(s.paths.Storage))
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting file from storage", Operation: op, Err: err}
	}

	if !database {
		return nil
	}

	err = s.filesRepo.Delete(file.Id)
	if err != nil {
		return err
	}

	return nil
}
