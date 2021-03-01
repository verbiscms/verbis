package builder

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Sqlbuilder struct {
	string         string
	selectstmt     string
	wherestmt      string
	whereinstmt    string
	fromstmt       string
	deletefromstmt string
	leftjoinstmt   string
	limitstmt      string
	offsetstmt     string
	orderbystmt    string
	args           []string
	skip           []string
	Dialect        string //Can be postgres, mysql or (more to come)
}

func Builder() *Sqlbuilder {
	return &Sqlbuilder{
		Dialect: "whatever",
	}
}

func (s *Sqlbuilder) From(fromstmt string) *Sqlbuilder {
	s.fromstmt = s.formatSchema(fromstmt)

	return s
}

func (s *Sqlbuilder) DeleteFrom(fromstmt string) *Sqlbuilder {
	s.deletefromstmt = s.formatSchema(fromstmt)

	return s
}

func (s *Sqlbuilder) SelectRaw(selectstmt string) *Sqlbuilder {
	re := regexp.MustCompile(`\r?\n`)
	selectstmt = re.ReplaceAllString(selectstmt, " ")

	s.selectstmt += selectstmt + `, `
	return s
}

func (s *Sqlbuilder) Select(selectstmt ...string) *Sqlbuilder {

	for _, ss := range selectstmt {
		s.selectstmt += s.formatSchema(ss) + `, `
	}

	return s
}

func (s *Sqlbuilder) Where(table string, operator string, input interface{}) *Sqlbuilder {

	value := fmt.Sprintf("%v", input)
	operator = strings.ToUpper(operator)
	value = strings.TrimSuffix(value, `'`)
	value = strings.TrimSuffix(value, `"`)
	value = strings.TrimSuffix(value, "`")
	value = strings.TrimPrefix(value, `'`)
	value = strings.TrimPrefix(value, `"`)
	value = strings.TrimPrefix(value, "`")

	switch operator {
	case `BETWEEN`:
		re := regexp.MustCompile("and|AND|And")
		vp := re.Split(value, -1)

		value = ``

		for _, v := range vp {
			value += sanitiseString(`'`+strings.TrimSpace(v)+`'`) + ` AND `
		}

		value = strings.TrimSuffix(value, ` AND `)
	default:
		value = sanitiseString(`'` + value + `'`)
	}

	s.wherestmt += s.formatSchema(table) + " " + operator + " " + value + ` AND `

	return s
}

func (s *Sqlbuilder) WhereRaw(wherestmt string) *Sqlbuilder {
	s.wherestmt += wherestmt + ` AND `
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

	s.leftjoinstmt += `LEFT JOIN ` + table + ` AS ` + as + ` ON ` + on + ` `
	return s
}

func (s *Sqlbuilder) LeftJoinExtended(table string, as string, on string, additionalQuery string) *Sqlbuilder {

	table = s.formatSchema(table)
	on = s.formatJoinOn(on)

	s.leftjoinstmt += `LEFT JOIN ` + table + ` AS "` + as + `" ON ` + on + ` ` + additionalQuery + ` `
	return s
}

func (s *Sqlbuilder) Limit(limit int) *Sqlbuilder {
	s.limitstmt = `LIMIT ` + strconv.Itoa(limit) + ` `

	return s
}

func (s *Sqlbuilder) Offset(offset int) *Sqlbuilder {
	s.offsetstmt = `OFFSET ` + strconv.Itoa(offset) + ` `

	return s
}

func (s *Sqlbuilder) OrderBy(column string, diretion string) *Sqlbuilder {

	s.orderbystmt = `ORDER BY "` + column + `" ` + diretion

	return s
}

func (s *Sqlbuilder) Reset() *Sqlbuilder {
	s.string = ``
	s.selectstmt = ``
	s.orderbystmt = ``
	s.whereinstmt = ``
	s.limitstmt = ``
	s.fromstmt = ``
	s.leftjoinstmt = ``
	s.wherestmt = ``
	s.offsetstmt = ``

	return s
}

func (s *Sqlbuilder) Count() string {
	sqlquery := s.Build()

	countQuery := `SELECT COUNT(*) AS rowcount FROM (` + sqlquery + `) AS rowdata`

	return countQuery
}

