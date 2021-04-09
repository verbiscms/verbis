// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	mocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"testing"
)

// EventTestSuite defines the helper used for event
// testing.
type EventTestSuite struct {
	test.HandlerSuite
}

// TestAuth
//
// Assert testing has begun.
func TestEvent(t *testing.T) {
	suite.Run(t, &EventTestSuite{})
}

func (t *EventTestSuite) Test_EventExecuteHTML() {
	tt := map[string]struct {
		mock func(m *mocks.TemplateExecutor, buf *bytes.Buffer)
		want interface{}
	}{
		"Success": {
			func(m *mocks.TemplateExecutor, buf *bytes.Buffer) {
				m.On("Execute", buf, mock.Anything, nil).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("tpl"))
					t.NoError(err)
				}).Return(mock.Anything, nil)
			},
			"tpl",
		},
		"Error": {
			func(m *mocks.TemplateExecutor, buf *bytes.Buffer) {
				m.On("Execute", buf, mock.Anything, nil).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("tpl"))
					t.NoError(err)
				}).Return(mock.Anything, fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			mt := &mocks.TemplateHandler{}
			m := &mocks.TemplateExecutor{}
			mt.On("Prepare", mock.Anything).Return(m)

			var buf = &bytes.Buffer{}
			test.mock(m, buf)

			d := &deps.Deps{}
			d.SetTmpl(mt)
			e := Event{
				Deps: d,
			}

			got, err := e.executeHTML("test", nil)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}
