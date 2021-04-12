// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"bytes"
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	"os"
)

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
type event struct {
	Deps      *deps.Deps
	Template  string
	Subject   string
	PlainText string
}

type TplData struct {
	Options *domain.Options
	Site    domain.Site
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
type Dispatcher interface {
	Dispatch(data interface{}, recipients []string, a mail.Attachments) error
}

var (
	WrongTypeErr = errors.New("wrong type passed to dispatch")
)

const (
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	MailTplExtension = ".cms"
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	MasterTplLayout = "layout"
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	SubjectPrefix = "Verbis - "
)

// Send
//
//
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (e *event) send(data interface{}, r []string, a mail.Attachments) error {
	html, err := e.executeHTML(e.Template, data)
	if err != nil {
		return err
	}

	go e.Deps.Mail.Send(&mail.Transmission{
		Recipients:  r,
		Subject:     e.Subject,
		HTML:        html,
		PlainText:   e.PlainText,
		Attachments: a,
	})

	return nil
}

// executeHTML
//
// Executes the events HTML file by preparing the
// template and executing the data.
// Returns errors.INTERNAL if the render failed.
func (e *event) executeHTML(file string, data interface{}) (string, error) {
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

// commonTplData
//
//
func (e *event) commonTplData() *TplData {
	return &TplData{
		Options: e.Deps.Options,
		Site:    e.Deps.Site.Global(),
	}
}
