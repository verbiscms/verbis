package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/spf13/cast"
)

func (t *TemplateManager) getCategory(id interface{}) *domain.Category {
	i, err := cast.ToIntE(id)
	if err != nil {
		return nil
	}

	category, err := t.store.Categories.GetById(i)
	if err != nil {
		return nil
	}

	return &category
}

func (t *TemplateManager) getCategoryByName(name interface{}) *domain.Category {
	n, err := cast.ToStringE(name)
	if err != nil {
		return nil
	}

	category, err := t.store.Categories.GetByName(n)
	if err != nil {
		return nil
	}
	return &category
}

func (t *TemplateManager) getCategoryByParent(id interface{}) *domain.Category {
	category := t.getCategory(id)
	if category == nil || category.ParentId == nil {
		return nil
	}
	return t.getCategory(category.ParentId)
}

func (t *TemplateManager) getCategories(query map[string]interface{}) (map[string]interface{}, error) {
	p, err := http.GetTemplateParams(query)
	if err != nil {
		return nil, err
	}

	categories, total, err := t.store.Categories.Get(p.Params)
	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"Categories":      categories,
		"Pagination": http.NewPagination().Get(p.Params, total),
	}, nil
}
