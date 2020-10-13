package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Params struct {
	Page 			int
	Limit 			int
	OrderBy 		string
	OrderDirection 	string
	Filters 		map[string][]Filter
}

type Filter struct {
	Operator string
	Value    interface{}  `json:"value"`
}

// PaginationAllLimit defines how many items will be returned if
// the limit is set to list all
const (
	PaginationAllLimit = 999999
)

// Get query Parameters
func GetParams(g *gin.Context) Params {

	// Get page and set default
	var page int
	pageStr := g.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page == 0 {
		page = 1
	}

	// Get limit & calculate if list all
	var limit int
	limitStr := g.Query("limit")
	if limitStr == "all" {
		limit = PaginationAllLimit
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			limit = PaginationAllLimit
		}
		if limit == 0 {
			limit = 15
		}
	}

	// Get order and set defaults
	order := g.Query("order")
	if order == "" {
		order = "id,asc"
	}

	// Get order and set defaults
	orderArr := strings.Split(order, ",")
	var orderParams [3]string
	if len(orderArr) == 1 {
		orderParams[0] = orderArr[0]
		orderParams[1] = "ASC"
	} else {
		orderParams[0] = orderArr[0]
		orderParams[1] = orderArr[1]
	}

	// Get the filters
	filtersParam := g.Query("filter")
	var filters map[string][]Filter
	if filtersParam != "" {
		err := json.Unmarshal([]byte(filtersParam), &filters)
		if err != nil {
			filters = nil
		}
	}

	return Params{
		Page:  page,
		Limit: limit,
		OrderBy: orderArr[0],
		OrderDirection: orderArr[1],
		Filters: filters,
	}
}


