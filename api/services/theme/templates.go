// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"fmt"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"os"
	"strings"
)

// Templates
//
// Retrieves all templates stored within the templates
// directory of the theme path.
//
// Returns ErrNoTemplates in any error case.
// Returns errors.NOTFOUND if no templates were found.
// Returns errors.INTERNAL if the template path is invalid.
func (t *theme) Templates(theme string) (domain.Templates, error) {
	const op = "SiteRepository.Templates"

	tplDir := t.themesPath + string(os.PathSeparator) + theme + string(os.PathSeparator) + t.config.TemplateDir
	tplDir = strings.ReplaceAll(tplDir, "//", "/")

	files, err := t.walkMatch(tplDir, "*"+t.config.FileExtension)
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
