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

// deleteItem
//
// Removes possible file combinations from the file
// system.
func deleteItem(item domain.Media, uploadPath string) {
	const op = "client.Delete"

	items := item.PossibleFiles()
	for _, v := range items {
		path := uploadPath + string(os.PathSeparator) + v

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}

		err = os.Remove(path)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting file with the path: " + v, Operation: op, Err: err})
		}
	}
}
