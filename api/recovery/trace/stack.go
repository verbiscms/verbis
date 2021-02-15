// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	// How many lines before and after the calling function
	// to retrieve.
	LineLimit = 60
)

// Tracer represents the functionality for obtaining a new
// stack.
type Tracer interface {
	Trace(depth int, skip int) Stack
}

// Trace implements the trace method to obtain the stack
type trace struct{}

// Return a new tracer
func New() *trace {
	return &trace{}
}

// Stack defines the slice of file lines for recovery
type Stack []*File

// Append a file to the stack trace
func (s *Stack) Append(file *File) {
	*s = append(*s, file)
}

// Prepend a file to the stack trace (useful for templates)
func (s *Stack) Prepend(file *File) {
	if len(*s) == 0 {
		*s = append(*s, file)
		return
	}
	*s = append([]*File{file}, *s...)
}

// Find a file in the stack by function.
func (s *Stack) Find(fn string) *File {
	for _, v := range *s {
		if v.Function == fn {
			return v
		}
	}
	return nil
}

// Stack
//
// Returns a slice of FileStack's by traversing the caller.
// using the depth and traverse arguments to loop over.
// If there was an error reading the file, or the
// runtime.Caller function failed, it will not
// be appended to the stack.
func (t *trace) Trace(depth int, skip int) Stack {
	var stack Stack

	t.Test()

	for c := skip; c < depth; c++ {
		t, file, line, ok := runtime.Caller(c)

		if !ok {
			continue
		}

		contents, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}

		stack.Append(&File{
			File:     file,
			Line:     line,
			Function:     runtime.FuncForPC(t).Name(),
			Contents: string(contents),
			Language: Language(file),
		})
	}

	return stack
}

func (t *trace) Test() {
	pc, fspec, line, _ := runtime.Caller(1)
	fmt.Printf("☛ %v || %s [%s:%d]\n", "heklo", runtime.FuncForPC(pc).Name(), fspec, line)
}

// language
//
// Returns the language used in the file for syntax
// highlighting.
func Language(path string) string {
	ext := filepath.Ext(path)
	switch ext {
	case ".go":
		return "go"
	case ".s":
		return "assembly"
	default:
		return "handlebars"
	}
}

// FileStack defines the stack used for the error page
type File struct {
	File     string
	Line     int
	Function     string
	Contents string
	Language string
}

// FileLine defines the error for templating it includes the
// line & content of the error file.
type FileLine struct {
	Line    int
	Content string
}

// Vendor
//
// Determines if a file is Verbis specific or vendor.
func (f *File) Vendor() bool {
	if f.Language == "handlebars" {
		return false
	}
	return !strings.Contains(f.Function, "verbis")
}

// Lines
//
// Splits the file into a array of lines by separating
// them by a new line.
func (f *File) Lines() []*FileLine {
	lines := strings.Split(f.Contents, "\n")

	diff := LineLimit / 2

	var fileLines []*FileLine
	counter := 0
	for i := f.Line - diff; i < f.Line+diff; i++ {
		if i >= 0 && i < len(lines) {
			fileLines = append(fileLines, &FileLine{
				Line:    i + 1,
				Content: lines[i],
			})
		}
		counter++
	}

	return fileLines
}

