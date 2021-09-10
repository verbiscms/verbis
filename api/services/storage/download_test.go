// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"archive/zip"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	repo "github.com/verbiscms/verbis/api/mocks/store/files"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func (t *StorageTestSuite) TestStorage_Download() {
	file := domain.File{
		ID:         0,
		UUID:       upload.UUID,
		URL:        "test.txt",
		Name:       "test.txt",
		BucketID:   "test.txt",
		Mime:       "text/plain; charset=utf-8",
		SourceType: domain.MediaSourceType,
		Provider:   domain.StorageLocal,
	}

	tt := map[string]struct {
		mock     func(m *mocks.Service, r *repo.Repository)
		writeErr bool
		want     interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("List", params.Params{LimitAll: true}).
					Return(domain.Files{file}, 1, nil)
				c := &mocks.StowContainer{}
				m.On("BucketByFile", file).Return(c, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(ioutil.NopCloser(strings.NewReader("test")), nil)
				c.On("Item", mock.Anything).Return(item, nil)
			},
			false,
			"test",
		},
		"Repo Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("List", params.Params{LimitAll: true}).
					Return(nil, 0, fmt.Errorf("error"))
			},
			false,
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)

			fileName := filepath.Join(t.T().TempDir(), "verbis-test.zip")
			open, err := os.Create(fileName)
			t.NoError(err)

			defer func() {
				open.Close()
				os.Remove(fileName)
			}()

			got := s.Download(open)
			if got != nil {
				t.Contains(got.Error(), test.want)
				return
			}

			reader, err := zip.OpenReader(fileName)
			t.NoError(err)
			defer reader.Close()

			openFile, err := reader.Open("storage/test.txt")
			if test.writeErr {
				fmt.Println(err)
				t.Error(err)
				return
			} else {
				t.NoError(err)
			}

			contents, err := ioutil.ReadAll(openFile)
			t.NoError(err)
			t.Equal(test.want, string(contents))
		})
	}
}
