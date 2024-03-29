// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	rocket "github.com/mouuff/go-rocket-update/pkg/updater"
	"github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/version"
	"runtime"
	"strconv"
	"time"
)

// LatestVersion obtains the latest remote version from
// Github. The function panics if it encountered
// and error obtaining the version.
func (s *Sys) LatestVersion() string {
	const op = "System.LatestVersion"
	remote, err := s.client.GetLatestVersion()
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error obtaining remote version", Operation: op, Err: err}).Panic()
	}
	return remote
}

// HasUpdate determines if there is an update available
// from GitHub. The function panics if it encountered
// and error obtaining the version.
func (s *Sys) HasUpdate() bool {
	remote := version.Must(s.LatestVersion())
	return s.version.LessThan(remote)
}

// Update updates the Verbis executable and runs any DB
// migrations. It calls Restart() upon a successful
// update.
// Returns errors.INVALID if Verbis is already up to date.
// Returns errors.INTERNAL if it could not be updated.
func (s *Sys) Update(restart bool) (string, error) {
	const op = "System.Update"

	if !s.Installed || s.Driver == nil {
		return "", &errors.Error{Code: errors.CONFLICT, Message: "Verbis not installed, cannot update", Operation: op, Err: fmt.Errorf("verbis not installed")}
	}

	ver := s.LatestVersion()

	logger.Info("Attempting to update Verbis to version: " + ver)
	logger.Info("Updating executable")

	s.client.Provider = &provider.Github{
		RepositoryURL: api.Repo,
		ArchiveName:   fmt.Sprintf("verbis_%s_%s_%s.zip", s.LatestVersion(), runtime.GOOS, runtime.GOARCH),
	}

	code, err := s.client.Update()
	if err != nil {
		logger.Error("Executable updated failed, rolling back")
		err := s.client.Rollback()
		if err != nil {
			logger.Panic(err)
		}
		switch code {
		case rocket.UpToDate:
			return "", &errors.Error{Code: errors.INVALID, Message: "Verbis is up to date", Operation: op, Err: err}
		default:
			return "", &errors.Error{Code: errors.INTERNAL, Message: "Error updating Verbis with status code: " + strconv.Itoa(int(code)), Operation: op, Err: err}
		}
	}

	logger.Info("Updating database")

	err = s.Driver.Migrate(version.SemVer)
	if err != nil {
		logger.Error("Database updated failed, rolling back")
		rollBackErr := s.client.Rollback()
		if rollBackErr != nil {
			logger.Panic(err)
		}
		return "", err
	}

	logger.Info("Successfully updated Verbis, restarting system...")

	if restart {
		go func() {
			time.Sleep(1 * time.Second)
			err := s.Restart()
			if err != nil {
				// TODO: Send callback to webhook.
				logger.WithError(err)
			}
		}()
	}

	return ver, nil
}
