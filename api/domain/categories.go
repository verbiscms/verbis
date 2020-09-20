package domain

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	Id				int			`db:"id" json:"id"`
	UUID 			uuid.UUID	`db:"uuid" json:"uuid"`
	Slug			string 		`db:"slug" json:"slug" binding:"required,max=150,alphanum"`
	Name			string 		`db:"name" json:"name" binding:"required,max=150,alphanum"`
	Description		*string 	`db:"description" json:"description,max=500"`
	Hidden 			*bool 		`db:"hidden" json:"hidden"`
	ParentId 		*int 		`db:"parent_id" json:"parent_id"`
	PageTemplate	*string 	`db:"page_template" json:"page_template,max=150,alphanum"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
}