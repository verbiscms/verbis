package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"testing"
)

func TestGetField(t *testing.T) {
	f := newTestSuite(`{"text": "content"}`)

	tpl := `{{ getField "text" }}`
	runt(t, f, tpl, "content")

	tpl2 := `{{ getField "wrongval" }}`
	runt(t, f, tpl2, "")
}

func TestGetField_Post(t *testing.T) {
	f := newTestSuite(`{"text": "content"}`)

	mockPosts := mocks.PostsRepository{}
	mockPost := domain.Post{
		Id: 1,
		Fields: map[string]interface{}{
			"poststext": "postcontent",
		},
	}

	mockPosts.On("GetById", 1).Return(mockPost, nil)
	f.store.Posts = &mockPosts

	tpl := `{{ getField "text" 1 }}`
	runt(t, f, tpl, "")
}

func TestGetField_No_Post(t *testing.T) {
	f := newTestSuite(`{}`)

	mockPosts := mocks.PostsRepository{}
	f.store.Posts = &mockPosts
	mockPosts.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("No post"))

	tpl := `{{ getField "posttext" 1 }}`
	runt(t, f, tpl, "")
}

func TestGetField_Invalid_Json(t *testing.T) {
	f := newTestSuite("{}")

	mockPosts := mocks.PostsRepository{}
	mockPost := domain.Post{
		Id: 1,
		Fields: map[string]interface{}{
			"test": "content",
		},
	}

	mockPosts.On("GetById", 1).Return(mockPost, nil)
	f.store.Posts = &mockPosts

	tpl := `{{ getField "text" 1 }}`
	runt(t, f, tpl, "")
}

func TestCheckFieldType(t *testing.T) {

}

func TestHasField(t *testing.T) {
	f := newTestSuite(`{"text": "content"}`)

	tpl := `{{ hasField "text" }}`
	runt(t, f, tpl, true)

	tpl2 := `{{ hasField "wrongval" }}`
	runt(t, f, tpl2, false)
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

	f := newTestSuite(str)

	tpl := `{{ getRepeater "wrongval" }}`
	runt(t, f, tpl, "[]")

	tpl2 := `{{ getRepeater "repeater" }}`
	runt(t, f, tpl2, "[map[text1:content text2:content] map[text1:content text2:content]]")
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

	f := newTestSuite(str)

	tpl := `{{ getFlexible "wrongval" }}`
	runt(t, f, tpl, "[]")

	tpl2 := `{{ getFlexible "flexible" }}`
	runt(t, f, tpl2, "[map[fields:map[text:content text2:content] type:block1] map[fields:map[repeater:[map[text:content text2:content]] text:content text1:content text2:content] type:block2]]")
}
