// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"bytes"
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl"
	"os"
)

type Event struct {
	Deps      *deps.Deps
	template  string
	subject   string
	plaintext string
}

type Dispatcher interface {
	Dispatch(d Data, recipients []string, a mail.Attachments) error
}

const (
	MailTplExtension = ".cms"
	MasterTplLayout  = "layout"
)

type Data map[string]interface{}

func (d Data) Exists(key string) bool {
	_, ok := d[key]
	return ok
}

// Send
//
//
func (e *Event) send(d Data, r []string, a mail.Attachments) error {
	html, err := e.executeHTML(e.template, d)
	if err != nil {
		return err
	}

	e.Deps.Mail.Send(&mail.Transmission{
		Recipients:  r,
		Subject:     e.subject,
		HTML:        html,
		PlainText:   e.plaintext,
		Attachments: a,
	})

	return nil
}

// ExecuteHTML
//
// Execute the mail HTML files
// Returns errors.INTERNAL if the render failed
func (e *Event) executeHTML(file string, data interface{}) (string, error) {
	root := e.Deps.Paths.Web + string(os.PathSeparator) + "mail"

	exec := e.Deps.Tmpl().Prepare(&tpl.Config{
		Root:      root,
		Extension: MailTplExtension,
		Master:    MasterTplLayout,
	})

	var buf bytes.Buffer
	_, err := exec.Execute(&buf, file, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
