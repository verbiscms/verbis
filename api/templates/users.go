package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/spf13/cast"
)

func (t *TemplateManager) getUser(id interface{}) *domain.User {
	i, err := cast.ToIntE(id)
	if err != nil {
		return nil
	}

	user, err := t.store.User.GetById(i)
	if err != nil {
		return nil
	}

	user.HideCredentials()

	return &user
}

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

	users.HideCredentials()

	return map[string]interface{}{
		"Users":      users,
		"Pagination": http.NewPagination().Get(p.Params, total),
	}, nil
}
