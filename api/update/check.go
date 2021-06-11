// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package update

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Updater interface {
	Update() (int64, error)
	HasUpdate() bool
	Validate() error
	RollBack() error
}

type Update struct {
	Paths paths.Paths
	pkg   *updater.Updater
}

var (
	ErrLatestVersion = errors.New("at latest version")
	ZipDir           = "verbis/build/"
	Folders          = []string{
		paths.Admin,
		paths.API,
	}
)

// UpdateStatus represents the status after Update{}.Update() was called
type UpdateStatus int

const (
	// Unknown update status (something went wrong)
	Unknown UpdateStatus = iota
	// UpToDate means the Verbis is already up to date
	UpToDate = 1
	// Validation failed means V
	ValidationFailed = 2
	// Updated means the software have been updated
	Updated = 3
)

func New(paths paths.Paths) *Update {
	u := &Update{
		Paths: paths,
	}

	u.pkg = &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: "github.com/" + api.Repo,
			ArchiveName:   u.zipFile(),
		},
		Version: "v0.0.0",
	}

	return u
}

// Update - TODO
//
//
func (u *Update) Update() (UpdateStatus, int64, error) {
	const op = "Update.Update"

	err := u.Validate()
	if err == ErrLatestVersion {
		return UpToDate, 0, err
	} else if err != nil {
		return ValidationFailed, 0, err
	}

	_, err = u.pkg.Update()
	if err != nil {
		err := u.RollBack()
		if err != nil {
			return 0, 0, &errors.Error{Code: errors.INTERNAL, Message: "Error updating executable", Operation: op, Err: err}
		}
		return Unknown, 0, &errors.Error{Code: errors.INTERNAL, Message: "Error updating executable", Operation: op, Err: err}
	}

	err = u.backup()
	if err != nil {
		return Unknown, 0, &errors.Error{Code: errors.INTERNAL, Message: "Error backing up Verbis", Operation: op, Err: err}
	}

	fileCount, err := u.walk()
	if err != nil {
		return Unknown, 0, &errors.Error{Code: errors.INTERNAL, Message: "Error copying folders", Operation: op, Err: err}
	}

	err = u.cleanup()
	if err != nil {
		return 0, 0, err
	}

	return Updated, fileCount, err
}

// Validate determines if Verbis can be updated by
// performing health checks. If the current user
// is a super admin or Verbis does not require
// and update, errors.INVALID will be
// returned.
func (u *Update) Validate() error {
	const op = "Update.Validate"

	// Check if the user is not a super admin
	if api.SuperAdmin {
		return &errors.Error{Code: errors.INVALID, Message: "Error updating Verbis", Operation: op, Err: fmt.Errorf("only built verbis installations can be updated")}
	}

	// Check if the current version is up to date
	if !u.HasUpdate() {
		return &errors.Error{Code: errors.INVALID, Message: "Version number: " + version.Version + " not required to be updated", Operation: op, Err: ErrLatestVersion}
	}

	return nil
}

// HasUpdate determines if Verbis has an update to be run.
// This is done by comparing the version number on
// the latest GitHub release.
// Logs errors.INVALID if there was an error obtaining
// the version number.
func (u *Update) HasUpdate() bool {
	const op = "Update.HasUpdate"

	update, err := u.pkg.CanUpdate()
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Error obtaining version number", Operation: op, Err: err})
		return false
	}

	return update
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
// Rollback
func (u *Update) RollBack() error {
	const op = "Update.RollBack"

	return nil

	//_ = u.deleteFolders()
	//
	//for _, v := range Folders {
	//	err := files.CopyDir(u.backupPath+v, u.Paths.Base)
	//	if err != nil {
	//		return u.backupPath, &errors.Error{Code: errors.INTERNAL, Message: "Error copying folder: " + v, Operation: op, Err: err}
	//	}
	//}
	//
	//err := u.pkg.Rollback()
	//if err != nil {
	//	return u.backupPath, &errors.Error{Code: errors.INTERNAL, Message: "Error copying executable", Operation: op, Err: err}
	//}
	//
	//return u.backupPath, nil
}

// zipFile retrieves the release zip from GitHub by using
// the runtime GOOS and GOARCH.
// Example zip: verbis_0.0.1_darwin_amd64.zip
func (u *Update) zipFile() string {
	return fmt.Sprintf("verbis_%s_%s_%s.zip", version.Version, runtime.GOOS, runtime.GOARCH)
}

//
func (u *Update) walk() (int64, error) {
	const op = "Provider.Walk"
	err := u.pkg.Provider.Open()
	if err != nil {
		return 0, &errors.Error{Code: errors.INTERNAL, Message: "Error opening github Provider", Operation: op, Err: err}
	}

	defer func(Provider provider.Provider) {
		err := Provider.Close()
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error closing update provider", Operation: op, Err: err})
		}
	}(u.pkg.Provider)

	var filesCount int64 = 0
	err = u.pkg.Provider.Walk(func(info *provider.FileInfo) error {

		if strings.Contains(info.Path, "verbisexec") {
			fmt.Println(info.Path)
		}

		cleaned := strings.Replace(info.Path, ZipDir, "", 1)
		parts := strings.Split(cleaned, "/")

		if len(parts) == 0 {
			return nil
		}

		for _, v := range Folders {
			if parts[0] != strings.ReplaceAll(v, string(os.PathSeparator), "") {
				continue
			}
			destPath := u.Paths.Base + string(os.PathSeparator) + cleaned
			err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
			if err != nil {
				return err
			}
			err = u.pkg.Provider.Retrieve(info.Path, destPath)
			if err != nil {
				return err
			}
			filesCount += 1
		}
		return nil
	})

	return filesCount, nil
}

// verifyUpdate verifies if the executable is installed
// correctly we are going to run the newly installed
// program by running it with -version.
func (u *Update) verifyUpdate() error {
	latestVersion, err := u.pkg.GetLatestVersion()
	if err != nil {
		return err
	}

	executable, err := u.pkg.GetExecutable()
	if err != nil {
		return err
	}

	cmd := exec.Cmd{
		Path: executable,
		Args: []string{executable, "-version"},
	}

	// Should be replaced with Output() as soon as test project is updated
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	strOutput := string(output)
	if !strings.Contains(strOutput, latestVersion) {
		return errors.New("Version not found in program output")
	}

	return nil
}
