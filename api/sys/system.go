// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	sm "github.com/hashicorp/go-version"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	rocket "github.com/mouuff/go-rocket-update/pkg/updater"
	"github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/version"
)

// System represents cor functions for interacting with
// Verbis.
type System interface {
	Restart() error
	Update(restart bool) (string, error)
	LatestVersion() string
	HasUpdate() bool
	Installer
}

type Installer interface {
	Install(db domain.InstallVerbis) error
	ValidateInstall(step int, install domain.InstallVerbis) error
}

// Sys defines the base and core functionality for Verbis,
// such as restarting and updating the system.
type Sys struct {
	// The path of the current executable.
	ExecutablePath string
	Driver         database.Driver
	Installed      bool
	client         *rocket.Updater
	version        *sm.Version
}

// New creates a new system type, used for restarting
// and manipulating the system.
func New(db database.Driver, installed bool) *Sys {
	exec, err := execPath()
	if err != nil {
		logger.Panic(err)
	}

	u := &rocket.Updater{
		Provider: &provider.Github{
			RepositoryURL: api.Repo,
		},
		Version: version.String(),
	}

	s := &Sys{
		Driver:         db,
		ExecutablePath: exec,
		Installed:      installed,
		client:         u,
		version:        version.SemVer,
	}

	return s
}
