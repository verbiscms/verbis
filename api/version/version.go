// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/hashicorp/go-version"
)

// Version is  The main version number that is being run
// at the moment.
var Version = "v0.0.1"

// Prerelease A pre-release marker for the version. If this is ""
// (empty string) then it means that it is a final release.
// Otherwise, this is a pre-release such as "dev".
var Prerelease = ""

// SemVer is an instance of version.Version.
var SemVer *version.Version

func init() {
	SemVer = version.Must(version.NewVersion(Version))
	api.App.Version = String()
}

// Header is the header name used to send the current
// verbis version in http requests.
// TODO, Implement in Responses
const Header = "Verbis-Version"

// String returns the complete version string, including
// prerelease,
func String() string {
	if Prerelease != "" { //nolint
		return fmt.Sprintf("%s-%s", Version, Prerelease)
	}
	return Version
}
