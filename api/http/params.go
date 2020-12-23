package http

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/errors"
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
	Page            int    `json:"page"`
	Limit           int    `json:"limit"`
	LimitAll        bool   `json:"all"`
	OrderBy         string `json:"order_by"`
	OrderDirection string `json:"order_direction"`
	Filters        map[string][]Filter `json:"-"`
}

// TemplateParams represents the template params for use with the
// posts function.
type TemplateParams struct {
	Params
	Resource       string `json:"resource"`
	Category       string `json:"category"`
}

// Filter represents the searching fields for searching through records.
type Filter struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// PaginationAllLimit defines how many items will be returned if
// the limit is set to list all
const (
	PaginationDefault      = 15
	PaginationDefaultOrder = "id,ASC"
)

// NewParams - create a new parameter type
func NewParams(g *gin.Context) *Params {
	return &Params{
		gin: g,
	}
}

// Get query Parameters for http API routes
func (p *Params) Get() Params {

	// Get page and set default
	var page int
	pageStr := p.gin.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page == 0 {
		page = 1
	}

	// Get limit & calculate if list all
	var limit int
	var limitAll bool
	limitStr := p.gin.Query("limit")
	if limitStr == "all" {
		limitAll = true
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			limit = PaginationDefault
		}
		if limit == 0 || limitStr == "" {
			limit = PaginationDefault
		}
	}

	// Get order and set defaults
	order := p.gin.Query("order")
	if order == "" {
		order = PaginationDefaultOrder
	}

	// Get order and set defaults
	orderArr := strings.Split(order, ",")
	var orderParams [2]string

	if len(orderArr) != 2 {
		orderParams[0] = "id"
		orderParams[1] = "ASC"
	} else {
		if orderArr[1] == "" {
			orderParams[0] = "id"
			orderParams[1] = "ASC"
		} else {
			orderParams[0] = orderArr[0]
			orderParams[1] = orderArr[1]
		}
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
		LimitAll:       limitAll,
		OrderBy:        orderParams[0],
		OrderDirection: orderParams[1],
		Filters:        filters,
	}
}

// GetTemplateParams query Parameters for templates
func GetTemplateParams(query map[string]interface{}) (TemplateParams, error) {
	const op = "http.GetTemplateParams"

	data, err := json.Marshal(query)
	if err != nil {
		return TemplateParams{}, &errors.Error{Code: errors.INVALID, Message: "Could not convert query to Template Params", Operation: op, Err: err}
	}

	var params TemplateParams
	if err = json.Unmarshal(data, &params); err != nil {
		return TemplateParams{}, &errors.Error{Code: errors.INVALID, Message: "Could not convert query to Template Params", Operation: op, Err: err}
	}

	// Set default page
	if params.Page == 0 {
		params.Page = 1
	}

	//Set default limit
	if params.Limit == 0 {
		params.Limit = PaginationDefault
	}

	// Set default resource
	if params.Resource == "" {
		params.Resource = "all"
	}

	// Set default order by
	if params.OrderBy == "" {
		params.OrderBy = "published_at"
	}

	// Set default order direction
	if params.OrderDirection == "" {
		params.OrderDirection = "desc"
	}

	return params, nil
}
