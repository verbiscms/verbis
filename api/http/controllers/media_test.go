package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
)

// getMediaMock is a helper to obtain a mock user controller
// for testing.
func getMediaMock(m models.MediaRepository) *MediaController {
	return &MediaController{
		store: &models.Store{
			Media: m,
		},
	}
}

// Test_NewMedia - Test construct
func Test_NewMedia(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &MediaController{
		store:  &store,
		config: config,
	}
	got := newMedia(&store, config)
	assert.Equal(t, got, want)
}

// TestMediaController_Get - Test Get route
func TestMediaController_Get(t *testing.T) {

	media := []domain.Media{
		{Id:          123, Url:         "/logo.svg"},
		{Id:          124, Url:         "/logo.png"},
	}
	pagination := http.Params{Page: 1, Limit: 15, OrderBy: "id", OrderDirection: "asc", Filters: nil}

	tt := map[string]struct {
		name       string
		want       string
		status     int
		message    string
		mock func(u *mocks.MediaRepository)
	}{
		"Success": {
			want:       `[{"id":123,"uuid":"00000000-0000-0000-0000-000000000000","url":"/logo.svg","title":null,"alt":null,"description":null,"file_size":0,"file_name":"","sizes":null,"type":"","user_id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},{"id":124,"uuid":"00000000-0000-0000-0000-000000000000","url":"/logo.png","title":null,"alt":null,"description":null,"file_size":0,"file_name":"","sizes":null,"type":"","user_id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`,
			status:     200,
			message:    "Successfully obtained media",
			mock: func(u *mocks.MediaRepository) {
				u.On("Get", pagination).Return(media, 1, nil)
			},
		},
		"Not Found": {
			want:       `{}`,
			status:     200,
			message:    "no media found",
			mock: func(u *mocks.MediaRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no media found"})
			},
		},
		"Conflict": {
			want:       `{}`,
			status:     400,
			message:    "conflict",
			mock: func(u *mocks.MediaRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			want:       `{}`,
			status:     400,
			message:    "invalid",
			mock: func(u *mocks.MediaRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(u *mocks.MediaRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.MediaRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/users", "/users", nil, func(g *gin.Context) {
				getMediaMock(mock).Get(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestMediaController_GetById - Test GetByID route
func TestMediaController_GetById(t *testing.T) {

	media := domain.Media{
		Id:          123,
		Url:         "/logo.svg",
	}

	tt := map[string]struct {
		want       string
		status     int
		message    string
		mock func(m *mocks.MediaRepository)
		url string
	}{
		"Success": {
			want:       `{"alt":null,"created_at":"0001-01-01T00:00:00Z","description":null,"file_name":"","file_size":0,"id":123,"sizes":null,"title":null,"type":"","updated_at":"0001-01-01T00:00:00Z","url":"/logo.svg","user_id":0,"uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully obtained media item with ID: 123",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 123).Return(media, nil)
			},
			url: "/media/123",
		},
		"Invalid ID": {
			want:       `{}`,
			status:     400,
			message:    "Pass a valid number to obtain the media item by ID",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 123).Return(domain.Media{}, fmt.Errorf("error"))
			},
			url: "/media/wrongid",
		},
		"Not Found": {
			want:       `{}`,
			status:     200,
			message:    "no media items found",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 123).Return(domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: "no media items found"})
			},
			url: "/media/123",
		},
		"Internal Error": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 123).Return(domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/media/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.MediaRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", test.url, "/media/:id", nil, func(g *gin.Context) {
				getMediaMock(mock).GetById(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestMediaController_Update - Test Update route
func TestMediaController_Update(t *testing.T) {

	media := domain.Media{
		Id:          123,
		Url:         "/logo.svg",
	}

	mediaBadValidation := &domain.Media{}

	tt := map[string]struct {
		want       string
		status     int
		message    string
		input interface{}
		mock func(u *mocks.MediaRepository)
		url string
	}{
		"Success": {
			want:       `{"alt":null,"created_at":"0001-01-01T00:00:00Z","description":null,"file_name":"","file_size":0,"id":123,"sizes":null,"title":null,"type":"","updated_at":"0001-01-01T00:00:00Z","url":"/logo.svg","user_id":0,"uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully updated media item with ID: 123",
			input: media,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", &media).Return(nil)
			},
			url: "/media/123",
		},
		"Validation Failed": {
			want:       `{}`,
			status:     400,
			message:    "Validation failed",
			input: `{"id": "wrongid"}`,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", mediaBadValidation).Return(fmt.Errorf("error"))
			},
			url: "/media/123",
		},
		"Invalid ID": {
			want:       `{}`,
			status:     400,
			message:    "A valid ID is required to update the media item",
			input: media,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", media).Return(fmt.Errorf("error"))
			},
			url: "/media/wrongid",
		},
		"Not Found": {
			want:       `{}`,
			status:     400,
			message:    "not found",
			input: media,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", &media).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/media/123",
		},
		"Internal": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			input: media,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", &media).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/media/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.MediaRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("PUT", test.url, "/media/:id", bytes.NewBuffer(body), func(g *gin.Context) {
				getMediaMock(mock).Update(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}


// TestMediaController_Upload - Test Upload route
func TestMediaController_Upload(t *testing.T) {

	rr := newTestSuite(t)

	path := "/Users/ainsley/Desktop/Reddico/apis/verbis/api/test/testdata/images/gopher.svg"

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	part, err := w.CreateFormFile("file", filepath.Base(path))
	assert.NoError(t, err)

	file, err := os.Open(path)
	assert.NoError(t, err)
	defer file.Close()

	_, err = io.Copy(part, file)
	assert.NoError(t, err)
	assert.NoError(t, w.Close())



	m := &mocks.MediaRepository{}
	m.On("Upload", body).Return(domain.Media{}, nil)
	m.On("Validate", mock.Anything).Return(nil)

	u := &mocks.UserRepository{}
	u.On("CheckToken", mock.Anything).Return(domain.User{}, nil)

	rr.NewRequest("POST", "/media", body)
	rr.gin.Request.Header.Set("Content-Type", w.FormDataContentType())


	mc := &MediaController{
		store: &models.Store{
			Media: m,
			User: u,
		},
	}




	fmt.Println(rr.Body())
	fmt.Println(rr.recorder.Body.String())


}

// TestMediaController_Delete - Test Delete route
func TestMediaController_Delete(t *testing.T) {

	tt := map[string]struct {
		want       string
		status     int
		message    string
		mock func(u *mocks.MediaRepository)
		url string
	}{
		"Success": {
			want:       `{}`,
			status:     200,
			message:    "Successfully deleted media item with ID: 123",
			mock: func(u *mocks.MediaRepository) {
				u.On("Delete", 123).Return(nil)
			},
			url: "/media/123",
		},
		"Invalid ID": {
			want:       `{}`,
			status:     400,
			message:    "A valid ID is required to delete a media item",
			mock:  func(m *mocks.MediaRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/media/wrongid",
		},
		"Not Found": {
			want:       `{}`,
			status:     400,
			message:    "not found",
			mock:  func(m *mocks.MediaRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/media/123",
		},
		"Conflict": {
			want:       `{}`,
			status:     400,
			message:    "conflict",
			mock:  func(m *mocks.MediaRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			url: "/media/123",
		},
		"Internal": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock:  func(m *mocks.MediaRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/media/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.MediaRepository{}
			test.mock(mock)

			rr.RequestAndServe("DELETE", test.url, "/media/:id", nil, func(g *gin.Context) {
				getMediaMock(mock).Delete(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}