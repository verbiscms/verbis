// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
)

// FieldService defines methods for obtaining fields for the front end templates.
type FieldService interface {
	// GetField returns the value of a specific field.
	// Returns errors.NOTFOUND if the field was not found by the given key.
	GetField(name string, args ...interface{}) interface{}
	// GetFieldObject returns the raw object of a specific field.
	// Returns errors.NOTFOUND if the field was not found by the given key.
	GetFieldObject(name string, args ...interface{}) domain.PostField
	// GetFields Returns all of the fields for the current post,
	// or post ID given.
	GetFields(args ...interface{}) Fields
	// GetLayout returns a field group layout by the given name.
	// Logs errors.NOTFOUND if the layout could not be found by the given name.
	GetLayout(name string, args ...interface{}) domain.Field
	// GetLayouts returns all field group layouts.
	GetLayouts(args ...interface{}) domain.FieldGroups
	// GetRepeater Returns the collection of children from the given
	// key and returns a new Repeater.
	// Returns errors.NOTFOUND if the field was not found by the given key.
	// Returns errors.INVALID if the field type is not a repeater or the name could not be cast.
	GetRepeater(input interface{}, args ...interface{}) Repeater
	// GetFlexible Returns the collection of Layouts from the
	// given key and returns a new Flexible.
	// Returns errors.INVALID if the field type is not flexible content.
	// Returns errors.INTERNAL if the layouts could not be cast to a string slice.
	GetFlexible(input interface{}, args ...interface{}) Flexible
}

const (
	// SEPARATOR is the separator that defines the split between
	// field keys for repeaters and flexible content.
	SEPARATOR = "|"
)

// Service defines the helper for obtaining fields for
// front end templates.
type Service struct {
	// Used for obtaining categories, media items, posts and
	// users from the database when resolving fields.
	deps *deps.Deps
	// The original post ID.
	postID int
	// The slice of domain.PostField to create repeaters,
	// flexible content and resolving normal fields.
	fields domain.PostFields
	// The slice of domain.FieldGroup to iterate over
	// groups and layouts.
	layout domain.FieldGroups
}

// NewService construct, creates a new slice of post fields and
// slice of layouts.
func NewService(d *deps.Deps, p *domain.PostDatum) *Service {
	fields := make(domain.PostFields, 0)
	if p.Fields != nil {
		fields = p.Fields
	}
	layouts := make(domain.FieldGroups, 0)
	if p.Layout != nil {
		layouts = p.Layout
	}
	return &Service{
		deps:   d,
		postID: p.Id,
		fields: fields,
		layout: layouts,
	}
}
