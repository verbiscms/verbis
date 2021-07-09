// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
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
	logWriter bytes.Buffer
}

// TestStorage asserts testing has begun.
func TestStorage(t *testing.T) {
	suite.Run(t, new(StorageTestSuite))
}

// BeforeTest assigns the logger to a buffer.
func (t *StorageTestSuite) BeforeTest(suiteName, testName string) Storage {
	b := bytes.Buffer{}
	t.logWriter = b
	logger.SetOutput(&t.logWriter)

	s := Storage{}

	return s
}

func (t *StorageTestSuite) TestNew() {

}

var (
	key = "5855fe24-e0c5-11eb-ba80-0242ac130004"

	u = domain.Upload{
		UUID:       uuid.Must(uuid.Parse(key)),
		Path:       "/uploads/2020/01/test.txt",
		Size:       100,
		Contents:   strings.NewReader("test"),
		Private:    false,
		SourceType: domain.MediaSourceType,
	}
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

type MockIOReaderError struct{}

func (m MockIOReaderError) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func (m MockIOReaderError) Close() error {
	return nil
}

type MockIOSeekerError struct{}

func (m MockIOSeekerError) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (m MockIOSeekerError) Seek(offset int64, whence int) (int64, error) {
	return 0, fmt.Errorf("Error")
}

type MockIOReaderSeekerError struct{}

func (m MockIOReaderSeekerError) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func (m MockIOReaderSeekerError) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}
