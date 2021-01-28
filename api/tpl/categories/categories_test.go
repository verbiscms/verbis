package categories

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

func Test_GetCategory(t *testing.T) {

	category := domain.Category{Id: 1, Name: "cat"}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.CategoryRepository)
		want  interface{}
	}{
		"Success": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(category, nil)
			},
			want: category,
		},
		"Not Found": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(category, nil)
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			categoryMock := mocks.CategoryRepository{}
			test.mock(&categoryMock)

			ns := Namespace{deps: &deps.Deps{
				Store: &models.Store{
					Categories: &categoryMock,
				},
			}}

			got := ns.getCategory(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
