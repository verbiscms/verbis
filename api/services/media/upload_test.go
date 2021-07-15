// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/common/paths"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	resizer "github.com/ainsleyclark/verbis/api/mocks/services/media/resizer"
	webp "github.com/ainsleyclark/verbis/api/mocks/services/webp"
	storage "github.com/ainsleyclark/verbis/api/mocks/storage"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/media"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"path/filepath"
	"time"
)

func (t *MediaServiceTestSuite) TestClient_Upload() {
	tt := map[string]struct {
		input   string
		opts    *domain.Options
		mock    func(r *repo.Repository, s *storage.Bucket)
		resizer func(r *resizer.Resizer)
		want    interface{}
	}{
		"SVG": {
			filepath.Join(t.TestDataPath, "gopher.svg"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.svg").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(SVGFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: SVGFile, FileId: 1}).Return(domain.Media{UserId: 1, File: SVGFile}, nil)
			},
			nil,
			domain.Media{UserId: 1, File: SVGFile},
		},
		"JPG": {
			filepath.Join(t.TestDataPath, "gopher.jpg"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.jpg").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(JPGFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: JPGFile, FileId: 1}).Return(domain.Media{UserId: 1, File: JPGFile}, nil)
			},
			nil,
			domain.Media{UserId: 1, File: JPGFile},
		},
		"PNG": {
			filepath.Join(t.TestDataPath, "gopher.png"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.png").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(PNGFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: PNGFile, FileId: 1}).Return(domain.Media{UserId: 1, File: PNGFile}, nil)
			},
			nil,
			domain.Media{UserId: 1, File: PNGFile},
		},
		"Open Error": {
			"",
			&domain.Options{},
			nil,
			nil,
			"Error opening file",
		},
		"Upload Error": {
			filepath.Join(t.TestDataPath, "gopher.svg"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.svg").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(domain.File{}, &errors.Error{Message: "error"})
			},
			nil,
			"error",
		},
		"Resize Error": {
			filepath.Join(t.TestDataPath, "gopher.jpg"),
			opts,
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.jpg").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(JPGFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: JPGFile, FileId: 1}).Return(domain.Media{}, &errors.Error{Message: "error"})
			},
			func(r *resizer.Resizer) {
				r.On("Resize", mock.Anything, mock.Anything, mock.Anything).Return(nil, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Repo Error": {
			filepath.Join(t.TestDataPath, "gopher.jpg"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.jpg").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(JPGFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: JPGFile, FileId: 1}).Return(domain.Media{}, &errors.Error{Message: "error"})
			},
			nil,
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(&domain.ThemeConfig{}, test.opts, test.mock)

			var mt = &multipart.FileHeader{}
			if test.input != "" {
				multi, _ := t.ToMultiPartE(test.input)
				mt = multi
			}

			r := &resizer.Resizer{}
			if test.resizer != nil {
				test.resizer(r)
			}
			s.resizer = r

			got, err := s.Upload(mt, 1)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestClient_Dir() {
	now := time.Now()
	tt := map[string]struct {
		input *domain.Options
		want  string
	}{
		"Date": {
			&domain.Options{MediaOrganiseDate: true},
			filepath.Join(paths.Uploads, now.Format("2006"), now.Format("01")),
		},
		"Prefix": {
			&domain.Options{MediaOrganiseDate: false},
			paths.Uploads,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, test.input, nil)
			got := s.dir()
			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestClient_CleanFileName() {
	tt := map[string]struct {
		name string
		ext  string
		mock func(r *repo.Repository, s *storage.Bucket)
		want string
	}{
		"Simple": {
			"gopher",
			".png",
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.png").Return(false)
			},
			"gopher",
		},
		"Remove Characters": {
			"g£&*@oph£$er",
			".png",
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.png").Return(false)
			},
			"gopher",
		},
		"Exists": {
			"gopher",
			".png",
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.png").Return(true)
				s.On("Exists", "gopher-1.png").Return(false)
			},
			"gopher-1",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, nil, test.mock)
			got := s.cleanFileName(test.name, test.ext)
			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestClient_Resize() {
	var m domain.MediaSizes

	tt := map[string]struct {
		file    domain.File
		multi   func() multipart.File
		mock    func(r *repo.Repository, s *storage.Bucket)
		resizer func(r *resizer.Resizer)
		want    interface{}
	}{
		"JPG": {
			JPGFile,
			func() multipart.File {
				part := t.ToMultiPart(filepath.Join(t.TestDataPath, "gopher.jpg"))
				open, err := part.Open()
				if err != nil {
					t.Fail(err.Error())
				}
				return open
			},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(JPGFile, nil)
			},
			func(r *resizer.Resizer) {
				r.On("Resize", mock.Anything, mock.Anything, mock.Anything).Return(bytes.NewReader([]byte("test")), nil)
			},
			domain.MediaSizes{"thumbnail": domain.MediaSize{FileId: 1, SizeKey: "thumbnail", SizeName: "gopher-300x300.jpg", Width: 300, Height: 300, File: JPGFile}},
		},
		"PNG": {
			PNGFile,
			func() multipart.File {
				part := t.ToMultiPart(filepath.Join(t.TestDataPath, "gopher.png"))
				open, err := part.Open()
				if err != nil {
					t.Fail(err.Error())
				}
				return open
			},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(JPGFile, nil)
			},
			func(r *resizer.Resizer) {
				r.On("Resize", mock.Anything, mock.Anything, mock.Anything).Return(bytes.NewReader([]byte("test")), nil)
			},
			domain.MediaSizes{"thumbnail": domain.MediaSize{FileId: 1, SizeKey: "thumbnail", SizeName: "gopher-300x300.png", Width: 300, Height: 300, File: JPGFile}},
		},
		"Cant Resize": {
			SVGFile,
			func() multipart.File {
				return nil
			},
			nil,
			nil,
			m,
		},
		"Resize Error": {
			JPGFile,
			func() multipart.File {
				part := t.ToMultiPart(filepath.Join(t.TestDataPath, "gopher.jpg"))
				open, err := part.Open()
				if err != nil {
					t.Fail(err.Error())
				}
				return open
			},
			nil,
			func(r *resizer.Resizer) {
				r.On("Resize", mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error"))
			},
			"error",
		},
		"Upload Error": {
			JPGFile,
			func() multipart.File {
				part := t.ToMultiPart(filepath.Join(t.TestDataPath, "gopher.jpg"))
				open, err := part.Open()
				if err != nil {
					t.Fail(err.Error())
				}
				return open
			},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(domain.File{}, fmt.Errorf("error"))
			},
			func(r *resizer.Resizer) {
				r.On("Resize", mock.Anything, mock.Anything, mock.Anything).Return(bytes.NewReader([]byte("test")), nil)
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			r := &resizer.Resizer{}

			s := t.Setup(nil, opts, test.mock)

			if test.resizer != nil {
				test.resizer(r)
			}
			s.resizer = r

			file := test.multi()
			if file != nil {
				defer file.Close()
			}

			got, err := s.resize(test.file, file)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestClient_TopWebP() {
	tt := map[string]struct {
		input   domain.Media
		mock    func(r *repo.Repository, s *storage.Bucket)
		webp    func(e *webp.Execer)
		options *domain.Options
		want    string
	}{
		"Success": {
			domain.Media{File: PNGFile, Sizes: domain.MediaSizes{"thumbnail": domain.MediaSize{File: PNGFile}}},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", PNGFile.Url).Return([]byte("test"), PNGFile, nil)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(domain.File{}, nil)
			},
			func(e *webp.Execer) {
				e.On("Convert", bytes.NewReader([]byte("test")), 0).Return(bytes.NewReader([]byte("test")), nil)
			},
			&domain.Options{MediaConvertWebP: true},
			"Successfully converted to WebP image with the path",
		},
		"Find Error": {
			domain.Media{File: PNGFile},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", PNGFile.Url).Return(nil, domain.File{}, fmt.Errorf("find error"))
			},
			nil,
			&domain.Options{MediaConvertWebP: true},
			"find error",
		},
		"Convert Error": {
			domain.Media{File: PNGFile},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", PNGFile.Url).Return([]byte("test"), PNGFile, nil)
			},
			func(e *webp.Execer) {
				e.On("Convert", bytes.NewReader([]byte("test")), 0).Return(nil, fmt.Errorf("convert error"))
			},
			&domain.Options{MediaConvertWebP: true},
			"convert error",
		},
		"Storage Error": {
			domain.Media{File: PNGFile},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", PNGFile.Url).Return([]byte("test"), PNGFile, nil)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(domain.File{}, fmt.Errorf("storage error"))
			},
			func(e *webp.Execer) {
				e.On("Convert", bytes.NewReader([]byte("test")), 0).Return(bytes.NewReader([]byte("test")), nil)
			},
			&domain.Options{MediaConvertWebP: true},
			"storage error",
		},
		"OptionsBAD Permitted": {
			domain.Media{},
			nil,
			nil,
			&domain.Options{MediaConvertWebP: false},
			"",
		},
		"Cant Resize": {
			domain.Media{File: SVGFile},
			nil,
			nil,
			&domain.Options{MediaConvertWebP: true},
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, test.options, test.mock)
			w := &webp.Execer{}

			if test.webp != nil {
				test.webp(w)
			}
			s.webp = w

			s.toWebP(test.input)
			t.Contains(t.LogWriter.String(), test.want)
			t.Reset()
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
//		opts   domain.OptionsBAD
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
//			domain.OptionsBAD{},
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
//			domain.OptionsBAD{},
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
//			domain.OptionsBAD{},
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
//			domain.OptionsBAD{
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
//			domain.OptionsBAD{
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
//			domain.OptionsBAD{},
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
