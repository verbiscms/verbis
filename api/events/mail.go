// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"bytes"
	"fmt"
	client "github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/tpl"
	"os"
)

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
type Dispatcher interface {
	Dispatch(data interface{}, recipients []string, a client.Attachments) error
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
type mail struct {
	FromAddress string
	FromName    string
	Driver      string
	Deps        *deps.Deps
	Client      client.Mailer
	event
}

type event struct {
	Template  string
	Subject   string
	Plaintext string
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

// HealthCheck
//
//
func HealthCheck(env *environment.Env) error {
	_, err := newMailer(&deps.Deps{Env: env}, event{})
	if err != nil {
		return err
	}
	return nil
}

// newMailer
//
//
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

	client, err := client.NewClient(d.Env.MailDriver, cfg)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error validating mail configuration", Operation: op, Err: err}
	}

	return &mail{
		Client:      client,
		Deps:        d,
		event:       ev,
		FromAddress: d.Env.MailFromAddress,
		FromName:    d.Env.MailFromName,
		Driver:      d.Env.MailDriver,
	}, nil
}

// Send
//
//
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (m *mail) Send(data interface{}, r []string, a client.Attachments) {
	const op = "mail.Send"

	html, err := m.executeHTML(m.event.Template, data)
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	if m == nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Mailer is nil", Operation: op, Err: fmt.Errorf("nil mail")}).Error()
		return
	}

	if m.Client == nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "mail client is nil", Operation: op, Err: fmt.Errorf("nil mail client")}).Error()
		return
	}

	tm := &client.Transmission{
		Recipients:  r,
		Subject:     m.Subject,
		HTML:        html,
		PlainText:   m.Plaintext,
		Attachments: a,
	}

	_, err = m.Client.Send(tm)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "mail sending failed", Operation: op, Err: err}).Error()
		return
	}

	logger.Debug("Successfully sent email with the subject: " + m.event.Subject)
}

// executeHTML
//
// Executes the events HTML file by preparing the
// template and executing the data.
// Returns errors.INTERNAL if the render failed.
func (m *mail) executeHTML(file string, data interface{}) (string, error) {
	root := m.Deps.Paths.Web + string(os.PathSeparator) + "mail"

	exec := m.Deps.Tmpl().Prepare(&tpl.Config{
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
