// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/fields/layout"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/spf13/cast"
)

// GetGroup
//
func (s *Service) GetLayout(name string, args ...interface{}) domain.Field {
	l, err := layout.ByName(name, s.handleLayoutArgs(args))
	if err != nil {
		logger.WithError(err).Error()
		return domain.Field{}
	}
	return l
}

// GetGroups
//
func (s *Service) GetLayouts(args ...interface{}) domain.FieldGroups {
	return s.handleLayoutArgs(args)
}

// handleLayoutArgs
//
func (s *Service) handleLayoutArgs(args []interface{}) domain.FieldGroups {
	switch len(args) {
	case 1:
		return s.getLayoutByPost(args[0])
	default:
		return s.layout
	}
}

// getLayoutByPost
//
// Returns the layout by Post with the given ID.
// Logs errors.INVALID if the id failed to be cast to an int.
// Logs if the post if was not found or there was an error obtaining/formatting the post.
func (s *Service) getLayoutByPost(id interface{}) domain.FieldGroups {
	const op = "FieldsService.getFieldsByPost"

	i, err := cast.ToIntE(id)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Unable to cast Post ID to integer", Operation: op, Err: err}).Error()
		return nil
	}

	post, err := s.deps.Store.Posts.GetByID(i, true)
	if err != nil {
		logger.WithError(err).Error()
		return nil
	}

	return post.Layout
}
