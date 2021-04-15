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

// ChangedPassword defines the event instance for config
// resetting passwords, Password and User are required
// for dispatch.
type ChangedPassword struct {
	mail     *mail
	Password string
	User     domain.UserPart
}

// Creates a new ChangedPassword.
func NewChangedPassword(d *deps.Deps) *ChangedPassword {
	e := event{
		Subject:           SubjectPrefix + "Password Change",
		Template:          "new-password",
		PlainTextTemplate: "new-password",
		PreHeader:         SubjectPrefix + "Reset your password within a Verbis installation.",
	}

	mailer, err := newMailer(d, e)
	if err != nil {
		logger.WithError(err).Error()
	}

	return &ChangedPassword{
		mail: mailer,
	}
}

// Dispatches the ResetPassword event.
func (r *ChangedPassword) Dispatch(data interface{}, recipients []string, attachments client.Attachments) error {
	const op = "Events.ChangedPassword.Dispatch"

	cp, ok := data.(ChangedPassword)
	if !ok {
		return &errors.Error{Code: errors.INTERNAL, Message: "ChangedPassword should be passed to dispatch", Operation: op, Err: ErrWrongType}
	}

	go r.mail.Send(cp, recipients, attachments)

	return nil
}
