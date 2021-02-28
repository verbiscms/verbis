// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"regexp"
	"strings"
)

type StoreConfig struct {
	DB      *sqlx.DB
	Config  *domain.ThemeConfig
	Paths   paths.Paths
	Options OptionsRepository
}

// Store defines all of the repositories used to interact with the database
type Store struct {
	Auth       AuthRepository
	Categories CategoryRepository
	Fields     FieldsRepository
	Forms      FormRepository
	Media      MediaRepository
	Options    OptionsRepository
	Posts      PostsRepository
	Redirects  RedirectRepository
	Roles      RoleRepository
	Site       SiteRepository
	User       UserRepository
	Config     *domain.ThemeConfig
}

// Create a new database instance, connect to database.
func New(cfg *StoreConfig) *Store {

	cfg.Options = newOptions(cfg)

	return &Store{
		Auth:       newAuth(cfg),
		Categories: newCategories(cfg),
		Forms:      newForms(cfg),
		Fields:     newFields(cfg),
		Media:      newMedia(cfg),
		Options:    newOptions(cfg),
		Posts:      newPosts(cfg),
		Redirects:  newRedirects(cfg),
		Roles:      newRoles(cfg),
		Site:       newSite(cfg),
		User:       newUser(cfg),
	}
}

// filterRows takes in the filters from the params set in http.Params
// If there is no filters set, an empty string will be returned.
// Returns errors.INVALID if the operator or column name was not found.
func filterRows(db *sqlx.DB, filters map[string][]params.Filter, table string) (string, error) {
	const op = "Model.filterRows"

	q := ""
	operators := []string{"=", ">", ">=", "<", "<=", "<>", "LIKE", "IN", "NOT LIKE", "like", "in", "not like"}

	if len(filters) != 0 {

		counter := 0
		for column, v := range filters {

			// Strip tags
			column = stripAlphaNum(strings.ToLower(column))

			// Check if the column exists before continuing
			var exists bool
			err := db.QueryRow("SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ? AND COLUMN_NAME = ?", table, column).Scan(&exists)
			if !exists || err != nil {
				return "", &errors.Error{
					Code:      errors.INVALID,
					Message:   fmt.Sprintf("The %s search query does not exist", column),
					Operation: op,
					Err:       fmt.Errorf("the %s search query does not exists when searching for %s", column, table)}
			}

			var fTable string
			if table != "" {
				fTable = table + "."
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
						Err:       fmt.Errorf("the %s operator does not exists when searching for the %s", operator, fTable)}
				}

				if counter > 0 {
					q += fmt.Sprintf(" OR ")
				} else if index > 0 {
					q += fmt.Sprintf(" AND ")
				} else {
					q += fmt.Sprintf(" WHERE ")
				}

				q += fmt.Sprintf("(%s%s %s '%s')", fTable, column, operator, value)
			}

			counter++
		}
	}

	return q, nil
}

// stripAlphaNum - Strip characters and return alpha numeric string for
// database processing.
func stripAlphaNum(text string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9 =<>%.@/!+_']+")
	return reg.ReplaceAllString(text, "")
}
