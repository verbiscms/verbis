package domain

import "time"

type (
	// Redirects defines the data used for redirecting http
	// requests.
	Redirect struct {
		Id        int        `db:"id" json:"id"` //nolint
		From      string     `db:"from_path" json:"from_path" binding:"required"`
		To        string     `db:"to_path" json:"to_path" binding:"required"`
		Code      int        `db:"code" json:"code" binding:"required,numeric"`
		CreatedAt *time.Time `db:"created_at" json:"created_at"`
		UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	}
	// Redirects represents the slice of Redirect's.
	Redirects []Redirect
)
