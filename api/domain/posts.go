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
	"strconv"
	"strings"
	"time"
)

type PostData struct {
	Post     `json:"post"`
	Author   UserPart     `json:"author"`
	Category *Category    `json:"category"`
	Layout   []FieldGroup `json:"layout,omitempty"`
	Fields   []PostField  `json:"fields,omitempty"`
}

type Post struct {
	Id                int         `db:"id" json:"id" binding:"numeric"`
	UUID              uuid.UUID   `db:"uuid" json:"uuid"`
	Slug              string      `db:"slug" json:"slug" binding:"required,max=150"`
	Title             string      `db:"title" json:"title" binding:"required,max=500"`
	Status            string      `db:"status" json:"status,omitempty"`
	Resource          *string     `db:"resource" json:"resource,max=150"`
	PageTemplate      string      `db:"page_template" json:"page_template,omitempty" binding:"max=150"`
	PageLayout        string      `db:"layout" json:"layout,omitempty" binding:"max=150"`
	CodeInjectionHead *string     `db:"codeinjection_head" json:"codeinjection_head,omitempty"`
	CodeInjectionFoot *string     `db:"codeinjection_foot" json:"codeinjection_foot,omitempty"`
	UserId            int         `db:"user_id" json:"-"`
	PublishedAt       *time.Time  `db:"published_at" json:"published_at"`
	CreatedAt         *time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         *time.Time  `db:"updated_at" json:"updated_at"`
	SeoMeta           PostOptions `db:"options" json:"options"`
}

// TplPost defines the Post data for
// templates.
type PostTemplate struct {
	Post
	Author   UserPart
	Category *Category
	Fields   []PostField
}

func (p *PostData) Tpl() PostTemplate {
	return PostTemplate{
		Post:     p.Post,
		Author:   p.Author,
		Category: p.Category,
		Fields:   p.Fields,
	}
}

type PostCreate struct {
	Post
	Author   int         `json:"author,omitempty" binding:"numeric"`
	Category *int        `json:"category,omitempty" binding:"omitempty,numeric"`
	Fields   []PostField `json:"fields,omitempty"`
}

type PostField struct {
	Id            int         `db:"id" json:"-"`
	PostId        int         `db:"post_id" json:"-"`
	UUID          uuid.UUID   `db:"uuid" json:"uuid" binding:"required"`
	Type          string      `db:"type" json:"type"`
	Name          string      `db:"name" json:"name"`
	Key           string      `db:"field_key" json:"key"`
	Value         interface{} `json:"-"`
	OriginalValue FieldValue  `db:"value" json:"value"`
}

type FieldValue string

func (f PostField) TypeIsInArray(arr []string) bool {
	for _, v := range arr {
		if v == f.Type {
			return true
		}
	}
	return false
}

func (f FieldValue) Array() []string {
	return strings.Split(string(f), ",")
}

func (f FieldValue) IsEmpty() bool {
	return string(f) == ""
}

func (f FieldValue) String() string {
	return string(f)
}

func (f FieldValue) Int() (int, error) {
	const op = "FieldValue.Int"
	i, err := strconv.Atoi(f.String())
	if err != nil {
		return 0, &errors.Error{Code: errors.INVALID, Message: "Unable to cast FieldValue to an integer", Operation: op, Err: err}
	}
	return i, nil
}

type PostOptions struct {
	Id       int       `json:"-"`
	PageId   int       `json:"-" binding:"required|numeric"`
	Meta     *PostMeta `db:"meta" json:"meta"`
	Seo      *PostSeo  `db:"seo" json:"seo"`
	EditLock string    `db:"edit_lock" json:"edit_lock"`
}

type PostMeta struct {
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Twitter     PostTwitter  `json:"twitter,omitempty"`
	Facebook    PostFacebook `json:"facebook,omitempty"`
}

type PostTwitter struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImageId     int    `json:"image_id,numeric,omitempty"`
}

type PostFacebook struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImageId     int    `json:"image_id,numeric,omitempty"`
}

func (m *PostMeta) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scan not supported")
	}
	if bytes == nil || value == nil {
		return nil
	}
	return json.Unmarshal(bytes, &m)
}

func (m *PostMeta) Value() (driver.Value, error) {
	j, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal to domain.PostMeta")
	}
	return driver.Value(j), nil
}

type PostSeo struct {
	Public         bool   `json:"public"`
	ExcludeSitemap bool   `json:"exclude_sitemap"`
	Canonical      string `json:"canonical"`
}

func (m *PostSeo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scan not supported")
	}
	if bytes == nil || value == nil {
		return nil
	}
	return json.Unmarshal(bytes, &m)
}

func (m *PostSeo) Value() (driver.Value, error) {
	j, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal to domain.PostSeo")
	}
	return driver.Value(j), nil
}
