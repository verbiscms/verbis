package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)


type PostData struct {
	Post 			Post						`json:"post"`
	Author 			PostAuthor 					`json:"author"`
	Categories 		[]PostCategory 				`json:"categories"`
	Layout 			*[]FieldGroup				`json:"layout"`
}

type Post struct {
	Id				int							`db:"id" json:"id" binding:"numeric"`
	UUID 			uuid.UUID					`db:"uuid" json:"uuid"`
	Slug			string 						`db:"slug" json:"slug" binding:"required,max=150"`
	Title			string 						`db:"title" json:"title" binding:"required,max=500"`
	Status			string						`db:"status" json:"status,omitempty"`
	Resource		*string 					`db:"resource" json:"resource,max=150"`
	PageTemplate	string						`db:"page_template" json:"page_template,omitempty" binding:"max=150"`
	Layout			string						`db:"layout" json:"layout,omitempty" binding:"max=150"`
	Fields			*json.RawMessage 			`db:"fields" json:"fields"`
	CodeInjectHead	*string 					`db:"codeinjection_head" json:"codeinjection_head,omitempty"`
	CodeInjectFoot	*string 					`db:"codeinjection_foot" json:"codeinjection_foot,omitempty"`
	UserId 			int 						`db:"user_id" json:"-"`
	PublishedAt		*time.Time					`db:"published_at" json:"published_at"`
	CreatedAt		*time.Time					`db:"created_at" json:"created_at"`
	UpdatedAt		*time.Time					`db:"updated_at" json:"updated_at"`
	SeoMeta			PostSeoMeta					`db:"options" json:"options"`
}

type PostAuthor struct {
	Id					int				`json:"id" binding:"required,numeric"`
	UUID 				uuid.UUID		`db:"uuid" json:"uuid"`
	FirstName			string 			`json:"first_name"`
	LastName			string 			`json:"last_name"`
	Email				string			`json:"email"`
	Password			string			`json:"-"`
	Website				string			`json:"website"`
	Facebook			*string			`json:"facebook"`
	Twitter				*string			`json:"twitter"`
	Linkedin			*string			`json:"linked_in"`
	Instagram			*string			`json:"instagram"`
	ProfilePictureID	*int			`json:"profile_picture_id"`
	Token				string			`json:"-"`
	TokenLastUsed		*time.Time		`json:"-"`
	Role				UserRole 		`json:"role"`
	EmailVerifiedAt		*time.Time		`json:"email_verified_at"`
	CreatedAt			time.Time		`json:"created_at"`
	UpdatedAt			time.Time		`json:"updated_at"`
}

type PostCategory struct {
	Id				int			`json:"id"`
	Slug			string 		`json:"slug"`
	Name			string 		`json:"name"`
	Description		*string 	`json:"description"`
	Hidden 			*bool 		`json:"hidden"`
	ParentId 		*int 		`json:"parent_id"`
	PageTemplate	*string 	`json:"page_template"`
	UpdatedAt		time.Time	`json:"updated_at"`
	CreatedAt		time.Time	`json:"created_at"`
}

type PostSeoMeta struct {
	Id				int							`json:"-"`
	PageId			int							`json:"-" binding:"required|numeric"`
	Seo				*json.RawMessage			`db:"seo" json:"seo"`
	Meta			*json.RawMessage			`db:"meta" json:"meta"`
}

type PostMeta struct {
	Title   			string `json:"title,omitempty"`
	Description 		string `json:"description,omitempty"`
	Twitter struct {
		Title       	string `json:"title,omitempty"`
		Description 	string `json:"description,omitempty"`
		ImageId      	int `json:"image_id,omitempty"`
	} `json:"twitter,omitempty"`
	Facebook struct {
		Title       	string `json:"title,omitempty"`
		Description 	string `json:"description,omitempty"`
		ImageId      	int `json:"image_id,omitempty"`
	} `json:"facebook,omitempty"`
}

type PostCreate struct {
	Post
	Author 			int    		`json:"author,omitempty" binding:"numeric"`
	Categories		[]int		`json:"categories,omitempty" binding:"unique"`
}

//xss in golang implementation





