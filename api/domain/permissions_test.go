// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRBAC_Get(t *testing.T) {
	rbac := Rbac{OwnerRoleID: RbacGroup{PermissionSettings: {ViewMethod: {Allow: false}}}}

	tt := map[string]struct {
		input int
		want  interface{}
	}{
		"Found": {
			OwnerRoleID,
			RbacGroup{PermissionSettings: {ViewMethod: {Allow: false}}},
		},
		"Not Found": {
			BannedRoleID,
			"no permission group found",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := rbac.Get(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestPermissionGroup_Enforce(t *testing.T) {
	tt := map[string]struct {
		input  RbacGroup
		group  string
		method string
		want   interface{}
	}{
		"No Group": {
			RbacGroup{},
			PermissionSettings,
			ViewMethod,
			ErrNoGroupFound,
		},
		"No Method": {
			RbacGroup{
				PermissionSettings: {},
			},
			PermissionSettings,
			ViewMethod,
			ErrNoMethodFound,
		},
		"Not Allowed": {
			RbacGroup{
				PermissionSettings: {
					ViewMethod: Permission{Allow: false},
				},
			},
			PermissionSettings,
			ViewMethod,
			ErrPermissionDenied,
		},
		"Allowed": {
			RbacGroup{
				PermissionSettings: {
					ViewMethod: Permission{Allow: true},
				},
			},
			PermissionSettings,
			ViewMethod,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Enforce(test.group, test.method)
			assert.Equal(t, test.want, got)
		})
	}
}
