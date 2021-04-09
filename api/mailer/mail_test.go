// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mailer

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/go-mail"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	tt := map[string]struct {
		input *environment.Env
		want  interface{}
	}{
		"Nil Environment": {
			nil,
			"Environment can't be nil",
		},
		"SparkPost": {
			&environment.Env{
				MailDriver:      mail.SparkPost,
				MailFromAddress: "hello@verbiscms.com",
				MailFromName:    "name",
				SparkpostAPIKey: "key",
				SparkpostURL:    "https://api.eu.sparkpost.com",
			},
			mail.SparkPost,
		},
		"MailGun": {
			&environment.Env{
				MailDriver:      mail.MailGun,
				MailFromAddress: "hello@verbiscms.com",
				MailFromName:    "name",
				MailGunAPIKey:   "key",
				MailGunURL:      "https://api.eu.sparkpost.com",
				MailGunDomain:   "domain",
			},
			mail.MailGun,
		},
		"SendGrid": {
			&environment.Env{
				MailDriver:      mail.SendGrid,
				MailFromAddress: "hello@verbiscms.com",
				MailFromName:    "name",
				SendGridAPIKey:  "key",
			},
			mail.SendGrid,
		},
		"No Driver": {
			&environment.Env{
				MailDriver: "wrong",
			},
			"No mail driver exists",
		},
		"New Client Error": {
			&environment.Env{
				MailDriver: mail.SparkPost,
			},
			"Error validating mail configuration",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := New(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Equal(t, test.want, got.Driver)
			assert.NotNil(t, got.client)
			assert.Equal(t, test.input.MailFromAddress, got.FromAddress)
			assert.Equal(t, test.input.MailFromName, got.FromName)
			assert.Equal(t, test.input, got.Env)
			assert.NotNil(t, got.Paths)
		})
	}
}

type mockMailError struct{}

func (m *mockMailError) Send(t *mail.Transmission) (mail.Response, error) {
	return mail.Response{}, fmt.Errorf("error")
}

func TestMail_Send_Error(t *testing.T) {
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)

	m := Mail{
		client: &mockMailError{},
	}

	m.Send(&mail.Transmission{})

	assert.Contains(t, buf.String(), "Mail.Send: error")
}

type mockMailSuccess struct{}

func (m *mockMailSuccess) Send(t *mail.Transmission) (mail.Response, error) {
	return mail.Response{}, nil
}

func TestMail_Send_Success(t *testing.T) {
	logger.Init(&environment.Env{
		AppDebug: "true",
	})

	buf := &bytes.Buffer{}
	logger.SetOutput(buf)

	m := Mail{
		client: &mockMailSuccess{},
	}

	m.Send(&mail.Transmission{})

	assert.Contains(t, buf.String(), "Successfully sent email with the subject")
}
