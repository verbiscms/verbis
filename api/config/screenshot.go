// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"os"
)

const (
	// The default screenshot name within the theme's
	// directory.
	ScreenshotName = "screenshot"

	ScreenshotURL = "/themes/"
)

var (
	ScreenshotExtensions = []string{
		".png",
		".svg",
		".jpg",
	}
)

// findScreenshot
//
//
func findScreenshot(path string) (string, error) {
	const op = "Config.FindScreenshot"

	for _, v := range ScreenshotExtensions {
		name := path + string(os.PathSeparator) + ScreenshotName + v
		info, err := os.Stat(name)
		if err != nil {
			continue
		}
		return ScreenshotURL + info.Name(), nil
	}

	return "", &errors.Error{Code: errors.NOTFOUND, Message: "No screenshot found from the theme", Operation: op, Err: fmt.Errorf("no theme screenshot found")}
}
