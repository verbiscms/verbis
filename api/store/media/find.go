// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"database/sql"
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// Find
//
// Returns a media item by searching with the given ID.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the media item was not found by the given ID.
func (s *Store) Find(id int) (domain.Media, error) {
	const op = "MediaStore.Find"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("id", "=", id).
		Limit(1)

	var media domain.Media
	err := s.DB().Get(&media, q.Build())
	if err == sql.ErrNoRows {
		return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No media item exists with the ID: %d", id), Operation: op, Err: err}
	} else if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return media, nil
}

// FindByName
//
// Returns a media item by searching with the given name.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the media was not found by the given name.
func (s *Store) FindByName(name string) (domain.Media, error) {
	const op = "MediaStore.FindByPost"

	q := s.Builder().
		From(s.Schema()+TableName).
		Where("file_name", "=", name).
		Limit(1)

	var media domain.Media
	err := s.DB().Get(&media, q.Build())
	if err == sql.ErrNoRows {
		return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: "No media item exists with the name: " + name, Operation: op, Err: err}
	} else if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	return media, nil
}

// FindByURL
//
// Returns a media by searching with the given url.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the media item was not found by the given url.
func (s *Store) FindByURL(url string) (domain.Media, error) {
	const op = "MediaStore.FindBySlug"

	var media domain.Media

	// Test normal size.
	q := s.Builder().
		From(s.Schema()+TableName).
		Where("url", "=", url).
		Limit(1)

	err := s.DB().Get(&media, q.Build())
	if err == nil {
		return media, nil
	}

	// Test sizes.
	sq := s.Builder().
		From(s.Schema() + TableName).
		WhereRaw("WHERE sizes LIKE '%" + url + "%' LIMIT 1")

	err = s.DB().Get(&media, sq.Build())
	if err != nil {
		return domain.Media{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	}

	for _, v := range media.Sizes {
		if v.Url == url {
			return media, nil
		}
	}

	return domain.Media{}, &errors.Error{Code: errors.NOTFOUND, Message: "Error getting media item with the url: " + url, Operation: op, Err: fmt.Errorf("no media item exists")}
}
