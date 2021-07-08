// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/logger"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/files"
	_ "github.com/graymeta/stow/azure"
	_ "github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/s3"
	"github.com/stretchr/testify/suite"
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

func (t *StorageTestSuite) TestStorage_Exists() {

}

func (t *StorageTestSuite) TestStorage_Delete() {
	tt := map[string]struct {
		mock func(r *mocks.Repository)
		want interface{}
	}{
		"Resource": {
			func(r *mocks.Repository) {

			},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			fmt.Println(test)
		})
	}
}
