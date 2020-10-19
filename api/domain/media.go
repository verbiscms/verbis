package domain

import (
	"encoding/json"
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
	Sizes 			*json.RawMessage 	`db:"sizes" json:"sizes"`
	Type			string 				`db:"type" json:"type"`
	UserID			int					`db:"user_id" json:"user_id"`
	CreatedAt		time.Time			`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time			`db:"updated_at" json:"updated_at"`
}

type MediaSizes map[string]MediaSize

type MediaSizeDB struct {
	FilePath		string 				`db:"file_path" json:"file_path"`
	MediaSize
}

type MediaSize struct {
	UUID 			uuid.UUID			`db:"uuid" json:"uuid"`
	Url 			string				`db:"url" json:"url"`
	Name			string 				`db:"name" json:"name"`
	SizeName 		string 				`db:"size_name" json:"size_name"`
	FileSize		int 				`db:"file_size" json:"file_size"`
	Width			int 				`db:"width" json:"width"`
	Height			int 				`db:"height" json:"height"`
	Crop			bool 				`db:"crop" json:"crop"`
}

