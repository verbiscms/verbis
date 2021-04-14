// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/errors"
	mailer "github.com/ainsleyclark/verbis/api/mailer"
)

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
type Event struct {
	mailer    mailer.Mailer
	template  string
	subject   string
	plaintext string
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
type Dispatcher interface {
	Dispatch(data interface{}, recipients []string, a mail.Attachments) error
}

var (
	WrongTypeErr = errors.New("wrong type passed to dispatch")
)

const (
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	SubjectPrefix = "Verbis - "
)

// Send
//
//
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (e *Event) send(data interface{}, r []string, a mail.Attachments) error {
	html, err := e.mailer.ExecuteHTML(e.template, data)
	if err != nil {
		return err
	}

	go e.mailer.Send(&mail.Transmission{
		Recipients:  r,
		Subject:     e.subject,
		HTML:        html,
		PlainText:   e.plaintext,
		Attachments: a,
	})

	return nil
}
