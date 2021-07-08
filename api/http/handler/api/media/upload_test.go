// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	service "github.com/ainsleyclark/verbis/api/mocks/services/media"
	users "github.com/ainsleyclark/verbis/api/mocks/store/users"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"net/http"
)

func (t *MediaTestSuite) TestMedia_Upload() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		files   int
		mock    func(s *service.Library, u *users.Repository, multi []multipart.FileHeader)
		url     string
	}{
		"Success": {
			mediaItemPublic,
			http.StatusOK,
			"Successfully uploaded media item",
			1,
			func(s *service.Library, u *users.Repository, multi []multipart.FileHeader) {
				s.On("Validate", &multi[0]).Return(nil)
				u.On("FindByToken", mock.Anything).Return(domain.User{}, nil)
				s.On("Upload", &multi[0], mock.Anything).Return(mediaItem, nil)
			},
			"/media",
		},
		"No Form": {
			nil,
			http.StatusBadRequest,
			"No files attached to the upload",
			0,
			func(s *service.Library, u *users.Repository, multi []multipart.FileHeader) {
				s.On("Upload", multipart.FileHeader{}, "").Return(domain.Media{}, nil)
				s.On("Validate", multipart.FileHeader{}).Return(nil)
			},
			"/media",
		},
		"No Files": {
			nil,
			http.StatusBadRequest,
			"Attach a file to the request to be uploaded",
			0,
			func(s *service.Library, u *users.Repository, multi []multipart.FileHeader) {
				s.On("Upload", multipart.FileHeader{}, "").Return(domain.Media{}, nil)
				s.On("Validate", multipart.FileHeader{}).Return(nil)
			},
			"/media",
		},
		"Too Many Files": {
			nil,
			http.StatusBadRequest,
			"Files are only permitted to be uploaded one at a time",
			5,
			func(s *service.Library, u *users.Repository, multi []multipart.FileHeader) {
				s.On("Upload", &multi[0], "").Return(domain.Media{}, nil)
				s.On("Validate", &multi[0]).Return(nil)
			},
			"/media",
		},
		"Invalid": {
			nil,
			http.StatusUnsupportedMediaType,
			"invalid",
			1,
			func(s *service.Library, u *users.Repository, multi []multipart.FileHeader) {
				s.On("Validate", &multi[0]).Return(&errors.Error{Code: errors.INVALID, Message: "invalid"})
				u.On("FindByToken", mock.Anything).Return(domain.User{}, nil)
			},
			"/media",
		},
		"Token Error": {
			nil,
			http.StatusUnauthorized,
			"You must be logged in to uploaded media items",
			1,
			func(s *service.Library, u *users.Repository, multi []multipart.FileHeader) {
				u.On("FindByToken", mock.Anything).Return(domain.User{}, fmt.Errorf("error"))
			},
			"/media",
		},
		"Upload Error": {
			nil,
			http.StatusInternalServerError,
			"error",
			1,
			func(s *service.Library, u *users.Repository, multi []multipart.FileHeader) {
				s.On("Validate", &multi[0]).Return(nil)
				u.On("FindByToken", mock.Anything).Return(domain.User{}, nil)
				s.On("Upload", &multi[0], mock.Anything).Return(mediaItem, &errors.Error{Code: errors.INVALID, Message: "error"})
			},
			"/media",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			request, multi := t.UploadRequest(test.files, "https://google.com/upload", t.ImagePath())
			t.Context.Request = request

			if name == "No Form" {
				t.RequestAndServe(http.MethodGet, test.url, "/media", nil, func(ctx *gin.Context) {
					t.SetupUpload(multi, test.mock).Upload(t.Context)
				})
			} else {
				t.SetupUpload(multi, test.mock).Upload(t.Context)
			}

			t.RunT(test.want, test.status, test.message)
		})
	}
}
