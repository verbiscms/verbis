// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/environment"
	"testing"
)

func TestGCP(t *testing.T) {
	json := "json"
	UtilTestProvider(&environment.Env{
		GCPJson:      "json",
		GCPProjectId: "secret",
	}, &gcp{json: &json}, t)
}
