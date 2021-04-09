// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
)

// ResetPassword defines the event instance for resetting passwords
type ResetPassword struct {
	ev        *Event
	file      string
	subject   string
	plainText string
}

func NewResetPassword(e *Event) *ResetPassword {
	return &ResetPassword{
		ev:        e,
		file:      "reset-password",
		subject:   "Verbis - Reset Password",
		plainText: "",
	}
}

// Dispatch
//
//
func (r *ResetPassword) Dispatch() error {
	err := r.Validate()
	if err != nil {
		return err
	}

	err = r.ev.send(r.file, r.subject, r.plainText, nil)
	if err != nil {
		return err
	}

	return nil
}

// Validate
//
//
func (r *ResetPassword) Validate() error {
	const op = "Events.ResetPassword.Validate"

	if !r.ev.Data.Exists("Token") {
		return &errors.Error{Code: errors.INVALID, Message: "Token cannot be empty to send reset password event.", Operation: op, Err: fmt.Errorf("token must not be empty")}
	}

	if !r.ev.Data.Exists("User") {
		return &errors.Error{Code: errors.INVALID, Message: "User cannot be empty to send reset password event.", Operation: op, Err: fmt.Errorf("user must not be empty")}
	}

	return nil
}
