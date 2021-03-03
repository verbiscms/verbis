// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mail

import (
	"bytes"
	"fmt"
	sp "github.com/SparkPost/gosparkpost"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/forms"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"html/template"
)

type Mailer struct {
	client       sp.Client
	Transmission Sender
	FromAddress  string
	FromName     string
	Env          *environment.Env
	Paths        paths.Paths
}

type Sender struct {
	To          []string
	Subject     string
	HTML        string
	Attachments forms.Attachments
}

type Data map[string]interface{}

// New, Create a new mailable instance using environment variables.
func New() (*Mailer, error) {
	const op = "mail.New"
	env, _ := environment.Load()
	m := &Mailer{
		Env:   env,
		Paths: paths.Get(),
	}
	if err := m.load(); err != nil {
		return &Mailer{}, err
	}
	return m, nil
}

// Load the mailer and connect to sparkpost
// Returns errors.INTERNAL if the new mailer instance could not be created
func (m *Mailer) load() error {
	const op = "mail.Load"

	// TODO this is temporary
	mailConf := m.Env.MailConfig()
	config := &sp.Config{
		BaseUrl:    mailConf.SparkpostURL,
		ApiKey:     mailConf.SparkpostAPIKey,
		ApiVersion: 1,
	}

	var client sp.Client
	err := client.Init(config)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not create a new mailer instance", Operation: op, Err: err}
	}
	m.client = client

	m.FromAddress = mailConf.FromAddress
	m.FromName = mailConf.FromName

	return nil
}

// Create a Transmission using an inline Recipient List
// and inline email Content.
// Returns errors.INVALID if the mail failed to send via sparkpost.
func (m *Mailer) Send(t *Sender) {
	const op = "mail.Send"

	content := sp.Content{
		HTML:    t.HTML,
		From:    m.FromAddress,
		Subject: t.Subject,
	}

	if len(t.Attachments) != 0 {
		var att []sp.Attachment
		for _, v := range t.Attachments {
			att = append(att, sp.Attachment{
				MIMEType: v.MIMEType,
				Filename: v.Filename,
				B64Data:  *v.B64Data,
			})
		}
		content.Attachments = att
	}

	tx := &sp.Transmission{
		Recipients: t.To,
		Content:    content,
	}

	id, _, err := m.client.Send(tx)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Mail sending failed: %s", id), Operation: op, Err: err}).Error()
		return
	}

	// TODO: Nil pointer dereference here for logging?
	fmt.Println("Email successfully sent")
}

// Execute the mail HTML files
// Returns errors.INTERNAL if the render failed
func (m *Mailer) ExecuteHTML(file string, data interface{}) (string, error) {
	const op = "mail.ExecuteHTML"
	path := m.Paths.Web + "/mail/" + file
	tmpl, err := RenderTemplate("main", data, m.Paths.Web+"/mail/main-layout.html", path)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to render the template: %s", path), Operation: op, Err: err}
	}
	return tmpl, nil
}

// RenderTemplate executes the html and returns a string
// Returns errors.INTERNAL if the template failed to be created
// or be executed.
func RenderTemplate(layout string, data interface{}, files ...string) (string, error) {
	const op = "html.RenderTemplate"

	t, err := template.New("").ParseFiles(files...)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Unable to create a new template", Operation: op, Err: err}
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, layout, data); err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Unable to render the template", Operation: op, Err: err}
	}

	return tpl.String(), nil
}
