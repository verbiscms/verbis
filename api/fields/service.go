package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
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
	SEPARATOR = "|"
)

// Service
//
// Defines the helper for obtaining fields for front end templates.
type Service struct {
	// Used for obtaining categories, media items, posts and
	// users from the database when resolving fields.
	store *models.Store
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
func NewService(s *models.Store, d domain.PostData) *Service {
	fields := make([]domain.PostField, 0)
	if d.Fields != nil {
		fields = d.Fields
	}

	layouts := make([]domain.FieldGroup, 0)
	if d.Layout != nil {
		layouts = d.Layout
	}

	return &Service{
		store:  s,
		postId: d.Post.Id,
		fields: fields,
		layout: layouts,
	}
}
