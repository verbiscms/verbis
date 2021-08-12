// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package breadcrumbs

import (
	"bytes"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/tpl/embedded"
	"github.com/verbiscms/verbis/api/verbis"
	"html/template"
)

var (
	// TemplateName is the name of the file containing
	// the markup of breadcrumbs
	TemplateName = "breadcrumbs"
	// EmbeddedExtension is extension of the embedded
	// template files.
	EmbeddedExtension = ".cms"
)

// Get
//
// Obtains the breadcrumbs for the post in a struct
// verbis.Breadcrumbs
//
// Example: {{ $crumbs := breadcrumbs }}
func (ns *Namespace) Get() verbis.Breadcrumbs {
	return ns.crumbs
}

// HTML
//
// Returns the breadcrumbs already constructed as
// HTML data.
//
// Example: {{ $crumbs := breadcrumbsHTML }}
func (ns *Namespace) HTML() template.HTML {
	const op = "Templates.Breadcrumbs"

	path := TemplateName + EmbeddedExtension

	file, err := embedded.Static.ReadFile(path)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error reading static file: " + path, Operation: op, Err: err}).Error()
		return ""
	}

	var b bytes.Buffer
	err = ns.deps.Tmpl().ExecuteTpl(&b, string(file), map[string]interface{}{
		"breadcrumbs": ns.crumbs,
	})

	if err != nil {
		logger.WithError(err).Error()
	}

	return template.HTML(b.String())
}