func (s *Sqlbuilder) Build() string {

	//build selects
	if s.deletefromstmt == `` {
		if s.selectstmt == `` {
			s.string = `SELECT * `
		} else {
			s.string = `SELECT ` + strings.TrimSuffix(s.selectstmt, `, `) + ` `
		}
	}

	//build from
	if s.fromstmt == `` {
		if s.deletefromstmt != `` {
			s.string += `DELETE FROM ` + strings.TrimSuffix(s.deletefromstmt, `.`) + ` `
		} else {
			return ``
		}
	} else {
		s.string += `FROM ` + strings.TrimSuffix(s.fromstmt, `.`) + ` `
	}

	//left joins
	s.string += s.leftjoinstmt + ` `

	//where
	if s.wherestmt != `` {
		s.string += `WHERE ` + strings.TrimSuffix(s.wherestmt, ` AND `) + ` `
	}

	//orderby
	if s.orderbystmt != `` {
		s.string += s.orderbystmt + ` `
	}

	//limit and offset
	s.string += s.limitstmt
	s.string += s.offsetstmt

	space := regexp.MustCompile(`\s+`)
	s.string = space.ReplaceAllString(s.string, " ")

	returnString := s.string

	return returnString
}

func (s *Sqlbuilder) Skip(skip []string) *Sqlbuilder {
	s.skip = skip
	return s
}

func (s *Sqlbuilder) Args(args []string) *Sqlbuilder {
	s.args = args
	return s
}

func (s *Sqlbuilder) BuildInsert(table string, data interface{}, additionalQuery ...string) string {
	dbCols, dbVals, err := mapStruct(data, s.args, s.skip, false)
	if err != nil {
		return ""
	}

	additional := ""
	if len(additionalQuery) == 1 {
		additional = additionalQuery[0]
	}

	sql := "INSERT INTO " + s.formatSchema(table) + " (" + strings.Join(dbCols, ", ") + ") VALUES (" + strings.Join(dbVals, ", ") + ") " + additional

	return sql
}

func (s *Sqlbuilder) BuildUpdate(table string, data interface{}, additionalQuery ...string) string {

	dbCols, dbVals, err := mapStruct(data, s.args, s.skip, true)
	if err != nil {
		return ""
	}

	setString := ""
	sql := ""

	additional := ""
	if len(additionalQuery) == 1 {
		additional = additionalQuery[0]
	}

	for i, col := range dbCols {
		setString += col + ` = ` + dbVals[i] + `, `
	}
	setString = strings.TrimSuffix(setString, `, `) + ` `

	if setString != "" {
		sql = "UPDATE " + s.formatSchema(table) + ` SET ` + setString

		if s.wherestmt != `` {
			sql += `WHERE ` + strings.TrimSuffix(s.wherestmt, ` AND `) + ` `
		}

		sql += additional

		return sql
	}

	return sql
}

/**
Helpers
*/

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

func (s *Sqlbuilder) formatJoinOn(joinstmt string) string {
	joinParts := strings.Split(joinstmt, "=")
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

func mapStruct(data interface{}, args []string, skip []string, update bool) (dbCols []string, dbVals []string, error error) {
	fields := reflect.TypeOf(data)
	values := reflect.ValueOf(data)

	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)

		// TODO this needs to be dynamic in options
		//val, exists := field.Tag.Lookup("sqlb")
		val, exists := field.Tag.Lookup("db")

		if stringInSlice(skip, val) {
			continue
		}

		if exists {
			dbCols = append(dbCols, "\""+val+"\"")
		} else {
			dbCols = append(dbCols, "\""+toSnakeCase(field.Name)+"\"")
		}

		v := compareFields(value)

		if val == "created_at" {
			v = "NOW()"
		}

		if val == "updated_at" {
			v = "NOW()"
		}

		if update && val == "updated_at" {
			continue
		}

		if stringInSlice(args, val) {
			v = "?"
		}

		dbVals = append(dbVals, v)
	}

	return dbCols, dbVals, nil
}

func compareFields(value reflect.Value) string {
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
			v = fmt.Sprintf("NULL")
		} else {
			v = compareFields(reflect.ValueOf(value.Elem().Interface()))
		}
	case reflect.Bool:
		if value.Bool() {
			v = "TRUE"
		} else {
			v = "FALSE"
		}
		//default:
		//	return dbCols, dbVals, errors.New("type: " + value.Kind().String() + " unsupported")
	}

	return v
}

func stringInSlice(slice []string, needle string) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
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
