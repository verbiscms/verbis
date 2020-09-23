package logger

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	log "github.com/sirupsen/logrus"
	"os"
)

// init will determine if SuperAdmin and set logging levels
// dependant on environment variables.
func Init() error {

	// Only log panics if app debug is set to true
	if environment.IsDebug() {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.PanicLevel)
	}

	// Log json to file if environment is in production
	if !api.SuperAdmin {
		log.SetFormatter(&log.TextFormatter{
			ForceColors: true,
			DisableColors: false,
			FullTimestamp: true,
			TimestampFormat: "02-01-2006 15:04:05v",
		})
	} else {
		var filename string = paths.Storage() + "/logs/cms.log"
		// Create the log file if doesn't exist. And append to it if it already exists.
		f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
		if err != nil {
			return err
		}

		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat:"2006-01-02 15:04:05",
		})

		log.SetOutput(f)
	}

	return nil
}
