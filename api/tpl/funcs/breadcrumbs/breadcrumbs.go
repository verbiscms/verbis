// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package breadcrumbs

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/tpl/embedded"
	"html/template"
)

const (
	// TemplateName is the name of the file containing
	// the markup of breadcrumbs
	TemplateName = "breadcrumbs"
	// EmbeddedExtension is extension of the embedded
	// template files.
	EmbeddedExtension = ".cms"
)

// Get
//
// Obtains the breadcrumbs for the post.
//
// Example: {{ $crumbs := breadcrumbs }}
func (ns *Namespace) Get(args ...interface{}) interface{} {
	const op = "Templates.Breadcrumbs"

	if len(args) == 0 {
		return ns.crumbs
	}

	if args[0] == true {
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
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error parsing template", Operation: op, Err: err}).Error()
		}

		return template.HTML(b.String())
	}

	return ns.crumbs
}
