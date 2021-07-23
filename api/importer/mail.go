// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/verbiscms/verbis/api/domain"
)

func SendNewPassword(user domain.UserPart, password string, site domain.Site) error {
	//dispatcher := events.NewChangedPassword(&events.NewChangedPassword(&deps.Deps{}))
	//dispatcher.Dispatch()
	//if err != nil {
	//	return err
	//}
	//return mailer.Send(user, password, site)
	return nil
}
