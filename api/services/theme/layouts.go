// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"fmt"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"path/filepath"
)

// Layouts satisfies the Theme interface by retrieving all
// template layouts from the active theme.
func (t *Theme) Layouts() (domain.Layouts, error) {
	const op = "SiteRepository.GetLayouts"

	path, cfg, err := t.getActiveTheme()
	if err != nil {
		return nil, err
	}

	layoutDir := filepath.Join(path, cfg.LayoutDir)
	files, err := t.walkMatch(layoutDir, "*"+cfg.FileExtension)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error getting layouts with the path: %s", layoutDir), Operation: op, Err: ErrNoLayouts}
	}

	var layouts domain.Layouts
	for _, file := range files {
		layouts = append(layouts, domain.Layout{
			Key:  file,
			Name: t.fileName(file),
		})
	}

	if len(layouts) == 0 {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No layouts available", Operation: op, Err: ErrNoLayouts}
	}

	return layouts, nil
}
