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

// getPagination
//
// Gets the page query parameter and returns, if the page
// query param wasn't found or the string could
// not be cast to an integer, it will return 1.
//
// Example: {{ paginationPage }}
//func (ns *Namespace) getPaginationPage() int {
//	page := t.gin.Query("page")
//	if page == "" {
//		return 1
//	}
//	pageInt, err := cast.ToIntE(page)
//	if err != nil {
//		return 1
//	}
//	return pageInt
//}

//func (t *TplTestSuite) TestGetPagination() {
//	g, _ := gin.CreateTestContext(httptest.NewRecorder())
//	g.Request, _ = http.NewRequest("GET", "/get?page=123", nil)
//	t.gin = g
//	tpl := "{{ paginationPage }}"
//	t.RunT(tpl, 123)
//}
//
//func (t *TplTestSuite) TestGetPagination_NoPage() {
//	tpl := "{{ paginationPage }}"
//	t.RunT(tpl, 1)
//}
//
//func (t *TplTestSuite) TestGetPagination_ConvertString() {
//	g, _ := gin.CreateTestContext(httptest.NewRecorder())
//	g.Request, _ = http.NewRequest("GET", "/get?page=wrongval", nil)
//	t.gin = g
//	tpl := "{{ paginationPage }}"
//	t.RunT(tpl, "1")
//}
