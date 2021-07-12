// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	storage "github.com/ainsleyclark/verbis/api/mocks/storage"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/media"
	"github.com/gookit/color"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"path/filepath"
)

var (
	SVGFile = domain.File{
		Id:       1,
		Url:      "/uploads/gopher.svg",
		Name:     "gopher.svg",
		BucketId: "/uploads/gopher.svg",
		Mime:     "image/svg+xml",
		Private:  false,
	}
	PNGFile = domain.File{
		Id:       1,
		Url:      "/uploads/gopher.png",
		Name:     "gopher.png",
		BucketId: "/uploads/gopher.png",
		Mime:     "image/png",
		Private:  false,
	}
)

func (t *MediaServiceTestSuite) TestClient_Upload() {
	tt := map[string]struct {
		input string
		opts  *domain.Options
		mock  func(r *repo.Repository, s *storage.Bucket)
		want  interface{}
	}{
		"SVG": {
			filepath.Join(t.TestDataPath, "/gopher.svg"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.svg").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(SVGFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: SVGFile, FileId: 1}).Return(domain.Media{UserId: 1, File: SVGFile}, nil)
			},
			domain.Media{UserId: 1, File: SVGFile},
		},
		"PNG": {
			filepath.Join(t.TestDataPath, "/gopher.png"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.png").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(PNGFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: PNGFile, FileId: 1}).Return(domain.Media{UserId: 1, File: PNGFile}, nil)
			},
			domain.Media{UserId: 1, File: PNGFile},
		},
		"Upload Error": {
			filepath.Join(t.TestDataPath, "/gopher.svg"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.svg").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(domain.File{}, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Open Error": {
			"",
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.svg").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(domain.File{}, &errors.Error{Message: "error"})
			},
			"Error opening file",
		},
		//"PNG Sizes": {
		//	filepath.Join(t.TestDataPath, "/gopher.png"),
		//	&domain.Options{
		//		MediaSizes: map[string]domain.MediaSize{
		//			"Thumbnail": domain.MediaSize{
		//				SizeKey:  "300",
		//				SizeName: "300",
		//			},
		//		},
		//	},
		//	func(r *repo.Repository, s *storage.Bucket) {
		//		s.On("Exists", "gopher.png").Return(false)
		//		s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(PNGFile, nil)
		//		r.On("Create", domain.Media{UserId: 1, File: PNGFile, FileId: 1}).Return(domain.Media{UserId: 1, File: PNGFile}, nil)
		//	},
		//	domain.Media{UserId: 1, File: PNGFile},
		//},
	}

	for name, test := range tt {
		t.Run(name, func() {
			var mt = &multipart.FileHeader{}
			if test.input != "" {
				mt = t.FileToMultiPart(test.input)
			}

			c := t.Setup(&domain.ThemeConfig{}, test.opts, test.mock)

			got, err := c.Upload(mt, 1)
			if err != nil {
				color.Red.Println(errors.Message(err))
				t.Contains(errors.Message(err), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

//func (t *MediaServiceTestSuite) TestClient_Upload() {
//	size := domain.MediaSize{
//		SizeName: "Test Size",
//		Width:    100,
//		Height:   100,
//		Crop:     true,
//	}
//
//	tt := map[string]struct {
//		input  string
//		cfg    domain.ThemeConfig
//		opts   domain.Options
//		want   domain.Media
//		err    string
//	}{
//		"SVG": {
//			t.MediaPath + "/gopher.svg",
//			domain.ThemeConfig{
//				Media: domain.MediaConfig{
//					UploadPath: "/uploads",
//				},
//			},
//			domain.Options{},
//			domain.Media{
//				Url:      "/uploads/gopher.svg",
//				FilePath: "",
//				FileSize: 7623,
//				FileName: "gopher.svg",
//				Sizes:    domain.MediaSizes{},
//				Mime:     "image/svg+xml",
//			},
//			"",
//		},
//		"PNG": {
//			t.MediaPath + "/gopher.png",
//			domain.ThemeConfig{
//				Media: domain.MediaConfig{
//					UploadPath: "/uploads",
//				},
//			},
//			domain.Options{},
//			domain.Media{
//				Url:      "/uploads/gopher.png",
//				FilePath: "",
//				FileSize: 163677,
//				FileName: "gopher.png",
//				Sizes:    domain.MediaSizes{},
//				Mime:     "image/png",
//			},
//			"",
//		},
//		"JPG": {
//			t.MediaPath + "/gopher.jpg",
//			domain.ThemeConfig{
//				Media: domain.MediaConfig{
//					UploadPath: "/uploads",
//				},
//			},
//			domain.Options{},
//			domain.Media{
//				Url:      "/uploads/gopher.jpg",
//				FilePath: "",
//				FileSize: 162023,
//				FileName: "gopher.jpg",
//				Sizes:    domain.MediaSizes{},
//				Mime:     "image/jpeg",
//			},
//			"",
//		},
//		"PNG Size": {
//			t.MediaPath + "/gopher.png",
//			domain.ThemeConfig{
//				Media: domain.MediaConfig{
//					UploadPath: "/uploads",
//				},
//			},
//			domain.Options{
//				MediaSizes: domain.MediaSizes{
//					"fileToWebP": size,
//				},
//			},
//			domain.Media{
//				Url:      "/uploads/gopher.png",
//				FilePath: "",
//				FileSize: 163677,
//				FileName: "gopher.png",
//				Sizes: domain.MediaSizes{
//					"fileToWebP": domain.MediaSize{
//						Url:      "/uploads/gopher-100x100.png",
//						Name:     "gopher-100x100.png",
//						SizeName: "Test Size",
//						Width:    100,
//						Height:   100,
//						Crop:     true,
//					},
//				},
//				Mime: "image/png",
//			},
//			"",
//		},
//		"JPG Size": {
//			t.MediaPath + "/gopher.jpg",
//			domain.ThemeConfig{
//				Media: domain.MediaConfig{
//					UploadPath: "/uploads",
//				},
//			},
//			domain.Options{
//				MediaSizes: domain.MediaSizes{
//					"fileToWebP": size,
//				},
//			},
//			domain.Media{
//				Url:      "/uploads/gopher.jpg",
//				FilePath: "",
//				FileSize: 162023,
//				FileName: "gopher.jpg",
//				Sizes: domain.MediaSizes{
//					"fileToWebP": domain.MediaSize{
//						Url:      "/uploads/gopher-100x100.jpg",
//						Name:     "gopher-100x100.jpg",
//						SizeName: "Test Size",
//						Width:    100,
//						Height:   100,
//						Crop:     true,
//					},
//				},
//				Mime: "image/jpeg",
//			},
//			"",
//		},
//		"Nil": {
//			"",
//			domain.ThemeConfig{},
//			domain.Options{},
//			domain.Media{},
//			"Error opening file with the name",
//		},
//	}
//
//	for name, fileToWebP := range tt {
//		t.Run(name, func() {
//			c := t.Setup(fileToWebP.cfg, fileToWebP.opts)
//			c.exists = fileToWebP.exists
//
//			var mt = &multipart.FileHeader{}
//			if fileToWebP.input != "" {
//				mt = t.File(fileToWebP.input)
//			}
//
//			got, err := c.Upload(mt)
//			if err != nil {
//				t.Contains(errors.Message(err), fileToWebP.err)
//				return
//			}
//
//			defer func() {
//				c.Delete(got)
//			}()
//
//			t.Equal(fileToWebP.want.Url, got.Url)
//			t.Equal(fileToWebP.want.FilePath, got.FilePath)
//			t.Equal(fileToWebP.want.FileSize, got.FileSize)
//			t.Equal(fileToWebP.want.FileName, got.FileName)
//			t.Equal(fileToWebP.want.Mime, got.Mime)
//
//			if fileToWebP.want.Sizes != nil {
//				for k, v := range fileToWebP.want.Sizes {
//					t.Equal(v.Url, got.Sizes[k].Url)
//					t.Equal(v.Name, got.Sizes[k].Name)
//					t.Equal(v.SizeName, got.Sizes[k].SizeName)
//					t.Equal(v.Width, got.Sizes[k].Width)
//					t.Equal(v.Height, got.Sizes[k].Height)
//					t.Equal(v.Crop, got.Sizes[k].Crop)
//				}
//			}
//		})
//	}
//}
