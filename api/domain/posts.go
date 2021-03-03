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
	"github.com/jmoiron/sqlx/types"
	"time"
)

type (
	// Post defines the main page entity of Verbis.
	Post struct {
		Id                int           `db:"id" json:"id" binding:"numeric"` //nolint
		UUID              uuid.UUID     `db:"uuid" json:"uuid"`
		Slug              string        `db:"slug" json:"slug" binding:"required,max=150"`
		Title             string        `db:"title" json:"title" binding:"required,max=500"`
		Status            string        `db:"status" json:"status,omitempty"`
		Resource          *string       `db:"resource" json:"resource"`
		PageTemplate      string        `db:"page_template" json:"page_template,omitempty" binding:"max=150"`
		PageLayout        string        `db:"layout" json:"layout,omitempty" binding:"max=150"`
		CodeInjectionHead *string       `db:"codeinjection_head" json:"codeinjection_head,omitempty"`
		CodeInjectionFoot *string       `db:"codeinjection_foot" json:"codeinjection_foot,omitempty"`
		UserId            int           `db:"user_id" json:"-"` //nolint
		IsArchive         types.BitBool `db:"archive" json:"archive"`
		PublishedAt       *time.Time    `db:"published_at" json:"published_at"`
		CreatedAt         *time.Time    `db:"created_at" json:"created_at"`
		UpdatedAt         *time.Time    `db:"updated_at" json:"updated_at"`
		SeoMeta           PostOptions   `db:"options" json:"options"`
	}
	// Posts represents the slice of Post's.
	Posts []Post
	// PostDatum defines the post including author, category,
	// layout and field information.
	PostDatum struct {
		Post     `json:"post"`
		Author   UserPart     `json:"author"`
		Category *Category    `json:"category"`
		Layout   []FieldGroup `json:"layout,omitempty"`
		Fields   PostFields   `json:"fields,omitempty"`
	}
	// PostData represents the slice of PostDatum's.
	PostData []PostDatum
	// PostField defines the individual field that is attached
	// to a post.
	PostField struct {
		Id            int         `db:"id" json:"-"`      //nolint
		PostId        int         `db:"post_id" json:"-"` //nolint
		UUID          uuid.UUID   `db:"uuid" json:"uuid" binding:"required"`
		Type          string      `db:"type" json:"type"`
		Name          string      `db:"name" json:"name"`
		Key           string      `db:"field_key" json:"key"`
		Value         interface{} `json:"-"`
		OriginalValue FieldValue  `db:"value" json:"value"`
	}
	// PostFields represents the slice of PostField's.
	PostFields []PostField
	// PostCreate defines the data when a post is created.
	PostCreate struct {
		Post
		Author   int         `json:"author,omitempty" binding:"numeric"`
		Category *int        `json:"category,omitempty" binding:"omitempty,numeric"`
		Fields   []PostField `json:"fields,omitempty"`
	}
	// PostOptions defines the global post options that
	// includes post meta and post seo information.
	PostOptions struct {
		Id       int       `json:"-"`                            //nolint
		PageId   int       `json:"-" binding:"required|numeric"` //nolint
		Meta     *PostMeta `db:"meta" json:"meta"`
		Seo      *PostSeo  `db:"seo" json:"seo"`
		EditLock string    `db:"edit_lock" json:"edit_lock"`
	}
	// PostMeta defines the global meta information for the
	// post used when calling the VerbisHeader.
	PostMeta struct {
		Title       string       `json:"title,omitempty"`
		Description string       `json:"description,omitempty"`
		Twitter     PostTwitter  `json:"twitter,omitempty"`
		Facebook    PostFacebook `json:"facebook,omitempty"`
	}
	// PostTwitter defines the twitter meta information
	// used when calling the VerbisHeader.
	PostTwitter struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		ImageId     int    `json:"image_id,omitempty" binding:"numeric"` //nolint
	}
	// PostFacebook defines the opengraph meta information
	// used when calling the VerbisHeader.
	PostFacebook struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		ImageId     int    `json:"image_id,omitempty" binding:"numeric"` //nolint
	}
	// PostSeo defines the options for Seo on the post,
	// including if the post is indexable, if it
	// should appear in the sitemap and any
	// canonical overrides.
	PostSeo struct {
		Public         bool   `json:"public"`
		ExcludeSitemap bool   `json:"exclude_sitemap"`
		Canonical      string `json:"canonical"`
	}
	// PostTemplate defines the Post data for templates when they
	// are used in the front end.
	PostTemplate struct {
		Post
		Author   UserPart
		Category *Category
		Fields   []PostField
	}
)

// IsPublic
//
// Determines if the post is published.
func (p *Post) IsPublic() bool {
	return p.Status == "published"
}

// HasCategories
//
// Determines if a post has any resources attached
// to it.
func (p *PostDatum) HasResource() bool {
	return p.Resource != nil
}

// HasCategories
//
// Determines if a post has any categories attached
// to it.
func (p *PostDatum) HasCategories() bool {
	return p.Category != nil
}

// IsHomepage
//
// Determines if the post is the homepage by comparing
// the domain options.
func (p *PostDatum) IsHomepage(id int) bool {
	return id == p.Post.Id
}

// Tpl
//
// Converts a PostDatum to a PostTemplate and hides
// layouts.
func (p *PostDatum) Tpl() PostTemplate {
	return PostTemplate{
		Post:     p.Post,
		Author:   p.Author,
		Category: p.Category,
		Fields:   p.Fields,
	}
}

// TypeIsInSlice
//
// Determines if the given field values is in the slice
// passed.
func (f *PostField) TypeIsInSlice(arr []string) bool {
	for _, v := range arr {
		if v == f.Type {
			return true
		}
	}
	return false
}

// Scan
//
// Scanner for PostMeta. unmarshal the PostMeta when
// the entity is pulled from the database.
func (m *PostMeta) Scan(value interface{}) error {
	const op = "Domain.PostMeta.Scan"
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok || bytes == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Scan unsupported for PostMeta", Operation: op, Err: fmt.Errorf("scan not supported")}
	}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error unmarshalling into PostMeta", Operation: op, Err: err}
	}
	return nil
}

// Value
//
// Valuer for PostMeta. marshal the PostMeta when
// the entity is inserted to the database.
func (m *PostMeta) Value() (driver.Value, error) {
	const op = "Domain.PostMeta.Value"
	j, err := json.Marshal(m)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling PostMeta", Operation: op, Err: err}
	}
	return driver.Value(j), nil
}

// Scan
//
// Scanner for PostSeo. unmarshal the PostSeo when
// the entity is pulled from the database.
func (m *PostSeo) Scan(value interface{}) error {
	const op = "Domain.PostSeo.Scan"
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok || bytes == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Scan unsupported for PostSeo", Operation: op, Err: fmt.Errorf("scan not supported")}
	}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error unmarshalling into PostSeo", Operation: op, Err: err}
	}
	return nil
}

// Value
//
// Valuer for PostSeo. marshal the PostSeo when
// the entity is inserted to the database.
func (m *PostSeo) Value() (driver.Value, error) {
	const op = "Domain.PostSeo.Value"
	j, err := json.Marshal(m)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling PostSeo", Operation: op, Err: err}
	}
	return driver.Value(j), nil
}
