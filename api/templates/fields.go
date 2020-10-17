package templates

import "fmt"

/*
 * Fields
 * Functions for templates for Fields associated with the post
 */

// Get field based on input return nothing if found
func (t *TemplateFunctions) getField(field string) interface{} {
	if _, found := t.fields[field]; found {
		return t.fields[field]
	} else {
		return ""
	}
}

// Get all fields for template
func (t *TemplateFunctions) getFields() map[string]interface{} {
	return t.fields
}

// Determine if the given field exists
func (t *TemplateFunctions) hasField(field string) bool {
	if _, found := t.fields[field]; found {
		return true
	}
	return false
}

// Get repeater field
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

// Get flexible content field
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

func (t *TemplateFunctions) getSubField(field string, test string) map[string]interface{} {
	fmt.Println(test)
	fmt.Println(field)
	return nil
}

