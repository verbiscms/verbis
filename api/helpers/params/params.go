// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params

import (
	"encoding/json"
	"strconv"
)

const (
	// DefaultPage defines the page number if none is set.
	DefaultPage = 1
	// DefaultLimit defines how many items will be returned if
	// the limit is set to list all.
	DefaultLimit = 15
	// DefaultOrderBy defines the default order by if an error
	// occurred.
	DefaultOrderBy = "id"
	// DefaultOrderDirection defines the default order direction
	// if an error occurred.
	DefaultOrderDirection = "DESC"
)

// Params represents the http params for interacting with the DB
type Params struct {
	Page           int      `json:"page"`
	Limit          int      `json:"limit"`
	LimitAll       bool     `json:"all"`
	OrderBy        string   `json:"order_by"`
	OrderDirection string   `json:"order_direction"`
	Filters        Filters  `json:"-"`
	defaults       Defaults `json:"-"`
	Stringer       `json:"-"`
}

// Stringer defines the method for obtaining parameters.
type Stringer interface {
	Param(string) string
}

// Filters represents the map and slice of filters
type Filters map[string][]Filter

// Filter represents the searching fields for searching through records.
type Filter struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// Defaults represents the default configuration for obtaining params.
type Defaults struct {
	Page           int         `json:"page"`
	Limit          interface{} `json:"limit"`
	OrderBy        string      `json:"order_by"`
	OrderDirection string      `json:"order_direction"`
}

// NewParams - create a new parameter type
func New(str Stringer, def Defaults) *Params {
	p := &Params{
		Stringer: str,
		defaults: def,
	}
	return p
}

// Get
//
// Get query Parameters for http API routes and
// query loops in templates.
func (p *Params) Get() Params {
	limit, limitAll := p.limit()
	order := p.order()
	return Params{
		Page:           p.page(),
		Limit:          limit,
		LimitAll:       limitAll,
		OrderBy:        order[0],
		OrderDirection: order[1],
		Filters:        p.filter(),
	}
}

// page
//
// Obtain the page parameter and set a default if there
// was an error converting the page or the page number
// is set to 0.
func (p *Params) page() int {
	var page int
	pageStr := p.Param("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = p.defaults.Page
	}

	if page <= 0 {
		page = 1
	}

	return page
}

// limit
//
// Obtain the limit parameter and set a default if there
// was an error converting the limit or the page number
// is set to 0. Returns true if limit is set to "all"
func (p *Params) limit() (int, bool) {
	limitStr := p.Param("limit")
	if limitStr == "all" {
		return 0, true
	}

	limit, err := strconv.Atoi(limitStr)
	defLimit, ok := p.defaults.Limit.(int)
	if !ok || err != nil || defLimit == 0 {
		return DefaultLimit, false
	}

	if limit == 0 || limitStr == "" || defLimit == 0 {
		return defLimit, false
	}

	return limit, false
}

// order
//
// Obtain the order array (order by and order direction)
// set defaults if there is none set.
func (p *Params) order() []string {
	order := []string{p.defaults.OrderBy, p.defaults.OrderDirection}

	orderBy := p.Param("order_by")
	if orderBy != "" {
		order[0] = orderBy
	}

	orderDirection := p.Param("order_direction")
	if orderDirection != "" {
		order[1] = orderDirection
	}

	if order[0] == "" {
		order[0] = DefaultOrderBy
	}

	if order[1] == "" {
		order[1] = DefaultOrderDirection
	}

	return order
}

// filter
//
// Obtain the map of filters by unmarshalling into a
// Filter, if an error occurred, filters will be
// set to nil.
func (p *Params) filter() map[string][]Filter {
	filtersParam := p.Param("filter")

	var filters map[string][]Filter
	if filtersParam != "" {
		err := json.Unmarshal([]byte(filtersParam), &filters)
		if err != nil {
			filters = nil
		}
	}

	return filters
}
