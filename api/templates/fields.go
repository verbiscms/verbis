package templates

import (
	"encoding/json"
	"reflect"
)

// If field is type of media
// If field is type of post
// Get the field

// getField - Get field based on input return nothing if found
func (t *TemplateFunctions) getField(field string, id ...int) interface{} {

	// Check if the post ID was passed, and assign
	fields := t.fields
	if len(id) > 0 {
		post, err := t.store.Posts.GetById(id[0])

		if err != nil {
			return ""
		}

		var m map[string]interface{}
		if err := json.Unmarshal(*post.Fields, &m); err != nil {
			return ""
		}

		fields = m
	}

	val, found := fields[field]
	if !found {
		return ""
	}

	val = t.checkFieldType(val)

	return val
}

func (t *TemplateFunctions) checkFieldType(field interface{}) interface{} {
	if reflect.TypeOf(field).String() == "map[string]interface {}" {
		m := field.(map[string]interface{})
		fieldType, found := m["type"]

		if found {
			if fieldType == "image" {
				field = t.getMedia(m["id"].(float64))
			}
			if fieldType == "post" {

			}
		}
	}
	return field
}

// getFields - Get all fields for template
func (t *TemplateFunctions) getFields(id ...int) map[string]interface{} {
	fields := t.fields
	if len(id) > 0 {
		post, err := t.store.Posts.GetById(id[0])

		if err != nil {
			return nil
		}

		var m map[string]interface{}
		if err := json.Unmarshal(*post.Fields, &m); err != nil {
			return nil
		}

		fields = m
	}

	return fields
}

// hasField - Determine if the given field exists
func (t *TemplateFunctions) hasField(field string) bool {
	if _, found := t.fields[field]; found {
		return true
	}
	return false
}

// getRepeater
func (t *TemplateFunctions) getRepeater(field string) []map[string]interface{} {
	if _, found := t.fields[field]; found {
		fields := t.fields[field].([]interface{})
		var f []map[string]interface{}
		for _, v := range fields {
			f = append(f, v.(map[string]interface{}))
		}
		return f
	}
	return nil
}

// getFlexible
func (t *TemplateFunctions) getFlexible(field string) []map[string]interface{} {
	if _, found := t.fields[field]; found {
		fields, ok := t.fields[field].([]interface{})
		if !ok {
			return nil
		}
		var f []map[string]interface{}
		for _, v := range fields {
			f = append(f, v.(map[string]interface{}))
		}
		return f
	}
	return nil
}


func (t *TemplateFunctions) getSubField(field string, layout map[string]interface{}) interface{} {
	block := layout["fields"]
	fields, ok := block.(map[string]interface{})
	if !ok {
		return nil
	}
	val, found := fields[field]
	if !found {
		return nil
	}
	return val
}



