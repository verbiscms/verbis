// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mailer

import (
	"fmt"
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/logger"
	"os"
)

//
type Mailer interface {
	Send(t *mail.Transmission)
}

//
type Mail struct {
	client      mail.Mailer
	FromAddress string
	FromName    string
	Driver      string
	Env         *environment.Env
	Paths       paths.Paths
}

const (
	// LayoutPath defines where the main layout tpl file
	// resides.
	LayoutPath = "mail" + string(os.PathSeparator) + "main-layout.html"
)

// New
//
//
func New(env *environment.Env) (*Mail, error) {
	const op = "Mail.New"

	cfg := mail.Config{}

	if env == nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Environment can't be nil", Operation: op, Err: fmt.Errorf("nil enviroment")}
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

	_, err := m.client.Send(t)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Mail sending failed", Operation: op, Err: err}).Error()
		return
	}

	logger.Debug("Successfully sent email with the subject: " + t.Subject)
}
