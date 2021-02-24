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
		Id          int       `db:"id" json:"id"`
		UUID        uuid.UUID `db:"uuid" json:"uuid"`
		Slug        string    `db:"slug" json:"slug" binding:"required,max=150"`
		Name        string    `db:"name" json:"name" binding:"required,max=150"`
		Description *string   `db:"description" json:"description,max=500"`
		Resource    string    `db:"resource" json:"resource" binding:"required,max=150"`
		ParentId    *int      `db:"parent_id" json:"parent_id" binding:"omitempty,numeric"`
		ArchiveId   *int      `db:"archive_id" json:"archive_id" binding:"omitempty,numeric"`
		UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
		CreatedAt   time.Time `db:"created_at" json:"created_at"`
	}
	// Categories represents the slice of Category's.
	Categories []Category
)
