// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

type (
	// Role defines the role a user has, from the pivot
	// table.
	Role struct {
		ID          int    `db:"id" json:"id" binding:"required,numeric" validation_key:"role_id"`
		Name        string `db:"name" json:"name"`
		Description string `db:"description" json:"description"`
	}
	// Roles represents the slice of Role's.
	Roles []Role
)

const (
	// BannedRoleID is the default banned role ID.
	BannedRoleID = 1
	// ContributorRoleID is the default contributor role ID.
	ContributorRoleID = 2
	// AuthorRoleID is the default author role ID.
	AuthorRoleID = 3
	// EditorRoleID is the default editor role ID.
	EditorRoleID = 4
	// AdminRoleID is the default admin role ID.
	AdminRoleID = 5
	// OwnerRoleID is the default owner role ID.
	OwnerRoleID = 6
)
