package events

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mail"
)

// ResetPassword defines the event instance for resetting passwords
type ResetPassword struct {
	mailer *mail.Mailer
}

// NewResetPassword creates a new reset password event.
func NewResetPassword() (*ResetPassword, error) {
	const op = "events.NewResetPassword"

	m, err := mail.New()
	if err != nil {
		return &ResetPassword{}, err
	}

	return &ResetPassword{
		mailer: m,
	}, nil
}

// Send the reset password event.
func (e *ResetPassword) Send(u *domain.User, url string, token string, title string) error {
	const op = "events.ResetPassword.Send"

	data := mail.Data{
		"AppUrl":    url,
		"AppTitle":  title,
		"AdminPath": e.mailer.Config.Admin.Path,
		"Token":     token,
		"UserName":  u.FirstName,
	}

	html, err := e.mailer.ExecuteHTML("reset-password.html", data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tm := mail.Sender{
		To:      []string{u.Email},
		Subject: "Reset password",
		HTML:    html,
	}

	e.mailer.Send(&tm)

	return nil
}
