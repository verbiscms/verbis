package url

import (
	"github.com/spf13/cast"
	"golang.org/x/net/html"
)

// Query
//
// Gets the page query parameter and returns, if the page
// query param wasn't found or the string could
// not be cast to an integer, it will return 1.
//
// Example: {{ paginationPage }}
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

	return html.EscapeString(val[0])
}

// Pagination
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
	return cast.ToInt(page)
}
