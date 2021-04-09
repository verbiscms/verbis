// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"fmt"
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// ChangedPassword defines the event instance for new passwords reset by the system
type ChangedPassword struct {
	ev *Event
}

// Reset Password
//
//
func NewChangedPassword(e *Event) *ResetPassword {
	e.subject = "Verbis - Reset Password"
	e.template = "reset-password"
	e.plaintext = ""
	return &ResetPassword{
		ev: e,
	}
}

// Dispatch
//
//
func (r *ChangedPassword) Dispatch(d Data, recipients []string, attachments mail.Attachments) error {
	err := r.Validate(d)
	if err != nil {
		return err
	}

	err = r.ev.send(d, recipients, attachments)
	if err != nil {
		return err
	}

	return nil
}

// Validate
//
//
func (r *ChangedPassword) Validate(d Data) error {
	const op = "Events.ResetPassword.Validate"

	if !d.Exists("Password") {
		return &errors.Error{Code: errors.INVALID, Message: "Token cannot be empty to send reset password event.", Operation: op, Err: fmt.Errorf("token must not be empty")}
	}

	if !d.Exists("User") {
		return &errors.Error{Code: errors.INVALID, Message: "User cannot be empty to send reset password event.", Operation: op, Err: fmt.Errorf("user must not be empty")}
	}

	_, ok := d["Password"].(string)
	if !ok {
		return &errors.Error{Code: errors.INVALID, Message: "Token must be a string", Operation: op, Err: fmt.Errorf("token must be a string")}
	}

	_, ok = d["User"].(domain.UserPart)
	if !ok {
		return &errors.Error{Code: errors.INVALID, Message: "User must be a domain.UserPart", Operation: op, Err: fmt.Errorf("user must be a user part")}
	}

	return nil
}
