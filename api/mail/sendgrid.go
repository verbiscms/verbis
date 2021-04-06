// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mail

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type mandrill struct {
	cfg    Config
	client *sendgrid.Client
}

// newMandrill
//
//
func (m *mandrill) newMandrill(cfg Config) (*mandrill, error) {
	return &mandrill{
		cfg:    cfg,
		client: sendgrid.NewSendClient(cfg.ApiKey),
	}, nil
}

// Send
//
//
func (m *mandrill) Send(t Transmission) (Response, error) {
	sender := mail.NewV3Mail()

	// Add from
	from := mail.NewEmail(m.cfg.FromName, m.cfg.FromAddress)
	sender.SetFrom(from)

	// Add subject
	sender.Subject = t.Subject

	// Add to
	p := mail.NewPersonalization()
	var to []*mail.Email
	for _, recipient := range t.To {
		to = append(to, mail.NewEmail("", recipient))
	}
	p.AddTos(to...)

	// Add Plain Text
	content := mail.NewContent("text/plain", t.PlainText)
	sender.AddContent(content)

	// Add HTML
	html := mail.NewContent("text/html", t.HTML)
	sender.AddContent(html)

	response, err := m.client.Send(sender)
	if err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: response.StatusCode,
		Body:       response.Body,
		Headers:    response.Headers,
	}, nil
}
