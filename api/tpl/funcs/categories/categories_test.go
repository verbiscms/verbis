package categories

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	vhttp "github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/tpl/params"
	"github.com/stretchr/testify/assert"
	"testing"
)

type noStringer struct{}

func Setup() (*Namespace, *mocks.CategoryRepository) {
	mock := &mocks.CategoryRepository{}
	return &Namespace{deps: &deps.Deps{
		Store: &models.Store{
			Categories: mock,
		},
	}}, mock
}

func TestNamespace_Find(t *testing.T) {

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
		"Nil": {
			input: nil,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(category, nil)
			},
			want: nil,
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

func TestNamespace_ByName(t *testing.T) {

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
		"Nil": {
			input: nil,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetByName", "cat").Return(category, nil)
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup()
			test.mock(mock)
			got := ns.ByName(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Parent(t *testing.T) {

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
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"Nil Parent": {
			input: 1,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"Nil": {
			input: nil,
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			t.Run(name, func(t *testing.T) {
				ns, mock := Setup()
				test.mock(mock)
				got := ns.Parent(test.input)
				assert.Equal(t, test.want, got)
			})
		})
	}
}

func TestNamespace_List(t *testing.T) {

	categories := []domain.Category{
		{Id: 1, Name: "cat1"},
		{Id: 1, Name: "cat2"},
	}

	p := vhttp.Params{
		Page: 1,
		Limit: 15,
		LimitAll: false,
		OrderBy: OrderBy,
		OrderDirection: OrderDirection,
	}

	tt := map[string]struct {
		input params.Query
		mock  func(m *mocks.CategoryRepository)
		want  interface{}
	}{
		"Success": {
			input: params.Query{"limit": 15},
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", p).Return(categories, 2, nil)
			},
			want: Categories{
				Categories: categories,
				Pagination: &vhttp.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 2,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Nil": {
			input: nil,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", p).Return(categories, 2, nil)
			},
			want: Categories{
				Categories: categories,
				Pagination: &vhttp.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 2,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Not Found": {
			input: params.Query{"limit": 15},
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", p).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
			want: nil,
		},
		"Internal Error": {
			input: params.Query{"limit": 15},
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", p).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal error"})
			},
			want: "internal error",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup()
			test.mock(mock)
			got, err := ns.List(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}