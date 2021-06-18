// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/ainsleyclark/verbis/api/version/internal"
	"path/filepath"
)

// init adds v0.0.1 add to the registry.
func init() {
	internal.Registry.AddUpdate(&internal.Update{
		Version:       "v0.0.1",
		MigrationPath: filepath.Join("v1", "v0.0.1.sql"),
		Callback: func() error {
			return nil
		},
		Stage: internal.Minor,
	})
}
