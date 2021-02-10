// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	// TplPrefix defines the prefix for template searching
	// such as "errors-500.html"
	TplPrefix = "error"
	// The extension of the web error file
	VerbisErrorExtension = "html"
	// The main layout of the web error file
	VerbisErrorLayout = "layouts/main"
)

// getError
//
// Converts an interface{} to a Verbis internal error ready
// for processing and to return to the recovery page.
func getError(e interface{}) *errors.Error {
	switch v := e.(type) {
	case errors.Error:
		return &v
	case *errors.Error:
		return v
	case error:
		return &errors.Error{Code: errors.TEMPLATE, Message: v.Error(), Operation: "", Err: v}
	default:
		return &errors.Error{Code: errors.TEMPLATE, Message: "Internal Verbis error, please report", Operation: "Internal", Err: fmt.Errorf("%v", e)}
	}
}

// tplLineNumber
//
// Returns the line number of the template that broke.
// If the line number could not be retrieved using
// a regex find, -1 will be returned.
func tplLineNumber(err interface{}) int {
	e := getError(err)
	reg := regexp.MustCompile(`:\d+:`)
	lc := string(reg.Find([]byte(e.Error())))
	line := strings.ReplaceAll(lc, ":", "")

	i, ok := strconv.Atoi(line)
	if ok != nil {
		return -1
	}
	return i
}

// tplFileContents
//
// Gets the file contents of the errored template..
// Logs errors.NOTFOUND if the path could not be found.
func tplFileContents(path string) string {
	const op = "Recovery.tplFileContents"

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.WithFields(log.Fields{
			"error": &errors.Error{Code: errors.NOTFOUND, Message: "Could not get the file contents with the path: " + path, Operation: op, Err: err},
		})
		return ""
	}

	return string(contents)
}

// resolver
//
//
func (r *Recover) resolveErrorPage(custom bool) (string, tpl.TemplateExecutor, bool) {
	code := strconv.Itoa(r.config.Code)
	if r.config.Code == 0 || code == "0" {
		code = "500"
	}

	path := ""
	e := r.deps.Tmpl().Prepare(tpl.Config{
		Root:      r.deps.Paths.Theme + "/" + r.deps.Theme.TemplateDir,
		Extension: r.deps.Theme.FileExtension,
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

// verbisErrorResolver
//
// Returns a new tpl.TemplateExecutor when no other
// custom error templates have been found.
func (r *Recover) verbisErrorResolver() (string, tpl.TemplateExecutor) {
	return "templates/error", r.deps.Tmpl().Prepare(tpl.Config{
		Root:      r.deps.Paths.Web,
		Extension: VerbisErrorExtension,
		Master:    VerbisErrorLayout,
	})
}
