// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE downloadFile.

package storage

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	repo "github.com/verbiscms/verbis/api/mocks/store/files"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var downloadFile = domain.File{
	ID:         0,
	UUID:       upload.UUID,
	URL:        "test.txt",
	Name:       "test.txt",
	BucketID:   "test.txt",
	Mime:       "text/plain; charset=utf-8",
	SourceType: domain.MediaSourceType,
	Provider:   domain.StorageLocal,
}

var mockDownloadSuccess = func(m *mocks.Service, r *repo.Repository) {
	r.On("List", params.Params{LimitAll: true}).
		Return(domain.Files{downloadFile}, 1, nil)
	c := &mocks.StowContainer{}
	m.On("BucketByFile", downloadFile).Return(c, nil)

	item := &mocks.StowItem{}
	item.On("Open").Return(ioutil.NopCloser(strings.NewReader("test")), nil)
	c.On("Item", mock.Anything).Return(item, nil)
}

func (t *StorageTestSuite) TestStorage_Download() {
	tt := map[string]struct {
		mock func(m *mocks.Service, r *repo.Repository)
		err  bool
		want interface{}
	}{
		"Success": {
			mockDownloadSuccess,
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
		"File Bytes Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("List", params.Params{LimitAll: true}).
					Return(domain.Files{downloadFile}, 1, nil)
				m.On("BucketByFile", downloadFile).
					Return(nil, &errors.Error{Message: "error"})
			},
			true,
			"error",
		},
		"Create Error": {
			func(m *mocks.Service, r *repo.Repository) {
				f := domain.File{BucketID: ""}
				const maxUint16 = 1<<16 + 10
				for i := 0; i < maxUint16; i++ {
					f.BucketID += "t"
				}

				r.On("List", params.Params{LimitAll: true}).
					Return(domain.Files{f}, 1, nil)
				c := &mocks.StowContainer{}
				m.On("BucketByFile", f).Return(c, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(ioutil.NopCloser(strings.NewReader("test")), nil)
				c.On("Item", mock.Anything).Return(item, nil)
			},
			true,
			"Error creating zip downloadFile",
		},
		"Write Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("List", params.Params{LimitAll: true}).
					Return(domain.Files{downloadFile}, 1, nil)
				c := &mocks.StowContainer{}
				m.On("BucketByFile", downloadFile).Return(c, nil)

				item := &mocks.StowItem{}
				var buf []byte
				item.On("Open").Return(ioutil.NopCloser(bytes.NewReader(buf)), nil)
				c.On("Item", mock.Anything).Return(item, nil)
			},
			true,
			"Error writing zip downloadFile",
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
				t.Reset()
			}()

			got := s.Download(open)
			if got != nil {
				t.Contains(got.Error(), test.want)
				return
			}

			reader, err := zip.OpenReader(fileName)
			t.NoError(err)
			defer reader.Close()

			if test.err {
				t.Contains(t.LogWriter.String(), errors.Message(err))
				return
			}

			openFile, err := reader.Open("storage/test.txt")
			t.NoError(err)
			contents, err := ioutil.ReadAll(openFile)
			t.NoError(err)
			t.Equal(test.want, string(contents))
		})
	}
}

type mockZipWriter struct{}

type mockIOWriterErr struct{}

func (m *mockIOWriterErr) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write error")
}

func (m *mockZipWriter) Create(name string) (io.Writer, error) {
	return &mockIOWriterErr{}, nil
}

func (t *StorageTestSuite) TestStorage_Download_WriteError() {
	s := t.Setup(mockDownloadSuccess)

	wg := sync.WaitGroup{}
	wg.Add(1)
	c := make(chan bool, 1)
	c <- true
	go s.addDownloadToZip(&mockZipWriter{}, downloadFile, c, &wg)
	wg.Wait()

	t.Contains(t.LogWriter.String(), "write error")
}
