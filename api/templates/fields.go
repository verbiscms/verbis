package templates

import "fmt"

type FieldsHandler interface {
	getField(field string) interface{}
	getFields() map[string]interface{}
	hasField(field string) bool
	getRepeater(field string) []map[string]interface{}
	getFlexible(field string) []map[string]interface{}
	getSubField(field string, test string) map[string]interface{}
}

type Fields struct {
	fields map[string]interface{}
}

// getField - Get field based on input return nothing if found
func (f *Fields) getField(field string) interface{} {
	if _, found := f.fields[field]; found {
		return f.fields[field]
	} else {
		return ""
	}

	// If field is type of media
	// If field is type of post
}

// getFields - Get all fields for template
func (f *Fields) getFields() map[string]interface{} {
	return f.fields
}

// hasField - Determine if the given field exists
func (f *Fields) hasField(field string) bool {
	if _, found := f.fields[field]; found {
		return true
	}
	return false
}

// getRepeater
func (f *Fields) getRepeater(field string) []map[string]interface{} {
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
func (f *Fields) getFlexible(field string) []map[string]interface{} {
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

func (f *Fields) getSubField(field string, test string) map[string]interface{} {
	fmt.Println(test)
	fmt.Println(field)
	return nil
}



