package templates

import "github.com/ainsleyclark/verbis/api/domain"

// getPost obtains the post by ID and returns a domain.Post type
// or nil if not found.
func (t *TemplateFunctions) getPost(id float64) *domain.Post {
	p, err := t.store.Posts.GetById(int(id))
	if err != nil {
		return nil
	}
	return &p
}