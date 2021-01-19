package tpl

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/spf13/cast"
)

// getUser
//
// Obtains the user by ID and returns a domain.UserPart type
// or nil if not found.
//
// Example: {{ user 123 }}
func (t *TemplateManager) getUser(id interface{}) interface{} {
	i, err := cast.ToIntE(id)
	if err != nil {
		return nil
	}

	user, err := t.store.User.GetById(i)
	if err != nil {
		return nil
	}

	return user.HideCredentials()
}

// getUsers
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
//         <h2>{{ $user.Name }}</h2>
//     {{ end }}
//     {{ else }}
//         <h4>No users found</h4>
// {{ end }}
func (t *TemplateManager) getUsers(query map[string]interface{}) (map[string]interface{}, error) {
	p, err := http.GetTemplateParams(query)
	if err != nil {
		return nil, err
	}

	users, total, err := t.store.User.Get(p.Params)
	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"Users":      users.HideCredentials(),
		"Pagination": http.NewPagination().Get(p.Params, total),
	}, nil
}
