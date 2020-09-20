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
	if params.Limit == 0 {
		pages = int(math.Round(float64(total / params.Limit)))
	} else {
		pages = int(math.Round(float64(total / params.Limit + 1)))
	}

	// Calculate prev and next variables
	var next bool = false
	var prev bool = false
	if params.Page + 1 < pages {
		next = true
	}
	if params.Page > 0 {
		prev = true
	}
	if params.Page >= total {
		prev = true
	}

	// Construct pagination meta
	var pagination *Pagination
	if (Params{}) == params {
		pagination = nil
	} else {
		pagination = &Pagination{
			Page:  params.Page,
			Pages: pages,
			Limit: params.Limit,
			Total: total,
			Next:  &next,
			Prev:  &prev,
		}
	}

	return pagination
}
