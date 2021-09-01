// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_HidePassword(t *testing.T) {
	u := User{Password: "password"}
	u.HidePassword()
	assert.Equal(t, "", u.Password)
}

func TestUser_HideCredentials(t *testing.T) {
	up := UserPart{FirstName: "verbis"}
	u := User{Password: "password", UserPart: up}
	assert.Equal(t, up, u.HideCredentials())
}

func TestUsers_HideCredentials(t *testing.T) {
	up := UserPart{FirstName: "verbis"}
	u := Users{{Password: "password", UserPart: up}}
	want := UsersParts{up}
	assert.Equal(t, want, u.HideCredentials())
}

func TestUser_AssignPermissions(t *testing.T) {
	u := User{UserPart: UserPart{FirstName: "verbis", Role: Role{ID: OwnerRoleID}}}
	got := u.AssignPermissions()
	want := User{UserPart: UserPart{FirstName: "verbis", Role: Role{ID: OwnerRoleID, Permissions: Permissions[OwnerRoleID]}}}
	assert.Equal(t, want, got)
}

func TestUsers_AssignPermissions(t *testing.T) {
	u := Users{
		User{UserPart: UserPart{FirstName: "verbis", Role: Role{ID: OwnerRoleID}}},
		User{UserPart: UserPart{FirstName: "verbis", Role: Role{ID: AuthorRoleID}}},
	}
	got := u.AssignPermissions()
	want := Users{
		User{UserPart: UserPart{FirstName: "verbis", Role: Role{ID: OwnerRoleID, Permissions: Permissions[OwnerRoleID]}}},
		User{UserPart: UserPart{FirstName: "verbis", Role: Role{ID: AuthorRoleID, Permissions: Permissions[AuthorRoleID]}}},
	}
	assert.Equal(t, want, got)
}
