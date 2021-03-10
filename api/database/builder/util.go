// Copyright 2020 The GoQueryBuilder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builder

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/**
Helpers
*/

func (s *Sqlbuilder) getDelim() string {
	switch strings.ToLower(s.Dialect) {
	case "postgres":
		return `"`
	case "mysql":
		return "`"
	default:
		return `"`
	}
}

func (s *Sqlbuilder) formatSchema(schema string) string {
	schemaParts := strings.Split(schema, ".")
	finalSchemaStmt := ``

	var dialectFormat string

	switch strings.ToLower(s.Dialect) {
	case "postgres":
		dialectFormat = `"`
		break
	case "mysql":
		dialectFormat = "`"
		break
	default:
		dialectFormat = `"`
	}

	for _, v := range schemaParts {
		if v == `*` {
			finalSchemaStmt += `*`
		} else {
			part := strings.TrimSpace(v)
			if string(part[0]) == dialectFormat && string(part[len(part)-1]) == dialectFormat {
				finalSchemaStmt += part + `.`
			} else {
				finalSchemaStmt += dialectFormat + part + dialectFormat + `.`
			}

		}
	}

	return strings.TrimSuffix(finalSchemaStmt, `.`)
}

func (s *Sqlbuilder) formatJoinOn(joinStmt string) string {
	joinParts := strings.Split(joinStmt, "=")
	finalJoinStmt := ``

	for _, v := range joinParts {
		finalJoinStmt += s.formatSchema(v) + ` = `
	}

	return strings.TrimSuffix(finalJoinStmt, ` = `)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func mapStruct(data interface{}) (dbCols []string, dbVals []string, error error) {
	fields := reflect.TypeOf(data)
	values := reflect.ValueOf(data)

	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)

		val, exists := field.Tag.Lookup("sqlb")

		if exists {
			dbCols = append(dbCols, "\""+val+"\"")
		} else {
			dbCols = append(dbCols, "\""+toSnakeCase(field.Name)+"\"")
		}

		var v string

		v, ok := printInterface(value)
		if !ok {
			return dbCols, dbVals, errors.New("type: " + value.Kind().String() + " unsupported")
		}
		dbVals = append(dbVals, v)
	}

	return dbCols, dbVals, nil
}

func printInterface(value reflect.Value) (string, bool) {
	var v string

	switch value.Kind() {
	case reflect.String:
		v = "'" + sanitiseString(value.String()) + "'"
	case reflect.Int:
		v = strconv.FormatInt(value.Int(), 10)
	case reflect.Int8:
		v = strconv.FormatInt(value.Int(), 10)
	case reflect.Int32:
		v = strconv.FormatInt(value.Int(), 10)
	case reflect.Int64:
		v = strconv.FormatInt(value.Int(), 10)
	case reflect.Float64:
		v = fmt.Sprintf("%f", value.Float())
	case reflect.Float32:
		v = fmt.Sprintf("%f", value.Float())
	case reflect.Slice:
		v = fmt.Sprintf("%v", value)
	case reflect.Array:
		v = fmt.Sprintf("%v", value)
	case reflect.Ptr:
		if value.IsNil() {
			v = "NULL"
		} else {
			val, ok := printInterface(reflect.ValueOf(value.Elem().Interface()))
			if !ok {
				return "", false
			}
			return val, true
		}
	case reflect.Bool:
		if value.Bool() {
			v = "TRUE"
		} else {
			v = "FALSE"
		}
	default:
		return "", false
	}

	// Sanity checking
	if v == "'?'" {
		v = "?"
	}

	if v == "'NOW()'" {
		v = "NOW()"
	}

	return v, true
}

func sanitiseString(str string) string {

	if len(str) > 0 {
		rebuildSingles := false

		if string(str[0]) == "'" && string(str[len(str)-1]) == "'" {
			rebuildSingles = true
		}

		str = strings.TrimSuffix(strings.TrimPrefix(str, "'"), "'")
		str = strings.ReplaceAll(str, "'", "''")

		if rebuildSingles {
			str = "'" + str + "'"
		}
	}

	return str
}
