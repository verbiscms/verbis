// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/services/forms"
)

// FormSend defines the event instance for emailing forms
type FormSend struct {
	mailer *MailOld
}

type FormSendData struct {
	Site   domain.Site
	Form   *domain.Form
	Values domain.FormValues
}

// NewVerifyEmail creates a new verify email event
func NewFormSend() (*FormSend, error) {
	const op = "events.NewFormSend"

	m, err := NewOld()
	if err != nil {
		return &FormSend{}, err
	}

	return &FormSend{
		mailer: m,
	}, nil
}

// Send the verify email event.
func (e *FormSend) Send(f *FormSendData, attachments forms.Attachments) error {
	const op = "events.VerifyEmail.Send"

	fv := make(domain.FormValues)
	for _, v := range f.Form.Fields {
		val, ok := f.Values[v.Key]
		if !ok {
			continue
		}
		if v.Type != "file" {
			fv[v.Label.String()] = val
		}
	}
	f.Values = fv

	html, err := e.mailer.ExecuteHTML("form-send.html", &f)
	if err != nil {
		return err
	}

	tm := Sender{
		To:          f.Form.GetRecipients(),
		Subject:     f.Form.EmailSubject,
		HTML:        html,
		Attachments: attachments,
	}

	e.mailer.Send(&tm)

	return nil
}
