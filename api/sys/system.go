// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/version"
	sm "github.com/hashicorp/go-version"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
	"runtime"
)

// System represents cor functions for interacting with
// Verbis.
type System interface {
	Restart() error
	Update(restart bool) (string, error)
	LatestVersion() string
	HasUpdate() bool
}

// Sys defines the base and core functionality for Verbis,
// such as restarting and updating the system.
type Sys struct {
	// The path of the current executable.
	ExecutablePath string
	Driver         database.Driver
	updater        *updater.Updater
	version        *sm.Version
}

// New creates a new system type, used for restarting
// and manipulating the system.
func New(db database.Driver) *Sys {
	exec, err := execPath()
	if err != nil {
		logger.Panic(err)
	}

	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: api.Repo,
		},
		Version: version.String(),
	}

	s := &Sys{
		Driver:         db,
		ExecutablePath: exec,
		updater:        u,
		version:        version.SemVer,
	}

	s.updater.Provider = &provider.Github{
		RepositoryURL: api.Repo,
		ArchiveName:   fmt.Sprintf("verbis_%s_%s_%s.zip", s.LatestVersion(), runtime.GOOS, runtime.GOARCH),
	}

	return s

}
