package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
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