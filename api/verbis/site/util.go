// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"os"
	"path/filepath"
	"strings"
)

// walkMatch
//
// Walk through root and return array of strings
// to the file path.
func (s *Site) walkMatch(root, pattern string) ([]string, error) {
	const op = "SiteRepository.walkMatch"

	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		matched, err := filepath.Match(pattern, filepath.Base(path))
		if err != nil {
			return err
		} else if matched {
			template := strings.Replace(path, root+"/", "", 1)
			template = strings.Replace(template, s.config.FileExtension, "", -1)
			matches = append(matches, template)
		}
		return nil
	})

	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Unable to find page templates", Err: err, Operation: op}
	}

	return matches, nil
}

// fileName
//
// Cleans the file name to a friendly string for
// page templates and layouts.
func (s *Site) fileName(file string) string {
	return strings.Title(strings.ToLower(strings.ReplaceAll(file, "-", " ")))
}
