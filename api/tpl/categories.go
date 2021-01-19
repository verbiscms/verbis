package tpl

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/spf13/cast"
)

// getCategory
//
// Obtains the category by ID and returns a domain.Category type
// or nil if not found.
//
// Example: {{ category 123 }}
func (t *TemplateManager) getCategory(id interface{}) interface{} {
	i, err := cast.ToIntE(id)
	if err != nil {
		return nil
	}

	category, err := t.store.Categories.GetById(i)
	if err != nil {
		return nil
	}

	return category
}

// getCategoryByName
//
// Obtains the category by name and returns a domain.Category type
// or nil if not found.
//
// Example: {{ categoryByName "sports" }}
func (t *TemplateManager) getCategoryByName(name interface{}) interface{} {
	n, err := cast.ToStringE(name)
	if err != nil {
		return nil
	}

	category, err := t.store.Categories.GetByName(n)
	if err != nil {
		return nil
	}

	return category
}

// getCategoryParent
//
// Obtains the category by parent and returns a domain.Category type
// or nil if not found.
//
// Example: {{ categoryByParent "sports" }}
func (t *TemplateManager) getCategoryParent(id interface{}) interface{} {
	i, err := cast.ToIntE(id)
	if err != nil {
		return nil
	}

	category, err := t.store.Categories.GetParent(i)
	if err != nil {
		return nil
	}

	return category
}

// getCategories
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
//         <h2>{{ $category.Name }}</h2>
//     {{ end }}
//     {{ else }}
//         <h4>No categories found</h4>
// {{ end }}
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
		"Categories": categories,
		"Pagination": http.NewPagination().Get(p.Params, total),
	}, nil
}
