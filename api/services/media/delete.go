// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
)

// Delete satisfies the Library to remove possible testMedia
// item combinations from the file system, if the file
// does not exist (user moved) it will be
// skipped and logged out.
func (s *Service) Delete(id int) error {
	// Find the original testMedia item
	item, err := s.repo.Find(id)
	if err != nil {
		return err
	}

	// Remove from the database
	err = s.repo.Delete(item.Id)
	if err != nil {
		return err
	}

	// Delete the files from storage
	go s.deleteFiles(item)

	return nil
}

// delete files removes all files from the storage bucket.
func (s *Service) deleteFiles(item domain.Media) {
	const op = "Service.DeleteFiles"

	// Remove original file
	err := s.storage.Delete(item.File.Id)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting original testMedia item: " + item.File.Url, Operation: op, Err: err}).Error()
	}
	logger.Info("Deleted original testMedia item: " + item.File.Url)

	// Delete original WebP
	s.deleteWebP(item.File)

	// Remove testMedia sizes
	for _, size := range item.Sizes {
		// Delete original
		err = s.storage.Delete(size.File.Id)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting testMedia size: " + size.File.Url, Operation: op, Err: err}).Error()
		}
		logger.Info("Deleted testMedia size: " + size.File.Url)

		// Delete sized WebP
		s.deleteWebP(size.File)
	}
}

// deleteWebP removes any webp images from the bucket.
func (s *Service) deleteWebP(file domain.File) {
	const op = "Service.DeleteWebP"

	_, webp, err := s.storage.Find(file.Url + domain.WebPExtension)
	if err != nil {
		return
	}

	err = s.storage.Delete(webp.Id)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting webp image: " + webp.Url, Operation: op, Err: err}).Error()
		return
	}

	logger.Info("Deleted WebP file: " + webp.Url)
}
