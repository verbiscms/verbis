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

// getPostsMock is a helper to obtain a mock posts controller
// for testing.
func getPostsMock(m models.PostsRepository) *PostsController {
	return &PostsController{
		store: &models.Store{
			Posts: m,
		},
	}
}

// Test_NewPosts - Test construct
func Test_NewPosts(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &PostsController{
		store:  &store,
		config: config,
	}
	got := newPosts(&store, config)
	assert.Equal(t, got, want)
}

// TestPostsController_Get - Test Get route
func TestPostsController_Get(t *testing.T) {

	posts := []domain.Post{
		{Id:          123, Slug:        "/post", Title:        "post"},
		{Id:          124, Slug:        "/post1", Title:        "post1"},
	}
	postData := []domain.PostData{
		{Post:       posts[0]},
		{Post:       posts[1]},
	}
	pagination := http.Params{Page: 1, Limit: 15, OrderBy: "id", OrderDirection: "asc", Filters: nil}

	tt := map[string]struct {
		name       string
		want       string
		status     int
		message    string
		mock func(m *mocks.PostsRepository)
	}{
		"Success": {
			want:       `[{"author":{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"","id":0,"instagram":null,"last_name":"","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"},"category":null,"layout":null,"post":{"created_at":null,"fields":null,"id":123,"options":{"meta":null,"seo":null},"published_at":null,"resource":null,"slug":"/post","title":"post","updated_at":null,"uuid":"00000000-0000-0000-0000-000000000000"}},{"author":{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"","id":0,"instagram":null,"last_name":"","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"},"category":null,"layout":null,"post":{"created_at":null,"fields":null,"id":124,"options":{"meta":null,"seo":null},"published_at":null,"resource":null,"slug":"/post1","title":"post1","updated_at":null,"uuid":"00000000-0000-0000-0000-000000000000"}}]`,
			status:     200,
			message:    "Successfully obtained posts",
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", pagination, "").Return(posts, 2, nil)
				m.On("FormatMultiple", posts).Return(postData, nil)
			},
		},
		"Not Found": {
			want:       `{}`,
			status:     200,
			message:    "no posts found",
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", pagination, "").Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no posts found"})
				m.On("FormatMultiple", posts).Return(postData, nil)
			},
		},
		"Conflict": {
			want:       `{}`,
			status:     400,
			message:    "conflict",
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", pagination, "").Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
				m.On("FormatMultiple", posts).Return(postData, nil)
			},
		},
		"Invalid": {
			want:       `{}`,
			status:     400,
			message:    "invalid",
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", pagination, "").Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
				m.On("FormatMultiple", posts).Return(postData, nil)
			},
		},
		"Internal Error": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", pagination, "").Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
				m.On("FormatMultiple", posts).Return(postData, nil)
			},
		},
		"Format Error": {
			want:       `{}`,
			status:     500,
			message:    "format",
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", pagination, "").Return(posts, 2, nil)
				m.On("FormatMultiple", posts).Return(postData, &errors.Error{Code: errors.INTERNAL, Message: "format"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.PostsRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/posts", "/posts", nil, func(g *gin.Context) {
				getPostsMock(mock).Get(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestPostsController_GetById - Test GetByID route
func TestPostsController_GetById(t *testing.T) {

	post := domain.Post{Id:          123, Slug:        "/post", Title:        "post"}
	postData := domain.PostData{Post:     domain.Post{Id:          123, Slug:        "/post", Title:        "post"}  }

	tt := map[string]struct {
		want       string
		status     int
		message    string
		mock func(m *mocks.PostsRepository)
		url string
	}{
		"Success": {
			want:       `{"author":{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"","id":0,"instagram":null,"last_name":"","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"},"category":null,"layout":null,"post":{"created_at":null,"fields":null,"id":123,"options":{"meta":null,"seo":null},"published_at":null,"resource":null,"slug":"/post","title":"post","updated_at":null,"uuid":"00000000-0000-0000-0000-000000000000"}}`,
			status:     200,
			message:    "Successfully obtained post with ID: 123",
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 123).Return(post, nil)
				m.On("Format", post).Return(postData, nil)
			},
			url: "/posts/123",
		},
		"Invalid ID": {
			want:       `{}`,
			status:     400,
			message:    "Pass a valid number to obtain the post by ID",
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 123).Return(domain.Post{}, fmt.Errorf("error"))
				m.On("Format", post).Return(postData, nil)
			},
			url: "/posts/wrongid",
		},
		"Not Found": {
			want:       `{}`,
			status:     200,
			message:    "no posts found",
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 123).Return(domain.Post{}, &errors.Error{Code: errors.NOTFOUND, Message: "no posts found"})
				m.On("Format", post).Return(postData, nil)
			},
			url: "/posts/123",
		},
		"Internal Error": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 123).Return(domain.Post{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
				m.On("Format", post).Return(postData, nil)
			},
			url: "/posts/123",
		},
		"Format Error": {
			want:       `{}`,
			status:     500,
			message:    "format",
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 123).Return(post, nil)
				m.On("Format", post).Return(postData, &errors.Error{Code: errors.INTERNAL, Message: "format"})
			},
			url: "/posts/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.PostsRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", test.url, "/posts/:id", nil, func(g *gin.Context) {
				getPostsMock(mock).GetById(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestPostsController_Create - Test Create route
func TestPostsController_Create(t *testing.T) {

	post := domain.Post{Id:          123, Slug:        "/post", Title:        "post"}
	postBadValidation := domain.Post{Id:          123, Title:        "post"}

	tt := map[string]struct {
		want       string
		status     int
		message    string
		input interface{}
		mock func(m *mocks.PostsRepository)
	}{
		"Success": {
			want:       `{"archive_id":null,"created_at":"0001-01-01T00:00:00Z","description":null,"id":123,"name":"post","parent_id":null,"resource":"test","slug":"/cat","updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully created post with ID: 123",
			input: post,
			mock: func(m *mocks.PostsRepository) {
				m.On("Create", &post).Return(post, nil)
			},
		},
		"Validation Failed": {
			want:       `{"errors":[{"key":"slug","message":"Slug is required.","type":"required"}]}`,
			status:     400,
			message:    "Validation failed",
			input: postBadValidation,
			mock: func(m *mocks.PostsRepository) {
				m.On("Create", &postBadValidation).Return(domain.Post{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			want:       `{}`,
			status:     400,
			message:    "invalid",
			input: post,
			mock: func(m *mocks.PostsRepository) {
				m.On("Create", &post).Return(domain.Post{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			want:       `{}`,
			status:     400,
			message:    "conflict",
			input: post,
			mock: func(m *mocks.PostsRepository) {
				m.On("Create", &post).Return(domain.Post{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			input: post,
			mock: func(m *mocks.PostsRepository) {
				m.On("Create", &post).Return(domain.Post{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.PostsRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("POST", "/posts", "/posts", bytes.NewBuffer(body), func(g *gin.Context) {
				getPostsMock(mock).Create(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestPostsController_Update - Test Update route
func TestPostsController_Update(t *testing.T) {

	post := domain.Post{Id:          123, Slug:        "/post", Title:        "post"}
	postBadValidation := domain.Post{Id:          123, Title:        "post"}

	tt := map[string]struct {
		want       string
		status     int
		message    string
		input interface{}
		mock func(m *mocks.PostsRepository)
		url string
	}{
		"Success": {
			want:       `{"archive_id":null,"created_at":"0001-01-01T00:00:00Z","description":null,"id":123,"name":"post","parent_id":null,"resource":"test","slug":"/cat","updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:     200,
			message:    "Successfully updated post with ID: 123",
			input: post,
			mock: func(m *mocks.PostsRepository) {
				m.On("Update", &post).Return(post, nil)
			},
			url: "/posts/123",
		},
		"Validation Failed": {
			want:       `{"errors":[{"key":"slug","message":"Slug is required.","type":"required"}]}`,
			status:     400,
			message:    "Validation failed",
			input: postBadValidation,
			mock: func(m *mocks.PostsRepository) {
				m.On("Update", postBadValidation).Return(domain.Post{}, fmt.Errorf("error"))
			},
			url: "/posts/123",
		},
		"Invalid ID": {
			want:       `{}`,
			status:     400,
			message:    "A valid ID is required to update the post",
			input: post,
			mock: func(m *mocks.PostsRepository) {
				m.On("Update", postBadValidation).Return(domain.Post{}, fmt.Errorf("error"))
			},
			url: "/posts/wrongid",
		},
		"Not Found": {
			want:       `{}`,
			status:     400,
			message:    "not found",
			input: post,
			mock: func(m *mocks.PostsRepository) {
				m.On("Update", &post).Return(domain.Post{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/posts/123",
		},
		"Internal": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			input: post,
			mock: func(m *mocks.PostsRepository) {
				m.On("Update", &post).Return(domain.Post{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/posts/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.PostsRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("PUT", test.url, "/posts/:id", bytes.NewBuffer(body), func(g *gin.Context) {
				getPostsMock(mock).Update(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestPostsController_Delete - Test Delete route
func TestPostsController_Delete(t *testing.T) {

	tt := map[string]struct {
		want       string
		status     int
		message    string
		mock func(m *mocks.PostsRepository)
		url string
	}{
		"Success": {
			want:       `{}`,
			status:     200,
			message:    "Successfully deleted post with ID: 123",
			mock: func(m *mocks.PostsRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/posts/123",
		},
		"Invalid ID": {
			want:       `{}`,
			status:     400,
			message:    "A valid ID is required to delete a post",
			mock:  func(m *mocks.PostsRepository) {
				m.On("Delete", 123).Return(nil)
			},
			url: "/posts/wrongid",
		},
		"Not Found": {
			want:       `{}`,
			status:     400,
			message:    "not found",
			mock:  func(m *mocks.PostsRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			url: "/posts/123",
		},
		"Conflict": {
			want:       `{}`,
			status:     400,
			message:    "conflict",
			mock:  func(m *mocks.PostsRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			url: "/posts/123",
		},
		"Internal": {
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock:  func(m *mocks.PostsRepository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			url: "/posts/123",
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.PostsRepository{}
			test.mock(mock)

			rr.RequestAndServe("DELETE", test.url, "/posts/:id", nil, func(g *gin.Context) {
				getPostsMock(mock).Delete(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}