// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
	"time"
)

type (
	// Media defines the core media entity for Verbis.
	Media struct {
		Id          int        `db:"id" json:"id"`
		UUID        uuid.UUID  `db:"uuid" json:"uuid"`
		Url         string     `db:"url" json:"url"`
		Title       string     `db:"title" json:"title"`
		Alt         string     `db:"alt" json:"alt"`
		Description string     `db:"description" json:"description"`
		FilePath    string     `db:"file_path" json:"-"`
		FileSize    int        `db:"file_size" json:"file_size"`
		FileName    string     `db:"file_name" json:"file_name"`
		Sizes       MediaSizes `db:"sizes" json:"sizes"`
		Type        string     `db:"type" json:"type"`
		UserID      int        `db:"user_id" json:"user_id"`
		CreatedAt   time.Time  `db:"created_at" json:"created_at"`
		UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	}
	// MediaItems represents the slice of Media.
	MediaItems []Media
	// MediaSizes defines the map of MediaSizes, by key value
	// pair.
	MediaSizes map[string]MediaSize
	// MediaSize defines an individual media size that's
	// stored in the database.
	MediaSize struct {
		UUID     uuid.UUID `db:"uuid" json:"uuid"`
		Url      string    `db:"url" json:"url"`
		Name     string    `db:"name" json:"name"`
		SizeName string    `db:"size_name" json:"size_name"`
		FileSize int       `db:"file_size" json:"file_size"`
		Width    int       `db:"width" json:"width"`
		Height   int       `db:"height" json:"height"`
		Crop     bool      `db:"crop" json:"crop"`
	}
	// MediaSizeOptions defines the options for saving different
	// image sizes when uploaded.
	MediaSizeOptions struct {
		Name   string `db:"name" json:"name" binding:"required,numeric"`
		Width  int    `db:"width" json:"width" binding:"required,numeric"`
		Height int    `db:"height" json:"height" binding:"required,numeric"`
		Crop   bool   `db:"crop" json:"crop"`
	}
)

// Scan
//
// Scanner for MediaSize. unmarshal the MediaSize when
// the entity is pulled from the database.
func (m MediaSizes) Scan(value interface{}) error {
	const op = "Domain.MediaSizes.Scan"
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok || bytes == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Scan unsupported for MediaSize", Operation: op, Err: fmt.Errorf("scan not supported")}
	}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error unmarshalling into MediaSize", Operation: op, Err: err}
	}
	return nil
}

// Value
//
// Valuer for MediaSize. marshal the MediaSize when
// the entity is inserted to the database.
func (m MediaSizes) Value() (driver.Value, error) {
	const op = "Domain.MediaSizes.Value"
	if len(m) == 0 {
		return nil, nil
	}
	j, err := json.Marshal(m)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling MediaSizes", Operation: op, Err: err}
	}
	return driver.Value(j), nil
}
