package errors

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime"
)

// Application error codes.
const (
	CONFLICT   = "conflict"   // Action cannot be performed
	INTERNAL   = "internal"   // Internal error
	INVALID    = "invalid"    // Validation failed
	NOTFOUND   = "not_found"  // Entity does not exist
)

// Error defines a standard application error.
type Error struct {
	Code    	string
	Message 	string
	Operation 	string
	Err      	error
	Stack 		[]string
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Operation != "" {
		fmt.Fprintf(&buf, "%s: ", e.Operation)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	}


	return buf.String()
}

// ErrorCode returns the code of the root error, if available. Otherwise returns INTERNAL.
func ErrorCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return INTERNAL
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred."
}

// Error returns the string representation of the error message.
func ErrorLog(err error) string {
	var buf bytes.Buffer

	e := err.(*Error)

	// Print the current operation in our stack, if any.
	if e.Operation != "" {
		fmt.Fprintf(&buf, "%s: ", e.Operation)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			_, _ = fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}

	return buf.String()
}

// ErrorStack returns the stack from which the error was called from.
func ErrorStack(err error) []string {
	var stack []string
	for c := 2; c < 5; c++ {
		_, file, _, _ := runtime.Caller(c)
		stack = append(stack, file)
	}
	return stack
}

// Report the error to logging.
func ErrorReport(err error) {

	var returnErr string = ""
	if err.Error() != "" {
		returnErr = err.Error()
	}

	e := err.(*Error)

	log.WithFields(log.Fields{
		"code"		: ErrorCode(err),
		"message"	: ErrorLog(err),
		"operation" : e.Operation,
		"err"		: returnErr,
		"stack"		: ErrorStack(e),
	}).Error()
}
