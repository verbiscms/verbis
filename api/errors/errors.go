// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"bytes"
	"fmt"
)

// Application error codes.
const (
	CONFLICT = "conflict"  // Action cannot be performed
	INTERNAL = "internal"  // Internal error
	INVALID  = "invalid"   // Validation failed
	NOTFOUND = "not_found" // Entity does not exist
	TEMPLATE = "template"  // Templating error
)

// Global Error message when no message has been found.
const GlobalError = "An internal error has occurred."

// Error defines a standard application error.
type Error struct {
	Code      string   `json:"code"`
	Message   string   `json:"message"`
	Operation string   `json:"operation"`
	Err       error    `json:"error"`
}

// Error
//
// Returns the string representation of the error message.
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
	} else {
		if e.Code != "" {
			_, _ = fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}

	return buf.String()
}

// ErrorCode
//
// Returns the code of the root error, if available.
// Otherwise returns INTERNAL.
func Code(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return Code(e.Err)
	}
	return INTERNAL
}

// ErrorMessage
//
// Returns the human-readable message of the error, if
// available. Otherwise returns a generic error
// message.
func Message(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return Message(e.Err)
	}
	return GlobalError
}
