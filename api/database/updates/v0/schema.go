// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v0

import (
	"github.com/verbiscms/verbis/api/database/internal"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/version"
	"path/filepath"
)

// Version is the base version of the migration relative
// to the filepath.
const Version = "v0"

// init adds the migration to the updater.
func init() {
	err := internal.AddMigration(&internal.Migration{
		Version:      "v0.0.0",
		CallBackUp:   nil,
		CallBackDown: nil,
		Stage:        version.Major,
		SQLPath:      filepath.Join(Version, "mysql_schema.sql"),
		PostgresPath: filepath.Join(Version, "postgres_schema.sql"),
	})
	if err != nil {
		logger.Panic(err)
	}
}
