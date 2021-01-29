package params

import (
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/spf13/cast"
)

// Query represents the map of arguments passed to
// list functions in templates.
type Query map[string]interface{}

var (
	// Defaults represents the default params if
	// none were passed for templates.
	Defaults = http.Defaults{
		Page:           1,
		Limit:          15,
		OrderBy:        "created_at",
		OrderDirection: "desc",
	}
)

// Get
//
// Returns parameters for the store to used for obtaining
// multiple entities. If the orderBy or orderDirection
// arguments are not passed, defaults will be used.
func (q Query) Get(orderBy string, orderDirection string) http.Params {
	def := Defaults
	if orderBy != "" {
		def.OrderBy = orderBy
	}
	if orderDirection != "" {
		def.OrderDirection = orderDirection
	}
	params := http.NewParams(q, def)
	return params.Get()
}

// Param
//
// Is an implementation of a stringer to return
// parameters from the Query map.
func (q Query) Param(param string) string {
	val, ok := q[param]
	if !ok {
		return ""
	}
	s, err := cast.ToStringE(val)
	if err != nil {
		return ""
	}
	return s
}

// Default
//
// Sets or gets default parameters for the Query.
// If the parameter is not found, it will
// return the default string passed.
func (q Query) Default(param string, def string) interface{} {
	val, ok := q[param]
	if !ok {
		return def
	}
	return val
}
