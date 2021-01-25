package tpl

import (
	"github.com/gin-contrib/location"
	"github.com/spf13/cast"
)

func (t *TemplateManager) getBaseURL() string {
	return location.Get(t.gin).String()
}

func (t *TemplateManager) getScheme() string {
	return location.Get(t.gin).Scheme
}

func (t *TemplateManager) getHost() string {
	return location.Get(t.gin).Host
}

func (t *TemplateManager) getFullURL() string {
	return location.Get(t.gin).String() + t.gin.Request.URL.Path
}

func (t *TemplateManager) getURL() string {
	return t.gin.Request.URL.String()
}

func (t *TemplateManager) getQueryParams(i interface{}) string {
	key, err := cast.ToStringE(i)
	if err != nil {
		return ""
	}

	query := t.gin.Request.URL.Query()
	val, ok := query[key]
	if !ok {
		return ""
	}

	return val[0]
}
