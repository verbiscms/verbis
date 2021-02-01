package resolve

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *ResolverTestSuite) TestValue_Category() {

	tt := map[string]struct {
		value domain.FieldValue
		mock  func(c *mocks.CategoryRepository)
		want  interface{}
	}{
		"Category": {
			value: domain.FieldValue("1"),
			mock: func(c *mocks.CategoryRepository) {
				c.On("GetById", 1).Return(domain.Category{Name: "cat"}, nil)
			},
			want: domain.Category{Name: "cat"},
		},
		"Category Error": {
			value: domain.FieldValue("1"),
			mock: func(c *mocks.CategoryRepository) {
				c.On("GetById", 1).Return(domain.Category{}, fmt.Errorf("not found"))
			},
			want: "not found",
		},
		"Cast Error": {
			value: domain.FieldValue("wrongval"),
			mock:  func(c *mocks.CategoryRepository) {},
			want:  `strconv.Atoi: parsing "wrongval": invalid syntax`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			categoryMock := &mocks.CategoryRepository{}

			test.mock(categoryMock)
			v.deps.Store.Categories = categoryMock

			got, err := v.category(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *ResolverTestSuite) TestValue_CategoryResolve() {

	tt := map[string]struct {
		field domain.PostField
		mock  func(c *mocks.CategoryRepository)
		want  domain.PostField
	}{
		"Success": {
			field: domain.PostField{OriginalValue: "1,2,3", Type: "category"},
			mock: func(c *mocks.CategoryRepository) {
				c.On("GetById", 1).Return(domain.Category{Name: "cat1"}, nil)
				c.On("GetById", 2).Return(domain.Category{Name: "cat2"}, nil)
				c.On("GetById", 3).Return(domain.Category{Name: "cat3"}, nil)
			},
			want: domain.PostField{OriginalValue: "1,2,3", Type: "category", Value: []interface{}{
				domain.Category{Name: "cat1"},
				domain.Category{Name: "cat2"},
				domain.Category{Name: "cat3"},
			}},
		},
		"Trailing Comma": {
			field: domain.PostField{OriginalValue: "1,2,3,", Type: "category"},
			mock: func(c *mocks.CategoryRepository) {
				c.On("GetById", 1).Return(domain.Category{Name: "cat1"}, nil)
				c.On("GetById", 2).Return(domain.Category{Name: "cat2"}, nil)
				c.On("GetById", 3).Return(domain.Category{Name: "cat3"}, nil)
			},
			want: domain.PostField{OriginalValue: "1,2,3,", Type: "category", Value: []interface{}{
				domain.Category{Name: "cat1"},
				domain.Category{Name: "cat2"},
				domain.Category{Name: "cat3"},
			}},
		},
		"Leading Comma": {
			field: domain.PostField{OriginalValue: ",1,2,3", Type: "category"},
			mock: func(c *mocks.CategoryRepository) {
				c.On("GetById", 1).Return(domain.Category{Name: "cat1"}, nil)
				c.On("GetById", 2).Return(domain.Category{Name: "cat2"}, nil)
				c.On("GetById", 3).Return(domain.Category{Name: "cat3"}, nil)
			},
			want: domain.PostField{OriginalValue: ",1,2,3", Type: "category", Value: []interface{}{
				domain.Category{Name: "cat1"},
				domain.Category{Name: "cat2"},
				domain.Category{Name: "cat3"},
			}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			categoryMock := &mocks.CategoryRepository{}

			test.mock(categoryMock)
			v.deps.Store.Categories = categoryMock

			got := v.resolve(test.field)

			t.Equal(test.want, got)
		})
	}
}
