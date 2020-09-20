package events

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/mail"
)

type ResetPassword struct {
	mailer *mail.Mailer
}

// Create a new verify email event.
func NewResetPassword() (*ResetPassword, error) {

	m, err := mail.New()
	if err != nil {
		return &ResetPassword{}, err
	}

	return &ResetPassword{
		mailer: m,
	}, nil
}

// Send the verify email event.
func (e *ResetPassword) Send(u *domain.User, token string) error {

	tm := mail.Sender{
		To:      	[]string{u.Email},
		Subject: 	"Reset password",
		HTML: 		"<p>Reset password here</p>" +
			"<a href='" + environment.Env.AppUrl + "/" + config.Admin.Path + "/password/reset/" + token + "'>Reset</a>",
	}

	_, err := e.mailer.Send(&tm)
	if err != nil {
		return err
	}

	return nil
}
