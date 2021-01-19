package tpl

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	vhttp "github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *TplTestSuite) Test_GetCategory() {

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
			want: "",
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(category, nil)
			},
			want: "",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			categoryMock := mocks.CategoryRepository{}

			test.mock(&categoryMock)
			t.store.Categories = &categoryMock

			tpl := `{{ category . }}`

			t.RunTWithData(tpl, test.want, test.input)
		})
	}
}

func (t *TplTestSuite) Test_GetCategoryByName() {

	category := domain.Category{Id: 1, Name: "cat"}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.CategoryRepository)
		want  interface{}
	}{
		"Success": {
			input: "cat",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetByName", "cat").Return(category, nil)
			},
			want: category,
		},
		"Not Found": {
			input: "cat",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetByName", "cat").Return(domain.Category{}, fmt.Errorf("error"))
			},
			want: "",
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetByName", "cat").Return(category, nil)
			},
			want: "",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			categoryMock := mocks.CategoryRepository{}

			test.mock(&categoryMock)
			t.store.Categories = &categoryMock

			tpl := `{{ categoryByName . }}`
			t.RunTWithData(tpl, test.want, test.input)
		})
	}
}

func (t *TplTestSuite) Test_GetParentCategory() {

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.CategoryRepository)
		want  interface{}
	}{
		"Success": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
			},
			want: domain.Category{Id: 1, Name: "cat"},
		},
		"Not Found": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error")).Once()
			},
			want: "",
		},
		"Nil Parent": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetParent", 2).Return(domain.Category{}, fmt.Errorf("error"))
			},
			want: "",
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			want: "",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			categoryMock := mocks.CategoryRepository{}

			test.mock(&categoryMock)
			t.store.Categories = &categoryMock

			tpl := `{{ categoryParent . }}`
			t.RunTWithData(tpl, test.want, test.input)
		})
	}
}

func (t *TplTestSuite) Test_GetCategories() {

	categories := []domain.Category{
		{Id: 1, Name: "cat1"},
		{Id: 1, Name: "cat2"},
	}

	tt := map[string]struct {
		input map[string]interface{}
		mock  func(m *mocks.CategoryRepository)
		want  interface{}
	}{
		"Success": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}).Return(categories, 2, nil)
			},
			want: map[string]interface{}{
				"Categories": categories,
				"Pagination": &vhttp.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 2,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Failed Params": {
			input: map[string]interface{}{"limit": "wrongval"},
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}).Return(categories, 2, nil)
			},
			want: "cannot unmarshal string into Go struct field TemplateParams.limit",
		},
		"Not Found": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
			want: map[string]interface{}(nil),
		},
		"Internal Error": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal error"})
			},
			want: "internal error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			categoryMock := mocks.CategoryRepository{}

			test.mock(&categoryMock)
			t.store.Categories = &categoryMock

			c, err := t.getCategories(test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.EqualValues(test.want, c)
		})
	}
}
