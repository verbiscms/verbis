package tpl

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/spf13/cast"
)

// getPost
//
// Obtains the post by ID and returns a domain.Post type
// or nil if not found.
//
// Example: {{ post 123 }}
func (t *TemplateManager) getPost(id interface{}) interface{} {
	i, err := cast.ToIntE(id)
	if err != nil {
		return nil
	}

	p, err := t.store.Posts.GetById(i)
	if err != nil {
		return nil
	}

	fp, err := t.formatPost(p)
	if err != nil {
		return nil
	}

	return fp
}

// getPosts
//
// Accepts a dict (map[string]interface{}) and returns an
// array of domain.Post. It sets defaults if some of the param
// arguments are missing, and returns an error if the data
// could not be marshalled.

// Returns errors.TEMPLATE if the template post params failed to parse.
//
// Example:
// {{ $result := post (dict "limit" 10 "resource" "posts") }}
// {{ with $result.Posts }}
//     {{ range $post := . }}
//         <h2>{{ $post.Title }}</h2>
//         <a href="{{ $post.Slug }}">Read more</a>
//     {{ end }}
//     {{ else }}
//         <h4>No posts found</h4>
// {{ end }}
func (t *TemplateManager) getPosts(query map[string]interface{}) (map[string]interface{}, error) {
	p, err := http.GetTemplateParams(query)
	if err != nil {
		return nil, err
	}

	posts, total, err := t.store.Posts.NewGetTest(p.Params, p.Resource, "published")
	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"Posts":      posts,
		"Pagination": http.NewPagination().Get(p.Params, total),
	}, nil
}

// getPagination
//
// Gets the page query parameter and returns, if the page
// query param wasn't found or the string could
// not be cast to an integer, it will return 1.
//
// Example: {{ paginationPage }}
func (t *TemplateManager) getPaginationPage() int {
	page := t.gin.Query("page")
	if page == "" {
		return 1
	}
	pageInt, err := cast.ToIntE(page)
	if err != nil {
		return 1
	}
	return pageInt
}

// formatPost
//
// Format's from the posts store and creates a new ViewPost
// ready to be returned to the template. It removes
// layouts from the formatting as it is not
// needed in the frontend.
func (t *TemplateManager) formatPost(post domain.Post) (domain.ViewPost, error) {

	fp, err := t.store.Posts.Format(post)
	if err != nil {
		return domain.ViewPost{}, err
	}

	return fp.ViewPost(), nil
}
