package tpl

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/spf13/cast"
)

func (t *TemplateManager) getUser(id interface{}) domain.UserPart {
	i, err := cast.ToIntE(id)
	if err != nil {
		return domain.UserPart{}
	}

	user, err := t.store.User.GetById(i)
	if err != nil {
		return domain.UserPart{}
	}

	return user.HideCredentials()
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

	return map[string]interface{}{
		"Users":      users.HideCredentials(),
		"Pagination": http.NewPagination().Get(p.Params, total),
	}, nil
}
