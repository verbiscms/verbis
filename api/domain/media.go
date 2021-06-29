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
		Id          int             `db:"id" json:"id"` //nolint
		UUID        uuid.UUID       `db:"uuid" json:"uuid"`
		Url         string          `db:"url" json:"url"` //nolint
		Title       string          `db:"title" json:"title"`
		Alt         string          `db:"alt" json:"alt"`
		Description string          `db:"description" json:"description"`
		FilePath    string          `db:"file_path" json:"-"`
		FileSize    int64           `db:"file_size" json:"file_size"`
		FileName    string          `db:"file_name" json:"file_name"`
		Sizes       MediaSizes      `db:"sizes" json:"sizes"`
		Mime        Mime            `db:"mime" json:"mime"`
		UserId      int             `db:"user_id" json:"user_id"` //nolint
		Location    StorageProvider `db:"location" json:"location"`
		CreatedAt   time.Time       `db:"created_at" json:"created_at"`
		UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
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
		FileSize int64     `db:"file_size" json:"file_size"`
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
	// WebPExtension defines the extension used for webp images.
	WebPExtension = ".webp"
)

// UploadPath returns the upload path of the media item
// without the storage uploads path, for example:
// 2020/01/photo.jpg
func (m *Media) UploadPath() string {
	if !m.IsOrganiseYearMonth() {
		return m.UUID.String() + m.Extension()
	}
	return m.FilePath + string(os.PathSeparator) + m.UUID.String() + m.Extension()
}

// IsOrganiseYearMonth returns a bool indicating if the
// file has been saved a year month path, i.e 2020/01.
func (m *Media) IsOrganiseYearMonth() bool {
	return m.FilePath != ""
}

// Extension returns the extension of the file by stripping
// from the url.
func (m *Media) Extension() string {
	return filepath.Ext(m.Url)
}

// PossibleFiles Returns a the possible files saved to the
// system after the files have been uploaded. Note: This
// does not include the upload path.
func (m *Media) PossibleFiles() []string {
	files := []string{
		m.UploadPath(),
		m.UploadPath() + ".webp",
	}
	for _, v := range m.Sizes {
		path := m.FilePath + string(os.PathSeparator) + v.UUID.String() + m.Extension()
		files = append(files, path, path+WebPExtension)
	}
	return files
}

// Mime is a string representation of a MIME type.
type Mime string

// CanResize Returns true if the mime type is of JPG or
// PNG, determining if the image can be resized.
func (m Mime) CanResize() bool {
	return m.IsJPG() || m.IsPNG()
}

// IsJPG returns true if the mime type is of JPG.
func (m Mime) IsJPG() bool {
	return m == "image/jpeg" || m == "image/jp2"
}

// IsPNG returns true if the mime type is of PNG.
func (m Mime) IsPNG() bool {
	return m == "image/png"
}

// String is the stringer on Mime type.
func (m Mime) String() string {
	return string(m)
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
