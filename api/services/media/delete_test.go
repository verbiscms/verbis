// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/services/media"
	"github.com/google/uuid"
)

func (t *MediaServiceTestSuite) TestClient_Delete() {
	tt := map[string]struct {
		input     domain.Media
		mock      func(m *mocks.Library)
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
			func(m *mocks.Library) {

			},
			".png",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {

		})
	}
}
