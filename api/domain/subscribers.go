package domain

import "time"

type Subscriber struct {
	ID			int			`db:"id" json:"id"`
	Email		string		`db:"email" json:"email" binding:"required,email"`
	CreatedAt	time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt	time.Time	`db:"updated_at" json:"updated_at"`
}


