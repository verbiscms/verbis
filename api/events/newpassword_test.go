// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

func (t *EventTestSuite) Test_NewPasswordDispatch() {
	tt := map[string]struct {
		data  interface{}
		error bool
		want  interface{}
	}{
		"Success": {
			ChangedPassword{
				Password: "password",
				User:     domain.UserPart{},
			},
			false,
			nil,
		},
		"Validation failed": {
			event{},
			true,
			"ChangedPassword should be passed to dispatch",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			deps := t.Setup(test.error)
			dispatcher := NewChangedPassword(deps)
			err := dispatcher.Dispatch(test.data, []string{"hello@verbiscms.com"}, nil)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}
