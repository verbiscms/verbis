// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/graymeta/stow/azure"
	_ "github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/s3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	repo "github.com/verbiscms/verbis/api/mocks/store/files"
	options "github.com/verbiscms/verbis/api/mocks/store/options"
	"github.com/verbiscms/verbis/api/store/files"
	"io/ioutil"
	"strings"
	"testing"
)

// StorageTestSuite defines the helper used for field
// testing.
type StorageTestSuite struct {
	suite.Suite
}

// TestStorage asserts testing has begun.
func TestStorage(t *testing.T) {
	suite.Run(t, new(StorageTestSuite))
}

// BeforeTest discards
func (t *StorageTestSuite) BeforeTest(suiteName, testName string) {
	logger.SetOutput(ioutil.Discard)
}

// Setup the suite with the mock functions.
func (t *StorageTestSuite) Setup(mf func(s *mocks.Service, r *repo.Repository)) *Storage {
	m := &mocks.Service{}
	r := &repo.Repository{}
	if mf != nil {
		mf(m, r)
	}
	mc := &cache.Store{}
	mc.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Times(100)
	mc.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("errr")).Times(100)
	return &Storage{
		filesRepo: r,
		service:   m,
		env:       &environment.Env{},
		cache:     mc,
	}
}

// Setup the suite with mock functions including
// options.
func (t *StorageTestSuite) SetupOptions(mf func(m *mocks.Service, r *repo.Repository, o *options.Repository)) *Storage {
	m := &mocks.Service{}
	r := &repo.Repository{}
	o := &options.Repository{}
	mc := &cache.Store{}
	mc.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Times(100)
	mc.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("errr")).Times(100)
	if mf != nil {
		mf(m, r, o)
	}
	return &Storage{
		filesRepo:   r,
		optionsRepo: o,
		service:     m,
		cache:       mc,
		env:         &environment.Env{},
	}
}

const (
	// TestFileURL is the default file url used for
	// testing.
	TestFileURL = "/file.txt"
	// TestBucket is the default storage bucket used
	// for testing.
	TestBucket = "verbis-bucket"
)

var (
	// key is the default UUID used for testing.
	key = "5855fe24-e0c5-11eb-ba80-0242ac130004"
	// upload is the default upload used for testing.
	upload = domain.Upload{
		UUID:       uuid.Must(uuid.Parse(key)),
		Path:       "/uploads/2020/01/test.txt",
		Size:       100,
		Contents:   strings.NewReader("test"),
		Private:    false,
		SourceType: domain.MediaSourceType,
	}
	// fileLocal is the default local file used for
	// testing.
	fileLocal = domain.File{
		ID:         0,
		UUID:       upload.UUID,
		URL:        "/uploads/2020/01/test.txt",
		Name:       "test.txt",
		BucketID:   "uploads/2020/01/test.txt",
		Mime:       "text/plain; charset=utf-8",
		SourceType: domain.MediaSourceType,
		Provider:   domain.StorageLocal,
		Region:     "",
		Bucket:     "",
		FileSize:   100,
		Private:    false,
	}
	// fileRemote is the default remote file used for
	// testing.
	fileRemote = domain.File{
		ID:         0,
		UUID:       upload.UUID,
		URL:        "/uploads/2020/01/test.txt",
		Name:       "test.txt",
		BucketID:   "uploads/2020/01/test.txt",
		Mime:       "text/plain; charset=utf-8",
		SourceType: domain.MediaSourceType,
		Provider:   domain.StorageAWS,
		Region:     "",
		Bucket:     "bucket",
		FileSize:   100,
		Private:    false,
	}
	// filesSlice are the default files used for
	// testing.
	filesSlice = domain.Files{
		fileLocal, fileRemote,
	}
)

func (t *StorageTestSuite) TestNew() {
	tt := map[string]struct {
		input Config
		want  interface{}
	}{
		"Success": {
			Config{
				Environment: &environment.Env{},
				Options:     &options.Repository{},
				Files:       &files.Store{},
				Cache:       &cache.Store{},
				Paths:       paths.Paths{Storage: t.T().TempDir()},
			},
			nil,
		},
		"Error Folder": {
			Config{
				Environment: &environment.Env{},
				Options:     &options.Repository{},
				Files:       &files.Store{},
				Cache:       &cache.Store{},
				Paths:       paths.Paths{Storage: "/wrong"},
			},
			"Error creating storage folder",
		},
		"Error": {
			Config{},
			"Error, no Environment set",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got, err := New(test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.NotNil(got)
		})
	}
}

func (t *StorageTestSuite) TestConfig_Validate() {
	tt := map[string]struct {
		input Config
		want  interface{}
	}{
		"Valid": {
			Config{
				Environment: &environment.Env{},
				Options:     &options.Repository{},
				Files:       &files.Store{},
				Cache:       &cache.Store{},
				Paths:       paths.Paths{Storage: "test"},
			},
			nil,
		},
		"Nil Environment": {
			Config{},
			"Error, no Environment set",
		},
		"Nil Options": {
			Config{
				Environment: &environment.Env{},
			},
			"Error, no options repository set",
		},
		"Nil Files": {
			Config{
				Environment: &environment.Env{},
				Options:     &options.Repository{},
			},
			"Error, no files repository set",
		},
		"Nil Cache": {
			Config{
				Environment: &environment.Env{},
				Options:     &options.Repository{},
				Files:       &files.Store{},
			},
			"Error, no cache set",
		},
		"Nil Path": {
			Config{
				Environment: &environment.Env{},
				Options:     &options.Repository{},
				Files:       &files.Store{},
				Cache:       &cache.Store{},
			},
			"Error, no storage path set",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			err := test.input.Validate()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}

type mockIOReaderReadError struct{}

func (m mockIOReaderReadError) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (m mockIOReaderReadError) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func (m mockIOReaderReadError) Close() error {
	return nil
}
