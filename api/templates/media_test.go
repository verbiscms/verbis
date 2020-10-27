package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"testing"
)

func TestGetMedia(t *testing.T) {
	mockMedia := mocks.MediaRepository{}
	f := newTestSuite()

	mockMediaItem := domain.Media{
		Id:          1,
		Url:         "/uploads/test.jpg",
	}

	f.store.Media = &mockMedia
	mockMedia.On("GetById", 1).Return(mockMediaItem, nil)

	tpl := "{{ getMedia 1 }}"
	runt(t, f, tpl, mockMediaItem)
}

func TestGetMedia_NoItem(t *testing.T) {
	mockMedia := mocks.MediaRepository{}
	f := newTestSuite()

	f.store.Media = &mockMedia
	mockMedia.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("No media item"))

	_ = f.getMedia(1)

	tpl := "{{ getMedia 1 }}"
	runt(t, f, tpl, nil)
}

