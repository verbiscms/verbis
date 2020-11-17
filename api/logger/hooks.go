package logger

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	log "github.com/sirupsen/logrus"
	"io"
)

// WriterHook is a hook that writes logs of specified LogLevels to specified Writer
type WriterHook struct {
	Writer    io.Writer
	LogLevels []log.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *WriterHook) Fire(entry *log.Entry) error {

	if !environment.IsDebug() {
		if err := entry.Data["error"]; err != nil {
			e := entry.Data["error"].(errors.Error)

			m, err := json.Marshal(log.Fields{
				"level":     entry.Level,
				"code":      e.Code,
				"message":   e.Message,
				"operation": e.Operation,
				"err":       e.Err.Error(),
				"stack":     errors.Stack(&e),
				"time":      entry.Time.Format("2006-01-02 15:04:05"),
			})

			if err != nil {
				return err
			}

			str := string(m) + "\n"
			_, err = hook.Writer.Write([]byte(str))
			return err
		}
	}

	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *WriterHook) Levels() []log.Level {
	return hook.LogLevels
}
