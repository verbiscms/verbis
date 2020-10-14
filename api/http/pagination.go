package http

import (
	"math"
)

type Pagination struct {
	Page 		int					`json:"page"`
	Pages 		int					`json:"pages"`
	Limit 		int					`json:"limit"`
	Total 		int					`json:"total"`
	Next 		*bool 				`json:"next"`
	Prev 		*bool 				`json:"prev"`
}

// Get pagination parameters
func GetPagination(params Params, total int) *Pagination {

	// Calculate total pages
	var pages int
	pages = int(math.Ceil(float64(total) / float64(params.Limit)))

	// Calculate prev and next variables
	var next = false
	var prev = false
	if params.Page + 1 < pages {
		next = true
	}
	if params.Page > 1 {
		prev = true
	}

	// Construct pagination meta
	var pagination *Pagination
	pagination = &Pagination{
		Page:  params.Page,
		Pages: pages,
		Limit: params.Limit,
		Total: total,
		Next:  &next,
		Prev:  &prev,
	}

	return pagination
}
