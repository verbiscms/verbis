package importer

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/events"
)

func SendNewPassword(user domain.UserPart, password string, site domain.Site) error {
	mailer, err := events.NewChangedPassword()
	if err != nil {
		return err
	}
	return mailer.Send(user, password, site)
}