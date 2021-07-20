// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocal_Dial(t *testing.T) {
	UtilTestProviderDial(&environment.Env{}, &local{}, t)
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
