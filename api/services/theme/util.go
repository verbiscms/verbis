// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"github.com/verbiscms/verbis/api/common/files"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"os"
	"path/filepath"
	"strings"
)

// walkMatch walks through root and return array of strings
// to the file path.
func (t *Theme) walkMatch(root, pattern string) ([]string, error) {
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
			matches = append(matches, files.RemoveFileExtension(template))
		}
		return nil
	})

	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Unable to find files.", Err: err, Operation: op}
	}

	return matches, nil
}

// fileName Cleans the file name to a friendly string for
// page templates and layouts.
func (t *Theme) fileName(file string) string {
	return strings.Title(strings.ToLower(strings.ReplaceAll(file, "-", " ")))
}

// getActive theme retrieves the active theme name, and
// configuration for that theme.
func (t *Theme) getActiveTheme() (string, domain.ThemeConfig, error) {
	theme, err := t.options.GetTheme()
	if err != nil {
		return "", domain.ThemeConfig{}, err
	}

	cfg, err := t.Config()
	if err != nil {
		return "", domain.ThemeConfig{}, err
	}

	return filepath.Join(t.themesPath, theme), cfg, nil
}
