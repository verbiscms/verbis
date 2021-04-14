// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mailer

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/tpl"
	"os"
)

//
type Mailer interface {
	Send(t *mail.Transmission)
	ExecuteHTML(file string, data interface{}) (string, error)
}

//
type Mail struct {
	client      mail.Mailer
	FromAddress string
	FromName    string
	Driver      string
	Env         *environment.Env
	Tpl         tpl.TemplateHandler
	Paths       paths.Paths
}

const (
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	MailTplExtension = ".cms"
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	MasterTplLayout = "layout"
)

// New
//
//
func New(env *environment.Env, tpl tpl.TemplateHandler) (*Mail, error) {
	const op = "Mail.New"

	cfg := mail.Config{}

	if env == nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Environment can't be nil", Operation: op, Err: fmt.Errorf("nil environment")}
	}

	switch env.MailDriver {
	case mail.SparkPost:
		cfg = mail.Config{
			URL:    env.SparkpostURL,
			APIKey: env.SparkpostAPIKey,
		}
	case mail.MailGun:
		cfg = mail.Config{
			URL:    env.MailGunURL,
			APIKey: env.MailGunAPIKey,
			Domain: env.MailGunDomain,
		}
	case mail.SendGrid:
		cfg = mail.Config{
			APIKey: env.SendGridAPIKey,
		}
	default:
		return nil, &errors.Error{Code: errors.INVALID, Message: "No mail driver exists: " + env.MailDriver, Operation: op, Err: fmt.Errorf("invalid mail driver")}
	}

	cfg.FromName = env.MailFromName
	cfg.FromAddress = env.MailFromAddress

	client, err := mail.NewClient(env.MailDriver, cfg)
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error validating mail configuration", Operation: op, Err: err}
	}

	return &Mail{
		client:      client,
		FromAddress: env.MailFromAddress,
		FromName:    env.MailFromName,
		Driver:      env.MailDriver,
		Env:         env,
		Paths:       paths.Get(),
	}, nil
}

// Send
//
//
func (m *Mail) Send(t *mail.Transmission) {
	const op = "Mail.Send"

	if m == nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Mailer is nil", Operation: op, Err: fmt.Errorf("nil mail")}).Error()
		return
	}

	if m.client == nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Mail client is nil", Operation: op, Err: fmt.Errorf("nil mail client")}).Error()
		return
	}

	_, err := m.client.Send(t)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Mail sending failed", Operation: op, Err: err}).Error()
		return
	}

	logger.Debug("Successfully sent email with the subject: " + t.Subject)
}

// executeHTML
//
// Executes the events HTML file by preparing the
// template and executing the data.
// Returns errors.INTERNAL if the render failed.
func (m *Mail) ExecuteHTML(file string, data interface{}) (string, error) {
	root := m.Paths.Web + string(os.PathSeparator) + "mail"

	exec := m.Tpl.Prepare(&tpl.Config{
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
