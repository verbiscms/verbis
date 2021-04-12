// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// ResetPassword defines the event instance for config
// resetting passwords, Token and User are required
// for dispatch.
type ResetPassword struct {
	event *event
	Token string
	User  domain.UserPart
	*TplData
}

// NewResetPassword
//
// Creates a new ResetPassword.
func NewResetPassword(d *deps.Deps) *ResetPassword {
	e := &event{
		Deps:      d,
		Subject:   SubjectPrefix + "Reset Password",
		Template:  "reset-password",
		PlainText: "",
	}
	return &ResetPassword{
		event:   e,
		TplData: e.commonTplData(),
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
