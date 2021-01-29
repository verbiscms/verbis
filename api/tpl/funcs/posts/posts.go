package posts

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/tpl/params"
	"github.com/spf13/cast"
)

const (
	// The default order by field for the list function.
	OrderBy = "updated_at"
	// The default order direction field for the list function.
	OrderDirection = "desc"
)

// Find
//
// Obtains the post by ID and returns a domain.PostData type
// or nil if not found.
//
// Example: {{ post 123 }}
func (ns *Namespace) Find(id interface{}) interface{} {
	i, err := cast.ToIntE(id)
	if err != nil {
		return nil
	}

	post, err := ns.deps.Store.Posts.GetById(i, false)
	if err != nil {
		return nil
	}

	return post
}

// Posts defines the struct for returning
// posts and pagination back to the
// template.
type Posts struct {
	Posts []TplPost
	Pagination *http.Pagination
}

// Tpl
type TplPost struct {
	domain.Post
	Author   domain.UserPart
	Category *domain.Category
	Fields   []domain.PostField
}

// List
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
func (ns *Namespace) List(query params.Query) (interface{}, error) {
	p := query.Get(OrderBy, OrderDirection)

	resource := query.Default("resource", "")
	status := query.Default("status", "published")

	posts, total, err := ns.deps.Store.Posts.Get(p, false, resource.(string), status.(string))
	if errors.Code(err) == errors.NOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var tplPosts = make([]TplPost, len(posts))
	for i, v := range posts {
		tplPosts[i] = TplPost{
			Post:     v.Post,
			Author:   v.Author,
			Category: v.Category,
			Fields:   v.Fields,
		}
	}

	return Posts{
		Posts: tplPosts,
		Pagination: http.NewPagination().Get(p, total),
	}, nil
}

