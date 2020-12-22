package templates

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/spf13/cast"
	"strconv"
)

type ViewPost struct {
	Author   *domain.PostAuthor
	Category *domain.PostCategory
	domain.Post
}

// getPost
//
// Obtains the post by ID and returns a domain.Post type
// or nil if not found.
func (t *TemplateFunctions) getPost(i interface{}) *ViewPost {
	p, err := t.store.Posts.GetById(cast.ToInt(i))
	if err != nil {
		return nil
	}

	vp, err := t.formatPost(p)
	if err != nil {
		return nil
	}

	return &vp
}

// getPosts
//
// Accepts a dict (map[string]interface{}) and returns an
// array of domain.Post. It sets defaults if some of the param
// arguments are missing, and returns an error if the data
// could not be marshalled.
func (t *TemplateFunctions) getPosts(query map[string]interface{}) (map[string]interface{}, error) {

	type params struct {
		Page           int    `json:"page"`
		Limit          int    `json:"limit"`
		Resource       string `json:"resource"`
		OrderBy        string `json:"order_by"`
		OrderDirection string `json:"order_direction"`
		Category       string `json:"category"`
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
			Category:       "",
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
			Value:    "published",
		},
	}

	postParams := http.Params{
		Page:           tmplParams.Page,
		Limit:          tmplParams.Limit,
		OrderBy:        tmplParams.OrderBy,
		OrderDirection: tmplParams.OrderDirection,
		Filters:        f,
	}

	// Obtain the post and detect if it was not found,
	// return nil if so.
	posts, total, err := t.store.Posts.Get(postParams, tmplParams.Resource)

	pagination := http.NewPagination().Get(postParams, total)

	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var returnPosts []ViewPost
	for _, post := range posts {
		formattedPost, err := t.formatPost(post)
		if err != nil {

		}
		returnPosts = append(returnPosts, formattedPost)
		//if tmplParams.Category == "" {
		//	returnPosts = append(returnPosts, formattedPost)
		//}
		//if formattedPost.Category.Name == tmplParams.Category {
		//	returnPosts = append(returnPosts, formattedPost)
		//}
	}

	return map[string]interface{}{
		"Posts":      returnPosts,
		"Pagination": pagination,
	}, nil
}

func (t *TemplateFunctions) formatPost(post domain.Post) (ViewPost, error) {

	fp, err := t.store.Posts.Format(post)
	if err != nil {
		return ViewPost{}, err
	}

	return ViewPost{
		Author: fp.Author,
		Category: fp.Category,
		Post: fp.Post,
	}, nil
}

// getPagination gets the page query parameter and returns, if the
// page query param wasn't found or the string could not be cast
// to an integer, it will return 1.
func (t *TemplateFunctions) getPaginationPage() int {
	page := t.gin.Query("page")
	if page == "" {
		return 1
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return 1
	}
	return pageInt
}
