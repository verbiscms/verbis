// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
	"strings"
	"time"
)

// Form defines the form for sending form the API.
type Form struct {
	Id           int           `db:"id" json:"id" binding:"numeric"`
	UUID         uuid.UUID     `db:"uuid" json:"uuid"`
	Name         string        `db:"name" json:"name" binding:"required,max=500"`
	Fields       []FormField   `db:"fields" json:"fields"`
	EmailSend    types.BitBool `db:"email_send" json:"email_send"`
	EmailMessage string        `db:"email_message" json:"email_message"`
	EmailSubject string        `db:"email_subject" json:"email_subject"`
	StoreDB      types.BitBool `db:"store_db" json:"store_db"`
	Body         interface{}   `db:"-" json:"-"`
	CreatedAt    *time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time    `db:"updated_at" json:"updated_at"`
}

// FormField defines a field from the pivot table.
type FormField struct {
	Id         int           `db:"id" json:"id" binding:"numeric"`
	UUID       uuid.UUID     `db:"uuid" json:"uuid"`
	FormId     int           `db:"form_id" json:"-"`
	Key        string        `db:"key" json:"key" binding:"required"`
	Label      FormLabel     `db:"label" json:"label" binding:"required"`
	Type       string        `db:"type" json:"type" binding:"required"`
	Validation *string       `db:"validation" json:"validation"`
	Required   types.BitBool `db:"required" json:"required"`
	Options    DBMap         `db:"options" json:"options"`
}

// FormSubmission defines a submission of the form.
type FormSubmission struct {
	Id        int        `db:"id" json:"id" binding:"numeric"`
	UUID      uuid.UUID  `db:"uuid" json:"uuid"`
	FormId    int        `db:"form_id" json:"form_id"`
	Fields    DBMap      `db:"fields" json:"fields"`
	IpAddress string     `db:"ip_address" json:"ip_address"`
	UserAgent string     `db:"user_agent" json:"user_agent"`
	SentAt    *time.Time `db:"sent_at" json:"sent_at"`
}

// FormLabel defines the label/name for form fields.
type FormLabel string

// Name converts the label to a dynamic-struct friendly name.
func (f FormLabel) Name() string {
	s := strings.ReplaceAll(string(f), " ", "")
	return strings.Title(s)
}
