// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
)

// Version is  The main version number that is being run
// at the moment.
var Version = "0.0.0"

// Prerelease A pre-release marker for the version. If this is ""
// (empty string) then it means that it is a final release.
// Otherwise, this is a pre-release such as "dev".
var Prerelease = ""

// SemVer is an instance of version.Version.
var SemVer *version.Version

func init() {
	v, err := version.NewVersion(Version)
	if err != nil {
		log.Fatal(err)
	}
	SemVer = v
	api.App.Version = String()
}

// Header is the header name used to send the current
// verbis version in http requests.
const Header = "Verbis-Version"

// String returns the complete version string, including
// prerelease,
func String() string {
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s", Version, Prerelease)
	}
	return Version
}

// Must is an alias for version.Must
func Must(v string) *version.Version {
	return version.Must(version.NewVersion(v))
}
