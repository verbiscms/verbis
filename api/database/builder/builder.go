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

type Sqlbuilder struct {
	string         string
	selectStmt     string
	whereStmt      string
	whereinStmt    string
	fromStmt       string
	deletefromStmt string
	leftjoinStmt   string
	limitStmt      string
	offsetStmt     string
	orderbyStmt    string
	// Insert table
	insertStmt string
	// Update table
	updateStmt string
	// Slice of columns, key value for updates & inserts.
	columns [][2]string
	Dialect string //Can be postgres, mysql or (more to come)
}

func New(dialect string) *Sqlbuilder {
	return &Sqlbuilder{
		Dialect: dialect,
	}
}

func (s *Sqlbuilder) From(fromStmt string) *Sqlbuilder {
	s.fromStmt = s.formatSchema(fromStmt)

	return s
}

func (s *Sqlbuilder) DeleteFrom(fromStmt string) *Sqlbuilder {
	s.deletefromStmt = s.formatSchema(fromStmt)

	return s
}

func (s *Sqlbuilder) SelectRaw(selectStmt string) *Sqlbuilder {
	re := regexp.MustCompile(`\r?\n`)
	selectStmt = re.ReplaceAllString(selectStmt, " ")

	s.selectStmt += selectStmt + `, `
	return s
}

func (s *Sqlbuilder) Select(selectStmt ...string) *Sqlbuilder {

	for _, ss := range selectStmt {
		s.selectStmt += s.formatSchema(ss) + `, `
	}

	return s
}

func (s *Sqlbuilder) Where(table string, operator string, value interface{}) *Sqlbuilder {

	val, ok := printInterface(reflect.ValueOf(value))
	if !ok {
		// TODO: handle
		return s
	}

	operator = strings.ToUpper(operator)
	val = strings.TrimSuffix(val, `'`)
	val = strings.TrimSuffix(val, `"`)
	val = strings.TrimSuffix(val, "`")
	val = strings.TrimPrefix(val, `'`)
	val = strings.TrimPrefix(val, `"`)
	val = strings.TrimPrefix(val, "`")

	switch operator {
	case `BETWEEN`:
		re := regexp.MustCompile("and|AND|And")
		vp := re.Split(val, -1)

		val = ``

		for _, v := range vp {
			val += sanitiseString(`'`+strings.TrimSpace(v)+`'`) + ` AND `
		}

		val = strings.TrimSuffix(val, ` AND `)
	default:
		val = sanitiseString(`'` + val + `'`)
	}

	s.whereStmt += s.formatSchema(table) + " " + operator + " " + val + ` AND `

	return s
}

func (s *Sqlbuilder) WhereRaw(whereStmt string) *Sqlbuilder {
	s.whereStmt += whereStmt + ` AND `
	return s
}

//Accepts Slice of INT, FLOAT32, STRING, or a simple comma separated STRING
func (s *Sqlbuilder) WhereIn(column string, params interface{}) *Sqlbuilder {

	output := ""

	switch foo := params.(type) {
	case []int, []float32:
		output += "(" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(foo)), ", "), "[]") + ")"
		break
	case []string:
		output += "("
		for _, v := range foo {
			output += "'" + sanitiseString(v) + "', "
		}
		output = strings.TrimSuffix(output, ", ")
		output += ")"
		break
	case string:
		output = "(" + sanitiseString(foo) + ")"
		break
	default:
		output = ""
	}

	if output != "" {
		s.WhereRaw(column + ` IN ` + output)
	}

	return s
}

func (s *Sqlbuilder) WhereStringMatchAny(column string, params []string) *Sqlbuilder {

	output := ""

	output += "(array["
	for _, v := range params {
		output += "'%" + sanitiseString(strings.TrimSpace(v)) + "%', "
	}
	output = strings.TrimSuffix(output, ", ")
	output += "])"

	if output != "" {
		s.WhereRaw(column + ` ILIKE ANY ` + output)
	}

	return s
}

func (s *Sqlbuilder) WhereStringMatchAll(column string, params []string) *Sqlbuilder {

	output := ""

	output += "'"
	for _, v := range params {
		output += "%" + sanitiseString(strings.TrimSpace(v)) + "% "
	}
	output = strings.TrimSuffix(output, " ")
	output += "'"

	if output != "" {
		s.WhereRaw(column + ` ILIKE ` + output)
	}

	return s
}

func (s *Sqlbuilder) LeftJoin(table string, as string, on string) *Sqlbuilder {

	table = s.formatSchema(table)
	on = s.formatJoinOn(on)
	as = s.formatSchema(as)

	s.leftjoinStmt += `LEFT JOIN ` + table + ` AS ` + as + ` ON ` + on + ` `
	return s
}

