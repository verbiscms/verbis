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
		Sizes       MediaSizes `db:"-" json:"sizes"`
		UserId      int        `db:"user_id" json:"user_id"` //nolint
		FileId      int        `db:"file_id" json:"-"`       //nolint
		File        File       `db:"file" json:"file"`
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
		Id       int           `db:"id" json:"id"`      //nolint
		FileId   int           `db:"file_id" json:"-"`  //nolint
		MediaId  int           `db:"media_id" json:"-"` //nolint
		SizeKey  string        `db:"size_key" json:"-" binding:"required,numeric"`
		SizeName string        `db:"size_name" json:"name" binding:"required,numeric"`
		Width    int           `db:"width" json:"width" binding:"required,numeric"`
		Height   int           `db:"height" json:"height" binding:"required,numeric"`
		Crop     types.BitBool `db:"crop" json:"crop"`
		File     File          `db:"file" json:"file"`
	}
	// MediaPublic represents a media item sent back to the
	// frontend or API.
	MediaPublic struct {
		Id          int              `json:"id"` //nolint
		Title       string           `json:"title"`
		Alt         string           `json:"alt"`
		Description string           `json:"description"`
		Sizes       MediaSizesPublic `json:"sizes"`
		UserId      int              `json:"user_id"` //nolint
		Url         string           `json:"url"`     //nolint
		Name        string           `json:"name"`
		Path        string           `json:"path"`
		Mime        Mime             `json:"mime"`
		FileSize    int64            `json:"file_size"`
		CreatedAt   time.Time        `json:"created_at"`
		UpdatedAt   time.Time        `json:"updated_at"`
	}
	// MediaSizesPublic represents media sizes sent back to
	// the frontend or API.
	MediaSizesPublic map[string]MediaSizePublic
	// MediaSizePublic represents a media size sent back to
	// the frontend or API.
	MediaSizePublic struct {
		Name     string        `json:"name"`
		Width    int           `json:"width"`
		Height   int           `json:"height"`
		Crop     types.BitBool `json:"crop"`
		URL      string        `json:"url"`
		Path     string        `json:"path"`
		Mime     Mime          `json:"mime"`
		FileSize int64         `json:"file_size"`
	}
)

const (
	// WebPExtension defines the extension used for WebP images.
	WebPExtension = ".webp"
)

// Public converts a Media type to a public struct.
func (m *Media) Public() MediaPublic {
	var ms = make(MediaSizesPublic, len(m.Sizes))
	for _, v := range m.Sizes {
		ms[v.SizeKey] = MediaSizePublic{
			Name:     v.SizeName,
			Width:    v.Width,
			Height:   v.Height,
			Crop:     v.Crop,
			URL:      v.File.Url,
			Mime:     v.File.Mime,
			FileSize: v.File.FileSize,
		}
	}

	if len(ms) == 0 {
		ms = nil
	}

	return MediaPublic{
		Id:          m.Id,
		Title:       m.Title,
		Alt:         m.Alt,
		Description: m.Description,
		Sizes:       ms,
		UserId:      m.UserId,
		Url:         m.File.Url,
		Name:        m.File.Name,
		Mime:        m.File.Mime,
		FileSize:    m.File.FileSize,
		UpdatedAt:   m.UpdatedAt,
		CreatedAt:   m.CreatedAt,
	}
}

// Public converts a MediaItems type to a slice of
// public structs.
func (m MediaItems) Public() []MediaPublic {
	var mp = make([]MediaPublic, len(m))
	for i, v := range m {
		mp[i] = v.Public()
	}
	return mp
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
