package templates

import (
	"github.com/spf13/cast"
	"reflect"
)

// getField
//
// Get field based on input return empty string if not found.
// Uses check field type to obtain if the kind of field is a media or
// post item.
func (t *TemplateManager) getField(field string, id ...int) interface{} {

	// Check if the post ID was passed, and assign
	fields := t.fields
	if len(id) > 0 {
		post, err := t.store.Posts.GetById(id[0])
		if err != nil {
			return nil
		}
		fields = post.Fields
	}

	// Retrieve the value of the field
	val, found := fields[field]
	if !found {
		return nil
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
func (t *TemplateManager) checkFieldType(field interface{}) interface{} {
	typ := reflect.TypeOf(field)

	switch typ.Kind() {
		case reflect.Array, reflect.Slice:
			field = t.fieldSliceResolver(field.([]interface{}), field)
		case reflect.Map: {
			field = t.fieldMapResolver(field.(map[string]interface{}), field)
		}
	}

	return field
}

func (t *TemplateManager) fieldSliceResolver(i []interface{}, field interface{}) interface{} {
	for _, v := range i {
		typ := reflect.TypeOf(v)
		switch typ.Kind() {
		case reflect.Map:
			field = t.fieldMapResolver(v.(map[string]interface{}), field)
		}
	}

	return field
}

func (t *TemplateManager) fieldMapResolver(m map[string]interface{}, field interface{}) interface{} {
	if fieldType, found := m["type"]; found {
		field = t.resolveField(fieldType, m)
	} else {
		for _, v := range m {
			typ := reflect.TypeOf(v)
			switch typ.Kind() {
			case reflect.Map:
				field = t.fieldMapResolver(v.(map[string]interface{}), field)
			case reflect.Array, reflect.Slice:
				field = t.fieldSliceResolver(v.([]interface{}), field)
			}
		}
	}

	return field
}

// resolveField
func (t *TemplateManager) resolveField(typ interface{}, fields map[string]interface{}) interface{} {

	switch cast.ToString(typ) {
	case "category":
		return *t.getCategory(fields["id"])
	case "image":
		return *t.getMedia(fields["id"])
	case "post":
		return *t.getPost(fields["id"])
	case "user":
		return *t.getUser(fields["id"])
	default:
		return nil
	}
}

// getFields
//
// Get all fields for the given template.
func (t *TemplateManager) getFields(id ...int) map[string]interface{} {
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
func (t *TemplateManager) hasField(field string) bool {
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
func (t *TemplateManager) getRepeater(field string) []map[string]interface{} {
	if f, found := t.fields[field]; found {
		items := f.([]interface{})
		repeater := make([]map[string]interface{}, len(items))

		// Loop over items in the repeater
		for index, item := range items {
			repeater[index] = make(map[string]interface{})

			// Loop sub over elements in the repeater map
			if reflect.TypeOf(item).Kind() == reflect.Map {
				for k, sub := range item.(map[string]interface{}) {

					// Account for sub arrays
					if reflect.TypeOf(sub).Kind() == reflect.Slice {
						subSlice := sub.([]interface{})
						subRepeater := make([]interface{}, len(subSlice))
						for k1, sub2 := range subSlice {
							subRepeater[k1] = t.checkFieldType(sub2)
						}

						if len(subRepeater) == 1 {
							repeater[index][k] = subRepeater[0]
						} else {
							repeater[index][k] = subRepeater
						}
					} else {
						repeater[index][k] = t.checkFieldType(sub)
					}
				}
			}
		}

		return repeater
	}
	return nil
}

// getFlexible
//
// Accepts a field and checks to see if the flexible content
// exists. Build an array of map[string]interface to return to
// the template.
func (t *TemplateManager) getFlexible(field string) []map[string]interface{} {
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
func (t *TemplateManager) getSubField(field string, layout map[string]interface{}) interface{} {
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
