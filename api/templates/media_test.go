package templates

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"html"
	"html/template"
	"strings"
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

	tpl := "{{ getMedia 1 }}"
	runt(f, tpl, mockMediaItem)

	//_ = f.getMedia(1)
	//
	//mockMedia.AssertExpectations(t)
}

func TestGetMedia_NoItem(t *testing.T) {
	mockMedia := mocks.MediaRepository{}
	f := NewFunctions(nil, &models.Store{}, &domain.Post{})

	f.store.Media = &mockMedia
	mockMedia.On("GetById", 1).Return(domain.Media{}, fmt.Errorf("No media item"))

	_ = f.getMedia(1)

	mockMedia.AssertExpectations(t)
}

func runt(tf *TemplateFunctions, tpl string, expect interface{}) error {
	tt := template.Must(template.New("test").Funcs(tf.GetFunctions()).Parse(tpl))
	var b bytes.Buffer
	err := tt.Execute(&b, map[string]string{})
	if err != nil {
		fmt.Println(err)
	}
	got := strings.ReplaceAll(html.EscapeString(fmt.Sprintf("%v", expect)), "+", "&#43;")
	if got != b.String() {
		fmt.Println("fail")
	}
	return nil
}