// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	Colours         bool
	TimestampFormat string
	Debug           bool
	entry           *logrus.Entry
	buf             *bytes.Buffer
}

// Format
//
// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	if !f.Colours {
		color.Disable()
	}

	b := &bytes.Buffer{}
	f.buf = b
	f.entry = entry

	b.WriteString("[VERBIS] ")

	f.Time()
	f.StatusCode()
	f.Level()
	f.IP()
	f.Method()
	f.URL()
	f.Message()
	f.Error()
	f.Fields()

	str := b.String()
	str = strings.TrimSuffix(str, "|")
	str = strings.TrimSuffix(str, "|")
	str = strings.TrimSuffix(str, " ")
	str += "\n"

	return []byte(str), nil
}

// Time
//
// Prints the timestamp for the log, if no format is set
// on the formatter, time.StampMilli will be used.
func (f *Formatter) Time() {
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.StampMilli
	}
	f.buf.WriteString(f.entry.Time.Format(timestampFormat))
}

// StatusCode
//
// Prints the status code of the request, if there is none
// set the log is internal and "VRB" will be printed.
func (f *Formatter) StatusCode() {
	f.buf.WriteString(" | ")

	cc := color.Style{color.FgLightWhite, color.BgRed, color.OpBold}

	status, ok := f.entry.Data["status_code"]
	if !ok {
		cc = color.Style{color.FgLightWhite, color.BgBlack, color.OpBold}
		f.buf.WriteString(cc.Sprint("VRB"))
	}

	if codeInt, ok := status.(int); ok {
		if codeInt < http.StatusBadRequest {
			cc = color.Style{color.FgLightWhite, color.BgGreen, color.OpBold}
		}
	}

	if status != "" && status != nil {
		f.buf.WriteString(cc.Sprintf("%d", status))
	}

	f.buf.WriteString(" | ")
}

// Level
//
// Prints the entry level of the log entry in uppercase.
func (f *Formatter) Level() {
	cc := color.Style{} //nolint
	switch f.entry.Level {
	case logrus.DebugLevel:
		cc = color.Style{color.FgGray, color.OpBold}
	case logrus.WarnLevel:
		cc = color.Style{color.FgYellow, color.OpBold}
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		cc = color.Style{color.FgRed, color.OpBold}
	default:
		cc = color.Style{color.FgBlue, color.OpBold}
	}

	level := strings.ToUpper(f.entry.Level.String())
	if len(level) == 4 { //nolint
		f.buf.WriteString(cc.Sprintf("[%s] ", level))
		return
	}

	f.buf.WriteString(cc.Sprintf("[%s]", level))
}

// IP
//
// Prints the IP address if there is any.
func (f *Formatter) IP() {
	ip, ok := f.entry.Data["client_ip"].(string)
	if ok {
		f.buf.WriteString(fmt.Sprintf(" | %s | ", ip))
		return
	}
	f.buf.WriteString(" ")
}

// Method
//
// Prints the entry request method if there is one set.
func (f *Formatter) Method() {
	method, ok := f.entry.Data["request_method"].(string)
	if !ok {
		return
	}
	rc := color.Style{color.FgLightWhite, color.BgBlue, color.OpBold}
	f.buf.WriteString(rc.Sprintf("  %s   ", method))
}

// URL
//
// Prints the entry request url if there is one set.
func (f *Formatter) URL() {
	url, ok := f.entry.Data["request_url"].(string)
	if ok {
		f.buf.WriteString(fmt.Sprintf(" \"%s\" ", url))
	}
}

// Message
//
// Prints the entry message if there is obe set.
func (f *Formatter) Message() {
	_, method := f.entry.Data["request_method"].(string)
	_, ok := f.HasError()
	if ok && !method {
		return
	}

	msg, ok := f.entry.Data["message"].(string)
	if ok && msg != "" {
		f.buf.WriteString(fmt.Sprintf("| [msg] %s |", msg))
		return
	}
	if f.entry.Message != "" {
		f.buf.WriteString(fmt.Sprintf("| [msg] %s |", f.entry.Message))
	}
}

// Fields
//
// Prints the entry fields
func (f *Formatter) Fields() {
	fields, ok := f.entry.Data["fields"].(logrus.Fields)
	if !ok {
		return
	}
	f.buf.WriteString("| ")
	for k, v := range fields {
		f.buf.WriteString(fmt.Sprintf("%s: %s ", k, v))
	}
}

// HasError
//
// Determines if a verbis error has been logged.
func (f *Formatter) HasError() (interface{}, bool) {
	e, _ := f.entry.Data["error"]
	if e == nil {
		return nil, false
	}
	return e, true
}

// Error
//
// Prints out the error if there is one set. If the error
// is nil, nothing will be printed.
func (f *Formatter) Error() {
	e, ok := f.HasError()
	if !ok {
		return
	}

	f.buf.WriteString("|")

	err := errors.ToError(e)

	if err == nil {
		return
	}

	if err.Code != "" {
		f.buf.WriteString(color.Red.Sprintf(" [code] "))
		f.buf.WriteString(err.Code)
	}

	if err.Message != "" {
		f.buf.WriteString(color.Red.Sprintf(" [msg] "))
		f.buf.WriteString(err.Message)
	}

	if err.Operation != "" && f.Debug {
		f.buf.WriteString(color.Red.Sprintf(" [op] "))
		f.buf.WriteString(err.Operation)
	}

	if err.Err != nil {
		f.buf.WriteString(color.Red.Sprintf(" [error] "))
		f.buf.WriteString(err.Err.Error())
	}
}
