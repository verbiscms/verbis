// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/jmoiron/sqlx/types"
	"time"
)

type (
	// Media defines the core media entity for Verbis.
	Media struct {
		Id          int        `db:"id" json:"id"` //nolint
		Title       string     `db:"title" json:"title"`
		Alt         string     `db:"alt" json:"alt"`
		Description string     `db:"description" json:"description"`
		Sizes       MediaSizes `db:"sizes" json:"sizes"`
		UserId      int        `db:"user_id" json:"user_id"` //nolint
		StorageId   int        `db:"storage_id" json:"-"`    //nolint
		CreatedAt   time.Time  `db:"created_at" json:"created_at"`
		UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
		File
	}
	// MediaItems represents the slice of Media.
	MediaItems []Media
	// MediaSizes defines the map of MediaSizes, by key value
	// pair.
	MediaSizes map[string]MediaSize
	// MediaSize defines an individual media size that's
	// stored in the database.
	MediaSize struct {
		Id     int           `db:"id" json:"id"` //nolint
		FileId int           `db:"file_id" json:"file_id"`
		Key    string        `db:"size_key" json:"-" binding:"required,numeric"`
		Name   string        `db:"size_name" json:"name" binding:"required,numeric"`
		Width  int           `db:"width" json:"width" binding:"required,numeric"`
		Height int           `db:"height" json:"height" binding:"required,numeric"`
		Crop   types.BitBool `db:"crop" json:"crop"`
		File
	}
)

const (
	// WebPExtension defines the extension used for WebP images.
	WebPExtension = ".webp"
)

// PossibleFiles Returns a the possible files saved to the
// system after the files have been uploaded. Note: This
// does not include the upload path.
func (m *Media) PossibleFiles(prefix string) []string {
	files := []string{
		m.UploadPath(prefix),
		m.UploadPath(prefix) + WebPExtension,
	}
	for _, v := range m.Sizes {
		path := v.UploadPath(prefix)
		files = append(files, path, path+WebPExtension)
	}
	return files
}

// Scan implements the scanner for MediaSizes. unmarshal
// the MediaSizes when the entity is pulled from the
// database.
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

// Value implements the valuer for MediaSizes. marshal the
// MediaSizes when the entity is inserted to the
// database.
func (m MediaSizes) Value() (driver.Value, error) {
	const op = "Domain.MediaSizes.Value"
	if len(m) == 0 {
		return nil, nil
	}
	j, err := marshaller(m)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling MediaSizes", Operation: op, Err: err}
	}
	return driver.Value(j), nil
}
