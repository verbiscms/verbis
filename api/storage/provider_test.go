// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/mocks/storage/mocks"
)

func (t *StorageTestSuite) TestProvider_SetProvider() {
	tt := map[string]struct {
		input domain.StorageProvider
		mock  func(s *mocks.Service)
		want  interface{}
	}{
		"Local": {
			domain.StorageLocal,
			func(s *mocks.Service) {
				s.On("Provider", domain.StorageLocal).Return(&mocks.StowLocation{}, nil)
			},
			domain.StorageLocal,
		},
		"Remote": {
			domain.StorageAWS,
			func(s *mocks.Service) {
				s.On("Provider", domain.StorageAWS).Return(&mocks.StowLocation{}, nil)
			},
			domain.StorageAWS,
		},
		"Error": {
			domain.StorageLocal,
			func(s *mocks.Service) {
				s.On("Provider", domain.StorageLocal).Return(nil, fmt.Errorf("Error"))
			},
			"Error setting provider",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := &mocks.Service{}

			if test.mock != nil {
				test.mock(m)
			}

			loc := &mocks.StowLocation{}
			s := Storage{
				service:      m,
				stowLocation: loc,
			}

			err := s.SetProvider(test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}

			t.NotEqual(&s.stowLocation, &loc)
			t.Equal(test.input, s.ProviderName)
		})
	}
}
