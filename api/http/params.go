package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	// PaginationAllLimit defines how many items will be returned if
	// the limit is set to list all
	PaginationDefault      = 15
	// PaginationDefaultOrder defines the default order by for the
	// API
	PaginationDefaultOrder = "id,DESC"
)


// Parameterize defines the function for getting http params
type Parameterize interface {
	Get() Params
}

// Params represents the http params for interacting with the DB
type Params struct {
	//gin            *gin.Context
	Page           int                 `json:"page"`
	Limit          int                 `json:"limit"`
	LimitAll       bool                `json:"all"`
	OrderBy        string              `json:"order_by"`
	OrderDirection string              `json:"order_direction"`
	Filters        map[string][]Filter `json:"-"`
	defaults Defaults `json:"-"`
	Stringer  `json:"-"`
}

type Stringer interface {
	Param(string) string
}

// Filter represents the searching fields for searching through records.
type Filter struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Defaults struct {
	Page int
	Limit interface{}
	OrderBy string
	OrderDirection string
}


// NewParams - create a new parameter type
func NewParams(str Stringer, def Defaults) *Params {
	p := &Params{
		Stringer: str,
		defaults: def,
	}
	p.validateDefaults()
	return p
}

type apiParams struct {
	gin *gin.Context
}

func (a *apiParams) Param(q string) string {
	return a.gin.Query(q)
}

func ApiParams(g *gin.Context, def Defaults) *Params {
	p := &Params{
		Stringer: &apiParams{gin: g},
		defaults: def,
	}
	p.validateDefaults()
	return p
}


func (p *Params) validateDefaults() {
	if p.defaults.OrderBy == "" {
		panic("No default order by set")
	}
	if p.defaults.OrderDirection == "" {
		panic("No default order direction set")
	}
	if p.defaults.Limit == nil {
		panic("No default limit set")
	}
}

// Get query Parameters for http API routes
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

// Get page and set default
func (p *Params) page() int {
	var page int
	pageStr := p.Param("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page == 0 {
		page = p.defaults.Page
	}

	return page
}

// Get limit & calculate if list all
func (p *Params) limit() (int, bool) {
	limitStr := p.Param("limit")
	if limitStr == "all" {
		return 0, true
	}

	limit, err := strconv.Atoi(limitStr)
	defLimit, ok := p.defaults.Limit.(int)
	if !ok {
		return PaginationDefault, false
	}

	if err != nil || limit == 0 || limitStr == "" {
		return defLimit, false
	}

	return limit, false
}

// Get order and set defaults
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

	return order
}

// Get the filters
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