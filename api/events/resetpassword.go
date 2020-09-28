package events

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
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
func (e *ResetPassword) Send(u *domain.User, token string) error {
	const op = "events.ResetPassword.Send"

	tm := mail.Sender{
		To:      	[]string{u.Email},
		Subject: 	"Reset password",
		HTML: 		"<p>Reset password here</p>" +
			"<a href='" + environment.GetAppName() + "/" + config.Admin.Path + "/password/reset/" + token + "'>Reset</a>",
	}

	_, err := e.mailer.Send(&tm)
	if err != nil {
		return err
	}

	return nil
}
