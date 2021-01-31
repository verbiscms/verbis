package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getFormsMock is a helper to obtain a mock forms controller
// for testing.
func getFormsMock(m models.FormRepository) *Forms {
	return &Forms{
		store: &models.Store{
			Forms: m,
		},
	}
}

// Test_NewForms - Test construct
func Test_NewForms(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &Forms{
		store:  &store,
		config: config,
	}
	got := NewForms(&store, config)
	assert.Equal(t, got, want)
}

// TestForms_Get - Test Get route
func TestForms_Get(t *testing.T) {

	forms := []domain.Form{
		{Id: 123, Name: "Form"},
		{Id: 124, Name: "Form1"},
	}
	pagination := params.Params{Page: 1, Limit: 15, OrderBy: "id", OrderDirection: "ASC", Filters: nil}

	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		mock    func(m *mocks.FormRepository)
	}{
		"Success": {
			want:    `[{"created_at":null,"email_message":"","email_send":false,"email_subject":"","fields":null,"id":123,"name":"Form","store_db":false,"updated_at":null,"uuid":"00000000-0000-0000-0000-000000000000"},{"created_at":null,"email_message":"","email_send":false,"email_subject":"","fields":null,"id":124,"name":"Form1","store_db":false,"updated_at":null,"uuid":"00000000-0000-0000-0000-000000000000"}]`,
			status:  200,
			message: "Successfully obtained forms",
			mock: func(m *mocks.FormRepository) {
				m.On("Get", pagination).Return(forms, 1, nil)
			},
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "no forms found",
			mock: func(m *mocks.FormRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no forms found"})
			},
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			mock: func(m *mocks.FormRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			want:    `{}`,
			status:  400,
			message: "invalid",
			mock: func(m *mocks.FormRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(m *mocks.FormRepository) {
				m.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.FormRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/forms", "/forms", nil, func(g *gin.Context) {
				getFormsMock(mock).Get(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestForms_GetById - Test GetByID route
func TestForms_GetById(t *testing.T) {

	form := domain.Form{Id: 123, Name: "Form"}

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(m *mocks.FormRepository)
		url     string
	}{
		"Success": {
			want:    `{"created_at":null,"email_message":"","email_send":false,"email_subject":"","fields":null,"id":123,"name":"Form","store_db":false,"updated_at":null,"uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully obtained form with ID: 123",
			mock: func(m *mocks.FormRepository) {
				m.On("GetById", 123).Return(form, nil)
			},
			url: "/forms/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "Pass a valid number to obtain the form by ID",
			mock: func(m *mocks.FormRepository) {
				m.On("GetById", 123).Return(domain.Form{}, fmt.Errorf("error"))
			},
			url: "/forms/wrongid",
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "no forms found",
			mock: func(m *mocks.FormRepository) {
				m.On("GetById", 123).Return(domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: "no forms found"})
			},
			url: "/forms/123",
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(m *mocks.FormRepository) {
				m.On("GetById", 123).Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/forms/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.FormRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", test.url, "/forms/:id", nil, func(g *gin.Context) {
				getFormsMock(mock).GetById(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestForms_Create - Test Create route
func TestForms_Create(t *testing.T) {

	form := domain.Form{Id: 123, Name: "Form"}
	formBadValidation := domain.Form{Id: 123}

	tt := map[string]struct {
		want    string
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.FormRepository)
	}{
		"Success": {
			want:    `{"created_at":null,"email_message":"","email_send":false,"email_subject":"","fields":null,"id":123,"name":"Form","store_db":false,"updated_at":null,"uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully created form with ID: 123",
			input:   form,
			mock: func(m *mocks.FormRepository) {
				m.On("Create", &form).Return(form, nil)
			},
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"name","message":"Name is required.","type":"required"}]}`,
			status:  400,
			message: "Validation failed",
			input:   formBadValidation,
			mock: func(m *mocks.FormRepository) {
				m.On("Create", &formBadValidation).Return(domain.Form{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			want:    `{}`,
			status:  400,
			message: "invalid",
			input:   form,
			mock: func(m *mocks.FormRepository) {
				m.On("Create", &form).Return(domain.Form{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			input:   form,
			mock: func(m *mocks.FormRepository) {
				m.On("Create", &form).Return(domain.Form{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   form,
			mock: func(m *mocks.FormRepository) {
				m.On("Create", &form).Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.FormRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/forms", "/forms", bytes.NewBuffer(body), func(g *gin.Context) {
				getFormsMock(mock).Create(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestForms_Update - Test Update route
func TestForms_Update(t *testing.T) {

	form := domain.Form{Id: 123, Name: "Category"}
	formBadValidation := domain.Category{Id: 123}

	tt := map[string]struct {
		want    string
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.FormRepository)
		url     string
	}{
		"Success": {
			want:    `{"created_at":null,"email_message":"","email_send":false,"email_subject":"","fields":null,"id":123,"name":"Category","store_db":false,"updated_at":null,"uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully updated form with ID: 123",
			input:   form,
			mock: func(m *mocks.FormRepository) {
				m.On("Update", &form).Return(form, nil)
			},
			url: "/forms/123",
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"name","message":"Name is required.","type":"required"}]}`,
			status:  400,
			message: "Validation failed",
			input:   formBadValidation,
			mock: func(m *mocks.FormRepository) {
				m.On("Update", formBadValidation).Return(domain.Form{}, fmt.Errorf("error"))
			},
			url: "/forms/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "A valid ID is required to update the form",
			input:   form,
			mock: func(m *mocks.FormRepository) {
				m.On("Update", form).Return(domain.Form{}, fmt.Errorf("error"))
			},
			url: "/forms/wrongid",
		},
		"Not Found": {
			want:    `{}`,
			status:  400,
			message: "not found",
			input:   form,
			mock: func(m *mocks.FormRepository) {
				m.On("Update", &form).Return(domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/forms/123",
		},
		"Internal": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   form,
			mock: func(m *mocks.FormRepository) {
				m.On("Update", &form).Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/forms/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.FormRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("PUT", test.url, "/forms/:id", bytes.NewBuffer(body), func(g *gin.Context) {
				getFormsMock(mock).Update(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestForms_Delete - Test Delete route
func TestForms_Delete(t *testing.T) {

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(m *mocks.FormRepository)
		url     string
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "Successfully deleted form with ID: 123",
			mock: func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/forms/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "A valid ID is required to delete a form",
			mock: func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/forms/wrongid",
		},
		"Not Found": {
			want:    `{}`,
			status:  400,
			message: "not found",
			mock: func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/forms/123",
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			mock: func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			url: "/forms/123",
		},
		"Internal": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(m *mocks.FormRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/forms/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.FormRepository{}
			test.mock(mock)

			rr.RequestAndServe("DELETE", test.url, "/forms/:id", nil, func(g *gin.Context) {
				getFormsMock(mock).Delete(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TODO
func TestForms_Send(t *testing.T) {

}
