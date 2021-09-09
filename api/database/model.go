// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package database

import (
	"fmt"
	"github.com/verbiscms/verbis/api/common/params"
	verbisstrings "github.com/verbiscms/verbis/api/common/strings"
	"github.com/verbiscms/verbis/api/database/builder"
	"github.com/verbiscms/verbis/api/errors"
	"strings"
)

var (
	ErrQueryMessage = "Error executing sql query"
)

// FilterRows takes in the filters from the params set in http.Params
// If there is no filters set, an empty string will be returned.
// Returns errors.INVALID if the operator or column name was not found.
func FilterRows(driver Driver, query *builder.Sqlbuilder, filters map[string][]params.Filter, table string) error {
	const op = "Model.filterRows"

	operators := []string{"=", ">", ">=", "<", "<=", "<>", "LIKE", "IN", "NOT LIKE", "like", "in", "not like"}

	if len(filters) != 0 {
		counter := 0
		for column, v := range filters {
			// Strip tags
			column = mysqlRealEscapeString(strings.ToLower(column))

			// Check if the column exists before continuing
			var exists bool
			err := driver.DB().QueryRow("SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ? AND COLUMN_NAME = ?", table, column).Scan(&exists)
			if !exists || err != nil {
				return &errors.Error{
					Code:      errors.INVALID,
					Message:   fmt.Sprintf("The %s search query does not exist", column),
					Operation: op,
					Err:       fmt.Errorf("the %s search query does not exists when searching for %s", column, table)}
			}

			var fTable string
			if table != "" {
				fTable = table + "."
			}

			for _, filter := range v {
				// Strip tags
				operator := mysqlRealEscapeString(filter.Operator)
				value := mysqlRealEscapeString(filter.Value)

				// Account for like or not like values
				if operator == "like" || operator == "LIKE" || operator == "not like" || operator == "NOT LIKE" {
					value = "%" + value + "%"
				}

				// Check if the operator exists before continuing
				if opExists := verbisstrings.InSlice(operator, operators); !opExists {
					return &errors.Error{
						Code:      errors.INVALID,
						Message:   fmt.Sprintf("The %s operator does not exist", operator),
						Operation: op,
						Err:       fmt.Errorf("the %s operator does not exists when searching for the %s", operator, fTable)}
				}

				query.WhereRaw(fmt.Sprintf("(%s%s %s '%s')", fTable, column, operator, value))
			}
			counter++
		}
	}

	return nil
}

func mysqlRealEscapeString(sql string) string {
	dest := make([]byte, 0, 2*len(sql)) //nolint
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
		case '\n': /* Must be escaped for logs */
			escape = 'n'
		case '\r':
			escape = 'r'
		case '\\':
			escape = '\\'
		case '\'':
			escape = '\''
		case '"': /* Better safe than sorry */
			escape = '"'
		case '\032': //十进制26,八进制32,十六进制1a, /* This gives problems on Win32 */
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}

//// stripAlphaNum strips characters and returns an
//// alpha numeric string for database processing.
//func stripAlphaNum(text string) string {
//	reg := regexp.MustCompile("[^a-zA-Z0-9 =<>%.@/!+_']+")
//	return reg.ReplaceAllString(text, "")
//}
