package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	vhttp "github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/stretchr/testify/assert"

	"testing"
)

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
			f := newTestSuite()
			categoryMock := mocks.CategoryRepository{}

			test.mock(&categoryMock)
			f.store.Categories = &categoryMock

			tpl := `{{ category . }}`

			runtv(t, f, tpl, test.want, test.input)
		})
	}
}

func Test_GetCategoryByName(t *testing.T) {

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
			want: nil,
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetByName", "cat").Return(category, nil)
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			categoryMock := mocks.CategoryRepository{}

			test.mock(&categoryMock)
			f.store.Categories = &categoryMock

			tpl := `{{ categoryByName . }}`

			runtv(t, f, tpl, test.want, test.input)
		})
	}
}

func Test_GetParentCategory(t *testing.T) {

	p := 2
	category := domain.Category{Id: 1, Name: "cat", ParentId: &p}
	noParentCategory := domain.Category{Id: 1, Name: "cat", ParentId: nil}
	parentCategory := domain.Category{Id: 2, Name: "parent"}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.CategoryRepository)
		want  interface{}
	}{
		"Success": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(category, nil).Once()
				m.On("GetById", 2).Return(parentCategory, nil)
			},
			want: parentCategory,
		},
		"Not Found": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(domain.Category{}, fmt.Errorf("error")).Once()
			},
			want: nil,
		},
		"No Parent": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(noParentCategory, nil).Once()
			},
			want: nil,
		},
		"Nil Parent": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(category, nil).Once()
				m.On("GetById", 2).Return(domain.Category{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(domain.Category{}, fmt.Errorf("error")).Once()
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			categoryMock := mocks.CategoryRepository{}

			test.mock(&categoryMock)
			f.store.Categories = &categoryMock

			tpl := `{{ categoryByParent . }}`

			runtv(t, f, tpl, test.want, test.input)
		})
	}
}

func Test_GetCategories(t *testing.T) {

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
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			categoryMock := mocks.CategoryRepository{}

			test.mock(&categoryMock)
			f.store.Categories = &categoryMock

			c, err := f.getCategories(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.EqualValues(t, test.want, c)
		})
	}
}