func (s *Sqlbuilder) LeftJoinExtended(table string, as string, on string, additionalQuery ...string) *Sqlbuilder {

	table = s.formatSchema(table)
	on = s.formatJoinOn(on)

	var q string
	if len(additionalQuery) == 1 {
		q = additionalQuery[0]
	}

	s.leftjoinStmt += `LEFT JOIN ` + table + ` AS "` + as + `" ON ` + on + ` ` + q + ` `
	return s
}

func (s *Sqlbuilder) Limit(limit int) *Sqlbuilder {
	s.limitStmt = `LIMIT ` + strconv.Itoa(limit) + ` `

	return s
}

func (s *Sqlbuilder) Offset(offset int) *Sqlbuilder {
	s.offsetStmt = `OFFSET ` + strconv.Itoa(offset) + ` `

	return s
}

func (s *Sqlbuilder) OrderBy(column string, diretion string) *Sqlbuilder {

	s.orderbyStmt = `ORDER BY "` + column + `" ` + diretion

	return s
}

func (s *Sqlbuilder) Reset() *Sqlbuilder {
	s.string = ``
	s.selectStmt = ``
	s.orderbyStmt = ``
	s.whereinStmt = ``
	s.limitStmt = ``
	s.fromStmt = ``
	s.leftjoinStmt = ``
	s.whereStmt = ``
	s.offsetStmt = ``

	return s
}

func (s *Sqlbuilder) Count() string {
	sqlquery := s.Build()

	countQuery := `SELECT COUNT(*) AS rowcount FROM (` + sqlquery + `) AS rowdata`

	return countQuery
}

func (s *Sqlbuilder) Exists() string {
	sqlquery := s.Build()

	q := strings.TrimSuffix(sqlquery, " ")
	existsQuery := `SELECT EXISTS (` + q + `)`

	return existsQuery
}

func (s *Sqlbuilder) Build() string {

	// insert
	if s.insertStmt != `` {
		return s.buildInsert()
	}

	// update
	if s.updateStmt != `` {
		return s.buildUpdate()
	}

	//build selects
	if s.deletefromStmt == `` {
		if s.selectStmt == `` {
			s.string = `SELECT * `
		} else {
			s.string = `SELECT ` + strings.TrimSuffix(s.selectStmt, `, `) + ` `
		}
	}

	//build from
	if s.fromStmt == `` {
		if s.deletefromStmt != `` {
			s.string += `DELETE FROM ` + strings.TrimSuffix(s.deletefromStmt, `.`) + ` `
		} else {
			return ``
		}
	} else {
		s.string += `FROM ` + strings.TrimSuffix(s.fromStmt, `.`) + ` `
	}

	//left joins
	s.string += s.leftjoinStmt + ` `

	//where
	if s.whereStmt != `` {
		s.string += `WHERE ` + strings.TrimSuffix(s.whereStmt, ` AND `) + ` `
	}

	//orderby
	if s.orderbyStmt != `` {
		s.string += s.orderbyStmt + ` `
	}

	//limit and offset
	s.string += s.limitStmt
	s.string += s.offsetStmt

	space := regexp.MustCompile(`\s+`)
	s.string = space.ReplaceAllString(s.string, " ")

	returnString := strings.TrimSuffix(s.string, " ")

	return returnString
}

func (s *Sqlbuilder) BuildInsert(table string, data interface{}, additionalQuery ...string) (string, error) {
	dbCols, dbVals, err := mapStruct(data)
	if err != nil {
		return "", err
	}

	var q string
	if len(additionalQuery) == 1 {
		q = additionalQuery[0]
	}

	sql := "INSERT INTO " + s.formatSchema(table) + " (" + strings.Join(dbCols, ", ") + ") VALUES (" + strings.Join(dbVals, ", ") + ") " + q

	return sql, nil
}

func (s *Sqlbuilder) BuildUpdate(table string, data interface{}, additionalQuery string) (string, error) {

	dbCols, dbVals, err := mapStruct(data)
	if err != nil {
		return "", err
	}

	setString := ""
	sql := ""

	for i, col := range dbCols {
		setString += col + ` = ` + dbVals[i] + `, `
	}
	setString = strings.TrimSuffix(setString, `, `) + ` `

	if setString != "" {
		sql = "UPDATE " + s.formatSchema(table) + ` SET ` + setString

		if s.whereStmt != `` {
			sql += `WHERE ` + strings.TrimSuffix(s.whereStmt, ` AND `) + ` `
		}

		sql += additionalQuery

		return sql, nil
	}

	return sql, errors.New("sql build failed")
}
