// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

const (
	// InstallDatabaseStep is the step for validating the
	// database for installing the application.
	InstallDatabaseStep = iota + 1
	// InstallUserStep is the step for validating a new user
	// when installing the application.
	InstallUserStep
	// InstallSiteStep is the step for validating site
	// properties. when installing the application.
	InstallSiteStep
)

// ValidateInstall validates the particular stage of
// the installation of the app. If the stag is
// InstallDatabaseStep the database driver will be obtained
// and checked for errors.
// Returns errors.INVALID if the validation failed or there was
// an invalid step provided.
func (s *Sys) ValidateInstall(step int, install domain.InstallVerbis) error {
	const op = "System.ValidateInstall"

	switch step {
	case InstallDatabaseStep:
		err := validation.Validator().Struct(install.InstallDatabase)
		if err != nil {
			return err
		}
		_, _, err = s.getDatabase(install.InstallDatabase)
		if err != nil {
			return err
		}
		return nil
	case InstallUserStep:
		return validation.Validator().Struct(install.InstallUser)
	case InstallSiteStep:
		return validation.Validator().Struct(install.InstallSite)
	default:
		return &errors.Error{Code: errors.INVALID, Message: "Error validating install", Operation: op, Err: fmt.Errorf("invalid step provided: %d", step)}
	}
}
