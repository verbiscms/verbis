package events

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mail"
)

// ChangedPassword defines the event instance for new passwords reset by the system
type ChangedPassword struct {
	mailer *mail.Mailer
}

// NewPassword creates a new reset password event
func NewChangedPassword() (*ChangedPassword, error) {
	const op = "events.NewResetPassword"

	m, err := mail.New()
	if err != nil {
		return &ChangedPassword{}, err
	}

	return &ChangedPassword{
		mailer: m,
	}, nil
}

// Send the reset password event.
func (e *ChangedPassword) Send(u domain.UserPart, password string, site domain.Site) error {
	const op = "events.ResetPassword.Send"

	data := mail.Data{
		"AppUrl":    site.Url,
		"AppTitle":  site.Title,
		"AdminPath": e.mailer.Config.Admin.Path,
		"UserName":  u.FirstName,
		"Password":  password,
		"Email":     u.Email,
	}

	html, err := e.mailer.ExecuteHTML("new-password.html", data)
	if err != nil {
		return err
	}

	tm := mail.Sender{
		To:      []string{u.Email},
		Subject: fmt.Sprintf("New Login Details for %s", site.Title),
		HTML:    html,
	}

	e.mailer.Send(&tm)

	return nil
}
