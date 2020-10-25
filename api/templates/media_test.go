package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"testing"
)

func TestGetMedia(t *testing.T) {
	mockMedia := mocks.MediaRepository{}
	f := NewFunctions(nil, &models.Store{}, &domain.Post{})

	mockMediaItem := domain.Media{
		Id:          1,
		Url:         "/uploads/test.jpg",
	}

	f.store.Media = &mockMedia
	mockMedia.On("GetById", 1).Return(mockMediaItem, nil)

	_ = f.getMedia(1)

	mockMedia.AssertExpectations(t)
}

func TestGetMedia_NoItem(t *testing.T) {
	mockMedia := mocks.MediaRepository{}
	f := NewFunctions(nil, &models.Store{}, &domain.Post{})

	f.store.Media = &mockMedia
	mockMedia.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("No media item"))

	_ = f.getMedia(1)

	mockMedia.AssertExpectations(t)
}