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
	"os"
	"path/filepath"
	"time"
)

type (
	// Media defines the core media entity for Verbis.
	Media struct {
		Id          int        `db:"id" json:"id"` //nolint
		UUID        uuid.UUID  `db:"uuid" json:"uuid"`
		Url         string     `db:"url" json:"url"` //nolint
		Title       string     `db:"title" json:"title"`
		Alt         string     `db:"alt" json:"alt"`
		Description string     `db:"description" json:"description"`
		FilePath    string     `db:"file_path" json:"-"`
		FileSize    int        `db:"file_size" json:"file_size"`
		FileName    string     `db:"file_name" json:"file_name"`
		Sizes       MediaSizes `db:"sizes" json:"sizes"`
		Type        string     `db:"type" json:"type"`
		UserId      int        `db:"user_id" json:"user_id"` //nolint
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
		Url      string    `db:"url" json:"url"` //nolint
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

const (
	WebPExtension = ".webp"
)

// UploadPath
//
// Returns the upload path of the media item without the
// storage uploads path, for example:
// 2020/01/photo.jpg
func (m *Media) UploadPath() string {
	return m.FilePath + string(os.PathSeparator) + m.UUID.String() + m.Extension()
}

// Extension
//
// Returns the extension of the file by stripping from
// the URL.
func (m *Media) Extension() string {
	return filepath.Ext(m.Url)
}

// PossibleFiles
//
//
func (m *Media) PossibleFiles() []string {
	files := []string{
		m.UploadPath(),
		m.UploadPath() + ".webp",
	}
	for _, v := range m.Sizes {
		path := m.FilePath + string(os.PathSeparator) + v.UUID.String() + m.Extension()
		files = append(files, path)
		files = append(files, path + WebPExtension)
	}
	return files
}

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
