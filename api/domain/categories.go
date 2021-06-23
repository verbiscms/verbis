// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/google/uuid"
	"time"
)

type (
	// Category defines the groups used for categorising
	// individual posts.
	Category struct {
		Id          int       `sqlb:"id,skipAll" db:"id" json:"id"` //nolint
		UUID        uuid.UUID `sqlb:"uuid" db:"uuid" json:"uuid"`
		Slug        string    `sqlb:"slug" db:"slug" json:"slug" binding:"required,max=150"`
		Name        string    `sqlb:"name" db:"name" json:"name" binding:"required,max=150"`
		Description *string   `sqlb:"description" db:"description" json:"description" binding:"omitempty,max=500"`
		Resource    string    `sqlb:"resource" db:"resource" json:"resource" binding:"required,max=150"`
		ParentId    *int      `sqlb:"parent_id" db:"parent_id" json:"parent_id" binding:"omitempty,numeric"`    //nolint
		ArchiveId   *int      `sqlb:"archive_id" db:"archive_id" json:"archive_id" binding:"omitempty,numeric"` //nolint
		CreatedAt   time.Time `sqlb:"created_at,autoCreateTime" db:"created_at" json:"created_at"`
		UpdatedAt   time.Time `sqlb:"updated_at,autoUpdateTime" db:"updated_at" json:"updated_at"`
	}
	// Categories represents the slice of Category's.
	Categories []Category
)

// HasParent determines if a category has a parent
// ID attached to it.
func (c *Category) HasParent() bool {
	return c.ParentId != nil
}
