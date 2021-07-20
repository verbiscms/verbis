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
	storage "github.com/ainsleyclark/verbis/api/mocks/services/storage"
	webp "github.com/ainsleyclark/verbis/api/mocks/services/webp"
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
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(svgFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: svgFile, FileId: 1}).Return(domain.Media{UserId: 1, File: svgFile}, nil)
			},
			nil,
			domain.Media{UserId: 1, File: svgFile},
		},
		"JPG": {
			filepath.Join(t.TestDataPath, "gopher.jpg"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.jpg").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(jpgFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: jpgFile, FileId: 1}).Return(domain.Media{UserId: 1, File: jpgFile}, nil)
			},
			nil,
			domain.Media{UserId: 1, File: jpgFile},
		},
		"PNG": {
			filepath.Join(t.TestDataPath, "gopher.png"),
			&domain.Options{},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Exists", "gopher.png").Return(false)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(pngFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: pngFile, FileId: 1}).Return(domain.Media{UserId: 1, File: pngFile}, nil)
			},
			nil,
			domain.Media{UserId: 1, File: pngFile},
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
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(jpgFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: jpgFile, FileId: 1}).Return(domain.Media{}, &errors.Error{Message: "error"})
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
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(jpgFile, nil)
				r.On("Create", domain.Media{UserId: 1, File: jpgFile, FileId: 1}).Return(domain.Media{}, &errors.Error{Message: "error"})
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
			jpgFile,
			func() multipart.File {
				part := t.ToMultiPart(filepath.Join(t.TestDataPath, "gopher.jpg"))
				open, err := part.Open()
				if err != nil {
					t.Fail(err.Error())
				}
				return open
			},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(jpgFile, nil)
			},
			func(r *resizer.Resizer) {
				r.On("Resize", mock.Anything, mock.Anything, mock.Anything).Return(bytes.NewReader([]byte("test")), nil)
			},
			domain.MediaSizes{"thumbnail": domain.MediaSize{FileId: 1, SizeKey: "thumbnail", SizeName: "thumb", Width: 300, Height: 300, File: jpgFile}},
		},
		"PNG": {
			pngFile,
			func() multipart.File {
				part := t.ToMultiPart(filepath.Join(t.TestDataPath, "gopher.png"))
				open, err := part.Open()
				if err != nil {
					t.Fail(err.Error())
				}
				return open
			},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(jpgFile, nil)
			},
			func(r *resizer.Resizer) {
				r.On("Resize", mock.Anything, mock.Anything, mock.Anything).Return(bytes.NewReader([]byte("test")), nil)
			},
			domain.MediaSizes{"thumbnail": domain.MediaSize{FileId: 1, SizeKey: "thumbnail", SizeName: "thumb", Width: 300, Height: 300, File: jpgFile}},
		},
		"Cant Resize": {
			svgFile,
			func() multipart.File {
				return nil
			},
			nil,
			nil,
			m,
		},
		"Resize Error": {
			jpgFile,
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
			jpgFile,
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
			domain.Media{File: pngFile, Sizes: domain.MediaSizes{"thumbnail": domain.MediaSize{File: pngFile}}},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", pngFile.Url).Return([]byte("test"), pngFile, nil)
				s.On("Upload", mock.AnythingOfType("domain.Upload")).Return(domain.File{}, nil)
			},
			func(e *webp.Execer) {
				e.On("Convert", bytes.NewReader([]byte("test")), 0).Return(bytes.NewReader([]byte("test")), nil)
			},
			&domain.Options{MediaConvertWebP: true},
			"Successfully converted to WebP image with the path",
		},
		"Find Error": {
			domain.Media{File: pngFile},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", pngFile.Url).Return(nil, domain.File{}, fmt.Errorf("find error"))
			},
			nil,
			&domain.Options{MediaConvertWebP: true},
			"find error",
		},
		"Convert Error": {
			domain.Media{File: pngFile},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", pngFile.Url).Return([]byte("test"), pngFile, nil)
			},
			func(e *webp.Execer) {
				e.On("Convert", bytes.NewReader([]byte("test")), 0).Return(nil, fmt.Errorf("convert error"))
			},
			&domain.Options{MediaConvertWebP: true},
			"convert error",
		},
		"Storage Error": {
			domain.Media{File: pngFile},
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", pngFile.Url).Return([]byte("test"), pngFile, nil)
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
			domain.Media{File: svgFile},
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
