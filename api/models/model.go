package models

import (
	"github.com/ainsleyclark/verbis/api/database"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	Auth       		AuthRepository
	Categories 		CategoryRepository
	Fields     		FieldsRepository
	Media 			MediaRepository
	Options 		OptionsRepository
	Posts      		PostsRepository
	Roles      		RoleRepository
	Session 		SessionRepository
	Site      		SiteRepository
	User       		UserRepository
}

// Create a new database instance, connect to database.
func New(db *database.MySql) *Store {
	return &Store{
		Auth:       newAuth(db.Sqlx),
		Categories: newCategories(db.Sqlx),
		Fields:     newFields(db.Sqlx),
		Media: 		newMedia(db.Sqlx),
		Options:    newOptions(db.Sqlx),
		Posts:     	newPosts(db.Sqlx),
		Roles:      newRoles(db.Sqlx),
		Session: 	newSession(db.Sqlx),
		Site:     	newSite(db.Sqlx),
		User:       newUser(db.Sqlx),
	}
}


