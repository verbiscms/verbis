// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import "github.com/ainsleyclark/verbis/api/domain"

// Delete
//
// Satisfies the Library to remove possible media item
// combinations from the file system, if the file
// does not exist (user moved) it will be
// skipped.
//
// Logs errors.INTERNAL if the file could not be deleted.
func (s *Service) Delete(media domain.Media) {
	const op = "Service.Delete"

	//items := item.PossibleFiles(s.paths.Uploads)
	//for _, path := range items {
	//	err := s.storage.Delete(path)
	//
	//	// Exists func
	//	if err != nil {
	//		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting file with the path: " + path, Operation: op, Err: err})
	//	}
	//
	//	logger.Debug("Deleted file with the path: " + path)
	//}
}
