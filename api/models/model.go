package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/database"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"regexp"
)

// Store defines all of the repositories used to interact with the database
type Store struct {
	Auth       AuthRepository
	Categories CategoryRepository
	Fields     FieldsRepository
	Forms      FormRepository
	Media      MediaRepository
	Options    OptionsRepository
	Posts      PostsRepository
	Roles      RoleRepository
	Site       SiteRepository
	User       UserRepository
	Config     config.Configuration
}

// Create a new database instance, connect to database.
func New(db *database.MySql, config config.Configuration) *Store {
	return &Store{
		Auth:       newAuth(db.Sqlx, config),
		Categories: newCategories(db.Sqlx, config),
		Forms:      newForms(db.Sqlx, config),
		Fields:     newFields(db.Sqlx, config),
		Media:      newMedia(db.Sqlx, config),
		Options:    newOptions(db.Sqlx, config),
		Posts:      newPosts(db.Sqlx, config),
		Roles:      newRoles(db.Sqlx, config),
		Site:       newSite(db.Sqlx, config),
		User:       newUser(db.Sqlx, config),
		Config:     config,
	}
}

// filterRows takes in the filters from the params set in http.Params
// If there is no filters set, an empty string will be returned.
// Returns errors.INVALID if the operator or column name was not found.
func filterRows(db *sqlx.DB, filters map[string][]http.Filter, table string) (string, error) {
	const op = "Model.filterRows"

	q := ""
	operators := []string{"=", ">", ">=", "<", "<=", "<>", "LIKE", "IN", "NOT LIKE", "like", "in", "not like"}

	if len(filters) != 0 {
		counter := 0
		for column, v := range filters {

			// Strip tags
			column = stripAlphaNum(column)

			// Check if the column exists before continuing
			var exists bool
			_ = db.QueryRow(fmt.Sprintf("SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '%s' AND COLUMN_NAME = '%s'", table, column)).Scan(&exists)
			if !exists {
				return "", &errors.Error{
					Code:      errors.INVALID,
					Message:   fmt.Sprintf("The %s search query does not exist", column),
					Operation: op,
					Err:       fmt.Errorf("the %s search query does not exists when searching for %s", column, table)}
			}

			if table != "" {
				table = table + "."
			}

			for index, filter := range v {

				// Strip tags
				operator := stripAlphaNum(filter.Operator)
				value := stripAlphaNum(filter.Value)

				// Account for like or not like values
				if operator == "like" || operator == "LIKE" || operator == "not like" || operator == "NOT LIKE" {
					value = "%" + value + "%"
				}

				// Check if the operator exists before continuing
				if opExists := helpers.StringInSlice(operator, operators); !opExists {
					return "", &errors.Error{
						Code:      errors.INVALID,
						Message:   fmt.Sprintf("The %s operator does not exist", operator),
						Operation: op,
						Err:       fmt.Errorf("the %s operator does not exists when searching for the %s", operator, table)}
				}

				if counter > 0 {
					q += fmt.Sprintf(" OR ")
				} else if index > 0 {
					q += fmt.Sprintf(" AND ")
				} else {
					q += fmt.Sprintf(" WHERE ")
				}

				q += fmt.Sprintf("(%s%s %s '%s')", table, column, operator, value)
			}

			counter++
		}
	}

	return q, nil
}

// stripAlphaNum - Strip characters and return alpha numeric string for
// database processing.
func stripAlphaNum(text string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9 =<>%.@/!+']+")
	return reg.ReplaceAllString(text, "")
}
