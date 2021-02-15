package domain

import "time"

type Redirect struct {
	Id   int64   `db:"id" json:"id"`
	From string `db:"from_path" json:"from" binding:"required"`
	To   string `db:"to_path" json:"to" binding:"required"`
	Code int    `db:"code" json:"code" binding:"required,numeric"`
	CreatedAt         *time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         *time.Time  `db:"updated_at" json:"updated_at"`
}
