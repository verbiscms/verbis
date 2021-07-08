// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	storage "github.com/ainsleyclark/verbis/api/mocks/storage"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/media"
)

func (t *MediaServiceTestSuite) TestService_Delete() {
	tt := map[string]struct {
		mock func(r *repo.Repository, s *storage.Bucket)
		want interface{}
	}{
		"Success": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Find", MediaId).Return(domain.Media{Id: MediaId, File: domain.File{Id: 1}}, nil)
				r.On("Delete", MediaId).Return(nil)
				s.On("Delete", 1).Return(fmt.Errorf("error"))
			},
			nil,
		},
		"Find Error": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Find", MediaId).Return(domain.Media{}, fmt.Errorf("error"))
			},
			"error",
		},
		"Delete Error": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Find", MediaId).Return(domain.Media{Id: MediaId}, nil)
				r.On("Delete", MediaId).Return(fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, nil, test.mock)
			err := s.Delete(MediaId)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
		})
	}
}
