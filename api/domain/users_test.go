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
