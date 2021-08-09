// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
)

func (t *LoggerTestSuite) TestInit() {
	tt := map[string]struct {
		production bool
		env        environment.Env
		want       interface{}
	}{
		"Not Production": {
			false,
			environment.Env{AppDebug: "true"},
			"trace",
		},
		"Not Debug": {
			true,
			environment.Env{AppDebug: "false"},
			"info",
		},
		"Debug": {
			true,
			environment.Env{AppDebug: "true"},
			"debug",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			orig := api.Production
			defer func() {
				logger = nil
				logger = logrus.New()
				api.Production = orig
			}()
			api.Production = test.production
			Init(&test.env)
			t.Equal(test.want, logger.Level.String())
		})
	}
}

func (t *LoggerTestSuite) TestLogger() {
	tt := map[string]struct {
		fn   func()
		want string
	}{
		"Trace": {
			func() {
				Trace("trace")
			},
			"trace",
		},
		"Debug": {
			func() {
				Debug("debug")
			},
			"debug",
		},
		"Info": {
			func() {
				Info("info")
			},
			"info",
		},
		"Warn": {
			func() {
				Warn("warning")
			},
			"warning",
		},
		"Error": {
			func() {
				Error("error")
			},
			"error",
		},
		"With Field": {
			func() {
				WithField("test", "field").Error()
			},
			"field",
		},
		"With Fields": {
			func() {
				WithFields(logrus.Fields{"test": "field"}).Error()
			},
			"field",
		},
		"With Error": {
			func() {
				WithError(&errors.Error{Code: "code", Message: "message", Operation: "op", Err: fmt.Errorf("err")}).Error()
			},
			"[code] code [msg] message [op] op [error] err",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			buf := t.Setup()
			test.fn()
			t.Contains(buf.String(), test.want)
		})
	}
}

func (t *LoggerTestSuite) TestLogger_Fatal() {
	buf := t.Setup() // nolint
	defer func() {
		logger = logrus.New()
	}()
	logger.ExitFunc = func(i int) {}
	Fatal("fatal")
	t.Contains(buf.String(), "fatal")
}

func (t *LoggerTestSuite) TestLogger_Panic() {
	buf := t.Setup()
	t.Panics(func() {
		Panic("panic")
	})
	t.Contains(buf.String(), "panic")
}

func (t *LoggerTestSuite) TestLogger_SetOutput() {
	buf := &bytes.Buffer{}
	SetOutput(buf)
	t.Equal(buf, logger.Out)
}

func (t *LoggerTestSuite) TestSetLevel() {
	defer func() {
		logger = logrus.New()
	}()
	SetLevel(logrus.WarnLevel)
	t.Equal(logrus.WarnLevel, logger.GetLevel())
}

func (t *LoggerTestSuite) TestSetLogger() {
	defer func() {
		logger = logrus.New()
	}()
	l := logger
	SetLogger(l)
	t.Equal(l, logger)
}
