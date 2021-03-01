// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/tpl/embedded"
	"html/template"
	"strings"
)

const (
	// The extension of the embedded files.
	EmbeddedExtension = ".cms"
)

// TemplateMeta defines the helper for executing meta
// templates.
type TemplateMeta struct {
	Site          domain.Site
	Post          *domain.PostDatum
	Options       domain.Options
	FacebookImage string
	TwitterImage  string
	deps          *deps.Deps
}

// GetImage
//
// Is a helper function for the embedded meta templates.
// Returns an media item URL or an empty string if
// the media item did not exist.
func (tm *TemplateMeta) GetImage(id int) string {
	img, err := tm.deps.Store.Media.GetById(id)
	if err != nil {
		return ""
	}
	// TODO: This should be dynamic, not dependant on Site URL.
	return tm.deps.Site.Url + img.Url
}

// Header
//
// Header obtains all of the site and post wide Code Injection
// as well as any meta information from the page.
//
// Example: {{ verbisHead }}
func (ns *Namespace) Header() template.HTML {
	tm := &TemplateMeta{
		Site:    ns.deps.Site,
		Post:    ns.post,
		Options: *ns.deps.Options,
		deps:    ns.deps,
	}

	head := ns.executeTemplates(tm, []string{"meta", "opengraph", "twitter"})

	return template.HTML(head)
}

// MetaTitle
//
// metaTitle obtains the meta title from the post, if there is no
// title set on the post, it will look for the global title, if
// none, return empty string.
//
// Example: <title>Verbis - {{ metaTitle }}</title>
func (ns *Namespace) MetaTitle() string {
	postMeta := ns.post.SeoMeta.Meta

	if postMeta == nil {
		return ""
	}

	if postMeta.Title != "" {
		return postMeta.Title
	}

	if ns.deps.Options.MetaTitle != "" {
		return ns.deps.Options.MetaTitle
	}

	return ""
}

// Footer
//
// Obtains all of the site and post wide Code Injection
// Returns formatted HTML template for use after the
// closing `</body>`.
//
// Example: {{ verbisFoot }}
func (ns *Namespace) Footer() template.HTML {
	tm := &TemplateMeta{
		Post:    ns.post,
		Options: *ns.deps.Options,
	}

	foot := ns.executeTemplates(tm, []string{"footer"})

	return template.HTML(foot)
}

// executeTemplates
//
// Ranges over the templates passed and executes the embedded
// templates, logs if an error occurred or concatenates
// the meta and returns a string upon successful
// execution.
func (ns *Namespace) executeTemplates(tm *TemplateMeta, templates []string) string {
	const op = "Templates.Meta.executeTemplates"

	meta := ""
	for _, name := range templates {

		path := name + EmbeddedExtension
		file, err := embedded.Static.ReadFile(path)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error reading static file: " + path, Operation: op, Err: err}).Error()
		}

		var b bytes.Buffer
		err = ns.deps.Tmpl().ExecuteTpl(&b, string(file), tm)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error parsing template", Operation: op, Err: err}).Error()
		}

		meta += fmt.Sprintf("%s\n", b.String())
	}

	return strings.TrimRight(meta, "\n")
}
