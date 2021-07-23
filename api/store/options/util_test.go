// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"encoding/json"
	"github.com/verbiscms/verbis/api/errors"
)

func (t *OptionsTestSuite) TestStore_Unmarshal() {
	tt := map[string]struct {
		input json.RawMessage
		want  interface{}
	}{
		"Success": {
			json.RawMessage("\"test\""),
			nil,
		},
		"Error": {
			nil,
			"Error unmarshalling the option value",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil)
			_, err := s.unmarshal(test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Nil(err)
		})
	}
}

func (t *OptionsTestSuite) TestStore_Marshal() {
	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Success": {
			1,
			json.RawMessage("1"),
		},
		"Error": {
			make(chan int, 1),
			"Error marshalling the option value",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil)
			got, err := s.marshal(test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
