// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	vstrings "github.com/ainsleyclark/verbis/api/helpers/strings"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

type bucket struct {
	Config
	Location
}

// Find satisfies the Bucket interface by accepting an url
// and retrieving the file and byte contents of the file.
func (b *bucket) Find(path string) ([]byte, domain.File, error) {
	const op = "Storage.Find"

	file, err := b.Files.FindByURL(path)
	if err != nil {
		return nil, domain.File{}, err
	}

	bucket, err := b.Location.Get(file.Provider)
	if err != nil {
		return nil, domain.File{}, err
	}

	id := file.ID(b.paths.Storage)

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

// Find satisfies the Bucket interface by accepting an url
// and retrieving the file and byte contents of the file.
func (s *Storage) Find(path string) ([]byte, domain.File, error) {
	const op = "Storage.Find"

	file, err := s.repo.FindByURL(path)
	if err != nil {
		return nil, domain.File{}, err
	}

	bucket, err := s.getBucket(file)
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

// Upload Satisfies the Bucket interface by accepting a
// domain.Upload and inserting into the database and
// uploading to the bucket.
func (s *Storage) Upload(u domain.Upload) (domain.File, error) {
	const op = "Storage.Upload"

	err := u.Validate()
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INVALID, Message: "Validation failed", Operation: op, Err: err}
	}

	var (
		// UUID for the upload.
		key = uuid.New()
		// E.g. /2021/01/24d4ad32-53e7-4728-a2d5-35e297ac9abe.txt
		absPath = strings.TrimPrefix(filepath.Join(filepath.Dir(u.Path), key.String()+filepath.Ext(u.Path)), ".")
	)

	item, err := s.bucket.Put(absPath, u.Contents, u.Size, nil)
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INVALID, Message: "Error uploading file to storage provider", Operation: op, Err: err}
	}

	_, err = u.Contents.Seek(0, 0)
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error seeking bytes", Operation: op, Err: err}
	}

	m, err := mimetype.DetectReader(u.Contents)
	if err != nil {
		return domain.File{}, &errors.Error{Code: errors.INTERNAL, Message: "Error obtaining mime type", Operation: op, Err: err}
	}

	var (
		// E.g. uploads/2021/07/ea5101e3-9730-49cd-855b-a068524c6fd5.jpg
		id = item.ID()
		// E.g. bucket-name
		bucket = s.bucket.ID()
		// TODO E.g eu-west-2
		region = ""
	)

	if s.ProviderName == domain.StorageLocal {
		id = vstrings.TrimSlashes(strings.ReplaceAll(item.URL().Path, s.paths.Storage, ""))
		bucket = ""
		region = ""
	}

	f := domain.File{
		UUID:       key,
		Url:        "/" + vstrings.TrimSlashes(u.Path),
		Name:       path.Base(u.Path),
		BucketId:   id,
		Mime:       domain.Mime(m.String()),
		SourceType: u.SourceType,
		Provider:   s.ProviderName,
		Region:     region,
		Bucket:     bucket,
		FileSize:   u.Size,
		Private:    false,
	}

	create, err := s.repo.Create(f)
	if err != nil {
		return domain.File{}, err
	}

	return create, nil
}

// Exists satisfies the Bucket interface by accepting name
// and determining if a file exists by name.
func (s *Storage) Exists(name string) bool {
	return s.repo.Exists(name)
}

// Delete satisfies the Bucket interface by accepting an ID
// and deleting a file from the bucket and database.
func (s *Storage) Delete(id int) error {
	const op = "Storage.Delete"

	file, err := s.repo.Find(id)
	if err != nil {
		return err
	}

	bucket, err := s.getBucket(file)
	if err != nil {
		return &errors.Error{Code: errors.NOTFOUND, Message: "Error obtaining file from storage", Operation: op, Err: err}
	}

	err = bucket.RemoveItem(file.ID(s.paths.Storage))
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting file from storage", Operation: op, Err: err}
	}

	err = s.repo.Delete(file.Id)
	if err != nil {
		return err
	}

	return nil
}
