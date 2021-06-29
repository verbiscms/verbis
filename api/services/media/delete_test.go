// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/google/uuid"
	"os"
)

func (t *MediaServiceTestSuite) TestClient_Delete() {
	tt := map[string]struct {
		input     domain.Media
		extension string
	}{
		"PNG": {
			domain.Media{
				UUID:     uuid.New(),
				FileSize: 0,
				Url:      "/test.png",
				FileName: "test.png",
				Mime:     "image/png",
			},
			".png",
		},
		"JPG": {
			domain.Media{
				UUID:     uuid.New(),
				FileSize: 0,
				Url:      "/test.jpg",
				FileName: "test.jpg",
				Mime:     "image/jpeg",
			},
			".jpg",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			file := t.MediaPath + string(os.PathSeparator) + test.input.UUID.String() + test.extension
			teardown := t.DummyFile(file)
			defer teardown()

			webp := t.MediaPath + string(os.PathSeparator) + test.input.UUID.String() + test.extension + domain.WebPExtension
			teardownWebP := t.DummyFile(webp)
			defer teardownWebP()

			c := t.Setup(domain.ThemeConfig{}, domain.Options{})
			c.Delete(test.input)

			_, err := os.Stat(file)
			if !os.IsNotExist(err) {
				t.Fail("File wasn't deleted, cleaning up")
				err := os.Remove(file)
				if err != nil {
					t.Fail("Could not delete file! Clean up manually")
				}
			}

			_, err = os.Stat(webp)
			if !os.IsNotExist(err) {
				t.Fail("File wasn't deleted, cleaning up")
				err := os.Remove(file)
				if err != nil {
					t.Fail("Could not delete file! Clean up manually")
				}
			}
		})
	}
}
