package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// Parameterize defines the function for getting http params
type Parameterize interface {
	Get() Params
}

// Params represents the http params for interacting with the DB
type Params struct {
	gin            *gin.Context
	Page           int
	Limit          int
	OrderBy        string
	OrderDirection string
	Filters        map[string][]Filter
}

// Filter represents the searching fields for searching through records.
type Filter struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// PaginationAllLimit defines how many items will be returned if
// the limit is set to list all
const (
	PaginationAllLimit = 99999999999999
	PaginationDefault  = 15
)

// NewParams - create a new parameter type
func NewParams(g *gin.Context) *Params {
	return &Params{
		gin: g,
	}
}

// Get query Parameters
func (p *Params) Get() Params {

	fmt.Println(p.gin.Params)

	// Get page and set default
	var page int
	pageStr := p.gin.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page == 0 {
		page = 1
	}

	// Get limit & calculate if list all
	var limit int
	limitStr := p.gin.Query("limit")
	if limitStr == "all" {
		limit = PaginationAllLimit
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			limit = PaginationAllLimit
		}
		if limit == 0 || limitStr == "" {
			limit = PaginationDefault
		}
	}

	// Get order and set defaults
	order := p.gin.Query("order")
	if order == "" {
		order = "id,asc"
	}

	// Get order and set defaults
	orderArr := strings.Split(order, ",")
	var orderParams [3]string
	if len(orderArr) == 1 {
		orderParams[0] = "id"
		orderParams[1] = "ASC"
	} else {
		orderParams[0] = orderArr[0]
		orderParams[1] = orderArr[1]
	}

	// Get the filters
	filtersParam := p.gin.Query("filter")
	var filters map[string][]Filter
	if filtersParam != "" {
		err := json.Unmarshal([]byte(filtersParam), &filters)
		if err != nil {
			filters = nil
		}
	}

	return Params{
		Page:           page,
		Limit:          limit,
		OrderBy:        orderParams[0],
		OrderDirection: orderParams[1],
		Filters:        filters,
	}
}
