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
	apiPath  string
	testPath string
}

// TestFiles
//
// Assert testing has begun.
func TestFiles(t *testing.T) {
	suite.Run(t, &FileTestSuite{})
}

// SetupAllSuite
//
// Runs before test, set API path.
func (t *FileTestSuite) SetupSuite() {
	wd, err := os.Getwd()
	t.NoError(err)
	api := filepath.Join(filepath.Dir(wd), "../")
	t.apiPath = api
	t.testPath = api + TestPath
}

// The default test path.
const TestPath = "/test/testdata/media"

func (t *FileTestSuite) Test_Exists() {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Exists": {
			t.testPath + string(os.PathSeparator) + "gopher.jpg",
			true,
		},
		"Dir": {
			t.testPath,
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
			t.testPath,
			true,
		},
		"File": {
			t.testPath + string(os.PathSeparator) + "gopher.jpg",
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
