package http

import (
	"math"
)

// ParameterHandler defines the function for getting http params
type PaginationHandler interface  {
	Get() Params
}

// Pagination represents the data to be sent back from the API on
// get routes
type Pagination struct {
	Page  int         `json:"page"`
	Pages int         `json:"pages"`
	Limit int         `json:"limit"`
	Total int         `json:"total"`
	Next  interface{} `json:"next"`
	Prev  interface{} `json:"prev"`
}

// NewPagination - create a new pagination type
func NewPagination() *Pagination {
	return &Pagination{}
}

// Get pagination parameters
func (p *Pagination) GetPagination(params Params, total int) *Pagination {

	// Calculate total pages
	var pages int
	pages = int(math.Ceil(float64(total) / float64(params.Limit)))

	// Construct pagination meta
	var pagination *Pagination
	pagination = &Pagination{
		Page:  params.Page,
		Pages: pages,
		Limit: params.Limit,
		Total: total,
		Next:  false,
		Prev:  false,
	}

	// Calculate prev and next variables
	if params.Page < pages {
		pagination.Next = params.Page + 1
	}
	if params.Page > 1 {
		pagination.Prev = params.Page - 1
	}

	return pagination
}
