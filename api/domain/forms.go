// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
	"github.com/verbiscms/verbis/api/errors"
	"strings"
	"time"
)

type (
	// Form defines the data for sending data to the API
	// from the client side.
	Form struct {
		ID           int             `db:"id" json:"id" binding:"numeric"`
		UUID         uuid.UUID       `db:"uuid" json:"uuid"`
		Name         string          `db:"name" json:"name" binding:"required,max=500"`
		Fields       FormFields      `db:"fields" json:"fields"`
		Submissions  FormSubmissions `db:"-" json:"submissions"`
		EmailSend    types.BitBool   `db:"email_send" json:"email_send"`
		EmailMessage string          `db:"email_message" json:"email_message"`
		EmailSubject string          `db:"email_subject" json:"email_subject"`
		Recipients   string          `db:"recipients" json:"recipients"`
		StoreDB      types.BitBool   `db:"store_db" json:"store_db"`
		Body         interface{}     `db:"-" json:"-"`
		CreatedAt    time.Time       `db:"created_at" json:"created_at"`
		UpdatedAt    time.Time       `db:"updated_at" json:"updated_at"`
	}
	// Forms represents the slice of Form's.
	Forms []Form
	// FormField defines a individual field from the pivot
	// table.
	FormField struct {
		ID         int           `db:"id" json:"id" binding:"numeric"`
		UUID       uuid.UUID     `db:"uuid" json:"uuid"`
		FormID     int           `db:"form_id" json:"-"`
		Key        string        `db:"key" json:"key" binding:"required"`
		Label      FormLabel     `db:"label" json:"label" binding:"required"`
		Type       string        `db:"type" json:"type" binding:"required"`
		Validation string        `db:"validation" json:"validation"`
		Required   types.BitBool `db:"required" json:"required"`
		Options    DBMap         `db:"options" json:"options"`
	}
	// FormFields represents the slice of FormField's.
	FormFields []FormField
	// FormSubmission defines a submission of the of a form.
	FormSubmission struct {
		ID        int        `db:"id" json:"id" binding:"numeric"`
		UUID      uuid.UUID  `db:"uuid" json:"uuid"`
		FormID    int        `db:"form_id" json:"form_id"`
		Fields    FormValues `db:"fields" json:"fields"`
		IPAddress string     `db:"ip_address" json:"ip_address"`
		UserAgent string     `db:"user_agent" json:"user_agent"`
		SentAt    *time.Time `db:"sent_at" json:"sent_at"`
	}
	// FormSubmissions represents the slice of FormSubmission's.
	FormSubmissions []FormSubmission
	// FormLabel defines the label/name for form fields.
	FormLabel string
)

// GetRecipients Splits the recipients string and returns
// a slice of email addresses.
func (f *Form) GetRecipients() []string {
	return strings.FieldsFunc(f.Recipients, func(c rune) bool {
		return c == ','
	})
}

// Name converts the label to a dynamic-struct friendly
// name.
func (f FormLabel) Name() string {
	s := strings.ReplaceAll(string(f), " ", "")
	return strings.Title(s)
}

// String is thg stringer on the FormLabel type
func (f FormLabel) String() string {
	return string(f)
}

// FormValues represents the map of values.
type FormValues map[string]interface{}

// JSON Marshals the Form Values to bytes.
func (f FormValues) JSON() ([]byte, error) {
	const op = "FormValues.JSON"
	v, err := json.Marshal(f)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error processing the form fields for storing", Operation: op, Err: err}
	}
	return v, nil
}

// Scan implements the scanner for FormValues. unmarshal
// the FormValues when the entity is pulled from the
// database.
func (f FormValues) Scan(value interface{}) error {
	const op = "Domain.FormValues.Scan"
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok || bytes == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Scan unsupported for FormValues", Operation: op, Err: fmt.Errorf("scan not supported")}
	}
	err := json.Unmarshal(bytes, &f)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error unmarshalling into FormValues", Operation: op, Err: err}
	}
	return nil
}

// Value implements the valuer for FormValues. marshal the
// FormValues when the entity is inserted to the
// database.
func (f FormValues) Value() (driver.Value, error) {
	const op = "Domain.FormValues,Value"
	if len(f) == 0 {
		return nil, nil
	}
	j, err := marshaller(f)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling FormValues", Operation: op, Err: err}
	}
	return driver.Value(j), nil
}
