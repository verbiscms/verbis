package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
)

// FieldService defines methods for obtaining fields for the front end templates
type FieldService interface {
	GetField(name string, args ...interface{}) (interface{}, error)
	GetFields(args ...interface{}) (Fields, error)
	GetRepeater(name string, args ...interface{}) (Repeater, error)
	GetFlexible(name string, args ...interface{}) (Flexible, error)
}

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
	return &Service{
		store:  s,
		postId: d.Id,
		fields: *d.Fields,
		layout: *d.Layout,
	}
}
