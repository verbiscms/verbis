package logger

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
)

// logFiles defines the files to be written to when app debug is set to false
type logFiles struct {
	accessLog io.Writer
	errorLog io.Writer
}

// Init will determine if SuperAdmin and set logging levels
// dependant on environment variables.
func Init() error {

	// Set log level depending on SuperAdmin var
	if api.SuperAdmin {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Log json to file if environment is in production if not,
	// log to console.
	if environment.IsDebug() {
		log.SetFormatter(&Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			Colours: 		true,
		})
	} else {
		logs, err := getLogFiles()
		if err != nil {
			return err
		}

		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		setupLogs(logs)
	}

	return nil
}

// getLogFiles checks to see if the log files set out in the logs.yml file
// in the configuration is valid. If it is set to default (storage/logs)
// and the file does not exist, it will create them. If set to a
// custom path, will return an error if not found.
func getLogFiles() (*logFiles, error) {

	logFiles := logFiles{}

	// Access Log
	configAccess := config.Logs.AccessLog
	if configAccess == "default" {
		path := paths.Storage() + "/logs/access.log"
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			return nil, fmt.Errorf("Could not create the acccess.log file with the path %s", path)
		}

		logFiles.accessLog = f
	} else {
		if exists := files.Exists(configAccess); !exists {
			return nil, fmt.Errorf("The access log with the path %s, could not be found.", configAccess)
		} else {
			f, err := os.OpenFile(configAccess, os.O_APPEND|os.O_WRONLY, 0644)

			if err != nil {
				return nil, err
			}

			logFiles.accessLog = f
		}
	}

	// Error log
	configError := config.Logs.ErrorLog
	if configError == "default" {
		path := paths.Storage() + "/logs/error.log"
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			return nil, fmt.Errorf("Could not create the error.log file with the path %s", path)
		}

		logFiles.errorLog = f
	} else {
		if exists := files.Exists(configError); !exists {
			return nil, fmt.Errorf("The error log with the path %s, could not be found.", configError)
		} else {
			f, err := os.OpenFile(configAccess, os.O_APPEND|os.O_WRONLY, 0644)

			if err != nil {
				return nil, err
			}

			logFiles.errorLog = f
		}
	}

	return &logFiles, nil
}

// setupLogs adds hooks to send logs to different destinations depending on level
func setupLogs(logs *logFiles) {

	// Send all logs to nowhere by default
	log.SetOutput(ioutil.Discard)

	// Send logs with level higher than warning to stderr
	log.AddHook(&WriterHook{
		Writer: logs.errorLog,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})

	// Send info and debug logs to stdout
	log.AddHook(&WriterHook{
		Writer: logs.accessLog,
		LogLevels: []log.Level{
			log.InfoLevel,
			log.DebugLevel,
		},
	})
}

