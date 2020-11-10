package templates

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"strconv"
)

// getPost obtains the post by ID and returns a domain.Post type
// or nil if not found.
func (t *TemplateFunctions) getPost(id float64) *domain.Post {
	p, err := t.store.Posts.GetById(int(id))
	if err != nil {
		return nil
	}
	return &p
}

// getPosts accepts a dict (map[string]interface{}) and returns an
// array of domain.Post. It sets defaults if some of the param
// arguments are missing, and returns an error if the data
// could not be marshalled.
func (t *TemplateFunctions) getPosts(query map[string]interface{}) (map[string]interface{}, error) {

	type params struct {
		Page 			int
		Limit 			int
		Resource 		string
		OrderBy 		string
		OrderDirection 	string
	}

	data, _ := json.Marshal(query)
	var tmplParams params
	err := json.Unmarshal(data, &tmplParams)
	if err != nil {
		return nil, err
	}

	// Set nil
	if (params{} == tmplParams) {
		tmplParams = params{
			Page:           1,
			Limit:          15,
			Resource:       "all",
			OrderBy:        "published_at",
			OrderDirection: "desc",
		}
	}

	// Set default page
	if tmplParams.Page == 0 {
		tmplParams.Page = 1
	}

	//Set default limit
	if tmplParams.Limit == 0 {
		tmplParams.Limit = http.PaginationDefault
	}

	// Set default resource
	if tmplParams.Resource == "" {
		tmplParams.Resource = "all"
	}

	// Set default order by
	if tmplParams.OrderBy == "" {
		tmplParams.OrderBy = "published_at"
	}

	// Set default order direction
	if tmplParams.OrderDirection == "" {
		tmplParams.OrderDirection = "desc"
	}

	f := make(map[string][]http.Filter)
	f["status"] = []http.Filter{
		{
			Operator: "=",
			Value: "published",
		},
	}

	postParams := http.Params{
		Page:           tmplParams.Page,
		Limit:          tmplParams.Limit,
		OrderBy:        tmplParams.OrderBy,
		OrderDirection: tmplParams.OrderDirection,
		Filters: 		f,
	}

	// Obtain the post and detect if it was not found,
	// return nil if so.
	posts, total, err := t.store.Posts.Get(postParams, tmplParams.Resource)

	pagination := http.GetPagination(postParams, total)

	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"Items": posts,
		"Pagination": pagination,
	}, nil
}

// getPagination gets the page query paramater and returns, if the
// page query param wasn't found or the string could not be cast
// to an integer, it will return 1, along with an error if
// the cast failed.
func (t *TemplateFunctions) getPagination() (int, error) {
	page := t.gin.Query("page")
	if page == "" {
		return 1, nil
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return 1, err
	}
	return pageInt, nil
}