// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type noStringer struct{}

func Setup() (*Namespace, *mocks.MediaRepository) {
	mock := &mocks.MediaRepository{}
	return &Namespace{deps: &deps.Deps{
		Store: &models.Store{
			Media: mock,
		},
	}}, mock
}

func TestNamespace_Find(t *testing.T) {

	media := domain.Media{
		Id:  1,
		Url: "/uploads/test.jpg",
	}

	id := 1
	idFloat32 := float32(1)
	idFloat64 := float64(1)

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.MediaRepository)
		want  interface{}
	}{
		"Success": {
			1,
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			media,
		},
		"No Item": {
			1,
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("no media"))
			},
			nil,
		},
		"nil": {
			nil,
			func(m *mocks.MediaRepository) {
				m.On("GetById", nil).Return(domain.Media{}, fmt.Errorf("no media"))
			},
			nil,
		},
		"int": {
			id,
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			media,
		},
		"*int": {
			&id,
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			media,
		},
		"float32": {
			idFloat32,
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			media,
		},
		"*float32": {
			&idFloat32,
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			media,
		},
		"float64": {
			idFloat64,
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			media,
		},
		"*float64": {
			&idFloat64,
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			media,
		},
		"string": {
			"wrongval",
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("no media"))
			},
			nil,
		},
		"noStringer": {
			noStringer{},
			func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("no media"))
			},
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup()
			test.mock(mock)
			got := ns.Find(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
