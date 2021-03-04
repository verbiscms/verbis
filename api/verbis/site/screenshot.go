// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Screenshot
//
// TODO
//
func (s *Site) Screenshot(theme, file string) ([]byte, string, error) {
	const op = "SiteRepository.Screenshot"

	filePath := s.theme + string(os.PathSeparator) + theme + string(os.PathSeparator) + file
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, "", &errors.Error{Code: errors.NOTFOUND, Message: "Error finding screenshot with the path " + file, Operation: op, Err: err}
	}

	return b, mime.TypeByExtension(filepath.Ext(file)), nil
}