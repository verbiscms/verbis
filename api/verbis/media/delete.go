// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"os"
)

// Delete
//
//
func (c *Library) Delete(item domain.Media) {
	const op = "Client.Delete"

	items := item.PossibleFiles()
	go func() {
		for _, v := range items {
			err := os.Remove(v)
			if err != nil {
				logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting file with the path: " + v, Operation: op, Err: err})
			}
		}
	}()
}
