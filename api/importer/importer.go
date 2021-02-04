// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
	"net/url"
)

// Importer defines the method to migrate various CMS's content to Verbis.
type Importer interface {
	Import()
}

// ParseLink converts a link to of type *URL
// Returns errors.INVALID if the post link was unable to be parsed.
func ParseLink(link string) (string, error) {
	const op = "Importer.ParseLink"
	u, err := url.Parse(link)
	if err != nil {
		return "", &errors.Error{Code: errors.INVALID, Message: "Unable to parse post link", Operation: op, Err: err}
	}
	return u.Path, nil
}

// ParseUUID converts a string to of type uuid.UUID
// Returns errors.INVALID if the string was unable to be parsed.
func ParseUUID(u string) (uuid.UUID, error) {
	const op = "Importer.ParseUUID"
	id, err := uuid.Parse(u)
	if err != nil {
		return uuid.UUID{}, &errors.Error{Code: errors.INVALID, Message: "Could not pass UUID", Operation: op, Err: err}
	}
	return id, nil
}
