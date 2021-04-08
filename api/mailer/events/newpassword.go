// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mailer"
)

// ChangedPassword defines the event instance for new passwords reset by the system
type ChangedPassword struct {
	mailer *mailer.MailOld
}

// NewPassword creates a new reset password event
func NewChangedPassword() (*ChangedPassword, error) {
	const op = "events.NewResetPassword"

	m, err := mailer.NewOld()
	if err != nil {
		return &ChangedPassword{}, err
	}

	return &ChangedPassword{
		mailer: m,
	}, nil
}

// Send the reset password event.
func (e *ChangedPassword) Send(u domain.UserPart, password string, site domain.Site) error {
	const op = "events.ResetPassword.Send"

	data := mailer.Data{
		"AppUrl":    site.Url,
		"AppTitle":  site.Title,
		"AdminPath": "/admin",
		"UserName":  u.FirstName,
		"Password":  password,
		"Email":     u.Email,
	}

	html, err := e.mailer.ExecuteHTML("new-password.html", data)
	if err != nil {
		return err
	}

	tm := mailer.Sender{
		To:      []string{u.Email},
		Subject: fmt.Sprintf("NewOld Login Details for %s", site.Title),
		HTML:    html,
	}

	e.mailer.Send(&tm)

	return nil
}
