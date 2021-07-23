// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	store "github.com/ainsleyclark/verbis/api/store/categories"
	"github.com/ainsleyclark/verbis/api/tpl/params"
	"github.com/spf13/cast"
)

const (
	// The default order by field for the list function.
	OrderBy = "created_at"
	// The default order direction field for the list function.
	OrderDirection = "desc"
)

// Find
//
// Obtains the category by ID and returns a domain.Category type
// or nil if not found.
//
// Example: {{ category 123 }}
func (ns *Namespace) Find(id interface{}) interface{} {
	i, err := cast.ToIntE(id)
	if err != nil || id == nil {
		return nil
	}

	category, err := ns.deps.Store.Categories.Find(i)
	if err != nil {
		return nil
	}

	return category
}

// ByName
//
// Obtains the category by name and returns a domain.Category type
// or nil if not found.
//
// Example: {{ categoryByName "sports" }}
func (ns *Namespace) ByName(name interface{}) interface{} {
	n, err := cast.ToStringE(name)
	if err != nil || name == nil {
		return nil
	}

	category, err := ns.deps.Store.Categories.FindByName(n)
	if err != nil {
		return nil
	}

	return category
}

// ByParent
//
// Obtains the category by parent and returns a domain.Category type
// or nil if not found.
//
// Example: {{ categoryByParent "sports" }}
func (ns *Namespace) Parent(id interface{}) interface{} {
	i, err := cast.ToIntE(id)
	if err != nil || id == nil {
		return nil
	}

	category, err := ns.deps.Store.Categories.FindParent(i)
	if err != nil {
		return nil
	}

	return category
}

// Categories defines the struct for returning
// categories and pagination back to the
// template.
type Categories struct {
	Categories domain.Categories
	Pagination *pagination.Pagination
}

// List
//
// Accepts a dict (map[string]interface{}) and returns an
// array of domain.Category. It sets defaults if some of the param
// arguments are missing, and returns an error if the data
// could not be marshalled.

// Returns errors.TEMPLATE if the template post category failed to parse.
//
// Example:
// {{ $result := categories (dict "limit" 10) }}
// {{ with $result.Categories }}
//     {{ range $category := . }}
//         <h2>{{ $category.SizeName }}</h2>
//     {{ end }}
//     {{ else }}
//         <h4>No categories found</h4>
// {{ end }}
func (ns *Namespace) List(query params.Query) (interface{}, error) {
	p := query.Get(OrderBy, OrderDirection)

	resource := query.Default("resource", "")

	cfg := store.ListConfig{
		Resource: resource.(string),
	}

	categories, total, err := ns.deps.Store.Categories.List(p, cfg)
	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return Categories{
		Categories: categories,
		Pagination: pagination.Get(p, total),
	}, nil
}
