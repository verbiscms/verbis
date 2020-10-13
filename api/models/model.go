package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Store defines all of the repositories used to interact with the database
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

// filterRows takes in the filters from the params set in http.Params
// If there is no filters set, an empty string will be returned.
func filterRows(db *sqlx.DB, filters map[string][]http.Filter, table string) string {
	q := ""

	if table != "" {
		table = table + "."
	}
	var exists bool
	err := db.QueryRow(fmt.Sprintf("SELECT count(*) AS [Column Exists] FROM SYSOBJECTS INNER JOIN SYSCOLUMNS ON SYSOBJECTS.ID = SYSCOLUMNS.ID WHERE SYSOBJECTS.NAME = '%s' AND SYSCOLUMNS.NAME = '%s'", table, environment.GetDatabaseName())).Scan(&exists)

	fmt.Println(err)
	fmt.Println(exists)


	if len(filters) != 0 {
		counter := 0
		for column, v := range filters {
			for index, filter := range v {
				if counter > 0 {
					q += fmt.Sprintf(" OR ")
				} else if index > 0 {
					q += fmt.Sprintf(" AND ")
				} else {
					q += fmt.Sprintf(" WHERE ")
				}
				q += fmt.Sprintf("(%s%s %s '%s')", table, column, filter.Operator, filter.Value)
			}
			counter++
		}
	}

	return q
}