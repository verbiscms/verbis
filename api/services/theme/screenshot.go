// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"fmt"
	"github.com/verbiscms/verbis/api/common/mime"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// ScreenshotName is the default screenshot name within
	// the Theme's directory.
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

// Screenshot
//
// Finds a screenshot in the Theme directory based on the
// Theme passed (e.g. verbis) and the file passed
// (e.g. screenshot.png).
//
// Returns errors.NOTFOUND if there was not screenshot found.
func (t *Theme) Screenshot(theme, file string) ([]byte, domain.Mime, error) {
	const op = "SiteRepository.Screenshot"

	filePath := t.themesPath + string(os.PathSeparator) + theme + string(os.PathSeparator) + file
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, "", &errors.Error{Code: errors.NOTFOUND, Message: "Error finding screenshot with the path " + file, Operation: op, Err: err}
	}

	return b, domain.Mime(mime.TypeByExtension(filepath.Ext(file))), nil
}

// findScreenshot ranges over the allowed screenshot
// extensions and checks for a match, if the
// screenshot has been found, a url will
// be returned.
// Returns errors.NOTFOUND if no screenshot was found.
func findScreenshot(path string) (string, error) {
	const op = "Theme.FindScreenshot"

	for _, v := range ScreenshotExtensions {
		name := path + string(os.PathSeparator) + ScreenshotName + v
		info, err := os.Stat(name)
		if err != nil {
			continue
		}
		return ScreenshotURL + filepath.Base(path) + "/" + info.Name(), nil
	}

	return "", &errors.Error{Code: errors.NOTFOUND, Message: "No screenshot found from the Theme", Operation: op, Err: fmt.Errorf("no Theme screenshot found")}
}
