// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"github.com/verbiscms/verbis/api/errors"
	"os"
	"path/filepath"
)

const (
	// ScreenshotName is the default screenshot name within
	// the theme's directory.
	ScreenshotName = "screenshot"
	// ScreenshotURL is the url of the screenshot.
	ScreenshotURL = "/themes/"
)

var (
	// ScreenshotExtensions is the allowed extensions
	// that the function will scan for.
	ScreenshotExtensions = []string{
		".png",
		".svg",
		".jpg",
		".jpeg",
	}
)

// findScreenshot ranges over the allowed screenshot
// extensions and checks for a match, if the
// screenshot has been found, a url will
// be returned.
// Returns errors.NOTFOUND if no screenshot was found.
func findScreenshot(path string) (string, error) {
	const op = "Theme.findScreenshot"

	for _, v := range ScreenshotExtensions {
		name := path + string(os.PathSeparator) + ScreenshotName + v
		info, err := os.Stat(name)
		if err != nil {
			continue
		}
		return ScreenshotURL + filepath.Base(path) + "/" + info.Name(), nil
	}

	return "", &errors.Error{Code: errors.NOTFOUND, Message: "No screenshot found from the theme", Operation: op, Err: fmt.Errorf("no theme screenshot found")}
}
