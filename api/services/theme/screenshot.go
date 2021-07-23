// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"github.com/verbiscms/verbis/api/common/mime"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Screenshot
//
// Finds a screenshot in the theme directory based on the
// theme passed (e.g. verbis) and the file passed
// (e.g. screenshot.png).
//
// Returns errors.NOTFOUND if there was not screenshot found.
func (t *theme) Screenshot(theme, file string) ([]byte, domain.Mime, error) {
	const op = "SiteRepository.Screenshot"

	filePath := t.themesPath + string(os.PathSeparator) + theme + string(os.PathSeparator) + file
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, "", &errors.Error{Code: errors.NOTFOUND, Message: "Error finding screenshot with the path " + file, Operation: op, Err: err}
	}

	return b, domain.Mime(mime.TypeByExtension(filepath.Ext(file))), nil
}
