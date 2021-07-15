// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package files

import (
	"github.com/stretchr/testify/suite"
	"os"
	"path/filepath"
	"testing"
)

// FileTestSuite defines the helper used for file
// testing.
type FileTestSuite struct {
	suite.Suite
	TestPath string
}

// TestFiles asserts testing has begun.
func TestFiles(t *testing.T) {
	suite.Run(t, &FileTestSuite{})
}

// SetupSuite runs before test, to setup the test
// data path.
func (t *FileTestSuite) SetupSuite() {
	wd, err := os.Getwd()
	t.NoError(err)
	t.TestPath = filepath.Join(wd, "testdata")
}

// TestFile is the default test file.
const TestFile = "test.txt"

func (t *FileTestSuite) Test_Exists() {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Exists": {
			filepath.Join(t.TestPath, TestFile),
			true,
		},
		"Dir": {
			t.TestPath,
			false,
		},
		"Not Exists": {
			"",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := Exists(test.input)
			t.Equal(test.want, got)
		})
	}
}

func (t *FileTestSuite) Test_DirectoryExists() {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Exists": {
			t.TestPath,
			true,
		},
		"File": {
			filepath.Join(t.TestPath, TestFile),
			true,
		},
		"Not Exists": {
			"",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := DirectoryExists(test.input)
			t.Equal(test.want, got)
		})
	}
}

func (t *FileTestSuite) Test_RemoveFileExtension() {
	file := "gopher.jpg"
	got := RemoveFileExtension(file)
	want := "gopher"
	t.Equal(want, got)
}
