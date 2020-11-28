package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getPostsMock is a helper to obtain a mock options controller
// for testing.
func getOptionsMock(m models.OptionsRepository) *OptionsController {
	return &OptionsController{
		store: &models.Store{
			Options: m,
		},
	}
}

// Test_NewOptions - Test construct
func Test_NewOptions(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &OptionsController{
		store:  &store,
		config: config,
	}
	got := newOptions(&store, config)
	assert.Equal(t, got, want)
}

// TestOptionsController_Get - Test Get route
func TestOptionsController_Get(t *testing.T) {

	options := domain.OptionsDB{
		"test1": domain.OptionDB{ID: 123, Name: "test"},
		"test2": domain.OptionDB{ID: 124, Name: "test1"},
	}

	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		mock    func(m *mocks.OptionsRepository)
	}{
		"Success": {
			want:    `{"test1":{"id":123,"option_name":"test","option_value":null},"test2":{"id":124,"option_name":"test1","option_value":null}}`,
			status:  200,
			message: "Successfully obtained options",
			mock: func(m *mocks.OptionsRepository) {
				m.On("Get").Return(options, nil)
			},
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "no options found",
			mock: func(m *mocks.OptionsRepository) {
				m.On("Get").Return(nil, &errors.Error{Code: errors.NOTFOUND, Message: "no options found"})
			},
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			mock: func(m *mocks.OptionsRepository) {
				m.On("Get").Return(nil, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			want:    `{}`,
			status:  400,
			message: "invalid",
			mock: func(m *mocks.OptionsRepository) {
				m.On("Get").Return(nil, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(m *mocks.OptionsRepository) {
				m.On("Get").Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.OptionsRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/options", "/options", nil, func(g *gin.Context) {
				getOptionsMock(mock).Get(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestOptionsController_GetByName - Test GetByName route
func TestOptionsController_GetByName(t *testing.T) {

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(m *mocks.OptionsRepository)
		url     string
	}{
		"Success": {
			want:    `"testing"`,
			status:  200,
			message: "Successfully obtained option with name: test",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetByName", "test").Return("testing", nil)
			},
			url: "/options/test",
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "no option found",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetByName", "test").Return(nil, &errors.Error{Code: errors.NOTFOUND, Message: "no option found"})
			},
			url: "/options/test",
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(m *mocks.OptionsRepository) {
				m.On("GetByName", "test").Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/options/test",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.OptionsRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", test.url, "/options/:name", nil, func(g *gin.Context) {
				getOptionsMock(mock).GetByName(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestOptionsController_UpdateCreate - Test UpdateCreate route
func TestOptionsController_UpdateCreate(t *testing.T) {

	vOptions := domain.Options{
		SiteTitle:        "test",
		SiteDescription:  "test",
		SiteLogo:         "test",
		SiteUrl:          "http://verbiscms.com",
		GeneralLocale:    "test",
		MediaCompression: 10,
	}

	vOptionsBadValidation := domain.Options{
		SiteTitle:        "test",
		SiteDescription:  "test",
		SiteLogo:         "test",
		GeneralLocale:    "test",
		MediaCompression: 10,
	}
	jsonVOptions, err := json.Marshal(vOptions)
	if err != nil {
		t.Fatal(err)
	}

	dbOptions := domain.OptionsDB{}
	err = json.Unmarshal(jsonVOptions, &dbOptions)
	assert.NoError(t, err)

	tt := map[string]struct {
		want    string
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.OptionsRepository)
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "Successfully created/updated options",
			input:   vOptions,
			mock: func(m *mocks.OptionsRepository) {
				m.On("UpdateCreate", &dbOptions).Return(nil)
			},
		},
		"Validation Failed": {
			want:    `{"errors":[{"key":"site_url","message":"Site Url is required.","type":"required"}]}`,
			status:  400,
			message: "Validation failed",
			input:   vOptionsBadValidation,
			mock: func(m *mocks.OptionsRepository) {
				m.On("UpdateCreate", &dbOptions).Return(nil)
			},
		},
		"Validation Failed DB": {
			want:    `{}`,
			status:  400,
			message: "Validation failed",
			input:   "test",
			mock: func(m *mocks.OptionsRepository) {
				m.On("UpdateCreate", &dbOptions).Return(nil)
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   vOptions,
			mock: func(m *mocks.OptionsRepository) {
				m.On("UpdateCreate", &dbOptions).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.OptionsRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/posts", "/posts", bytes.NewBuffer(body), func(g *gin.Context) {
				getOptionsMock(mock).UpdateCreate(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}
