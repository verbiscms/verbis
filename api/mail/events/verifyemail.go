// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/mail"
	"strconv"
)

// VerifyEmail defines the event instance for verifying emails
type VerifyEmail struct {
	mailer *mail.Mailer
	config config.Configuration
}

// NewVerifyEmail creates a new verify email event.
func NewVerifyEmail(config config.Configuration) (*VerifyEmail, error) {
	const op = "events.NewResetPassword"

	m, err := mail.New()
	if err != nil {
		return &VerifyEmail{}, err
	}

	return &VerifyEmail{
		mailer: m,
		config: config,
	}, nil
}

// Send the verify email event.
func (e *VerifyEmail) Send(u *domain.User, title string) error {
	const op = "events.VerifyEmail.Send"

	md5String := encryption.MD5Hash(strconv.Itoa(u.Id) + u.Email)

	data := mail.Data{
		"AppUrl":    e.mailer.Env.AppName,
		"AppTitle":  title,
		"AdminPath": e.mailer.Config.Admin.Path,
		"Token":     md5String,
		"UserName":  u.FirstName,
	}

	html, err := e.mailer.ExecuteHTML("verify-email.html", data)
	if err != nil {
		return err
	}

	tm := mail.Sender{
		To:      []string{u.Email},
		Subject: "Thanks for signing up " + u.FirstName,
		HTML:    html,
	}

	e.mailer.Send(&tm)

	return nil
}
