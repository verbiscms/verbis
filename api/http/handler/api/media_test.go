package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	gohttp "net/http"
	"os"
	"path/filepath"
	"testing"
)

// getMediaMock is a helper to obtain a mock user controller
// for testing.
func getMediaMock(m models.MediaRepository) *Media {
	return &Media{
		store: &models.Store{
			Media: m,
		},
	}
}

// Test_NewMedia - Test construct
func Test_NewMedia(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &Media{
		store:  &store,
		config: config,
	}
	got := NewMedia(&store, config)
	assert.Equal(t, got, want)
}

// TestMedia_Get - Test Get route
func TestMedia_Get(t *testing.T) {

	media := []domain.Media{
		{Id: 123, Url: "/logo.svg"},
		{Id: 124, Url: "/logo.png"},
	}
	pagination := params.Params{Page: 1, Limit: 15, OrderBy: "id", OrderDirection: "ASC", Filters: nil}

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(u *mocks.MediaRepository)
	}{
		"Success": {
			want:    `[{"alt":"","created_at":"0001-01-01T00:00:00Z","description":"","file_name":"","file_size":0,"id":123,"sizes":null,"title":"","type":"","updated_at":"0001-01-01T00:00:00Z","url":"/logo.svg","user_id":0,"uuid":"00000000-0000-0000-0000-000000000000"},{"alt":"","created_at":"0001-01-01T00:00:00Z","description":"","file_name":"","file_size":0,"id":124,"sizes":null,"title":"","type":"","updated_at":"0001-01-01T00:00:00Z","url":"/logo.png","user_id":0,"uuid":"00000000-0000-0000-0000-000000000000"}]`,
			status:  200,
			message: "Successfully obtained media",
			mock: func(u *mocks.MediaRepository) {
				u.On("Get", pagination).Return(media, 1, nil)
			},
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "no media found",
			mock: func(u *mocks.MediaRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no media found"})
			},
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			mock: func(u *mocks.MediaRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			want:    `{}`,
			status:  400,
			message: "invalid",
			mock: func(u *mocks.MediaRepository) {
				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
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

// TestMedia_GetById - Test GetByID route
func TestMedia_GetById(t *testing.T) {

	media := domain.Media{
		Id:  123,
		Url: "/logo.svg",
	}

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(m *mocks.MediaRepository)
		url     string
	}{
		"Success": {
			want:    `{"alt":"","created_at":"0001-01-01T00:00:00Z","description":"","file_name":"","file_size":0,"id":123,"sizes":null,"title":"","type":"","updated_at":"0001-01-01T00:00:00Z","url":"/logo.svg","user_id":0,"uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully obtained media item with ID: 123",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 123).Return(media, nil)
			},
			url: "/media/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "Pass a valid number to obtain the media item by ID",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 123).Return(domain.Media{}, fmt.Errorf("error"))
			},
			url: "/media/wrongid",
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "no media items found",
			mock: func(m *mocks.MediaRepository) {
				m.On("GetById", 123).Return(domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: "no media items found"})
			},
			url: "/media/123",
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
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

// TestMedia_Update - Test Update route
func TestMedia_Update(t *testing.T) {

	media := domain.Media{
		Id:  123,
		Url: "/logo.svg",
	}

	mediaBadValidation := &domain.Media{}

	tt := map[string]struct {
		want    string
		status  int
		message string
		input   interface{}
		mock    func(u *mocks.MediaRepository)
		url     string
	}{
		"Success": {
			want:    `{"alt":"","created_at":"0001-01-01T00:00:00Z","description":"","file_name":"","file_size":0,"id":123,"sizes":null,"title":"","type":"","updated_at":"0001-01-01T00:00:00Z","url":"/logo.svg","user_id":0,"uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully updated media item with ID: 123",
			input:   media,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", &media).Return(nil)
			},
			url: "/media/123",
		},
		"Validation Failed": {
			want:    `{}`,
			status:  400,
			message: "Validation failed",
			input:   `{"id": "wrongid"}`,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", mediaBadValidation).Return(fmt.Errorf("error"))
			},
			url: "/media/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "A valid ID is required to update the media item",
			input:   media,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", media).Return(fmt.Errorf("error"))
			},
			url: "/media/wrongid",
		},
		"Not Found": {
			want:    `{}`,
			status:  400,
			message: "not found",
			input:   media,
			mock: func(m *mocks.MediaRepository) {
				m.On("Update", &media).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/media/123",
		},
		"Internal": {
			want:    `{}`,
			status:  500,
			message: "internal",
			input:   media,
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

// TestMedia_Upload - Test Upload route
func TestMedia_Upload(t *testing.T) {

	wd, err := os.Getwd()
	assert.NoError(t, err)
	apiPath := filepath.Join(filepath.Dir(wd), "../..")
	path := apiPath + "/test/testdata/spa/images/gopher.svg"

	tt := map[string]struct {
		want    string
		status  int
		message string
		files   int
		mock    func(u *mocks.MediaRepository, multi []multipart.FileHeader)
		url     string
	}{
		"Success": {
			want:    `{"alt":"","created_at":"0001-01-01T00:00:00Z","description":"","file_name":"","file_size":0,"id":0,"sizes":null,"title":"","type":"","updated_at":"0001-01-01T00:00:00Z","url":"","user_id":0,"uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully uploaded media item",
			files:   1,
			mock: func(m *mocks.MediaRepository, multi []multipart.FileHeader) {
				m.On("Upload", &multi[0], "").Return(domain.Media{}, nil)
				m.On("Validate", &multi[0]).Return(nil)
			},
			url: "/media",
		},
		"No Form": {
			want:    `{}`,
			status:  400,
			message: "No files attached to the upload",
			files:   0,
			mock: func(m *mocks.MediaRepository, multi []multipart.FileHeader) {
				m.On("Upload", multipart.FileHeader{}, "").Return(domain.Media{}, nil)
				m.On("Validate", multipart.FileHeader{}).Return(nil)
			},
			url: "/media",
		},
		"No Files": {
			want:    `{}`,
			status:  400,
			message: "Attach a file to the request to be uploaded",
			files:   0,
			mock: func(m *mocks.MediaRepository, multi []multipart.FileHeader) {
				m.On("Upload", multipart.FileHeader{}, "").Return(domain.Media{}, nil)
				m.On("Validate", multipart.FileHeader{}).Return(nil)
			},
			url: "/media",
		},
		"Too Many Files": {
			want:    `{}`,
			status:  400,
			message: "Files are only permitted to be uploaded one at a time",
			files:   5,
			mock: func(m *mocks.MediaRepository, multi []multipart.FileHeader) {
				m.On("Upload", &multi[0], "").Return(domain.Media{}, nil)
				m.On("Validate", &multi[0]).Return(nil)
			},
			url: "/media",
		},
		"Invalid": {
			want:    `{}`,
			status:  415,
			message: "invalid",
			files:   1,
			mock: func(m *mocks.MediaRepository, multi []multipart.FileHeader) {
				m.On("Upload", &multi[0], "").Return(domain.Media{}, nil)
				m.On("Validate", &multi[0]).Return(&errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
			url: "/media",
		},
		"Internal": {
			want:    `{}`,
			status:  500,
			message: "internal",
			files:   1,
			mock: func(m *mocks.MediaRepository, multi []multipart.FileHeader) {
				m.On("Upload", &multi[0], "").Return(domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
				m.On("Validate", &multi[0]).Return(nil)
			},
			url: "/media",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)

			request, multi := newfileUploadRequest(t, test.files, "https://google.com/upload", path)

			mock := &mocks.MediaRepository{}
			test.mock(mock, multi)
			rr.gin.Request = request

			if name == "No Form" {
				rr.RequestAndServe("POST", test.url, "/media", nil, func(g *gin.Context) {
					getMediaMock(mock).Upload(rr.gin)
				})
				return
			}

			getMediaMock(mock).Upload(rr.gin)
			rr.Run(test.want, test.status, test.message)
		})
	}
}

// newfileUploadRequest - is a helper for setting up test files for the upload
// endpoint. Creates a new file upload http request with optional
// extra params
func newfileUploadRequest(t *testing.T, filesAmount int, uri string, path string) (*gohttp.Request, []multipart.FileHeader) {
	file, err := os.Open(path)
	assert.NoError(t, err)
	defer file.Close()

	reqBody := bytes.Buffer{}
	var multi []multipart.FileHeader
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for i := 0; i < filesAmount; i++ {
		part, err := writer.CreateFormFile("file", filepath.Base(path))
		assert.NoError(t, err)
		_, err = io.Copy(part, file)
		assert.NoError(t, err)
	}

	err = writer.Close()
	assert.NoError(t, err)

	reqBody.Write(body.Bytes())

	if filesAmount != 0 {
		mr := multipart.NewReader(body, writer.Boundary())
		mt, err := mr.ReadForm(99999)
		assert.NoError(t, err)
		ft := mt.File["file"][0]
		multi = append(multi, *ft)
	}

	//&{map[] map[file:[0xc0002783c0 0xc000278410]]}

	req, err := gohttp.NewRequest("POST", uri, bytes.NewBuffer(reqBody.Bytes()))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, multi
}

// TestMedia_Delete - Test Delete route
func TestMedia_Delete(t *testing.T) {

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(u *mocks.MediaRepository)
		url     string
	}{
		"Success": {
			want:    `{}`,
			status:  200,
			message: "Successfully deleted media item with ID: 123",
			mock: func(u *mocks.MediaRepository) {
				u.On("Delete", 123).Return(nil)
			},
			url: "/media/123",
		},
		"Invalid ID": {
			want:    `{}`,
			status:  400,
			message: "A valid ID is required to delete a media item",
			mock: func(m *mocks.MediaRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/media/wrongid",
		},
		"Not Found": {
			want:    `{}`,
			status:  400,
			message: "not found",
			mock: func(m *mocks.MediaRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/media/123",
		},
		"Conflict": {
			want:    `{}`,
			status:  400,
			message: "conflict",
			mock: func(m *mocks.MediaRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			url: "/media/123",
		},
		"Internal": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(m *mocks.MediaRepository) {
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
