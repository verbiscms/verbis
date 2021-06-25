// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"mime/multipart"
	"os"
	"time"
)

func (t *MediaServiceTestSuite) TestClient_Upload() {
	size := domain.MediaSize{
		SizeName: "Test Size",
		Width:    100,
		Height:   100,
		Crop:     true,
	}

	tt := map[string]struct {
		input  string
		cfg    domain.ThemeConfig
		opts   domain.Options
		exists ExistsFunc
		want   domain.Media
		err    string
	}{
		"SVG": {
			t.MediaPath + "/gopher.svg",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					UploadPath: "/uploads",
				},
			},
			domain.Options{},
			exists,
			domain.Media{
				Url:      "/uploads/gopher.svg",
				FilePath: "",
				FileSize: 7623,
				FileName: "gopher.svg",
				Sizes:    domain.MediaSizes{},
				Mime:     "image/svg+xml",
			},
			"",
		},
		"PNG": {
			t.MediaPath + "/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					UploadPath: "/uploads",
				},
			},
			domain.Options{},
			exists,
			domain.Media{
				Url:      "/uploads/gopher.png",
				FilePath: "",
				FileSize: 163677,
				FileName: "gopher.png",
				Sizes:    domain.MediaSizes{},
				Mime:     "image/png",
			},
			"",
		},
		"JPG": {
			t.MediaPath + "/gopher.jpg",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					UploadPath: "/uploads",
				},
			},
			domain.Options{},
			exists,
			domain.Media{
				Url:      "/uploads/gopher.jpg",
				FilePath: "",
				FileSize: 162023,
				FileName: "gopher.jpg",
				Sizes:    domain.MediaSizes{},
				Mime:     "image/jpeg",
			},
			"",
		},
		"PNG Size": {
			t.MediaPath + "/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					UploadPath: "/uploads",
				},
			},
			domain.Options{
				MediaSizes: domain.MediaSizes{
					"test": size,
				},
			},
			exists,
			domain.Media{
				Url:      "/uploads/gopher.png",
				FilePath: "",
				FileSize: 163677,
				FileName: "gopher.png",
				Sizes: domain.MediaSizes{
					"test": domain.MediaSize{
						Url:      "/uploads/gopher-100x100.png",
						Name:     "gopher-100x100.png",
						SizeName: "Test Size",
						Width:    100,
						Height:   100,
						Crop:     true,
					},
				},
				Mime: "image/png",
			},
			"",
		},
		"JPG Size": {
			t.MediaPath + "/gopher.jpg",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					UploadPath: "/uploads",
				},
			},
			domain.Options{
				MediaSizes: domain.MediaSizes{
					"test": size,
				},
			},
			exists,
			domain.Media{
				Url:      "/uploads/gopher.jpg",
				FilePath: "",
				FileSize: 162023,
				FileName: "gopher.jpg",
				Sizes: domain.MediaSizes{
					"test": domain.MediaSize{
						Url:      "/uploads/gopher-100x100.jpg",
						Name:     "gopher-100x100.jpg",
						SizeName: "Test Size",
						Width:    100,
						Height:   100,
						Crop:     true,
					},
				},
				Mime: "image/jpeg",
			},
			"",
		},
		"Nil": {
			"",
			domain.ThemeConfig{},
			domain.Options{},
			exists,
			domain.Media{},
			"Error opening file with the name",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.cfg, test.opts)
			c.exists = test.exists

			var mt = &multipart.FileHeader{}
			if test.input != "" {
				mt = t.File(test.input)
			}

			got, err := c.Upload(mt)
			if err != nil {
				t.Contains(errors.Message(err), test.err)
				return
			}

			defer func() {
				c.Delete(got)
			}()

			t.Equal(test.want.Url, got.Url)
			t.Equal(test.want.FilePath, got.FilePath)
			t.Equal(test.want.FileSize, got.FileSize)
			t.Equal(test.want.FileName, got.FileName)
			t.Equal(test.want.Mime, got.Mime)

			if test.want.Sizes != nil {
				for k, v := range test.want.Sizes {
					t.Equal(v.Url, got.Sizes[k].Url)
					t.Equal(v.Name, got.Sizes[k].Name)
					t.Equal(v.SizeName, got.Sizes[k].SizeName)
					t.Equal(v.Width, got.Sizes[k].Width)
					t.Equal(v.Height, got.Sizes[k].Height)
					t.Equal(v.Crop, got.Sizes[k].Crop)
				}
			}
		})
	}
}

func (t *MediaServiceTestSuite) TestClient_Upload_DirError() {
	c := t.Setup(domain.ThemeConfig{}, domain.Options{})
	c.exists = exists
	mt := t.File(t.MediaPath + "/gopher.svg")

	c.paths.Uploads = "wrongpath"

	_, err := c.Upload(mt)

	want := "Error creating file"
	t.Contains(errors.Message(err), want)
}

func (t *MediaServiceTestSuite) TestUploader_SaveOriginal_Error() {
	file, _ := os.Open("") // Ignore on purpose
	u := uploader{File: file}
	_, err := u.SaveOriginal(t.T().TempDir())
	t.Error(err)
}

func (t *MediaServiceTestSuite) TestUploader_URL() {
	now := time.Now().Format("2006/01")

	tt := map[string]struct {
		input domain.Options
		cfg   domain.MediaConfig
		want  string
	}{
		"With Date": {
			domain.Options{MediaOrganiseDate: true},
			domain.MediaConfig{UploadPath: "uploads"},
			"/uploads/" + now,
		},
		"Without Date": {
			domain.Options{MediaOrganiseDate: false},
			domain.MediaConfig{UploadPath: "uploads"},
			"/uploads",
		},
		"Trailing Slash": {
			domain.Options{MediaOrganiseDate: true},
			domain.MediaConfig{UploadPath: "uploads/"},
			"/uploads/" + now,
		},
		"Leading Slash": {
			domain.Options{MediaOrganiseDate: true},
			domain.MediaConfig{UploadPath: "//uploads"},
			"/uploads/" + now,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			u := uploader{
				Options: &test.input,
				Config:  &domain.ThemeConfig{Media: test.cfg},
			}
			got := u.URL()
			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestUploader_FileSize() {
	tt := map[string]struct {
		input string
		want  int64
	}{
		"Exists": {
			t.MediaPath + "/gopher.svg",
			100,
		},
		"Not Found": {
			"",
			0,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			u := uploader{}
			got := u.FileSize(test.input)
			if test.want == 0 {
				t.Equal(test.want, got)
				return
			}
			t.GreaterOrEqual(got, test.want)
		})
	}
}
