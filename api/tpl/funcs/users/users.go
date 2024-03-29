// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/pagination"
	"github.com/verbiscms/verbis/api/tpl/params"
)

const (
	// The default order by field for the list function.
	OrderBy = "created_at"
	// The default order direction field for the list function.
	OrderDirection = "desc"
)

// Find
//
// Obtains the user by ID and returns a domain.UserPart type
// or nil if not found.
//
// Example: {{ user 123 }}
func (ns *Namespace) Find(id interface{}) interface{} {
	i, err := cast.ToIntE(id)
	if err != nil || id == nil {
		return nil
	}

	user, err := ns.deps.Store.User.Find(i)
	if err != nil {
		return nil
	}

	return user.HideCredentials()
}

// Categories defines the struct for returning
// categories and pagination back to the
// template.
type Users struct {
	Users      domain.UsersParts
	Pagination *pagination.Pagination
}

// List
//
// Accepts a dict (map[string]interface{}) and returns an
// array of domain.UserPart. It sets defaults if some of the param
// arguments are missing, and returns an error if the data
// could not be marshalled.

// Returns errors.TEMPLATE if the template user params failed to parse.
//
// Example:
// {{ $result := users (dict "limit" 10) }}
// {{ with $result.Users }}
//     {{ range $user := . }}
//         <h2>{{ $user.SizeName }}</h2>
//     {{ end }}
//     {{ else }}
//         <h4>No users found</h4>
// {{ end }}
func (ns *Namespace) List(query params.Query) (interface{}, error) {
	p := query.Get(OrderBy, OrderDirection)

	role := query.Default("role", "")

	users, total, err := ns.deps.Store.User.List(p, role.(string))
	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return Users{
		Users:      users.HideCredentials(),
		Pagination: pagination.Get(p, total),
	}, nil
}
