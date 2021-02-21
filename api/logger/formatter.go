// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	Colours         bool
	TimestampFormat string
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
	f.Url()
	f.Message()
	f.Error()
	f.Fields()

	str := b.String()
	str = strings.TrimSuffix(str, "|")
	str = strings.TrimSuffix(str, " ")
	str += "\n"

	return []byte(str), nil
}

// Time
//
//
func (f *Formatter) Time() {
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.StampMilli
	}
	f.buf.WriteString(f.entry.Time.Format(timestampFormat))
}

// StatusCode
//
//
func (f *Formatter) StatusCode() {
	f.buf.WriteString(" | ")

	cc := color.Style{color.FgLightWhite, color.BgRed, color.OpBold}

	status, ok := f.entry.Data["status_code"]
	if !ok {
		cc = color.Style{color.FgLightWhite, color.BgBlack, color.OpBold}
		f.buf.WriteString(cc.Sprint("VRB"))
	}

	if codeInt, ok := status.(int); ok {
		if codeInt < 400 {
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
//
func (f *Formatter) Level() {
	cc := color.Style{}
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
	f.buf.WriteString(cc.Sprintf("[%s]", level))
}

// IP
//
// Print the IP address if there is any.
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
// Print the request method if there is any.
func (f *Formatter) Method() {
	method, ok := f.entry.Data["request_method"].(string)
	if !ok {
		return
	}
	rc := color.Style{color.FgLightWhite, color.BgBlue, color.OpBold}
	f.buf.WriteString(rc.Sprintf("  %s   ", method))
}

// Url
//
// Print the request
func (f *Formatter) Url() {
	url, ok := f.entry.Data["request_url"].(string)
	if ok {
		f.buf.WriteString(fmt.Sprintf(" \"%s\"", url))
	}
}

// Message
//
//
func (f *Formatter) Message() {
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
//
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

// Error
//
//
func (f *Formatter) Error() {
	err, ok := f.entry.Data["error"]
	if !ok || err == nil {
		return
	}

	f.buf.WriteString("|")

	switch v := err.(type) {
	case *errors.Error:
		f.printError(v)
	case errors.Error:
		f.printError(&v)
	case error:
		f.printError(&errors.Error{Err: v})
	case string:
		f.printError(&errors.Error{Err: fmt.Errorf(v)})
	}
}

// printError
//
//
func (f *Formatter) printError(err *errors.Error) {
	if err.Code != "" {
		f.buf.WriteString(color.Red.Sprintf(" [code] %s", err.Code))
	}

	if err.Message != "" {
		f.buf.WriteString(color.Red.Sprintf(" [msg] %s", err.Message))
	}

	if err.Operation != "" && api.SuperAdmin {
		f.buf.WriteString(color.Red.Sprintf(" [op] %s", err.Operation))
	}

	if err.Err != nil {
		f.buf.WriteString(color.Red.Sprintf(" [error] %s", err.Err.Error()))
	}
}
