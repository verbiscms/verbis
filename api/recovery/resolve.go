// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"github.com/verbiscms/verbis/api/tpl"
	"strconv"
)

const (
	// TplPrefix defines the prefix for template searching
	// such as "errors-500.html".
	TplPrefix = "error"
	// VerbisErrorExtension is the extension of the web
	// error file.
	VerbisErrorExtension = ".html"
	// VerbisErrorLayout is the main layout of the web
	// error file.
	VerbisErrorLayout = "layouts/main"
)

// resolver Looks up custom error pages from the theme.
// It starts by looking at error-`code`.extension, if
// the tmpl does not exist it will continue to look
// for error.extension. Finally it uses the main
// verbis error pages.
// Returns the template path, the template execute and
// if the file is a custom template.
func (r *Recover) resolveErrorPage(custom bool) (string, tpl.TemplateExecutor, bool) {
	code := strconv.Itoa(r.config.Code)
	if r.config.Code == 0 || code == "0" {
		code = "500"
	}

	path := ""
	e := r.deps.Tmpl().Prepare(tpl.Config{
		Root:      r.deps.ThemePath() + "/" + r.deps.Config.TemplateDir,
		Extension: r.deps.Config.FileExtension,
	})

	// Look for `errors-404.cms` for example
	path = TplPrefix + "-" + code
	if e.Exists(path) && custom {
		return path, e, true
	}

	// Look for `error.cms` for example
	path = TplPrefix
	if e.Exists(TplPrefix) && custom {
		return path, e, true
	}

	// Return the native error page
	path, exec := r.verbisErrorResolver()
	return path, exec, false
}

// verbisErrorResolver Returns a new tpl.TemplateExecutor when
// no other custom error templates have been found.
func (r *Recover) verbisErrorResolver() (string, tpl.TemplateExecutor) {
	return "templates/error", r.deps.Tmpl().Prepare(tpl.Config{
		Extension: VerbisErrorExtension,
		Master:    VerbisErrorLayout,
		FS:        r.deps.FS.Web,
	})
}
