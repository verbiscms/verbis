// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// Delete satisfies the Library to remove possible media
// item combinations from the file system, if the file
// does not exist (user moved) it will be
// skipped and logged out.
func (s *Service) Delete(id int) error {
	// Find the original media item
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
	const op = "Media.DeleteFiles"

	// Remove original file
	err := s.storage.Delete(item.File.Id)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting original media item: " + item.File.Url, Operation: op, Err: err}).Error()
	}
	logger.Info("Deleted original media item: " + item.File.Url)

	// Delete original WebP
	s.deleteWebP(item.File, true)

	// Remove media sizes
	for _, size := range item.Sizes {
		// Delete original
		err = s.storage.Delete(size.File.Id)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting media size: " + size.File.Url, Operation: op, Err: err}).Error()
		}
		logger.Info("Deleted media size: " + size.File.Url)

		// Delete sized WebP
		s.deleteWebP(size.File, true)
	}

	err = s.repo.Delete(item.Id)
	if err != nil {
		logger.WithError(err).Error()
	}
}

// deleteWebP removes any webp images from the bucket.
func (s *Service) deleteWebP(file domain.File, log bool) {
	const op = "Media.DeleteWebP"

	url := file.Url + domain.WebPExtension

	_, webp, err := s.storage.Find(url)
	if err != nil {
		return
	}

	err = s.storage.Delete(webp.Id)
	if err != nil {
		if log {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error deleting webp image: " + webp.Url, Operation: op, Err: err}).Error()
		}
		return
	}

	logger.Info("Deleted WebP file: " + webp.Url)
}
