// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/suite"
	"io"
	"mime/multipart"
	gohttp "net/http"
	"os"
	"path/filepath"
	"testing"
)

// MediaTestSuite defines the helper used for media
// testing.
type MediaTestSuite struct {
	api.HandlerSuite
}

// TestCategories
//
// Assert testing has begun.
func TestMedia(t *testing.T) {
	suite.Run(t, &MediaTestSuite{
		HandlerSuite: api.TestSuite(),
	})
}

// Setup
//
// A helper to obtain a mock categories handler
// for testing.
func (t *MediaTestSuite) Setup(mf func(m *mocks.MediaRepository)) *Media {
	m := &mocks.MediaRepository{}
	if mf != nil {
		mf(m)
	}
	return &Media{
		Deps: &deps.Deps{
			Store: &models.Store{
				Media: m,
			},
		},
	}
}

// SetupUpload
//
// A helper to obtain a mock categories handler
// and uploads for testing.
func (t *MediaTestSuite) SetupUpload(files []multipart.FileHeader, mf func(m *mocks.MediaRepository, mfh []multipart.FileHeader)) *Media {
	m := &mocks.MediaRepository{}
	if mf != nil {
		mf(m, files)
	}
	return &Media{
		Deps: &deps.Deps{
			Store: &models.Store{
				Media: m,
			},
		},
	}
}

// ImagePath
//
// Returns a dummy image from test data.
func (t *MediaTestSuite) ImagePath() string {
	wd, err := os.Getwd()
	t.NoError(err)
	apiPath := filepath.Join(filepath.Dir(wd), "../../..")
	return apiPath + "/test/testdata/spa/images/gopher.svg"
}

var (
	// The default media item used for testing.
	mediaItem = domain.Media{
		Id: 123,
	}
	// The default media item with wrong validation used for testing.
	mediaBadValidation = &domain.Media{}
	// The default media items used for testing.
	mediaItems = []domain.Media{
		{
			Id:    1,
			Url:   "/uploads/1",
			Title: "title",
		},
		{
			Id:    1,
			Url:   "/uploads/1",
			Title: "title",
		},
	}
	// The default pagination used for testing.
	pagination = params.Params{
		Page:           api.DefaultParams.Page,
		Limit:          15,
		OrderBy:        api.DefaultParams.OrderBy,
		OrderDirection: api.DefaultParams.OrderDirection,
		Filters:        nil,
	}
)

// UploadRequest
//
// Is a helper for setting up test files for the upload
// endpoint. Creates a new file upload http request
//  with optional extra params.
func (t *MediaTestSuite) UploadRequest(filesAmount int, uri string, path string) (*gohttp.Request, []multipart.FileHeader) {
	file, err := os.Open(path)
	t.NoError(err)
	defer file.Close()

	reqBody := bytes.Buffer{}
	var multi []multipart.FileHeader
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for i := 0; i < filesAmount; i++ {
		part, err := writer.CreateFormFile("file", filepath.Base(path))
		t.NoError(err)
		_, err = io.Copy(part, file)
		t.NoError(err)
	}

	err = writer.Close()
	t.NoError(err)

	reqBody.Write(body.Bytes())

	if filesAmount != 0 {
		mr := multipart.NewReader(body, writer.Boundary())
		mt, err := mr.ReadForm(99999)
		t.NoError(err)
		ft := mt.File["file"][0]
		multi = append(multi, *ft)
	}

	req, err := gohttp.NewRequest("POST", uri, bytes.NewBuffer(reqBody.Bytes()))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, multi
}
