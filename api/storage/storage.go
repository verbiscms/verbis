// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	vstrings "github.com/ainsleyclark/verbis/api/helpers/strings"
	"github.com/ainsleyclark/verbis/api/store/files"
	"github.com/ainsleyclark/verbis/api/store/options"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/graymeta/stow"
	_ "github.com/graymeta/stow/azure"
	_ "github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/s3"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

// Client defines the methods used for interacting with
// the Verbis storage system. The client can be remote
// or work with the local file system dependant on
// what is set in the options table.
type Client interface {
	// Find looks up the file with the given URL and retrieves
	// the appropriate bucket to obtain the file contents.
	// It returns the byte value of the file as well as
	// the domain.File.
	// Returns errors.INTERNAL if the file could not be opened or read.
	// Returns errors.NOTFOUND if the file could not be retrieved from the bucket.
	Find(url string) ([]byte, domain.File, error)
	// Upload adds a domain.Upload to the database as well as
	// the bucket that is currently set in the env. The
	// file is seeked, the mime type is obtained and it
	// is inserted into the database and uploaded to
	// the bucket.
	// Returns errors.INVALID if the bucket could not be obtained.
	// Returns errors.INTERNAL if the contents couldn't be seeked or the mime
	// type could not be obtained.
	Upload(upload domain.Upload) (domain.File, error)
	// Delete removes an item from the the bucket. It first retrieves
	// the file by a lookup from the database, obtains the correct
	// bucket, then removes the file from the storage provider.
	// The file data will also be deleted from
	// the database.
	// Returns errors.INVALID if the file could not be deleted from the bucket.
	Delete(id int) error
	// Exists queries the database by the given name and
	// returns true if there was a match.
	Exists(name string) bool
	Provider
}

type Provider interface {
	ListBuckets() (domain.Buckets, error)
	SetProvider(location domain.StorageProvider) error
	SetBucket(id string) error
}

type Storage struct {
	ProviderName domain.StorageProvider
	provider     stow.Location
	bucket       stow.Container
	optsRepo     options.Repository
	opts         *domain.Options
	env          *environment.Env
	paths        paths.Paths
	repo         files.Repository
}

var (
	// ErrNoProvider is returned by New and SetProvider when
	// there is no match from the options table.
	ErrNoProvider = errors.New("invalid provider")
)

// New parse config
func New(env *environment.Env, opts options.Repository, repo files.Repository) (*Storage, error) {
	const op = "Storage.New"

	s := &Storage{
		env:      env,
		opts:     opts.Struct(),
		paths:    paths.Get(),
		repo:     repo,
		optsRepo: opts,
	}

	provider := s.opts.StorageProvider
	if provider == "" {
		provider = domain.StorageLocal
	}

	if !validate(provider) {
		return nil, &errors.Error{Code: errors.INVALID, Message: string("Error setting up storage with provider: " + provider), Operation: op, Err: ErrNoProvider}
	}

	err := s.SetProvider(provider)
	if err != nil {
		return nil, err
	}

	err = s.SetBucket(s.opts.StorageBucket)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Find satisfies the Client interface by accepting an url
// and retrieving the file and byte contents of the file.
func (s *Storage) Find(path string) ([]byte, domain.File, error) {
	const op = "Storage.FindByURL"

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

// Upload Satisfies the Client interface by accepting a
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
		// E.g eu-west-2
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

// Exists satisfies the Client interface by accepting name
// and determining if a file exists by name.
func (s *Storage) Exists(name string) bool {
	return s.repo.Exists(name)
}

// Delete satisfies the Client interface by accepting an ID
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

// InSlice checks if a string exists in a slice,
func validate(p domain.StorageProvider) bool {
	for _, sp := range domain.StorageProviders {
		if sp == p {
			return true
		}
	}
	return false
}
