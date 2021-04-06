// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mail

import (
	"github.com/mailgun/mailgun-go/v4"
	"golang.org/x/net/context"
	"time"
)

type mailGun struct {
	cfg    Config
	client *mailgun.MailgunImpl
}

// Init
//
//
func (m *mailGun) newMailGun(cfg Config) (*mailGun, error) {
	return &mailGun{
		cfg:    cfg,
		client: mailgun.NewMailgun(cfg.Domain, cfg.ApiKey),
	}, nil

}

// Send
//
//
func (m *mailGun) Send(t Transmission) (Response, string, error) {

	// The message object allows you to add attachments and Bcc recipients
	message := m.client.NewMessage(m.cfg.FromAddress, t.Subject, t.HTML, t.To...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	msg, id, err := m.client.Send(ctx, message)
	if err != nil {
		return Response{}, "", err
	}

	return Response{
		StatusCode: 0,
		Body:       "",
		Headers:    nil,
		Id:         id,
		Message:    msg,
	}, id, nil
}
