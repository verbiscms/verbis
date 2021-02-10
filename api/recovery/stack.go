// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"io/ioutil"
	"runtime"
	"strings"
)

const (
	// The amount of files in the stack to be retrieved
	StackDepth = 200
	// How many lines before and after the calling function
	// to retrieve
	LineLimit = 50
	// How many files to move up in the runtime.Caller
	// before obtaining the stack
	StackSkip = 2
)

// FileStack defines the stack used for the error page
type File struct {
	File     string
	Line     int
	Name     string
	Contents string
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

// FileLine defines the error for templating it includes the
// line & content of the error file.
type FileLine struct {
	Line    int
	Content string
}

// Stack
//
// Returns a slice of FileStack's by traversing the caller.
// using the depth and traverse arguments to loop over.
// If there was an error reading the file, or the
// runtime.Caller function failed, it will not
// be appended to the stack.
func GetStack(depth int, traverse int) Stack {
	var stack Stack

	for c := traverse; c < depth; c++ {
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
			Name:     runtime.FuncForPC(t).Name(),
			Contents: string(contents),
		})
	}

	return stack
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
		if i > 0 && i < len(lines) {
			fileLines = append(fileLines, &FileLine{
				Line:    i + 1,
				Content: lines[i],
			})
		}
		counter++
	}

	return fileLines
}
