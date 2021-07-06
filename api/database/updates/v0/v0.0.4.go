// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v0

import (
	"github.com/ainsleyclark/verbis/api/database/internal"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/version"
	"path/filepath"
)

// init adds the migration to the updater.
func init() {
	err := internal.AddMigration(&internal.Migration{
		Version:      "v0.0.4",
		CallBackUp:   nil,
		CallBackDown: nil,
		Stage:        version.Patch,
		SQLPath:      filepath.Join(Version, "v0.0.4.sql"),
	})
	if err != nil {
		logger.Panic(err)
	}
}
