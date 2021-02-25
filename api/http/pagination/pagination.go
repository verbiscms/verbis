// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pagination

import (
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"math"
)

// Pagination represents the data to be sent back from the
// API on list routes.
type Pagination struct {
	Page  int         `json:"page"`
	Pages int         `json:"pages"`
	Limit interface{} `json:"limit"`
	Total int         `json:"total"`
	Next  interface{} `json:"next"`
	Prev  interface{} `json:"prev"`
}

// Get
//
// Uses the parameters to return formatted pagination on
// list routes.
func Get(params params.Params, total int) *Pagination {

	// Calculate total pages
	var pages int
	pages = int(math.Ceil(float64(total) / float64(params.Limit)))

	// Set page to 1 if the user has passed "?limit=all"
	var limit interface{}
	if params.LimitAll {
		pages = 1
		limit = "all"
	} else {
		limit = params.Limit
	}

	// Construct pagination meta
	var pagination *Pagination
	pagination = &Pagination{
		Page:  params.Page,
		Pages: pages,
		Limit: limit,
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
