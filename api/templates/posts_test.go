package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	vhttp "github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetPost(t *testing.T) {

	post := domain.Post{
		Id:    1,
		Title: "test title",
		UserId: 1,
	}

	author := &domain.PostAuthor{}
	category := &domain.PostCategory{}

	viewData := ViewPost{
		Author:   author,
		Category: category,
		Post:     post,
	}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.PostsRepository)
		want  interface{}
	}{
		"Success": {
			input: 1,
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(post, nil)
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
			},
			want: viewData,
		},
		"Format Error": {
			input: 1,
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(post, nil)
				m.On("Format", post).Return(domain.PostData{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"Not Found": {
			input: 1,
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("error"))
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			postsMock := mocks.PostsRepository{}

			test.mock(&postsMock)
			f.store.Posts = &postsMock

			tpl := `{{ post .PostId }}`

			runtv(t, f, tpl, test.want, map[string]interface{}{"PostId": test.input})
		})
	}
}

func TestGetPosts_UnmarshalError(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := newTestSuite()

	p := vhttp.Params{
		Page:           1,
		Limit:          15,
		LimitAll:       false,
		OrderBy:        "published_at",
		OrderDirection: "desc",
		Filters:        map[string][]vhttp.Filter{
			"status": {
				{
					Operator: "=",
					Value:    "published",
				},
			},
		},
	}

	f.store.Posts = &mockPosts
	mockPosts.On("Get", p, "all").Return(nil, nil)

	tpl := `{{ $query := dict "534534" 5345345 }}{{ posts $query }}`
	runt(t, f, tpl, "")
}

func TestGetPosts_NilQuery(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := newTestSuite()

	p := vhttp.Params{
		Page:           1,
		Limit:          15,
		LimitAll:       false,
		OrderBy:        "published_at",
		OrderDirection: "desc",
		Filters:        map[string][]vhttp.Filter{
			"status": {
				{
					Operator: "=",
					Value:    "published",
				},
			},
		},
	}

	f.store.Posts = &mockPosts
	mockPosts.On("Get", p, "all").Return(nil, nil)

	tpl := `{{ $query := dict "wrongval" 123 }}{{ posts $query }}`
	runt(t, f, tpl, "")
}

func TestGetPosts_DefaultPage(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := newTestSuite()

	p := vhttp.Params{
		Page:           1,
		Limit:          15,
		LimitAll:       false,
		OrderBy:        "published_at",
		OrderDirection: "desc",
		Filters:        nil,
	}

	f.store.Posts = &mockPosts
	mockPosts.On("Get", p, "all").Return(nil, nil)

	tpl := `{{ $query := dict "limit" 15 }}{{ posts $query }}`
	runt(t, f, tpl, "")
}

func TestGetPosts_DefaultLimit(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := newTestSuite()

	p := vhttp.Params{
		Page:           1,
		Limit:          15,
		LimitAll:       false,
		OrderBy:        "published_at",
		OrderDirection: "desc",
		Filters:        map[string][]vhttp.Filter{
			"status": {
				{
					Operator: "=",
					Value:    "published",
				},
			},
		},
	}

	f.store.Posts = &mockPosts
	mockPosts.On("Get", p, "all").Return([]domain.Post{}, nil)

	tpl := `{{ $query := dict "page" 1 }}{{ posts $query }}`
	runt(t, f, tpl, "")
}

func TestGetPagination(t *testing.T) {
	f := newTestSuite()
	gin.SetMode(gin.ReleaseMode)
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get?page=123", nil)
	f.gin = g
	tpl := "{{ paginationPage }}"
	runt(t, f, tpl, 123)
}

func TestGetPagination_NoPage(t *testing.T) {
	f := newTestSuite()
	tpl := "{{ paginationPage }}"
	runt(t, f, tpl, 1)
}

func TestGetPagination_ConvertString(t *testing.T) {
	f := newTestSuite()
	gin.SetMode(gin.ReleaseMode)
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get?page=wrongval", nil)
	f.gin = g
	tpl := "{{ paginationPage }}"
	runt(t, f, tpl, "1")
}
