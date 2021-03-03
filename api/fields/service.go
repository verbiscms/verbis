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
	GetField(name string, args ...interface{}) interface{}
	GetFieldObject(name string, args ...interface{}) domain.PostField
	GetFields(args ...interface{}) Fields
	GetLayout(name string, args ...interface{}) domain.Field
	GetLayouts(args ...interface{}) domain.FieldGroups
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
// Defines the helper for obtaining fields for front end
// templates.
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

// NewService
//
// Construct, creates a new slice of post fields and slice
// of layouts.
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
