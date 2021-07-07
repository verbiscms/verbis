// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

//
//func (t *MediaServiceTestSuite) TestClient_Serve() {
//	tt := map[string]struct {
//		webp    bool
//		input   domain.Media
//		options domain.Options
//		want    interface{}
//	}{
//		"PNG": {
//			false,
//			domain.Media{
//				UUID:     uuid.MustParse("a66e03d9-cc64-4176-b226-47d0e7a2debb"),
//				FileSize: 0,
//				Url:      "/test.png",
//				FileName: "test.png",
//				Mime:     "image/png",
//			},
//			domain.Options{},
//			domain.Mime("image/png"),
//		},
//		"JPG": {
//			false,
//			domain.Media{
//				UUID:     uuid.MustParse("a66e03d9-cc64-4176-b226-47d0e7a2debb"),
//				FileSize: 0,
//				Url:      "/test.jpg",
//				FileName: "test.jpg",
//				Mime:     "image/jpg",
//			},
//			domain.Options{},
//			domain.Mime("image/jpg"),
//		},
//		"WebP": {
//			true,
//			domain.Media{
//				UUID:     uuid.MustParse("a66e03d9-cc64-4176-b226-47d0e7a2debb"),
//				FileSize: 0,
//				Url:      "/test.jpg",
//				FileName: "test.jpg",
//				Mime:     "image/jpg",
//			},
//			domain.Options{
//				MediaServeWebP: true,
//			},
//			domain.Mime("image/webp"),
//		},
//		"WebP Failed": {
//			true,
//			domain.Media{
//				UUID:     uuid.MustParse("2af11d5c-88ae-11eb-8dcd-0242ac130003"),
//				FileSize: 0,
//				Url:      "/test.jpg",
//				FileName: "test.jpg",
//				Mime:     "image/jpg",
//			},
//			domain.Options{
//				MediaServeWebP: true,
//			},
//			domain.Mime("image/jpg"),
//		},
//		"Not Found": {
//			false,
//			domain.Media{
//				UUID: uuid.UUID{},
//			},
//			domain.Options{},
//			"File does not exist with the path",
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(domain.ThemeConfig{}, test.options)
//			path := t.MediaPath + string(os.PathSeparator) + test.input.UUID.String() + filepath.Ext(test.input.FileName)
//			_, mime, err := s.Serve(test.input, path, test.webp)
//			if err != nil {
//				t.Contains(errors.Message(err), test.want)
//				t.Equal(domain.Mime(""), mime)
//				return
//			}
//			t.Equal(test.want, mime)
//		})
//	}
//}
