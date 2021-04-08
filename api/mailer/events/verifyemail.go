// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/mailer"
	"strconv"
)

// VerifyEmail defines the event instance for verifying emails
type VerifyEmail struct {
	mailer *mailer.MailOld
}

// NewVerifyEmail creates a new verify email event.
func NewVerifyEmail() (*VerifyEmail, error) {
	const op = "events.NewResetPassword"

	m, err := mailer.NewOld()
	if err != nil {
		return &VerifyEmail{}, err
	}

	return &VerifyEmail{
		mailer: m,
	}, nil
}

// Send the verify email event.
func (e *VerifyEmail) Send(u *domain.User, title string) error {
	const op = "events.VerifyEmail.Send"

	md5String := encryption.MD5Hash(strconv.Itoa(u.Id) + u.Email)

	data := mailer.Data{
		"AppUrl":    "Verbis",
		"AppTitle":  title,
		"AdminPath": "/admin",
		"Token":     md5String,
		"UserName":  u.FirstName,
	}

	html, err := e.mailer.ExecuteHTML("verify-email.html", data)
	if err != nil {
		return err
	}

	tm := mailer.Sender{
		To:      []string{u.Email},
		Subject: "Thanks for signing up " + u.FirstName,
		HTML:    html,
	}

	e.mailer.Send(&tm)

	return nil
}
