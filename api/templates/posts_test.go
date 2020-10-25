package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"testing"
)

func TestGetPost(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := NewFunctions(nil, &models.Store{}, &domain.Post{})

	mockPostItem := domain.Post{
		Id:          1,
		Title: 		"test title",
	}

	f.store.Posts = &mockPosts
	mockPosts.On("GetById", 1).Return(mockPostItem, nil)

	_ = f.getPost(1)

	mockPosts.AssertExpectations(t)
}

func TestGetPost_NoItem(t *testing.T) {
	mockPosts := mocks.PostsRepository{}
	f := NewFunctions(nil, &models.Store{}, &domain.Post{})

	f.store.Posts = &mockPosts
	mockPosts.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("No post item"))

	_ = f.getPost(1)

	mockPosts.AssertExpectations(t)
}