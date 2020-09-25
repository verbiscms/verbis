package logger

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)


// Formatter implements logrus.Formatter interface.
type Formatter struct {
	Colours bool
	TimestampFormat string
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}

	b.WriteString("[VERBIS] ")

	// Print the time
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.StampMilli
	}
	b.WriteString(entry.Time.Format(timestampFormat))

	// Get the response code colour
	code := entry.Data["status_code"].(int)
	cc := color.Style{}
	switch code {
		case 200: {
			cc = color.Style{color.FgLightWhite, color.BgGreen, color.OpBold}
			break
		}
		default: {
			cc = color.Style{color.FgLightWhite, color.BgRed, color.OpBold}
		}
	}

	// Print the response code
	b.WriteString(" |")
	if f.Colours {
		b.WriteString(cc.Sprintf(strconv.Itoa(code)))
	} else {
		b.WriteString(fmt.Sprintf(strconv.Itoa(code)))
	}
	b.WriteString("| ")

	// Get level the colour
	lc := color.Style{}
	switch entry.Level {
		case logrus.DebugLevel: {
			lc = color.Style{color.FgGray, color.OpBold}
			break
		}
		case logrus.WarnLevel: {
			lc = color.Style{color.FgYellow, color.OpBold}
			break
		}
		case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel: {
			lc = color.Style{color.FgRed, color.OpBold}
			break
		}
		default: {
			lc = color.Style{color.FgBlue, color.OpBold}
		}
	}

	// Print the level

	level := strings.ToUpper(entry.Level.String())
	if f.Colours {
		b.WriteString(lc.Sprintf("["))
		b.WriteString(lc.Sprintf(level))

		if entry.Level == logrus.InfoLevel {
			b.WriteString(lc.Sprintf("] "))
		} else {
			b.WriteString(lc.Sprintf("]"))
		}
	} else {
		b.WriteString(fmt.Sprintf("["))
		b.WriteString(fmt.Sprintf(level))

		if entry.Level == logrus.InfoLevel {
			b.WriteString(fmt.Sprintf("] "))
		} else {
			b.WriteString(fmt.Sprintf("]"))
		}
	}

	// Print the IP
	ip := entry.Data["client_ip"].(string)
	b.WriteString(fmt.Sprintf(" | %s | ", ip))

	// Print the method
	method := entry.Data["request_method"].(string)
	rc := color.Style{color.FgLightWhite, color.BgBlue, color.OpBold}

	if len(method) == 3 {
		if f.Colours {
			b.WriteString(rc.Sprintf("  %s   ", method))
		} else {
			b.WriteString(fmt.Sprintf("  %s   ", method))
		}
	} else {
		if f.Colours {
			b.WriteString(rc.Sprintf("  %s  ", method))
		} else {
			b.WriteString(fmt.Sprintf("  %s  ", method))
		}
	}

	// Print the url
	url := entry.Data["request_url"].(string)
	b.WriteString(fmt.Sprintf(" \"%s\" | ", url))

	// Print the message
	msg := entry.Data["message"].(string)
	b.WriteString(fmt.Sprintf("[msg] %s |", msg))

	// Print any errors if one is set
	errorData := entry.Data["error"].(errors.Error)
	if errorData.Error() != "" {

		if errorData.Code != errors.NOTFOUND {

			if errorData.Code != "" {
				if f.Colours {
					b.WriteString(color.Red.Sprintf(" [code] %s", errorData.Code))
				} else {
					b.WriteString(fmt.Sprintf(" [code] %s", errorData.Code))
				}

			}

			if errorData.Operation != "" {
				if api.SuperAdmin {
					if f.Colours {
						b.WriteString(color.Red.Sprintf(" [operation] %s", errorData.Operation))
					} else {
						b.WriteString(fmt.Sprintf(" [operation] %s", errorData.Operation))
					}
				}
			}

			if errorData.Err != nil {
				if f.Colours {
					b.WriteString(color.Red.Sprintf(" [error] %s", errorData.Err.Error()))
				} else {
					b.WriteString(fmt.Sprintf(" [error] %s", errorData.Err.Error()))
				}
			}
		}
	}

	b.WriteString("\n")

	return b.Bytes(), nil
}
