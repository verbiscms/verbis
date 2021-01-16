package resolve

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
	log "github.com/sirupsen/logrus"
)

type Value struct {
	store *models.Store
}

// TODO: If choice is with a key of `key` or `value`

var (
	Iterable = []string{
		"category",
		"image",
		"posts",
		"user",
		"tags",
	}
)

type valuer func(field domain.FieldValue) (interface{}, error)

type fieldValueMap map[string]valuer

// resolveField
//
// Determines if the given field values is a slice or array or singular
// and returns a resolved field value or a slice of interfaces.
func Field(field domain.PostField, store *models.Store) domain.PostField {
	exec := &Value{
		store: store,
	}
	resolved := exec.resolve(field)
	return resolved
}

func (v *Value) getMap() fieldValueMap {
	return fieldValueMap{
		"category": v.category,
		"checkbox": v.checkbox,
		"choice":   v.choice,
		"image":    v.media,
		"number":   v.number,
		"post":     v.post,
		"range":   	v.number,
		"user":     v.user,
	}
}

func (v *Value) resolve(field domain.PostField) domain.PostField  {
	original := field.OriginalValue

	if original.IsEmpty() && field.Key != "map" {
		field.Value = field.OriginalValue.String()
		return field
	}

	if !field.IsIterable(Iterable) {
		field.Value = v.execute(field.OriginalValue.String(), field.Type)
		return field
	}

	var items []interface{}
	for _, f := range original.Array() {
		items = append(items, v.execute(f, field.Type))
	}
	field.Value = items

	return field
}

func (v *Value) execute(value string, typ string) interface{} {
	fn, ok := v.getMap()[typ]
	if !ok {
		return value
	}

	val, err := fn(domain.FieldValue(value))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
	}

	return val
}
