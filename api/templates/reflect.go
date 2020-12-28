package templates

import (
	"fmt"
	"reflect"
)

// kindIs
//
// Returns true if the target and source types match.
//
// Example: {{ kindIs "int" 123 }} Returns true
func (t *TemplateManager) kindIs(target string, src interface{}) bool {
	return target == t.kindOf(src)
}

// kindOf
//
// Returns the type of the given `interface` as a string.
//
// Example: {{ kindOf 123 }} Returns `int`
func (t *TemplateManager) kindOf(src interface{}) string {
	return reflect.ValueOf(src).Kind().String()
}

// typeOf
//
// Returns the underlying type of the given value.
//
// Example: {{ typeOf .Post }} Returns `domain.Post`
func (t *TemplateManager) typeOf(src interface{}) string {
	return fmt.Sprintf("%T", src)
}

// typeIs
//
// Similar to `kindIs` but its used for types, instead of primitives.
//
// Example: {{ typeOf "domain.Post" .Post }} Returns true
func (t *TemplateManager) typeIs(target string, src interface{}) bool {
	return target == t.typeOf(src)
}

// typeIsLike
//
// Similar to `kindIs` but its used for types, instead of primitives.
//
// Example: {{ typeOf "domain.Post" .Post }} Returns true
func (t *TemplateManager) typeIsLike(target string, src interface{}) bool {
	s := t.typeOf(src)
	return target == s || "*"+target == s
}
