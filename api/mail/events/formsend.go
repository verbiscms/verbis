package events

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mail"
)

// FormSend defines the event instance for emailing forms
type FormSend struct {
	mailer *mail.Mailer
	config config.Configuration
}

// NewVerifyEmail creates a new verify email event
func NewFormSend(config config.Configuration) (*FormSend, error) {
	const op = "events.NewFormSend"

	m, err := mail.New()
	if err != nil {
		return &FormSend{}, err
	}

	return &FormSend{
		mailer: m,
		config: config,
	}, nil
}

// Send the verify email event.
func (e *FormSend) Send(form *domain.Form) error {
	const op = "events.VerifyEmail.Send"

	html, err := e.mailer.ExecuteHTML("form-send.html", nil)
	if err != nil {
		return err
	}

	tm := mail.Sender{
		To:      []string{"ainsley@reddico.co.uk"},
		Subject: form.EmailSubject,
		HTML:    html,
	}

	e.mailer.Send(&tm)

	return nil
}
