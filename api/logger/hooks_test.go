// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/sirupsen/logrus"
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
		err   interface{}
		want  interface{}
	}{
		"Error Entry": {
			&bytes.Buffer{},
			&logrus.Entry{
				Logger: &logrus.Logger{Formatter: &mockFormatErr{}},
			},
			&errors.Error{Code: errors.INTERNAL, Message: "Error obtaining the entry string", Operation: "logger.Hook.Fire", Err: fmt.Errorf("err")},
			"",
		},
		"Error Writer": {
			&mockWriterErr{},
			&logrus.Entry{
				Logger: &logrus.Logger{Formatter: &mockFormat{}},
			},
			&errors.Error{Code: errors.INTERNAL, Message: "Error writing entry to io.Writer", Operation: "logger.Hook.Fire", Err: fmt.Errorf("err")},
			"",
		},
		"Success": {
			buf,
			&logrus.Entry{
				Logger: &logrus.Logger{Formatter: &mockFormat{}},
			},
			nil,
			"test",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			h := t.SetupHooks(test.input)
			err := h.Fire(test.entry)
			if test.err != nil {
				t.Equal(test.err, err)
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
