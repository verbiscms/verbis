// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"bytes"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/tpl/embedded"
	"github.com/verbiscms/verbis/api/verbis/nav"
	"github.com/yosssi/gohtml"
	"html/template"
)

var (
	// TemplateName is the name of the file containing
	// the markup of breadcrumbs
	TemplateName = "nav"
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
func (ns *Namespace) Get(args nav.Args) (nav.Menu, error) {
	return ns.nav.Get(args)
}

// HTML
//
// Returns the breadcrumbs already constructed as
// HTML data.
//
// Example: {{ $crumbs := breadcrumbsHTML }}
func (ns *Namespace) HTML(args nav.Args) (template.HTML, error) {
	const op = "Templates.Nav.HTML"

	path := TemplateName + EmbeddedExtension

	file, err := embedded.Static.ReadFile(path)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Error reading static file: " + path, Operation: op, Err: err}
	}

	menu, err := ns.nav.Get(args)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = ns.deps.Tmpl().ExecuteTpl(&b, string(file), menu)

	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Error parsing template", Operation: op, Err: err}
	}

	return template.HTML(gohtml.Format(b.String())), nil
}
