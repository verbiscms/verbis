// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package res

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const (
	STACKDEPTH = 16
)

type Recovery struct {
	Err        errors.Error
	SubMessage string
	Path       string
	Line       int
	Contents   string
	Language   string
	Stack      []Stack
	Highlight  int
	config     config.Configuration
	theme      domain.ThemeConfig
}

// TemplateStack defines the stack used for the error page
type Stack struct {
	File    string
	Line    int
	Name    string
	Message string
}

// FileLine defines the error for templating it includes the
// line & content of the error file.
type FileLine struct {
	Line    int
	Content string
}

func Recover(g *gin.Context, err interface{}) {
	const op = "Server.Recover"

	// IMPORTANT: Do not set anything else to be err here
	cfg, ok := config.New()
	if ok != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get the configuration", Operation: op, Err: ok},
		}).Fatal()
	}

	rc := &Recovery{
		config: *cfg,
	}

	// Load up the Verbis error pages
	gvRecovery := goview.New(goview.Config{
		Root:         paths.Web(),
		Extension:    ".html",
		Master:       "layouts/main",
		DisableCache: true,
	})

	// Assign the error
	fmt.Println(err)
	rc.Err = *rc.setType(err)

	// Set the error for the logger & middleware
	g.Set("verbis_error", &rc.Err)

	// Get the stack
	rc.Stack = rc.getStack()

	// Get the file contents
	contents, err := rc.setFileContents()
	if err != nil {
		log.Panic(err)
	}
	rc.Contents = contents

	// Return the error page
	if err := gvRecovery.Render(g.Writer, http.StatusOK, "/templates/error", gin.H{
		"Error":         rc.Err,
		"Stack":         rc.Stack,
		"SubMessage":    rc.SubMessage,
		"RequestMethod": g.Request.Method,
		"File":          rc.getFileLines(rc.Contents, rc.Line, 10),
		"Highlight":     rc.Highlight,
		"LineNumber":    rc.Line,
		"FileLanguage":  rc.Language,
		"url":           g.Request.URL.Path,
		"Ip":            g.ClientIP(),
		"DataLength":    g.Writer.Size(),
	}); err != nil {
		log.Panic(err)
	}
}

// Get the type of error and return new Verbis error
func (r *Recovery) setType(err interface{}) *errors.Error {
	errType := reflect.TypeOf(err).String()

	var errData errors.Error
	var stack = r.getStack()
	if errType == "*logrus.Entry" {
		entry, ok := err.(*log.Entry)
		if !ok {
			return nil
		}
		errData = entry.Data["error"].(errors.Error)

		r.Line = stack[8].Line
		r.Path = stack[8].File
		r.Highlight = 7
		r.Language = "go"

	} else {
		e, ok := err.(error)
		if !ok {
			return nil
		}
		if strings.Contains(e.Error(), "ViewEngine") {
			if err := r.handleTemplate(e); err != nil {
				panic(err)
			}
			errData = errors.Error{
				Code:      errors.TEMPLATE,
				Message:   fmt.Sprintf("Could not render the template %s: ", r.Path),
				Operation: "RenderTemplate",
				Err:       fmt.Errorf(e.Error()),
			}
			r.SubMessage = e.Error()
		} else {
			errData = errors.Error{
				Code:    errors.INTERNAL,
				Message: "Internal Verbis error, please report.",
				Err:     fmt.Errorf(e.Error()),
			}
		}
		r.Highlight = -1
	}

	return &errData
}

// getStack obtains the stack details from the caller
func (r *Recovery) getStack() []Stack {
	var stack []Stack

	const stackDepth = STACKDEPTH
	for c := 0; c < stackDepth; c++ {
		t, file, line, ok := runtime.Caller(c)
		if ok {
			stack = append(stack, Stack{
				File: file,
				Line: line,
				Name: runtime.FuncForPC(t).Name(),
			})
		}
	}

	return stack
}

// getFileContents gets the file contents of the errored file.
// Returns INTERNAL if the path could not be found
func (r *Recovery) setFileContents() (string, error) {
	const op = "Recovery.getFileContents"

	exists := files.Exists(r.Path)
	if !exists {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not convert get file contents with path %s", r.Path), Operation: op, Err: err}
	}

	contents, err := ioutil.ReadFile(r.Path)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not convert get file contents with path %s", r.Path), Operation: op, Err: err}
	}

	return string(contents), nil
}

// getTemplate obtains the file path for the template and the line number
// if the errors is directly associated with a template, it the assigns
// the variables to the Recovery struct.
// Returns INTERNAL if the line number could not be converted to an integer.
func (r *Recovery) handleTemplate(err error) error {
	const op = "Recovery.getTemplate"

	var (
		file string
		line int
	)

	tmpl := helpers.StringsBetween(err.Error(), "name:", ",")
	lineStr := regexp.MustCompile("[0-9]+").FindAllString(err.Error(), -1)
	if len(lineStr) > 0 {
		l, err := strconv.Atoi(lineStr[0])
		if err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not convert %v to int", line), Operation: op, Err: err}
		}
		line = l
		file = paths.Theme() + "/" + tmpl + r.theme.FileExtension
	}

	r.Path = file
	r.Line = line
	r.Language = "handlebars"

	return nil
}

// Lines gets the range of lines of a file in between a limit
// Returns an array of file lines
func (r *Recovery) getFileLines(file string, line int, limit int) []FileLine {
	split := strings.Split(file, "\n")

	var fileLines []FileLine
	counter := line - (limit / 2)
	for i := 0; i < limit; i++ {
		if counter >= 0 && counter < len(split) {
			fileLines = append(fileLines, FileLine{
				Line:    counter + 1,
				Content: html.UnescapeString(strings.Replace(split[counter], " ", "&nbsp;", -1)),
			})
		}
		counter++
	}

	return fileLines
}
