package url

import (
	"github.com/spf13/cast"
)

func (ns *Namespace) Query(i interface{}) string {
	key, err := cast.ToStringE(i)
	if err != nil {
		return ""
	}

	query := ns.ctx.Request.URL.Query()
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
func (ns *Namespace) Pagination() int {
	page := ns.ctx.Query("page")
	if page == "" {
		return 1
	}
	pageInt, err := cast.ToIntE(page)
	if err != nil {
		return 1
	}
	return pageInt
}

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
