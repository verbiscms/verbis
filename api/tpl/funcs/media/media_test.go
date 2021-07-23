// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/media"
	"github.com/verbiscms/verbis/api/store"
	"testing"
)

type noStringer struct{}

func Setup() (*Namespace, *mocks.Repository) {
	mock := &mocks.Repository{}
	return &Namespace{deps: &deps.Deps{
		Store: &store.Repository{
			Media: mock,
		},
	}}, mock
}

func TestNamespace_Find(t *testing.T) {
	media := domain.Media{
		Id: 1,
		File: domain.File{
			Url: "/uploads/test.jpg",
		},
	}

	mediaPub := domain.MediaPublic{
		Id:  1,
		Url: "/uploads/test.jpg",
	}

	id := 1
	idFloat32 := float32(1)
	idFloat64 := float64(1)

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			1,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(media, nil)
			},
			mediaPub,
		},
		"No Item": {
			1,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Media{}, fmt.Errorf("no media"))
			},
			nil,
		},
		"nil": {
			nil,
			func(m *mocks.Repository) {
				m.On("Find", nil).Return(domain.Media{}, fmt.Errorf("no media"))
			},
			nil,
		},
		"int": {
			id,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(media, nil)
			},
			mediaPub,
		},
		"*int": {
			&id,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(media, nil)
			},
			mediaPub,
		},
		"float32": {
			idFloat32,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(media, nil)
			},
			mediaPub,
		},
		"*float32": {
			&idFloat32,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(media, nil)
			},
			mediaPub,
		},
		"float64": {
			idFloat64,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(media, nil)
			},
			mediaPub,
		},
		"*float64": {
			&idFloat64,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(media, nil)
			},
			mediaPub,
		},
		"string": {
			"wrongval",
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Media{}, fmt.Errorf("no media"))
			},
			nil,
		},
		"noStringer": {
			noStringer{},
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Media{}, fmt.Errorf("no media"))
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
