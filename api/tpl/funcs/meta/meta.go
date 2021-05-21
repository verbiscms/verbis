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
	"github.com/ainsleyclark/verbis/api/verbis"
	"html/template"
	"regexp"
	"strings"
)

const (
	// EmbeddedExtension is extension of the embedded
	// template files.
	EmbeddedExtension = ".cms"
)

// TemplateMeta defines the helper for executing meta
// templates.
type TemplateMeta struct {
	Site          domain.Site
	Post          *domain.PostDatum
	Options       domain.Options
	Breadcrumbs   verbis.Breadcrumbs
	FacebookImage string
	TwitterImage  string
	deps          *deps.Deps
}

// GetImage
//
// Is a helper function for the embedded meta templates.
// Returns an media item url or an empty string if
// the media item did not exist.
func (tm *TemplateMeta) GetImage(id int) string {
	img, err := tm.deps.Store.Media.Find(id)
	if err != nil {
		return ""
	}
	// TODO: This should be dynamic, not dependant on Site url.
	return tm.deps.Site.Global().Url + img.Url
}

// templates defines the slice of template files in the
// embedded dir to execute.
var templates = []string{
	"meta",
	"schema",
	"opengraph",
	"twitter",
}

// Header
//
// Header obtains all of the site and post wide Code Injection
// as well as any meta information from the page.
//
// Example: {{ verbisHead }}
func (ns *Namespace) Header() template.HTML {
	tm := &TemplateMeta{
		Post:        ns.post,
		Options:     *ns.deps.Options,
		Breadcrumbs: ns.crumbs,
		deps:        ns.deps,
		Site:        ns.deps.Site.Global(),
	}

	head := ns.executeTemplates(tm, templates)

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

// newLineRegex is the regex used to clean the meta
// of line breaks.
var newLineRegex = "\n\n"

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

	meta = strings.TrimRight(meta, "\n")

	regex, err := regexp.Compile(newLineRegex)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error compiling regex", Operation: op, Err: err}).Error()
		return meta
	}

	return regex.ReplaceAllString(meta, "\n")
}
