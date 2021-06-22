// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"github.com/ainsleyclark/updater"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/version"
)

// System represents cor functions for interacting with
// Verbis.
type System interface {
	Restart() error
	Update() (string, error)
	LatestVersion() string
	HasUpdate() bool
}

// Sys defines the base and core functionality for Verbis,
// such as restarting and updating the system.
type Sys struct {
	// The path of the current executable.
	ExecutablePath string
	updater        updater.Patcher
}

// New creates a new system type, used for restarting
// and manipulating the system.
func New(db database.Driver) *Sys {
	const op = "System.New"

	exec, err := execPath()
	if err != nil {
		logger.Panic(err)
	}

	u, err := updater.New(updater.Options{
		GithubURL: "https://github.com/ainsleyclark/verbis", // The URL of the Git Repos
		Version:   version.Version,                          // The currently running version
		Verify:    false,                                    // Updates will be verified by checking the new exec with -version
		DB:        db.DB().DB,                               // Pass in an sql.DB for a migration
	})

	if err != nil {
		logger.Panic(&errors.Error{Code: errors.INTERNAL, Message: "Error creating new Verbis updater", Operation: op, Err: err})
	}

	return &Sys{
		ExecutablePath: exec,
		updater:        u,
	}
}
