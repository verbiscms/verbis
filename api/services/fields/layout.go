// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/services/fields/layout"
)

// GetLayout returns a field group layout by the given name.
// Logs errors.NOTFOUND if the layout could not be found by the given name.
func (s *Service) GetLayout(name string, args ...interface{}) domain.Field {
	l, err := layout.ByName(name, s.handleLayoutArgs(args))
	if err != nil {
		logger.WithError(err).Error()
		return domain.Field{}
	}
	return l
}

// GetLayouts returns all field group layouts.
func (s *Service) GetLayouts(args ...interface{}) domain.FieldGroups {
	return s.handleLayoutArgs(args)
}

// handleLayoutArgs processes the args for finding layouts.
func (s *Service) handleLayoutArgs(args []interface{}) domain.FieldGroups {
	switch len(args) {
	case 1:
		return s.getLayoutByPost(args[0])
	default:
		return s.layout
	}
}

// getLayoutByPost returns the layout by post with the given ID.
// Logs errors.INVALID if the id failed to be cast to an int.
// Logs if the post if was not found or there was an error obtaining/formatting the post.
func (s *Service) getLayoutByPost(id interface{}) domain.FieldGroups {
	const op = "FieldsService.GetLayoutByPost"

	i, err := cast.ToIntE(id)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Unable to cast post ID to integer", Operation: op, Err: err}).Error()
		return nil
	}

	post, err := s.deps.Store.Posts.Find(i, true)
	if err != nil {
		logger.WithError(err).Error()
		return nil
	}

	return post.Layout
}
