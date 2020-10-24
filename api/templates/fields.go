package templates

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/models"
)

type fieldsHandler interface {
	getField(field string, id ...int) interface{}
	getFields(id ...int) map[string]interface{}
	hasField(field string) bool
	getRepeater(field string) []map[string]interface{}
	getFlexible(field string) []map[string]interface{}
	getSubField(field string, layout map[string]interface{}) interface{}
}

type fields struct {
	fields map[string]interface{}
	store *models.Store
}

// newFields - Construct
func newFields(f map[string]interface{}, s *models.Store) *fields {
	return &fields{
		fields: f,
		store: s,
	}
}

// If field is type of media
// If field is type of post
// Get the field

// getField - Get field based on input return nothing if found
func (f *fields) getField(field string, id ...int) interface{} {

	// Check if the post ID was passed, and assign
	fields := f.fields
	if len(id) > 0 {
		post, err := f.store.Posts.GetById(id[0])

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

	return val
}

// getFields - Get all fields for template
func (f *fields) getFields(id ...int) map[string]interface{} {
	fields := f.fields
	if len(id) > 0 {
		post, err := f.store.Posts.GetById(id[0])

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
func (f *fields) hasField(field string) bool {
	if _, found := f.fields[field]; found {
		return true
	}
	return false
}

// getRepeater
func (f *fields) getRepeater(field string) []map[string]interface{} {
	if _, found := f.fields[field]; found {
		fields := f.fields[field].([]interface{})
		var f []map[string]interface{}
		for _, v := range fields {
			f = append(f, v.(map[string]interface{}))
		}
		return f
	}
	return nil
}

// getFlexible
func (f *fields) getFlexible(field string) []map[string]interface{} {
	if _, found := f.fields[field]; found {
		fields, ok := f.fields[field].([]interface{})
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

func (f *fields) getSubField(field string, layout map[string]interface{}) interface{} {
	block := layout["fields"]
	fields, ok := block.(map[string]interface{})
	if !ok {
		return nil
	}
	if _, found := fields[field]; found {
		return fields[field]
	}
	return nil
}



