// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	client "github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// ResetPassword defines the event instance for config
// resetting passwords, Token and User are required
// for dispatch.
type ResetPassword struct {
	mail *mail
	User domain.UserPart
	URL  string
}

// NewResetPassword creates a new ResetPassword event.
func NewResetPassword(d *deps.Deps) *ResetPassword {
	e := event{
		Subject:           SubjectPrefix + "Reset Password",
		Template:          "reset-password",
		PlainTextTemplate: "reset-password",
		PreHeader:         SubjectPrefix + "Reset your password within a Verbis installation.",
	}

	mailer, err := newMailer(d, e)
	if err != nil {
		logger.WithError(err).Error()
	}

	return &ResetPassword{
		mail: mailer,
	}
}

// Dispatch the ResetPassword event.
func (r *ResetPassword) Dispatch(data interface{}, recipients []string, attachments client.Attachments) error {
	const op = "Events.ResetPassword.Dispatch"

	rp, ok := data.(ResetPassword)
	if !ok {
		return &errors.Error{Code: errors.INTERNAL, Message: "ResetPassword should be passed to dispatch", Operation: op, Err: ErrWrongType}
	}

	go r.mail.Send(rp, recipients, attachments)

	return nil
}
