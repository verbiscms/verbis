// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"fmt"
	client "github.com/ainsleyclark/go-mail"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// FormSend defines the event instance for config sending
// form data, Form and FormValues are required for
//dispatch.
type FormSend struct {
	mail   *mail
	Form   *domain.Form
	Values domain.FormValues
}

// NewFormSend creates a new FormSend event.
func NewFormSend(d *deps.Deps) *FormSend {
	e := event{
		Subject:           SubjectPrefix + "Form Submission",
		Template:          "form-send",
		PlainTextTemplate: "form-send",
		PreHeader:         SubjectPrefix + "New form submission within a Verbis installation.",
	}

	mailer, err := newMailer(d, e)
	if err != nil {
		logger.WithError(err).Error()
	}

	return &FormSend{
		mail: mailer,
	}
}

// Dispatch the FormSend event.
func (r *FormSend) Dispatch(data interface{}, recipients []string, attachments client.Attachments) error {
	const op = "Events.FormSend.Dispatch"

	fs, ok := data.(FormSend)
	if !ok {
		return &errors.Error{Code: errors.INTERNAL, Message: "FormSend should be passed to dispatch", Operation: op, Err: ErrWrongType}
	}

	if fs.Form == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Form cannot be nil", Operation: op, Err: fmt.Errorf("form is nil")}
	}

	fv := make(domain.FormValues)
	for _, v := range fs.Form.Fields {
		val, ok := fs.Values[v.Key]
		if !ok {
			continue
		}
		if v.Type != "file" {
			fv[v.Label.String()] = val
		}
	}
	fs.Values = fv

	go r.mail.Send(fs, recipients, attachments)

	return nil
}
