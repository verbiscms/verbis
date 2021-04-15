// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"bytes"
	"fmt"
	client "github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/tpl"
	"os"
	"text/template"
)

// Dispatcher describes the event for sending email. Data
// and recipients are required, but attachments are
// optional.
type Dispatcher interface {
	Dispatch(data interface{}, recipients []string, a client.Attachments) error
}

// mail describes the properties for sending emails
// through verbis.
type mail struct {
	FromAddress string
	FromName    string
	Driver      string
	Deps        *deps.Deps
	Client      client.Mailer
	TplRoot     string
	event
}

// event describes common properties of an event when a
// transmission is fired.
type event struct {
	Template          string // e.g. reset-password
	Subject           string // e.g. Reset Password
	PlainTextTemplate string // e.g. reset-password
	// PreHeader is the short summary text that follows the
	// subject line when viewing emails from the inbox.
	// Some clients do not support it.
	PreHeader string // e.g. "Reset your password within a Verbis installation."
}

// tplData defines the data for executing templates,
// including the unique data for each event.
type tplData struct {
	PreHeader string
	Options   *domain.Options
	Site      domain.Site
	Data      interface{}
}

var (
	// WrongTypeErr is returned by an event when the wrong
	// type is passed to Dispatch
	WrongTypeErr = errors.New("wrong type passed to dispatch")
)

const (
	// MailDir defines the directory where mail text and HTML
	// templates are stored.
	MailDir = "mail"
	// MailHtmlExtension defines the extension for executing
	// HTML templates.
	MailHtmlExtension = ".cms"
	// MailTextExtension defines the extension for executing
	// text templates.
	MailTextExtension = ".txt"
	// SubjectPrefix defines prefix attached to all emails
	// from Verbis.
	SubjectPrefix = "Verbis - "
	// MasterTplLayout defines the master layout for executing
	// HTML templates.
	MasterTplLayout = "layout"
)

// HealthCheck
//
// HealthCheck performs a check on the environment and
// mail client to see if mail can be sent with the
// current configuration.
// Returns errors.INVALID if the environment is nil, the
// mail driver does not exist or if there was an
// error creating a new client.
func HealthCheck(env *environment.Env) error {
	_, err := newMailer(&deps.Deps{Env: env}, event{})
	if err != nil {
		return err
	}
	return nil
}

// newMailer creates a new mailer instance by comparing
// the mail driver within the environment.
func newMailer(d *deps.Deps, ev event) (*mail, error) {
	const op = "Mail.NewMailer"

	cfg := client.Config{}

	if d.Env == nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Environment can't be nil", Operation: op, Err: fmt.Errorf("nil environment")}
	}

	switch d.Env.MailDriver {
	case client.SparkPost:
		cfg = client.Config{
			URL:    d.Env.SparkpostURL,
			APIKey: d.Env.SparkpostAPIKey,
		}
	case client.MailGun:
		cfg = client.Config{
			URL:    d.Env.MailGunURL,
			APIKey: d.Env.MailGunAPIKey,
			Domain: d.Env.MailGunDomain,
		}
	case client.SendGrid:
		cfg = client.Config{
			APIKey: d.Env.SendGridAPIKey,
		}
	default:
		return nil, &errors.Error{Code: errors.INVALID, Message: "No mail driver exists: " + d.Env.MailDriver, Operation: op, Err: fmt.Errorf("invalid mail driver")}
	}

	cfg.FromName = d.Env.MailFromName
	cfg.FromAddress = d.Env.MailFromAddress

	m, err := client.NewClient(d.Env.MailDriver, cfg)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error validating mail configuration", Operation: op, Err: err}
	}

	return &mail{
		Client:      m,
		Deps:        d,
		event:       ev,
		TplRoot:     d.Paths.Web + string(os.PathSeparator) + MailDir,
		FromAddress: d.Env.MailFromAddress,
		FromName:    d.Env.MailFromName,
		Driver:      d.Env.MailDriver,
	}, nil
}

// Validates the mail struct by checking if the client and
// client is nil.
// Returns errors.INTERNAL in both cases.
func (m *mail) Validate() error {
	const op = "Mail.Validate"

	if m == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Mailer is nil", Operation: op, Err: fmt.Errorf("nil mail")}
	}

	if m.Client == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Mail client is nil", Operation: op, Err: fmt.Errorf("nil mail client")}
	}

	return nil
}

// Sends the transmission from the email client. The mail
// instance is first validated, and HTML and PlainText
// is executed.
// Logs errors.INTERNAL in any instance of an error.
func (m *mail) Send(data interface{}, r []string, a client.Attachments) {
	const op = "mail.Send"

	err := m.Validate()
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	html, err := m.ExecuteHTML(m.event.Template, data)
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	plainText, err := m.ExecuteText(m.PlainTextTemplate, data)
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	tm := &client.Transmission{
		Recipients:  r,
		Subject:     m.Subject,
		HTML:        html,
		PlainText:   plainText,
		Attachments: a,
	}

	_, err = m.Client.Send(tm)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "mail sending failed", Operation: op, Err: err}).Error()
		return
	}

	logger.Info("Successfully sent email with the subject: " + m.event.Subject)
}

// ExecuteHTML Executes the events HTML file by preparing
// the template and executing the data.
// Returns errors.TEMPLATE if the file could not be
// rendered.
func (m *mail) ExecuteHTML(file string, data interface{}) (string, error) {
	exec := m.Deps.Tmpl().Prepare(&tpl.Config{
		Root:      m.TplRoot,
		Extension: MailHtmlExtension,
		Master:    MasterTplLayout,
	})

	var buf bytes.Buffer
	_, err := exec.Execute(&buf, file, m.GetTplData(data))
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// ExecuteText executes the plain text for the email.
// Returns errors.INTERNAL if the template could not be
// parsed or executed.
func (m *mail) ExecuteText(file string, data interface{}) (string, error) {
	const op = "Mail.ExecuteText"

	tmpl, err := template.New(file + MailTextExtension).ParseFiles(m.TplRoot + string(os.PathSeparator) + file + MailTextExtension)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Error parsing text template: " + file, Operation: op, Err: err}
	}

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, file+MailTextExtension, m.GetTplData(data))
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Error executing text template: " + file, Operation: op, Err: err}
	}

	return buf.String(), nil
}

// GetTplData returns the common template data for
// emails.
func (m *mail) GetTplData(data interface{}) tplData {
	return tplData{
		PreHeader: m.PreHeader,
		Options:   m.Deps.Options,
		Site:      m.Deps.Site.Global(),
		Data:      data,
	}
}
