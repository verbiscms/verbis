// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/verbiscms/verbis/api/environment"
	"testing"
)

func TestAzure(t *testing.T) {
	UtilTestProvider(&environment.Env{
		AzureKey:     "key",
		AzureAccount: "account",
	}, &azure{}, t)
}
