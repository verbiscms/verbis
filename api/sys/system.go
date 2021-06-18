// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"github.com/ainsleyclark/updater"
	"github.com/ainsleyclark/verbis/api"
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
func New() *Sys {
	const op = "System.New"

	exec, err := execPath()
	if err != nil {
		logger.Panic(err)
	}

	u, err := updater.New(&updater.Options{
		RepositoryURL: api.Repo,
		Version:       version.Version,
	})
	if err != nil {
		logger.Panic(&errors.Error{Code: errors.INTERNAL, Message: "Error creating new Verbis updater", Operation: op, Err: err})
	}

	return &Sys{
		ExecutablePath: exec,
		updater:        u,
	}
}
