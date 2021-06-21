// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"github.com/ainsleyclark/updater"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"runtime"
)

// LatestVersion obtains the latest remote version from
// Github. The function panics if it encountered
// and error obtaining the version.
func (s *Sys) LatestVersion() string {
	const op = "System.LatestVersion"
	hasUpdate, err := s.updater.LatestVersion()
	if err != nil {
		logger.Panic(&errors.Error{Code: errors.INTERNAL, Message: "Error obtaining remote version", Operation: op, Err: err})
	}
	return hasUpdate
}

// HasUpdate determines if there is an update available
// from GitHub. The function panics if it encountered
// and error obtaining the version.
func (s *Sys) HasUpdate() bool {
	const op = "System.HasUpdate"
	update, err := s.updater.HasUpdate()
	if err != nil {
		logger.Panic(&errors.Error{Code: errors.INTERNAL, Message: "Error obtaining remote version", Operation: op, Err: err})
	}
	return update
}

// Update updates the Verbis executable and runs any DB
// migrations. It calls Restart() upon a successful
// update.
//
// Returns errors.INVALID if Verbis is already up to date.
// Returns errors.INTERNAL if it could not be updated.
func (s *Sys) Update() (string, error) {
	const op = "System.Update"

	ver := s.LatestVersion()

	logger.Info("Attempting to update Verbis to version: " + ver)

	zip := fmt.Sprintf("verbis_%s_%s_%s.zip", s.LatestVersion(), runtime.GOOS, runtime.GOARCH)
	code, err := s.updater.Update(zip)
	if err != nil {
		switch code {
		case updater.UpToDate:
			return "", &errors.Error{Code: errors.INVALID, Message: "Verbis is up to date", Operation: op, Err: err}
		default:
			return "", &errors.Error{Code: errors.INTERNAL, Message: "Error updating Verbis", Operation: op, Err: err}
		}
	}

	logger.Info("Successfully updated Verbis, restarting system...")

	go func() {
		err := s.Restart()
		if err != nil {
			// TODO: Send callback to webhook.
			logger.WithError(err)
		}
	}()

	return ver, nil
}
