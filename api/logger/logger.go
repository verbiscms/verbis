// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
)

// logger is an alias for the the standard logger.
var logger = logrus.New()

// Init
//
// Init will determine if SuperAdmin and set logging levels
// dependant on environment variables.
func Init(env *environment.Env) {
	isDebug := env.IsDebug()

	// Set log level depending on SuperAdmin var
	if isDebug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logger.SetFormatter(&Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		Colours:         true,
		Debug:           isDebug,
	})

	// Send all logs to nowhere by default
	logger.SetOutput(ioutil.Discard)

	// Send logs with level higher than warning to stderr
	logger.AddHook(&WriterHook{
		Writer: os.Stderr,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	})

	// Send info and debug logs to stdout
	logger.AddHook(&WriterHook{
		Writer: os.Stderr,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})
}

// Trace
//
// Log a trace message with args.
func Trace(args ...interface{}) {
	logger.Trace(args...)
}

// Debug
//
// Log a debug message with args.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Info
//
// Log a info message with args.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn
//
// Log a warn message with args.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error
//
// Log a error message with args.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Fatal
//
// Log a fatal message with args.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic
//
// Log a panic message with args.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// WithField
//
//
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithFields(logrus.Fields{"fields": logrus.Fields{
		key: value,
	}})
}

// WithFields
//
//
func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(logrus.Fields{"fields": fields})
}

// WithError
//
//
func WithError(err interface{}) *logrus.Entry {
	return logger.WithField("error", err)
}

// SetOutput
//
//
func SetOutput(writer io.Writer) {
	logger.SetOutput(writer)
}
