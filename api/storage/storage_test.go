// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	filesMock "github.com/ainsleyclark/verbis/api/mocks/store/files"
	optionsMock "github.com/ainsleyclark/verbis/api/mocks/store/options"
	"github.com/ainsleyclark/verbis/api/store/files"
	"github.com/ainsleyclark/verbis/api/store/options"
	"github.com/google/uuid"
	_ "github.com/graymeta/stow/azure"
	_ "github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/s3"
	"github.com/stretchr/testify/suite"
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

func (t *StorageTestSuite) TestNew() {
	tt := map[string]struct {
		input  func() (Config, func())
		panics bool
		want   interface{}
	}{
		//"Test": {
		//	func() (Config, func()) {
		//		tmp := t.T().TempDir()
		//
		//		err := os.MkdirAll(filepath.Join(tmp, "storage"), os.ModePerm)
		//		if err != nil {
		//			t.Fail(err.Error())
		//		}
		//
		//		o := &optionsMock.Repository{}
		//		o.On("Struct").Return(&domain.Options{
		//			StorageProvider: domain.StorageLocal,
		//			StorageBucket:   "bucket",
		//		})
		//		o.On("Update", "storage_bucket", "").Return(nil)
		//
		//		return Config{
		//				Environment: &environment.Env{},
		//				Options:     o,
		//				Files:       &filesMock.Repository{},
		//			}, func() {
		//				os.RemoveAll(tmp)
		//			}
		//	},
		//	false,
		//	nil,
		//},
		"Bad Config": {
			func() (Config, func()) {
				return Config{}, nil
			},
			false,
			"Error",
		},
		"Bad Provider": {
			func() (Config, func()) {
				o := &optionsMock.Repository{}
				o.On("Struct").Return(&domain.Options{
					StorageProvider: "test",
					StorageBucket:   "",
				})
				return Config{
					Environment: &environment.Env{},
					Options:     o,
					Files:       &filesMock.Repository{},
				}, nil
			},
			false,
			"Error setting up storage with provider",
		},
		"Provider Error": {
			func() (Config, func()) {
				o := &optionsMock.Repository{}
				o.On("Struct").Return(&domain.Options{
					StorageProvider: domain.StorageLocal,
					StorageBucket:   "",
				})
				return Config{
					Environment: &environment.Env{},
					Options:     o,
					Files:       &filesMock.Repository{},
				}, nil
			},
			false,
			"Error setting provider",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			cfg, teardown := test.input()
			if teardown != nil {
				defer teardown()
			}
			got, err := New(cfg)
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
				Options:     &options.Store{},
				Files:       &files.Store{},
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
				Options:     &options.Store{},
			},
			"Error, no files repository set",
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

var (
	// TODO
	key = "5855fe24-e0c5-11eb-ba80-0242ac130004"
	u   = domain.Upload{
		UUID:       uuid.Must(uuid.Parse(key)),
		Path:       "/uploads/2020/01/test.txt",
		Size:       100,
		Contents:   strings.NewReader("test"),
		Private:    false,
		SourceType: domain.MediaSourceType,
	}
	// TODO
	fileLocal = domain.File{
		Id:         0,
		UUID:       u.UUID,
		Url:        "/uploads/2020/01/test.txt",
		Name:       "test.txt",
		BucketId:   "uploads/2020/01/test.txt",
		Mime:       "text/plain; charset=utf-8",
		SourceType: domain.MediaSourceType,
		Provider:   domain.StorageLocal,
		Region:     "",
		Bucket:     "",
		FileSize:   100,
		Private:    false,
	}
	// TODO
	fileRemote = domain.File{
		Id:         0,
		UUID:       u.UUID,
		Url:        "/uploads/2020/01/test.txt",
		Name:       "test.txt",
		BucketId:   "uploads/2020/01/test.txt",
		Mime:       "text/plain; charset=utf-8",
		SourceType: domain.MediaSourceType,
		Provider:   domain.StorageAWS,
		Region:     "",
		Bucket:     "bucket",
		FileSize:   100,
		Private:    false,
	}
)

type mockIOReaderError struct{}

func (m mockIOReaderError) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (m mockIOReaderError) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func (m mockIOReaderError) Close() error {
	return nil
}
