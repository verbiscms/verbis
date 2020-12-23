package templates

import (
	"reflect"
)

// getField
//
// Get field based on input return empty string if not found.
// Uses check field type to obtain if the kind of field is a media or
// post item.
func (t *TemplateFunctions) getField(field string, id ...int) interface{} {

	// Check if the post ID was passed, and assign
	fields := t.fields
	if len(id) > 0 {
		post, err := t.store.Posts.GetById(id[0])
		if err != nil {
			return ""
		}
		fields = post.Fields
	}

	// Retrieve the value of the field
	val, found := fields[field]
	if !found {
		return ""
	}

	// Check if the field is a post, or media item
	val = t.checkFieldType(val)

	return val
}

// checkFieldType
//
// Checks to see if the field passed to it was a post, or image.
// If it was get from the getMedia or getPost functions and assign the new
// value
func (t *TemplateFunctions) checkFieldType(field interface{}) interface{} {
	kind := reflect.TypeOf(field).String()

	if kind == "map[string]interface {}" {
		m := field.(map[string]interface{})
		fieldType, found := m["type"]
		if found {
			if fieldType == "image" {
				field = t.getMedia(m["id"].(float64))
			}
		}
	} else if kind == "[]interface {}" {
		i := field.([]interface{})

		var posts []ViewPost
		for _, v := range i {
			m := v.(map[string]interface{})
			fieldType, found := m["type"]
			if found {
				if fieldType == "post" {
					post := t.getPost(m["id"].(float64))
					if post != nil {
						posts = append(posts, *post)
					}
				}
			}
		}

		if len(posts) == 1 {
			return posts[0]
		} else if len(posts) > 1 {
			field = &posts
		}
	}

	return field
}

// getFields
//
// Get all fields for the given template.
func (t *TemplateFunctions) getFields(id ...int) map[string]interface{} {
	fields := t.fields
	if len(id) > 0 {
		post, err := t.store.Posts.GetById(id[0])
		if err != nil {
			return nil
		}
		fields = post.Fields
	}

	return fields
}

// hasField
//
// Determine if the given field exists
func (t *TemplateFunctions) hasField(field string) bool {
	if _, found := t.fields[field]; found {
		return true
	}
	return false
}

// getRepeater
//
// Accepts a field and checks to see if the repeater exists.
// If it exists, build an array of map[string]interface to return to
// the template.
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
//
// Accepts a field and checks to see if the flexible content
// exists. Build an array of map[string]interface to return to
// the template.
func (t *TemplateFunctions) getFlexible(field string) []map[string]interface{} {
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

// getSubField
//
// Looks for a given field from the input & compares against
// the layout. Returns the sub field value in the layout if found.
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
