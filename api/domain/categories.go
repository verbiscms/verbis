package domain

import (
	"github.com/google/uuid"
	"time"
)

// Category defines the categories for posts
type Category struct {
	Id				int			`db:"id" json:"id"`
	UUID 			uuid.UUID	`db:"uuid" json:"uuid"`
	Slug			string 		`db:"slug" json:"slug" binding:"required,max=150"`
	Name			string 		`db:"name" json:"name" binding:"required,max=150"`
	Description		*string 	`db:"description" json:"description,max=500"`
	Resource 		string		`db:"resource" json:"resource" binding:"required,max=150"`
	ParentId 		*int 		`db:"parent_id" json:"parent_id"`
	//ArchiveId		*int 		`db:"archive_id" json:"archive_id"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
}


