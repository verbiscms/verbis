// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// ChangedPassword defines the Event instance for config
// resetting passwords, Password and User are required
// for dispatch.
type ChangedPassword struct {
	event    *Event
	Password string
	User     domain.UserPart
}

// ChangedPassword
//
// Creates a new ChangedPassword.
func NewChangedPassword(mail mail.Mailer) *ChangedPassword {
	return &ChangedPassword{
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
func (r *ChangedPassword) Dispatch(data interface{}, recipients []string, attachments mail.Attachments) error {
	const op = "Events.ChangedPassword.Dispatch"

	cp, ok := data.(ChangedPassword)
	if !ok {
		return &errors.Error{Code: errors.INTERNAL, Message: "ChangedPassword should be passed to dispatch", Operation: op, Err: WrongTypeErr}
	}

	err := r.event.send(cp, recipients, attachments)
	if err != nil {
		return err
	}

	return nil
}
