// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"fmt"
	"github.com/verbiscms/verbis/api/errors"
)

// Rbac - Role Based Access Control defines a map of role
// ID's with their corresponding permission groups.
type Rbac map[int]RbacGroup

// RbacGroup defines a map of permission groups assigned
// to a request. This helps to separate out different end
// points. The groups are assigned below.
// For example: "settings": { ViewMethod: ... }
type RbacGroup map[string]RbacRequest

// RbacRequest defines a map of methods assigned to a permission.
// Methods are typically defined as View, Create, Update
// and Delete.
// For example: ViewMethod: {Allow: false},
type RbacRequest map[string]Permission

// Permission defines an enforcer for endpoints and whether
// or the user is allowed to access a route.
type Permission struct {
	Allow bool `json:"allow"`
	// TBC
}

// Methods are the strings defined for actions.
const (
	// ViewMethod defines the string for finding and
	// listing resources for web routes.
	ViewMethod = "view"
	// CreateMethod defines the string for creating
	// resources for web routes.
	CreateMethod = "create"
	// UpdateMethod defines the string for updating
	// resources for web routes.
	UpdateMethod = "update"
	// DeleteMethod defines the string for updating
	// resources for web routes.
	DeleteMethod = "delete"
)

// Permissions defines an enum of permission groups to
// be used to group web routes together,
var Permissions = struct {
	Posts        string
	Categories   string
	Media        string
	Integrations string
	Forms        string
	Users        string
	Settings     string
	System       string
}{
	Posts:        "posts",
	Categories:   "categories",
	Media:        "media",
	Integrations: "integrations",
	Forms:        "forms",
	Users:        "users",
	Settings:     "settings",
	System:       "system",
}

var (
	// ErrNoGroupFound is returned by enforce when no
	// group was matched.
	ErrNoGroupFound = errors.New("no permission group found")
	// ErrNoMethodFound is returned by enforce when no
	// method was matched.
	ErrNoMethodFound = errors.New("no permission method found")
	// ErrPermissionDenied is returned by enforce when the
	// user does not have access to a given route.
	ErrPermissionDenied = errors.New("forbidden, you do not have access to to this route")
)

// Get retrieves a RbacGroup by Role ID, if there was
// no group found it will return an error.
func (r Rbac) Get(roleID int) (RbacGroup, error) {
	p, ok := r[roleID]
	if !ok {
		return nil, fmt.Errorf("no permission group found with the role ID: %d", roleID)
	}
	return p, nil
}

// Enforce ensures that a user has the correct permissions to
// access a particular route by group and method name.
// If the user can pass, nil will be returned.
// Returns ErrNoGroupFound if the group passed is invalid.
// Returns ErrNoMethodFound if the method passed is invalid.
// Returns ErrPermissionDenied if the user does not have access to the route.
func (g RbacGroup) Enforce(group, method string) error {
	const op = "Permissions.Enforce"

	p, ok := g[group]
	if !ok {
		return ErrNoGroupFound
	}

	m, ok := p[method]
	if !ok {
		return ErrNoMethodFound
	}

	if !m.Allow {
		return ErrPermissionDenied
	}

	return nil
}

var PermissionMap = Rbac{
	// Contributor Permissions
	ContributorRoleID: RbacGroup{
		Permissions.Posts: {
			ViewMethod: {Allow: true},
			CreateMethod: {Allow: true},
			UpdateMethod: {Allow: true},
			DeleteMethod: {Allow: false},
		},
		Permissions.Categories: {
			ViewMethod: {Allow: true},
			CreateMethod: {Allow: false},
			UpdateMethod: {Allow: false},
			DeleteMethod: {Allow: false},
		},
		Permissions.Media: {
			ViewMethod: {Allow: true},
			CreateMethod: {Allow: true},
			UpdateMethod: {Allow: true},
			DeleteMethod: {Allow: false},
		},
		Permissions.Integrations: {
			ViewMethod: {Allow: false},
			CreateMethod: {Allow: false},
			UpdateMethod: {Allow: false},
			DeleteMethod: {Allow: false},
		},
		Permissions.Forms: {
			ViewMethod: {Allow: false},
			CreateMethod: {Allow: false},
			UpdateMethod: {Allow: false},
			DeleteMethod: {Allow: false},
		},
		Permissions.Users: {
			ViewMethod: {Allow: false},
			CreateMethod: {Allow: false},
			UpdateMethod: {Allow: false},
			DeleteMethod: {Allow: false},
		},
		Permissions.Settings: {
			ViewMethod: {Allow: false},
			CreateMethod: {Allow: false},
			UpdateMethod: {Allow: false},
			DeleteMethod: {Allow: false},
		},
		Permissions.System: {
			ViewMethod: {Allow: false},
			CreateMethod: {Allow: false},
			UpdateMethod: {Allow: false},
			DeleteMethod: {Allow: false},
		},
	},
	// Author Permissions
	AuthorRoleID: RbacGroup{
		Permissions.Settings: {
			ViewMethod: {Allow: false},
			// etc
		},
	},
	// Owner Permissions
	OwnerRoleID: RbacGroup{
		Permissions.Settings: {
			ViewMethod: {Allow: true},
			// etc
		},
	},
}
