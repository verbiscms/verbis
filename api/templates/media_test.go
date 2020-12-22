package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"testing"
)

func Test_GetMedia(t *testing.T) {

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
			input: 1,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			want: media,
		},
		"No Item": {
			input: 1,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("no media"))
			},
			want: nil,
		},
		"nil": {
			input: nil,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", nil).Return(domain.Media{}, fmt.Errorf("no media"))
			},
			want: nil,
		},
		"int": {
			input: id,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			want: media,
		},
		"*int": {
			input: &id,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			want: media,
		},
		"float32": {
			input: idFloat32,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			want: media,
		},
		"*float32": {
			input: &idFloat32,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			want: media,
		},
		"float64": {
			input: idFloat64,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			want: media,
		},
		"*float64": {
			input: &idFloat64,
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(media, nil)
			},
			want: media,
		},
		"string": {
			input: "wrongval",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(nil, fmt.Errorf("no media"))
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			mock := mocks.MediaRepository{}

			test.mock(&mock)
			f.store.Media = &mock

			tpl := `{{ media .MediaId }}`
			runtv(t, f, tpl, test.want, map[string]interface{}{"MediaId": test.input})
		})
	}
}
