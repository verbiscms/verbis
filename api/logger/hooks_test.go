// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/verbiscms/verbis/api/errors"
	"io"
)

type mockFormatErr struct{}

func (m *mockFormatErr) Format(entry *logrus.Entry) ([]byte, error) {
	return nil, fmt.Errorf("err")
}

type mockFormat struct{}

func (m *mockFormat) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte("test"), nil
}

type mockWriterErr struct{}

func (m *mockWriterErr) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("err")
}

func (t *LoggerTestSuite) TestWriterHook_Fire() {
	buf := &bytes.Buffer{}

	tt := map[string]struct {
		input io.Writer
		entry *logrus.Entry
		want  interface{}
	}{
		"Error Entry": {
			&bytes.Buffer{},
			&logrus.Entry{
				Logger: &logrus.Logger{Formatter: &mockFormatErr{}},
			},
			"Error obtaining the entry string",
		},
		"Error Writer": {
			&mockWriterErr{},
			&logrus.Entry{
				Logger: &logrus.Logger{Formatter: &mockFormat{}},
			},
			"Error writing entry to io.Writer",
		},
		"Success": {
			buf,
			&logrus.Entry{
				Logger: &logrus.Logger{Formatter: &mockFormat{}},
			},
			"test",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			h := t.SetupHooks(test.input)
			err := h.Fire(test.entry)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, buf.String())
		})
	}
}

func (t *LoggerTestSuite) TestWriterHook_Levels() {
	h := t.SetupHooks(nil)
	want := []logrus.Level{
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
	t.Equal(want, h.Levels())
}
