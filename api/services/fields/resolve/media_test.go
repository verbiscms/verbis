// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/media"
)

func (t *ResolverTestSuite) TestValue_Media() {
	tt := map[string]struct {
		value domain.FieldValue
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Media": {
			value: domain.FieldValue("1"),
			mock: func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Media{Title: "title"}, nil)
			},
			want: domain.MediaPublic{Title: "title"},
		},
		"Media Error": {
			value: domain.FieldValue("1"),
			mock: func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Media{}, fmt.Errorf("not found"))
			},
			want: "not found",
		},
		"Cast Error": {
			value: domain.FieldValue("wrongval"),
			mock:  func(m *mocks.Repository) {},
			want:  `strconv.Atoi: parsing "wrongval": invalid syntax`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			mediaMock := &mocks.Repository{}

			test.mock(mediaMock)
			v.deps.Store.Media = mediaMock

			got, err := v.media(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *ResolverTestSuite) TestValue_MediaResolve() {
	tt := map[string]struct {
		field domain.PostField
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			field: domain.PostField{OriginalValue: "1,2,3", Type: "image"},
			mock: func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Media{Title: "title1"}, nil)
				m.On("Find", 2).Return(domain.Media{Title: "title2"}, nil)
				m.On("Find", 3).Return(domain.Media{Title: "title3"}, nil)
			},
			want: domain.PostField{OriginalValue: "1,2,3", Type: "image", Value: []interface{}{
				domain.MediaPublic{Title: "title1"},
				domain.MediaPublic{Title: "title2"},
				domain.MediaPublic{Title: "title3"},
			}},
		},
		"Trailing Comma": {
			field: domain.PostField{OriginalValue: "1,2,3,", Type: "image"},
			mock: func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Media{Title: "title1"}, nil)
				m.On("Find", 2).Return(domain.Media{Title: "title2"}, nil)
				m.On("Find", 3).Return(domain.Media{Title: "title3"}, nil)
			},
			want: domain.PostField{OriginalValue: "1,2,3,", Type: "image", Value: []interface{}{
				domain.MediaPublic{Title: "title1"},
				domain.MediaPublic{Title: "title2"},
				domain.MediaPublic{Title: "title3"},
			}},
		},
		"Leading Comma": {
			field: domain.PostField{OriginalValue: ",1,2,3", Type: "image"},
			mock: func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Media{Title: "title1"}, nil)
				m.On("Find", 2).Return(domain.Media{Title: "title2"}, nil)
				m.On("Find", 3).Return(domain.Media{Title: "title3"}, nil)
			},
			want: domain.PostField{OriginalValue: ",1,2,3", Type: "image", Value: []interface{}{
				domain.MediaPublic{Title: "title1"},
				domain.MediaPublic{Title: "title2"},
				domain.MediaPublic{Title: "title3"},
			}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			mediaMock := &mocks.Repository{}

			test.mock(mediaMock)
			v.deps.Store.Media = mediaMock

			got := v.resolve(test.field)

			t.Equal(test.want, got)
		})
	}
}
