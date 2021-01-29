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
			1,
			func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(category, nil)
			},
			category,
		},
		"Not Found": {
			1,
			func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(category, nil)
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.CategoryRepository) {
				m.On("GetById", 1).Return(category, nil)
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

func TestNamespace_ByName(t *testing.T) {

	category := domain.Category{Id: 1, Name: "cat"}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.CategoryRepository)
		want  interface{}
	}{
		"Success": {
			"cat",
			func(m *mocks.CategoryRepository) {
				m.On("GetByName", "cat").Return(category, nil)
			},
			category,
		},
		"Not Found": {
			"cat",
			func(m *mocks.CategoryRepository) {
				m.On("GetByName", "cat").Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.CategoryRepository) {
				m.On("GetByName", "cat").Return(category, nil)
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.CategoryRepository) {
				m.On("GetByName", "cat").Return(category, nil)
			},
			nil,
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
			1,
			func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
			},
			domain.Category{Id: 1, Name: "cat"},
		},
		"Not Found": {
			1,
			func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"Nil Parent": {
			1,
			func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.CategoryRepository) {
				m.On("GetParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
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
		Page:           1,
		Limit:          15,
		LimitAll:       false,
		OrderBy:        OrderBy,
		OrderDirection: OrderDirection,
	}

	tt := map[string]struct {
		input params.Query
		mock  func(m *mocks.CategoryRepository)
		want  interface{}
	}{
		"Success": {
			params.Query{"limit": 15},
			func(m *mocks.CategoryRepository) {
				m.On("Get", p).Return(categories, 2, nil)
			},
			Categories{
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
			nil,
			func(m *mocks.CategoryRepository) {
				m.On("Get", p).Return(categories, 2, nil)
			},
			Categories{
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
			params.Query{"limit": 15},
			func(m *mocks.CategoryRepository) {
				m.On("Get", p).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
			nil,
		},
		"Internal Error": {
			params.Query{"limit": 15},
			func(m *mocks.CategoryRepository) {
				m.On("Get", p).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal error"})
			},
			"internal error",
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
