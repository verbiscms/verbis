package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/spf13/cast"
)

// ViewPost represents the data to be sent back to the template
// once acquired.
type ViewPost struct {
	Author   *domain.PostAuthor
	Category *domain.PostCategory
	domain.Post
}

// getPost
//
// Obtains the post by ID and returns a domain.Post type
// or nil if not found.
//
// Example: {{ post 123 }}
func (t *TemplateFunctions) getPost(i interface{}) *ViewPost {
	p, err := t.store.Posts.GetById(cast.ToInt(i))
	if err != nil {
		return nil
	}

	fp, err := t.formatPost(p)
	if err != nil {
		return nil
	}

	return &fp
}

// getPosts
//
// Accepts a dict (map[string]interface{}) and returns an
// array of domain.Post. It sets defaults if some of the param
// arguments are missing, and returns an error if the data
// could not be marshalled.
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
func (t *TemplateFunctions) getPosts(query map[string]interface{}) (map[string]interface{}, error) {

	p, err := http.GetTemplateParams(query)
	if err != nil {
		return nil, err
	}

	// Obtain the post and detect if it was not found,
	// return nil if so.
	posts, total, err := t.store.Posts.Get(p.Params, p.Resource, "published")
	pagination := http.NewPagination().Get(p.Params, total)

	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var returnPosts []ViewPost
	for _, post := range posts {
		formattedPost, err := t.formatPost(post)
		if err != nil {
			continue
		}

		if p.Category == "" {
			returnPosts = append(returnPosts, formattedPost)
			continue
		}

		if p.Category == formattedPost.Category.Name {
			returnPosts = append(returnPosts, formattedPost)
		}
	}

	return map[string]interface{}{
		"Posts":      returnPosts,
		"Pagination": pagination,
	}, nil
}

// getPagination
//
// Gets the page query parameter and returns, if the page
// query param wasn't found or the string could
// not be cast to an integer, it will return 1.
//
// Example: {{ paginationPage }}
func (t *TemplateFunctions) getPaginationPage() int {
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
func (t *TemplateFunctions) formatPost(post domain.Post) (ViewPost, error) {

	fp, err := t.store.Posts.Format(post)
	if err != nil {
		return ViewPost{}, err
	}

	return ViewPost{
		Author:   fp.Author,
		Category: fp.Category,
		Post:     fp.Post,
	}, nil
}
