// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/ainsleyclark/verbis/api/environment"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// L is an alias for the the standard logger.
var L = log.New()

// Init
//
// Init will determine if SuperAdmin and set logging levels
// dependant on environment variables.
func Init() error {
	addHooks()

	// Set log level depending on SuperAdmin var
	if environment.IsDebug() {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(&Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		Colours:         true,
	})

	return nil
}

// addHooks()
//
// setupLogs adds hooks to send logs to different destinations depending on level
func addHooks() {

	// Send all logs to nowhere by default
	log.SetOutput(ioutil.Discard)

	// Send logs with level higher than warning to stderr
	log.AddHook(&WriterHook{
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})

	// Send info and debug logs to stdout
	log.AddHook(&WriterHook{
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.InfoLevel,
			log.DebugLevel,
		},
	})
}
