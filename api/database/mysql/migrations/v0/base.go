// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v0

import (
	"fmt"
	"github.com/ainsleyclark/updater"
	"github.com/ainsleyclark/verbis/api/database/mysql/migrations"
	"github.com/ainsleyclark/verbis/api/logger"
	"path/filepath"
)

// Version is the base version of the migration relative
// to the filepath.
const Version = "v0"

// init adds the migration to the updater.
func init() {
	file, err := migrations.Migrations.Open(filepath.Join(Version, "base.sql"))
	if err != nil {
		logger.Panic(err)
		return
	}
	defer file.Close()

	err = updater.AddMigration(&updater.Migration{
		Version:      "v0.0.0",
		SQL:          file,
		CallBackUp:   nil,
		CallBackDown: nil,
		Stage:        updater.Patch,
	})

	if err != nil {
		logger.Panic(err)
	}
}

// init adds the migration to the updater.
func init() {
	file, err := migrations.Migrations.Open(filepath.Join(Version, "v0.0.1.sql"))
	if err != nil {
		logger.Panic(err)
		return
	}
	defer file.Close()

	err = updater.AddMigration(&updater.Migration{
		Version: "v0.0.1",
		SQL:     file,
		CallBackUp: func() error {
			fmt.Println("in v0.0.1")
			return nil
		},
		CallBackDown: func() error { return nil },
		Stage:        updater.Patch,
	})

	if err != nil {
		logger.Panic(err)
	}
}
