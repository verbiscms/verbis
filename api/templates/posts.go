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

// Get the post resource
func (t *TemplateFunctions) getResource() string {
	resource := t.post.Resource
	if resource == nil {
		return ""
	}
	return *resource
}


func (t *TemplateFunctions) getResources(query map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{}
}
