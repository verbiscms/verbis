// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// tplLineNumber
//
// Returns the line number of the template that broke. If
// the line number could not be retrieved using a regex
// find, -1 will be returned.
func tplLineNumber(err *errors.Error) int {
	if err == nil {
		return -1
	}

	reg := regexp.MustCompile(`:\d+:`)
	lc := string(reg.Find([]byte(err.Error())))
	line := strings.ReplaceAll(lc, ":", "")

	i, ok := strconv.Atoi(line)
	if ok != nil {
		return -1
	}
	return i
}

// tplFileContents
//
// Gets the file contents of the errored template.
// Logs errors.NOTFOUND if the path could not be found.
func tplFileContents(path string) string {
	const op = "Recovery.tplFileContents"

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.NOTFOUND, Message: "Could not get the file contents with the path: " + path, Operation: op, Err: err}).Error()
		return ""
	}

	return string(contents)
}
