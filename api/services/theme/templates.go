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

// Templates satisfies the Theme interface by retrieving all
// templates from the active theme.
func (t *Theme) Templates() (domain.Templates, error) {
	const op = "SiteRepository.Templates"

	path, cfg, err := t.getActiveTheme()
	if err != nil {
		return nil, err
	}

	tplDir := filepath.Join(path, cfg.TemplateDir)
	files, err := t.walkMatch(tplDir, "*"+cfg.FileExtension)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error getting templates with the path: %s", tplDir), Operation: op, Err: ErrNoTemplates}
	}

	var templates domain.Templates
	for _, file := range files {
		templates = append(templates, domain.Template{
			Key:  file,
			Name: t.fileName(file),
		})
	}

	if len(templates) == 0 {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No templates available", Operation: op, Err: ErrNoTemplates}
	}

	return templates, nil
}
