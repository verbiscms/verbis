package templates

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/test"
	"reflect"
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
		t.Errorf(test.Format(nil, field))
	}
}

func TestGetFields(t *testing.T) {
	f, err := helper(`{"text": "content", "text2": "content2"}`)
	if err != nil {
		t.Error(err)
	}

	m := map[string]interface{}{}
	if reflect.TypeOf(f.getFields()) != reflect.TypeOf(m) {
		t.Errorf(test.Format("map[string]interface{}", reflect.TypeOf(f.getFields())))
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
		t.Error(test.Format(`{"repeater": [{"text": "text1", "": "text2"}]}`, nil))
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
				 "type": "layoutkey1",
				 "fields": {
					"text": "text",
					"text2": ""
				 }
			},
			{
				"type": "layoutkey2",
				"fields": {
					"text": "default",
					"text1": "",
					"text2": "default",
					"repeater": [
						{
						  "text":"text",
						  "text2":"text"
						}
					]
				}
			}
      	]
   	}`

	_, err := helper(str)
	if err != nil {
		t.Error(err)
	}

}

func helper(fields string) (*Fields, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(fields), &m)
	if err != nil {
		fmt.Printf("Cannot unmarshal fields: %v", err)
		fmt.Println()
	}

	return &Fields{
		fields: m,
	}, nil
}
