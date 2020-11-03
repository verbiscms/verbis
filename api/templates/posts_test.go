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

func TestGetPost(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := newTestSuite()

	mockPostItem := domain.Post{
		Id:          1,
		Title: 		"test title",
	}

	f.store.Posts = &mockPosts
	mockPosts.On("GetById", 1).Return(mockPostItem, nil)

	tpl := "{{ getPost 1 }}"
	runt(t, f, tpl, mockPostItem)
}

func TestGetPost_NoItem(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := newTestSuite()

	f.store.Posts = &mockPosts
	mockPosts.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("No post item"))

	tpl := "{{ getPost 1 }}"
	runt(t, f, tpl, nil)
}

func TestGetPosts_UnmarshalError(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := newTestSuite()

	p := vhttp.Params{
		Page:           1,
		Limit:          15,
		OrderBy:        "published_at",
		OrderDirection: "desc",
		Filters:        nil,
	}

	f.store.Posts = &mockPosts
	mockPosts.On("Get", p, "all").Return(nil, nil)

	tpl := `{{ $query := dict "534534" 5345345 }}{{ getPosts $query }}`
	runt(t, f, tpl, "")
}

func TestGetPosts_NilQuery(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := newTestSuite()

	p := vhttp.Params{
		Page:           1,
		Limit:          15,
		OrderBy:        "published_at",
		OrderDirection: "desc",
		Filters:        nil,
	}

	f.store.Posts = &mockPosts
	mockPosts.On("Get", p, "all").Return(nil, nil)

	tpl := `{{ $query := dict "wrongval" 123 }}{{ getPosts $query }}`
	runt(t, f, tpl, "")
}

func TestGetPosts_NilDefault(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := newTestSuite()

	p := vhttp.Params{
		Page:           1,
		Limit:          123,
		OrderBy:        "published_at",
		OrderDirection: "desc",
		Filters:        nil,
	}

	f.store.Posts = &mockPosts
	mockPosts.On("Get", p, "all").Return(nil, nil)

	tpl := `{{ $query := dict "limit" 123 }}{{ getPosts $query }}`
	runt(t, f, tpl, "")
}

func TestGetPagination(t *testing.T) {
	f := newTestSuite()
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get?page=123", nil)
	f.gin = g
	tpl := "{{ getPagination }}"
	runt(t, f, tpl, 123)
}

func TestGetPagination_NoPage(t *testing.T) {
	f := newTestSuite()
	tpl := "{{ getPagination }}"
	runt(t, f, tpl, 1)
}

func TestGetPagination_ConvertString(t *testing.T) {
	f := newTestSuite()
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get?page=wrongval", nil)
	f.gin = g
	tpl := "{{ getPagination }}"
	runt(t, f, tpl, "")
}

