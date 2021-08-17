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
	"io/ioutil"
	"path/filepath"
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
// Returns a navigation menu with the arguments passed.
// If the menu does not exist, an error will occur.
//
// Example: {{ nav (dict "menu" "main-menu") }}
func (ns *Namespace) Get(args nav.Args) (nav.Menu, error) {
	return ns.nav.Get(args)
}

// HTML
//
// Returns a navigation menu already constructed as
// HTML data.
//
// Example: {{ navHTML (dict "menu" "main-menu") }}
func (ns *Namespace) HTML(args nav.Args) (template.HTML, error) {
	const op = "Templates.menusDB.HTML"

	menu, err := ns.nav.Get(args)
	if err != nil {
		return "", err
	}

	var (
		file   []byte
		tplErr error
	)

	if menu.Options.Partial != "" {
		file, tplErr = ioutil.ReadFile(filepath.Join(ns.themeGetter(), menu.Options.Partial))
	} else {
		file, tplErr = embedded.Static.ReadFile(TemplateName + EmbeddedExtension)
	}

	if tplErr != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Error reading file", Operation: op, Err: err}
	}

	var b bytes.Buffer
	err = ns.deps.Tmpl().ExecuteTpl(&b, string(file), menu)
	if err != nil {
		return "", err
	}

	return template.HTML(gohtml.Format(b.String())), nil
}
