// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mailer/events"
)

func SendNewPassword(user domain.UserPart, password string, site domain.Site) error {
	mailer, err := events.NewChangedPassword()
	if err != nil {
		return err
	}
	return mailer.Send(user, password, site)
}
