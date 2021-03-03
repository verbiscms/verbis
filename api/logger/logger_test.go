// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/sirupsen/logrus"
)

// func (t *LoggerTestSuite) TestInit() {
//
//	tt := map[string]struct {
//		entry *logrus.Entry
//		fn    func()
//		want  interface{}
//	}{
//		"Not Debug": {
//			&logrus.Entry{
//				Logger: &logrus.Logger{Formatter: &mockFormat{}},
//			},
//			func() {
//				os.Setenv("APP_DEBUG", "false")
//			},
//			"",
//		},
//		"Debug": {
//			&logrus.Entry{
//				Logger: &logrus.Logger{Formatter: &mockFormatErr{}},
//			},
//			func() {
//				os.Setenv("APP_DEBUG", "true")
//			},
//			"DEBUG",
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			// TODO Environment in STRUCT!
//
//			test.fn()
//			Init()
//			buf := t.Setup()
//			logrus.Debug("test")
//			t.Contains(buf.String(), test.want)
//		})
//	}
//}

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

//nolint
func (t *LoggerTestSuite) TestLogger_Fatal() {
	//buf := t.Setup()
	//
	//defer func() {
	//	logger = logrus.New()
	//}()
	//logger.Fatal = func(args ...interface{}) {
	//
	//}
	//
	//Fatal("fatal")
	//t.Contains(buf.String(), "fatal")
}

func (t *LoggerTestSuite) TestLogger_Panic() {
	buf := t.Setup()
	t.Panics(func() {
		Panic("panic")
	})
	t.Contains(buf.String(), "panic")
}
