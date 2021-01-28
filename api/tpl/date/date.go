package date

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
	"time"
)

// date
//
// Format date by `interface{}` input, a date can be
// a `time.Time` or an `int, int32, int64`.
//
// Example: {{ date "02/01/2006" now }}
func (ns *Namespace) date(fmt string, date interface{}) (string, error) {
	return ns.dateInZone(fmt, date, "Local")
}

// dateInZone
//
// Takes a format, the date and zone. Casts the
// date to the correct format.
//
// Returns errors.TEMPLATE if the the interface could not be cast
// to type time.Time
//
// Example: {{ dateInZone "02/01/2006" now "Europe/London" }}
func (ns *Namespace) dateInZone(format string, date interface{}, zone string) (string, error) {
	const op = "Templates.dateInZone"

	tm, err := cast.ToTimeE(date)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Cannot cast interfaace to time.Time"), Operation: op, Err: err}
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}

	return tm.In(loc).Format(format), nil
}

// ago
//
// Returns a duration from the given time input
// in seconds. a date can be a `time.Time` or
// an `int, int64`.
//
// Example: {{ ago .UpdatedAt }}
func (ns *Namespace) ago(i interface{}) string {
	tm, err := cast.ToTimeE(i)
	if err != nil {
		return "0s"
	}

	return time.Since(tm).Round(time.Second).String()
}

// duration
//
// Formats a given amount of seconds as a `time.Duration`
// For example `duration 85` will return `1m25s`.
//
// Example: {{ duration 85 }}
func (ns *Namespace) duration(sec interface{}) string {
	tm, err := cast.ToDurationE(sec)
	if err != nil {
		return ""
	}
	return (tm * time.Second).String()
}

// htmlDate
//
// Format's a date for inserting into a HTML date
// picker input field.
//
// Example: {{ htmlDate now }}
func (ns *Namespace) htmlDate(date interface{}) (string, error) {
	return ns.dateInZone("2006-01-02", date, "Local")
}

// htmlDateInZone
//
// Returns HTML date with a timezone
//
// Example: {{ htmlDateInZone now "Europe/London" }}
func (ns *Namespace) htmlDateInZone(date interface{}, zone string) (string, error) {
	return ns.dateInZone("2006-01-02", date, zone)
}
