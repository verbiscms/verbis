// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

//
//// MediaServiceTestSuite defines the helper used for media
//// library testing.
//type ImageTestSuite struct {
//	test.MediaSuite
//}
//
//// TestImageSuite asserts testing has begun.
//func TestImageSuite(t *testing.T) {
//	suite.Run(t, &ImageTestSuite{
//		MediaSuite: test.NewMediaSuite(),
//	})
//}
//
//func (t *ImageTestSuite) TestResize() {
//	tt := map[string]struct {
//		input domain.MediaSize
//		mock  func(m *mocks.Imager)
//		want  interface{}
//	}{
//		"decode Error": {
//			domain.MediaSize{},
//			func(m *mocks.Imager) {
//				m.On("decode").Return(fmt.Errorf("error"))
//			},
//			nil,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			m := &mocks.Imager{}
//			test.mock(m)
//			r := Resize{}
//
//			err := r.Resize(m, "d", test.input, 0)
//			if err != nil {
//				t.Contains(err.Error(), test.want)
//				return
//			}
//
//		})
//	}
//}
//
//func (t *ImageTestSuite) UtilTestImageSave(i Imager, ext string) {
//	image := t.Image()
//
//	// Test save success
//	tmp := t.T().TempDir()
//	path := filepath.Join(tmp, "verbis-test-image") + ext
//
//	err := i.Save(image, path, 0)
//	t.NoError(err)
//
//	// Test save error
//	err = i.Save(image, "wrong", 0)
//	t.Error(err)
//}
//
//type mockFileSeekErr struct{}
//
//func (m *mockFileSeekErr) Read(p []byte) (n int, err error) {
//	return 0, nil
//}
//
//func (m *mockFileSeekErr) ReadAt(p []byte, off int64) (n int, err error) {
//	return 0, nil
//}
//
//func (m *mockFileSeekErr) Seek(offset int64, whence int) (int64, error) {
//	return 0, fmt.Errorf("error")
//}
//
//func (m *mockFileSeekErr) Close() error {
//	return nil
//}
//
//func (t *ImageTestSuite) TestJPG_Save() {
//	t.UtilTestImageSave(&JPG{}, ".jpg")
//}
//
//func (t *ImageTestSuite) TestJPG_Decode() {
//	tt := map[string]struct {
//		input func() (multipart.File, func() error)
//		want  interface{}
//	}{
//		"Success": {
//			func() (multipart.File, func() error) {
//				m := t.File(filepath.Join(t.MediaPath, "gopher.jpg"))
//				file, _ := m.Open() // Ignore on purpose
//				return file, file.Close
//			},
//			nil,
//		},
//		"Seek Error": {
//			func() (multipart.File, func() error) {
//				return &mockFileSeekErr{}, func() error { return nil }
//			},
//			"error",
//		},
//		"decode Error": {
//			func() (multipart.File, func() error) {
//				m := t.File(filepath.Join(t.MediaPath, "gopher.png"))
//				file, _ := m.Open() // Ignore on purpose
//				return file, file.Close
//			},
//			"invalid JPEG format",
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			file, teardown := test.input()
//			defer teardown()
//
//			j := &JPG{File: file}
//			decode, err := j.decode()
//			if err != nil {
//				t.Contains(err.Error(), test.want)
//				return
//			}
//			t.NotNil(decode)
//			t.Nil(err)
//		})
//	}
//}
//
//func (t *ImageTestSuite) TestPNG_Save() {
//	t.UtilTestImageSave(&PNG{}, ".png")
//}
//
//func (t *ImageTestSuite) TestPNG_Decode() {
//	tt := map[string]struct {
//		input func() (multipart.File, func() error)
//		want  interface{}
//	}{
//		"Success": {
//			func() (multipart.File, func() error) {
//				m := t.File(filepath.Join(t.MediaPath, "gopher.png"))
//				file, _ := m.Open() // Ignore on purpose
//				return file, file.Close
//			},
//			nil,
//		},
//		"Seek Error": {
//			func() (multipart.File, func() error) {
//				return &mockFileSeekErr{}, func() error { return nil }
//			},
//			"error",
//		},
//		"decode Error": {
//			func() (multipart.File, func() error) {
//				m := t.File(filepath.Join(t.MediaPath, "gopher.jpg"))
//				file, _ := m.Open() // Ignore on purpose
//				return file, file.Close
//			},
//			"not a PNG file",
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			file, teardown := test.input()
//			defer teardown()
//
//			j := &PNG{File: file}
//			decode, err := j.decode()
//			if err != nil {
//				t.Contains(err.Error(), test.want)
//				return
//			}
//			t.NotNil(decode)
//			t.Nil(err)
//		})
//	}
//}
