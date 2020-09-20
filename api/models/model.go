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
	Subscriber 		SubscriberRepository
	Site      		SiteRepository
	User       		UserRepository
}

// Create a new database instance, connect to database.
func New(db *database.DB) (*Store, error) {

	auth := newAuth(db.Sqlx)
	categories := newCategories(db.Sqlx)
	media := newMedia(db.Sqlx)
	options := newOptions(db.Sqlx)
	roles := newRoles(db.Sqlx)
	session := newSession(db.Sqlx)
	seoMeta := newSeoMeta(db.Sqlx)
	subscriber := newSubscriber(db.Sqlx)
	user := newUser(db.Sqlx)
	site := newSite(db.Sqlx, options)
	fields := newFields(db.Sqlx, options)
	posts := newPosts(db.Sqlx, seoMeta, user, categories)

	s := &Store{
		Auth:       auth,
		Categories: categories,
		Fields:     fields,
		Media: 		media,
		Options:    options,
		Posts:     	posts,
		Roles:      roles,
		Session: 	session,
		Subscriber: subscriber,
		Site:     	site,
		User:       user,
	}

	return s, nil
}


