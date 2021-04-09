// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"bytes"
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gookit/color"
	"os"
)

type Event struct {
	Deps       *deps.Deps
	Data       Data
	Recipients []string
}

type Dispatcher interface {
	Dispatch() error
	Validate() error
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
func (e *Event) send(file, subject, plainText string, a mail.Attachments) error {
	html, err := e.executeHTML(file, e.Data)
	if err != nil {
		color.Red.Println(err)
		return err
	}

	color.Green.Println("Content:", html)

	//e.Deps.Mail.Send(&mail.Transmission{
	//	Recipients:  e.Recipients,
	//	Subject:     subject,
	//	HTML:        html,
	//	PlainText:   plainText,
	//	Attachments: a,
	//})

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
