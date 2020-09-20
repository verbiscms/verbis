package domain

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Options map[string]interface{}

type Option struct {
	ID				int					`db:"id" json:"id"`
	UUID 			uuid.UUID			`db:"uuid" json:"uuid"`
	Name			string 				`db:"option_name" json:"option_name" binding:"required"`
	Value			json.RawMessage		`db:"option_value" json:"option_value" binding:"required"`
}
