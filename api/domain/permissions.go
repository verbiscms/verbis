// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

type RBAC map[int]Group

type Request map[string]Permission

type Group map[string]Request

type Permission struct {
	Allow bool
	// TBC
}

const (
	ListMethod   = "list"
	FindMethod   = "find"
	CreateMethod = "create"
	UpdateMethod = "update"
	DeleteMethod = "delete"
)

var Permissions = RBAC{
	// Author Permissions
	AuthorRoleID: Group{
		"settings": {
			ListMethod: {Allow: false},
			// etc
		},
	},
	// Owner Permissions
	OwnerRoleID: Group{
		"settings": {
			ListMethod: {Allow: true},
			// etc
		},
	},
}
