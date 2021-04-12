// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

func (t *EventTestSuite) Test_FormSendDispatch() {
	tt := map[string]struct {
		data  interface{}
		error bool
		want  interface{}
	}{
		"Success": {
			FormSend{
				Form:   &domain.Form{Id: 1},
				Values: domain.FormValues{},
			},
			false,
			nil,
		},
		"Send Error": {
			FormSend{
				Form:   &domain.Form{Id: 1},
				Values: domain.FormValues{},
			},
			true,
			errors.GlobalError,
		},
		"Nil Form": {
			FormSend{
				Form:   nil,
				Values: domain.FormValues{},
			},
			true,
			"Form cannot be nil",
		},
		"Validation failed": {
			event{},
			true,
			"FormSend should be passed to dispatch",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			deps := t.Setup(test.error)
			dispatcher := NewFormSend(deps)
			err := dispatcher.Dispatch(test.data, []string{"hello@verbiscms.com"}, nil)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}
