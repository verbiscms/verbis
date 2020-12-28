package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/spf13/cast"
)

func (t *TemplateManager) getCategoryByName(name string) *domain.Category {
	c, err := t.store.Categories.GetByName(name)
	if err != nil {
		return nil
	}
	return &c
}

func (t *TemplateManager) getCategoryByID(i interface{}) *domain.Category {
	id := cast.ToInt(i)
	c, err := t.store.Categories.GetById(id)
	if err != nil {
		return nil
	}
	return &c
}
