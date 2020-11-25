package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getCategoriesMock is a helper to obtain a mock categories controller
// for testing.
func getCategoriesMock(m models.CategoryRepository) *CategoriesController {
	return &CategoriesController{
		store: &models.Store{
			Categories: m,
		},
	}
}

// Test_NewCategories - Test construct
func Test_NewCategories(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &CategoriesController{
		store:  &store,
		config: config,
	}
	got := newCategories(&store, config)
	assert.Equal(t, got, want)
}

// TestCategoriesController_Get - Test Get route
func TestCategoriesController_Get(t *testing.T) {

	categories := []domain.Category{
		{Id:          123, Slug:        "/cat", Name:        "Category"},
		{Id:          124, Slug:        "/cat1", Name:        "Category1"},
	}
	pagination := http.Params{Page: 1, Limit: 15, OrderBy: "id", OrderDirection: "asc", Filters: nil}

	tt := map[string]struct {
		name       string
		want       string
		status     int
		message    string
		mock func(m *mocks.CategoryRepository)
	}{
		"Success": {
			want:       `[{"archive_id":null,"created_at":"0001-01-01T00:00:00Z","description":null,"id":123,"name":"Category","parent_id":null,"resource":"","slug":"/cat","updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"},{"archive_id":null,"created_at":"0001-01-01T00:00:00Z","description":null,"id":124,"name":"Category1","parent_id":null,"resource":"","slug":"/cat1","updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}]`,
			status:     200,
			message:    "Successfully obtained categories",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(categories, 1, nil)
			},
		},
		"Not Found": {
			want:       `{}`,
			status:     200,
			message:    "no categories found",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
		},
		"Conflict": {
			want:       `{}`,
			status:     400,
			message:    "conflict",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			want:       `{}`,
			status:     400,
			message:    "invalid",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/categories", "/categories", nil, func(g *gin.Context) {
				getCategoriesMock(mock).Get(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestCategoriesController_GetById - Test GetByID route
func TestCategoriesController_GetById(t *testing.T) {

	category := domain.Category{Id:          123, Slug:        "/cat", Name:        "Category"}

	tt := map[string]struct {
		want       string
		status     int
		message    string
		mock func(m *mocks.CategoryRepository)
		url string
	}{
		"Success": {
			want:       `{"archive_id":null,"created_at":"0001-01-01T00:00:00Z","description":null,"id":123,"name":"Category","parent_id":null,"resource":"","slug":"/cat","updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully obtained category with ID: 123",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 123).Return(category, nil)
			},
			url: "/categories/123",
		},
		"Invalid ID": {
			want:       `{}`,
			status:     400,
			message:    "Pass a valid number to obtain the category by ID",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 123).Return(domain.Category{}, fmt.Errorf("error"))
			},
			url: "/categories/wrongid",
		},
		"Not Found": {
			want:       `{}`,
			status:     200,
			message:    "no categories found",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 123).Return(domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
			url: "/categories/123",
		},
		"Internal Error": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(m *mocks.CategoryRepository) {
				m.On("GetById", 123).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/categories/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", test.url, "/categories/:id", nil, func(g *gin.Context) {
				getCategoriesMock(mock).GetById(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestCategoriesController_Create - Test Create route
func TestCategoriesController_Create(t *testing.T) {

	category := domain.Category{Id:          123, Slug:        "/cat", Name:        "Category", Resource: "test"}
	categoryBadValidation := domain.Category{Id:          123,Name:        "Category", Resource: "test"}

	tt := map[string]struct {
		want       string
		status     int
		message    string
		input interface{}
		mock func(m *mocks.CategoryRepository)
	}{
		"Success": {
			want:       `{"archive_id":null,"created_at":"0001-01-01T00:00:00Z","description":null,"id":123,"name":"Category","parent_id":null,"resource":"test","slug":"/cat","updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully created category with ID: 123",
			input: category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(category, nil)
			},
		},
		"Validation Failed": {
			want:       `{"errors":[{"key":"slug","message":"Slug is required.","type":"required"}]}`,
			status:     400,
			message:    "Validation failed",
			input: categoryBadValidation,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			want:       `{}`,
			status:     400,
			message:    "invalid",
			input: category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			want:       `{}`,
			status:     400,
			message:    "conflict",
			input: category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			input: category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Create", &category).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/categories", "/categories", bytes.NewBuffer(body), func(g *gin.Context) {
				getCategoriesMock(mock).Create(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestCategoriesController_Update - Test Update route
func TestCategoriesController_Update(t *testing.T) {

	category := domain.Category{Id:          123, Slug:        "/cat", Name:        "Category", Resource: "test"}
	categoryBadValidation := domain.Category{Id:          123,Name:        "Category", Resource: "test"}

	tt := map[string]struct {
		want       string
		status     int
		message    string
		input interface{}
		mock func(m *mocks.CategoryRepository)
		url string
	}{
		"Success": {
			want:       `{"archive_id":null,"created_at":"0001-01-01T00:00:00Z","description":null,"id":123,"name":"Category","parent_id":null,"resource":"test","slug":"/cat","updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully updated category with ID: 123",
			input: category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", &category).Return(category, nil)
			},
			url: "/categories/123",
		},
		"Validation Failed": {
			want:       `{"errors":[{"key":"slug","message":"Slug is required.","type":"required"}]}`,
			status:     400,
			message:    "Validation failed",
			input: categoryBadValidation,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
			url: "/categories/123",
		},
		"Invalid ID": {
			want:       `{}`,
			status:     400,
			message:    "A valid ID is required to update the category",
			input: category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", categoryBadValidation).Return(domain.Category{}, fmt.Errorf("error"))
			},
			url: "/categories/wrongid",
		},
		"Not Found": {
			want:       `{}`,
			status:     400,
			message:    "not found",
			input: category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", &category).Return(domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/categories/123",
		},
		"Internal": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			input: category,
			mock: func(m *mocks.CategoryRepository) {
				m.On("Update", &category).Return(domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/categories/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("PUT", test.url, "/categories/:id", bytes.NewBuffer(body), func(g *gin.Context) {
				getCategoriesMock(mock).Update(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestCategoriesController_Delete - Test Delete route
func TestCategoriesController_Delete(t *testing.T) {

	tt := map[string]struct {
		want       string
		status     int
		message    string
		mock func(m *mocks.CategoryRepository)
		url string
	}{
		"Success": {
			want:       `{}`,
			status:     200,
			message:    "Successfully deleted category with ID: 123",
			mock: func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/categories/123",
		},
		"Invalid ID": {
			want:       `{}`,
			status:     400,
			message:    "A valid ID is required to delete a category",
			mock:  func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/categories/wrongid",
		},
		"Not Found": {
			want:       `{}`,
			status:     400,
			message:    "not found",
			mock:  func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/categories/123",
		},
		"Conflict": {
			want:       `{}`,
			status:     400,
			message:    "conflict",
			mock:  func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			url: "/categories/123",
		},
		"Internal": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock:  func(m *mocks.CategoryRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/categories/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.CategoryRepository{}
			test.mock(mock)

			rr.RequestAndServe("DELETE", test.url, "/categories/:id", nil, func(g *gin.Context) {
				getCategoriesMock(mock).Delete(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}