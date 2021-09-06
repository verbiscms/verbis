// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"testing"
)

func TestLocal_Dial(t *testing.T) {
	UtilTestProviderDial(&environment.Env{}, &local{}, t)
}

func TestLocal_Dial_Directory(t *testing.T) {
	l := local{path: t.TempDir()}
	_, err := l.Dial(&environment.Env{})
	assert.Nil(t, err)
}

func TestLocal_Dial_Directory_Error(t *testing.T) {
	l := local{path: "/test"}
	_, err := l.Dial(&environment.Env{})
	assert.Error(t, err)
}

func TestLocal_Info(t *testing.T) {
	l := local{}
	got := l.Info(nil)
	want := domain.StorageProviderInfo{
		Order:          1,
		Name:           LocalName,
		Connected:      true,
		Error:          false,
		EnvironmentSet: true,
	}
	assert.Equal(t, want, got)
}
