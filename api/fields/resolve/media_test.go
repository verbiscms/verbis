package resolve

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *ResolverTestSuite) TestFieldValue_Media() {

	tt := map[string]struct {
		value  domain.FieldValue
		mock   func(m *mocks.MediaRepository)
		want   interface{}
	}{
		"Media": {
			value: domain.FieldValue("1"),
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{Url: "image"}, nil)
			},
			want: domain.Media{Url: "image"},
		},
		"Media Error": {
			value: domain.FieldValue("1"),
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("not found"))
			},
			want: "not found",
		},
		"Cast Error": {
			value: domain.FieldValue("wrongval"),
			mock: func(m *mocks.MediaRepository) {},
			want: `strconv.Atoi: parsing "wrongval": invalid syntax`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			mediaMock := &mocks.MediaRepository{}

			test.mock(mediaMock)
			v.store.Media = mediaMock

			got, err := v.media(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}