package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Form defines the form for sending form the API.
type Form struct {
	Id        int              `db:"id" json:"id" binding:"numeric"`
	UUID      uuid.UUID        `db:"uuid" json:"uuid"`
	Name      string           `db:"name" json:"name" binding:"required,max=500"`
	Fields    FormFields       `db:"fields" json:"fields"`
	Event     *string           `db:"event" json:"event"`
	CreatedAt *time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time       `db:"updated_at" json:"updated_at"`
}

type FormFields map[string]FormField

type FormField struct {
	Name      	string           `db:"name" json:"name"`
	Key      	string           `db:"key" json:"key"`
	Type      	string           `db:"type" json:"type"`
	Validation 	string           `db:"validation" json:"validation"`
	Required 	bool           `db:"required" json:"required"`
}

func (m FormFields) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scan not supported")
	}
	if bytes == nil || value == nil {
		return nil
	}
	return json.Unmarshal(bytes, &m)
}

func (m FormFields) Value() (driver.Value, error) {
	j, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal to domain.FormFields")
	}
	return driver.Value(j), nil
}