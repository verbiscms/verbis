// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/mailer"
)

// ResetPassword defines the Event instance for config
// resetting passwords, Token and User are required
// for dispatch.
type ResetPassword struct {
	event *Event
	Token string
	User  domain.UserPart
}

// NewResetPassword
//
// Creates a new ResetPassword.
func NewResetPassword(mail mailer.Mailer) *ResetPassword {
	return &ResetPassword{
		event: &Event{
			subject:   SubjectPrefix + "Reset Password",
			template:  "reset-password",
			plaintext: "",
		},
	}
}

// Dispatch
//
// Dispatches the ResetPassword Event.
func (r *ResetPassword) Dispatch(data interface{}, recipients []string, attachments mail.Attachments) error {
	const op = "Events.ResetPassword.Dispatch"

	rp, ok := data.(ResetPassword)
	if !ok {
		return &errors.Error{Code: errors.INTERNAL, Message: "ResetPassword should be passed to dispatch", Operation: op, Err: WrongTypeErr}
	}

	err := r.event.send(rp, recipients, attachments)
	if err != nil {
		return err
	}

	return nil
}
