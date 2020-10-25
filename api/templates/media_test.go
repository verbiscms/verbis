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
	f := NewFunctions(nil, &models.Store{}, nil)
	
	mockMediaItem := domain.MediaPublic{
		Media: domain.Media{
			Id: 1,
		},
		Sizes: nil,
	}

	mockMedia.On("GetById", 1).Return(&mockMediaItem, nil)
	f.store.Media = &mockMedia

	m := f.getMedia(1)

	fmt.Println(m)
}
