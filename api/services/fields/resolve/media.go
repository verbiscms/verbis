// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

//
//
// media resolves a media from the given value.
// Returns the domain.Media if it was found and no error occurred.
// Returns errors.INVALID if the domain.FieldValue could not be cast to an integer.
func (v *Value) media(value domain.FieldValue) (interface{}, error) {
	const op = "FieldResolver.media"

	id, err := value.Int()
	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Unable to cast user ID to an integer", Operation: op, Err: err}
	}

	media, err := v.deps.Store.Media.Find(id)
	if err != nil {
		return nil, err
	}

	return media.Public(), nil
}
