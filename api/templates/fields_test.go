package templates

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mocks"
	"github.com/ainsleyclark/verbis/api/models"

	"github.com/ainsleyclark/verbis/api/test"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetField(t *testing.T) {
	f, err := helper(`{"text": "content"}`)
	if err != nil {
		t.Error(err)
	}

	if field := f.getField("text"); field == "" {
		t.Errorf(test.Format("content", nil))
	}

	if field := f.getField("wrongval"); field != "" {
		t.Errorf(test.Format("", field))
	}
}

func TestGetField_Post(t *testing.T) {
	f, err := helper("{}")
	if err != nil {
		t.Error(err)
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	data := []byte(`{"posttext": "postcontent"}`)
	mockPost := domain.Post{
		Id:     2,
		Fields: (*json.RawMessage)(&data),
	}

	posts := mocks.NewMockPostsRepository(controller)
	f.store.Posts = posts
	posts.EXPECT().GetById(2).Return(mockPost, nil)

	field := f.getField("posttext", 2)
	if field != "postcontent" {
		t.Errorf(test.Format("postcontent", field))
	}
}

func TestGetField_No_Post(t *testing.T) {
	f, err := helper("{}")
	if err != nil {
		t.Error(err)
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	posts := mocks.NewMockPostsRepository(controller)
	var mockErr = fmt.Errorf("No post")
	posts.EXPECT().GetById(gomock.Any()).Return(domain.Post{}, mockErr)
	f.store.Posts = posts

	field := f.getField("text", 1)

	if field != "" {
		t.Errorf(test.Format("", field))
	}
}

func TestGetField_Invalid_JSON(t *testing.T) {
	f, err := helper("{}")
	if err != nil {
		t.Error(err)
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	data := []byte(`"text "content"`)
	mockPost := domain.Post{
		Id:     1,
		Fields: (*json.RawMessage)(&data),
	}

	posts := mocks.NewMockPostsRepository(controller)
	posts.EXPECT().GetById(1).Return(mockPost, nil)
	f.store.Posts = posts

	field := f.getField("text", 1)
	if field != "" {
		t.Errorf(test.Format("", field))
	}
}


func TestHasField(t *testing.T) {
	f, err := helper(`{"text": "content"}`)
	if err != nil {
		t.Error(err)
	}

	if has := f.hasField("text"); !has {
		t.Errorf(test.Format(true, has))
	}

	if has := f.hasField("wrongval"); has {
		t.Errorf(test.Format(true, has))
	}
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

	f, err := helper(str)
	if err != nil {
		t.Error(err)
	}

	if field := f.getRepeater("wrongval"); len(field) != 0 {
		t.Error(test.Format(nil, field))
	}

	repeater := f.getRepeater("repeater")
	if repeater == nil {
		t.Error(test.Format(str, nil))
	}

	if len(repeater) != 2 {
		t.Error(test.Format("length of 2", len(repeater)))
	}

	if _, ok := repeater[0]["text1"]; !ok {
		t.Error(test.Format("text1", nil))
	}

	if _, ok := repeater[1]["text2"]; !ok {
		t.Error(test.Format("text2", nil))
	}
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

	f, err := helper(str)
	if err != nil {
		t.Error(err)
	}

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


func helper(str string) (*fields, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Printf("Cannot unmarshal fields: %v", err)
	}

	return newFields(m, &models.Store{}), nil
}
