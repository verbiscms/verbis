// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/environment"
	"io"
	"io/ioutil"
	"os"
)

var (
	// logger is an alias for the the standard logger.
	logger = logrus.New()
)

// Init will set up the logger and set logging levels
// dependant on environment variables.
func Init(env *environment.Env) {
	isDebug := env.IsDebug()

	// Set log level depending on Debug or Production.
	if !api.Production {
		logger.SetLevel(logrus.TraceLevel)
	} else if isDebug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	logger.SetFormatter(&Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		Colours:         true,
		Debug:           isDebug,
	})

	// Send all logs to nowhere by default.
	logger.SetOutput(ioutil.Discard)

	// Send logs with level higher than warning to stderr.
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
		Writer: os.Stdout,
		LogLevels: []logrus.Level{
			logrus.TraceLevel,
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})
}

// Trace logs a trace message with args.
func Trace(args ...interface{}) {
	logger.Trace(args...)
}

// Debug logs a debug message with args.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Info logs a info message with args.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn logs a warn message with args.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error logs a error message with args.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Fatal logs a fatal message with args.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic logs a panic message with args.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// WithField logs with field, sets a new map containing
// "fields".
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithFields(logrus.Fields{"fields": logrus.Fields{
		key: value,
	}})
}

// WithFields logs with fields, sets a new map containing
// "fields".
func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(logrus.Fields{"fields": fields})
}

// WithError - Logs with a Verbis error.
func WithError(err interface{}) *logrus.Entry {
	return logger.WithField("error", err)
}

// SetOutput sets the output of the logger to an io.Writer,
// useful for testing.
func SetOutput(writer io.Writer) {
	logger.SetOutput(writer)
}

// SetLevel sets the level of the logger.
func SetLevel(level logrus.Level) {
	logger.SetLevel(level)
}

// SetLogger sets the application logger.
func SetLogger(l *logrus.Logger) {
	logger = l
}
