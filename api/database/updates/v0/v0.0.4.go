// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v0

import (
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/database/internal"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/version"
	"io/fs"
	"os"
	"path/filepath"
)

// init adds the migration to the updater.
func init() {
	err := internal.AddMigration(&internal.Migration{
		Version: "v0.0.4",
		CallBackUp: func() error {
			err := filepath.Walk(paths.Get().Uploads, func(path string, info fs.FileInfo, err error) error {
				ext := filepath.Ext(path)
				if ext == ".webp" {
					err := os.Remove(path)
					if err != nil {
						logger.Trace(err)
					}
				}
				return nil
			})
			if err != nil {
				return err
			}
			return nil
		},
		CallBackDown: func() error {
			return nil
		},
		Stage:        version.Patch,
		SQLPath:      filepath.Join(Version, "v0.0.4.sql"),
	})
	if err != nil {
		logger.Panic(err)
	}
}
