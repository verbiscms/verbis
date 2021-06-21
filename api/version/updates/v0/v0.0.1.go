// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v0

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/version/internal"
	"github.com/gookit/color"
)

// init adds v0.0.2 add to the registry.
func init() {
	internal.Updates.AddUpdate(&internal.Update{
		Version:       "v0.0.2",
		MigrationPath: "v0.0.2.sql",
		CallBackUp: func() error {
			color.Red.Println("in v0.0.2 callback up")
			return fmt.Errorf("error here")
		},
		CallBackDown: func() error {
			color.Red.Println("in v0.0.2 callback down")
			return nil
		},
		Stage: internal.Minor,
	})
}

// init adds v0.0.1 add to the registry.
func init() {
	internal.Updates.AddUpdate(&internal.Update{
		Version:       "v0.0.1",
		MigrationPath: "v0.0.1.sql",
		CallBackUp: func() error {
			color.Red.Println("in v0.0.1 callback up")
			return nil
		},
		CallBackDown: func() error {
			color.Red.Println("in v0.0.1 callback down")
			return nil
		},
		Stage: internal.Minor,
	})
}
