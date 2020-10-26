package templates

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetField(t *testing.T) {
	f := helper(`{"text": "content"}`)

	if got := f.getField("text"); got == "" {
		assert.Equal(t, got, "content")
	}

	if got := f.getField("wrongval"); got != "" {
		assert.Equal(t, got, "")
	}
}

func TestGetField_Post(t *testing.T) {
	f := helper(`{"text": "content"}`)

	mockPosts := mocks.PostsRepository{}
	data := []byte(`{"posttext": "postcontent"}`)
	mockPost := domain.Post{
		Id:     1,
		Fields: (*json.RawMessage)(&data),
	}

	mockPosts.On("GetById", 1).Return(mockPost, nil)
	f.store.Posts = &mockPosts

	got := f.getField("posttext", 1)
	assert.Equal(t, got, "postcontent")
}

func TestGetField_No_Post(t *testing.T) {
	f := helper(`{}`)

	mockPosts := mocks.PostsRepository{}
	f.store.Posts = &mockPosts
	mockPosts.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("No post"))

	got := f.getField("posttext", 1)
	assert.Equal(t, got, "")
}

func TestGetField_Invalid_Json(t *testing.T) {
	f := helper("{}")

	mockPosts := mocks.PostsRepository{}
	data := []byte(`"text "content"`)
	mockPost := domain.Post{
		Id:     1,
		Fields: (*json.RawMessage)(&data),
	}

	mockPosts.On("GetById", 1).Return(mockPost, nil)
	f.store.Posts = &mockPosts

	got := f.getField("text", 1)
	assert.Equal(t, got, "")
}

func TestCheckFieldType(t *testing.T) {

}

func TestHasField(t *testing.T) {
	f := helper(`{"text": "content"}`)

	got := f.hasField("text")
	assert.Equal(t, got, true)

	got = f.hasField("wrongval")
	assert.Equal(t, got, false)
}

func TestGetRepeater(t *testing.T) {
	str := `{
		"repeater":[
			{
				"text1":"content",
				"text2":"content"
			},
			{
				 "text1":"content",
				 "text2":"content"
			}
		]
	}`

	f := helper(str)

	field := f.getRepeater("wrongval")
	assert.Equal(t, len(field), 0)

	repeater := f.getRepeater("repeater")

	assert.NotNil(t, repeater)
	assert.Equal(t, len(repeater), 2)
	assert.Equal(t, repeater[0]["text1"], "content")
	assert.Equal(t, repeater[0]["text2"], "content")
	assert.Equal(t, repeater[1]["text1"], "content")
	assert.Equal(t, repeater[1]["text2"], "content")
	assert.Nil(t, repeater[0]["wrongval"])
	assert.Nil(t, repeater[1]["wrongval"])
}

func TestGetFlexible(t *testing.T) {

	str := `{
		"flexible": [
			{
				 "type": "block1",
				 "fields": {
					"text": "content",
					"text2": "content"
				 }
			},
			{
				"type": "block2",
				"fields": {
					"text": "content",
					"text1": "content",
					"text2": "content",
					"repeater": [
						{
						  "text":"content",
						  "text2":"content"
						}
					]
				}
			}
      	]
   	}`

	f := helper(str)

	if field := f.getRepeater("wrongval"); len(field) != 0 {
		t.Error(test.Format(nil, field))
	}

	flexible := f.getRepeater("flexible")
	if flexible == nil {
		t.Error(test.Format(str, nil))
	}

	if len(flexible) != 2 {
		t.Error(test.Format("length of 2", len(flexible)))
	}
}

func helper(str string) *TemplateFunctions {
	data := []byte(str)
	p := domain.Post{
		Fields: (*json.RawMessage)(&data),
	}

	return NewFunctions(nil, &models.Store{}, &p)
}
