package logger

import (
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	log "github.com/sirupsen/logrus"
	"os"
)

func Init() error {

	// Only log panics if app debug is set to true
	if environment.IsDebug() {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.PanicLevel)
	}

	// Log json to file if environment is in production
	if environment.IsProduction() {
		var filename string = paths.Storage() + "/logs/cms.log"
		// Create the log file if doesn't exist. And append to it if it already exists.
		f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(f)
	} else {
		log.SetFormatter(&log.TextFormatter{
			ForceColors: true,
			DisableColors: false,
			FullTimestamp: true,
			TimestampFormat: "02-01-2006 15:04:05v",
		})
	}

	return nil
}
