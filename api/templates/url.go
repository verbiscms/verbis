package templates

import (
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/spf13/cast"
)

func (t *TemplateFunctions) getBaseURL() string {
	return location.Get(t.gin).String()
}

func (t *TemplateFunctions) getScheme() string {
	return location.Get(t.gin).Scheme
}

func (t *TemplateFunctions) getHost() string {
	return location.Get(t.gin).Host
}

func (t *TemplateFunctions) getFullURL() string {
	fmt.Println(t.gin.Request.Host)
	return location.Get(t.gin).String() + t.gin.Request.URL.Path
}

func (t *TemplateFunctions) getURL() string {
	return t.gin.Request.URL.String()
}

func (t *TemplateFunctions) getQueryParams(i interface{}) string {
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