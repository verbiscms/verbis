package fields

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
)

// FieldService defines methods for obtaining fields for the front end templates
type FieldService interface {
	GetField(name string, args ...interface{}) interface{}
	GetFieldObject(name string, args ...interface{}) domain.PostField
	GetFields(args ...interface{}) Fields
	GetLayout(name string, args ...interface{}) domain.Field
	GetLayouts(args ...interface{}) []domain.FieldGroup
	GetRepeater(input interface{}, args ...interface{}) Repeater
	GetFlexible(input interface{}, args ...interface{}) Flexible
}

const (
	// The separator that defines the split between field
	// keys for repeaters and flexible content.
	SEPARATOR = "|"
)

// Service
//
// Defines the helper for obtaining fields for front end templates.
type Service struct {
	// Used for obtaining categories, media items, posts and
	// users from the database when resolving fields.
	deps *deps.Deps
	// The original post to sort and filter the fields
	post domain.PostData
	// The original post ID.
	postId int
	// The slice of domain.PostField to create repeaters,
	// flexible content and resolving normal fields.
	fields []domain.PostField
	// The slice of domain.FieldGroup to iterate over
	// groups and layouts.
	layout []domain.FieldGroup
}

// NewService - Construct
func NewService(d *deps.Deps, p domain.PostData) *Service {
	fields := make([]domain.PostField, 0)
	if p.Fields != nil {
		fields = p.Fields
	}

	layouts := make([]domain.FieldGroup, 0)
	if p.Layout != nil {
		layouts = p.Layout
	}

	return &Service{
		deps:   d,
		postId: p.Id,
		fields: fields,
		layout: layouts,
	}
}
