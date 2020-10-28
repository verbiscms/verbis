package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Media defines the media entity for interacting with the database
// and sending data back
type Media struct {
	Id				int					`db:"id" json:"id"`
	UUID 			uuid.UUID			`db:"uuid" json:"uuid"`
	Url 			string				`db:"url" json:"url"`
	Title			*string 			`db:"title" json:"title"`
	Alt				*string 			`db:"alt" json:"alt"`
	Description		*string 			`db:"description" json:"description"`
	FilePath		string 				`db:"file_path" json:"-"`
	FileSize		int 				`db:"file_size" json:"file_size"`
	FileName		string 				`db:"file_name" json:"file_name"`
	Sizes 			MediaSizes		 	`db:"sizes" json:"sizes"`
	Type			string 				`db:"type" json:"type"`
	UserID			int					`db:"user_id" json:"user_id"`
	CreatedAt		time.Time			`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time			`db:"updated_at" json:"updated_at"`
}

type MediaSize struct {
	UUID 			uuid.UUID			`db:"uuid" json:"uuid"`
	Url 			string				`db:"url" json:"url"`
	Name			string 				`db:"name" json:"name"`
	SizeName 		string 				`db:"size_name" json:"size_name"`
	FileSize		int 				`db:"file_size" json:"file_size"`
	FilePath		string 				`db:"file_path" json:"-"`
	Width			int 				`db:"width" json:"width"`
	Height			int 				`db:"height" json:"height"`
	Crop			bool 				`db:"crop" json:"crop"`
}

type MediaSizeOptions struct {
	Name			string 				`db:"name" json:"name" binding:"required,numeric"`
	Width			int 				`db:"width" json:"width" binding:"required,numeric"`
	Height			int 				`db:"height" json:"height" binding:"required,numeric"`
	Crop			bool 				`db:"crop" json:"crop"`
}

type MediaSizes map[string]MediaSize

func (m MediaSizes) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scan not supported")
	}
	if bytes == nil {
		return nil
	}
	return json.Unmarshal(bytes, &m)
}

func (m MediaSizes) Value() (driver.Value, error) {
	if len(m) == 0 {
		return nil, nil
	}
	j, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal to domain.MediaSizes")
	}

	return driver.Value(j), nil
}
