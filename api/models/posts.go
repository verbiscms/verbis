package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// PostsRepository defines methods for Posts to interact with the database
type PostsRepository interface {
	Get(meta http.Params, resource string) ([]domain.Post, int, error)
	GetById(id int) (domain.Post, error)
	GetBySlug(slug string) (domain.Post, error)
	Create(p *domain.PostCreate) (domain.Post, error)
	Update(p *domain.PostCreate) (domain.Post, error)
	Delete(id int) error
	Exists(slug string) bool
	Total() (int, error)
}

// PostStore defines the data layer for Posts
type PostStore struct {
	db *sqlx.DB
	seoMetaModel 	SeoMetaRepository
	userModel 	 	UserRepository
	categoriesModel CategoryRepository
}

// newPosts - Construct
func newPosts(db *sqlx.DB) *PostStore {
	return &PostStore{
		db: db,
		seoMetaModel: newSeoMeta(db),
		userModel: newUser(db),
		categoriesModel: newCategories(db),
	}
}

// Get all posts
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no posts available.
func (s *PostStore) Get(meta http.Params, resource string) ([]domain.Post, int, error) {
	const op = "PostsRepository.Get"

	var p []domain.Post
	q := fmt.Sprintf("SELECT posts.*, seo_meta_options.seo 'options.seo', seo_meta_options.meta 'options.meta' FROM posts LEFT JOIN seo_meta_options ON posts.id = seo_meta_options.page_id")
	countQ := fmt.Sprintf("SELECT COUNT(*) FROM posts LEFT JOIN seo_meta_options ON posts.id = seo_meta_options.page_id")

	fmt.Println()

	// Get by resource
	if resource != "all" && resource != "" {
		resourceQ := fmt.Sprintf(" WHERE resource = '%s'", resource)
		q += resourceQ
		countQ += resourceQ
	}

	// Apply filters to total and original query
	filter, err := filterRows(s.db, meta.Filters, "posts")
	if err != nil {
		return nil, -1, err
	}
	q += filter
	countQ += filter

	// Apply pagination
	q += fmt.Sprintf(" ORDER BY posts.%s %s LIMIT %v OFFSET %v", meta.OrderBy, meta.OrderDirection, meta.Limit, (meta.Page - 1) * meta.Limit)

	// Select posts
	if err := s.db.Select(&p, q); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get posts", Operation: op, Err: err}
	}

	// Return not found error if no posts are available
	if len(p) == 0 {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No posts available", Operation: op}
	}

	// Count the total number of posts
	var total int
	if err := s.db.QueryRow(countQ).Scan(&total); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of posts", Operation: op, Err: err}
	}

	return p, total, nil
}

// GetById returns a post by Id
// Returns errors.NOTFOUND if the post was not found by the given Id.
func (s *PostStore) GetById(id int) (domain.Post, error) {
	const op = "PostsRepository.GetById"
	var p domain.Post
	if err := s.db.Get(&p, "SELECT posts.*, seo_meta_options.seo 'options.seo', seo_meta_options.meta 'options.meta' FROM posts LEFT JOIN seo_meta_options ON posts.id = seo_meta_options.page_id WHERE posts.id = ? LIMIT 1", id); err != nil {
		return domain.Post{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the post with the ID: %d", id), Operation: op}
	}
	return p, nil
}

// GetBySlug returns a a post by slug
// Returns errors.NOTFOUND if the post was not found by the given slug.
func (s *PostStore) GetBySlug(slug string) (domain.Post, error) {
	const op = "PostsRepository.GetBySlug"
	var p domain.Post
	if err := s.db.Get(&p, "SELECT * FROM posts WHERE slug = ?", slug); err != nil {
		return domain.Post{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get post with the slug %s", slug), Operation: op}
	}
	return p, nil
}

// Create a new post
// Returns errors.CONFLICT if the the post slug already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function
// could not get the newly created ID.
func (s *PostStore) Create(p *domain.PostCreate) (domain.Post, error) {
	const op = "PostsRepository.Create"

	if s.Exists(p.Slug) {
		return domain.Post{}, &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the post, the slug %v, already exists", p.Slug), Operation: op}
	}

	// Check if the author is set assign to owner if not.
	p.UserId = s.checkOwner(*p)

	// TODO: Work out why sql defaults arent working!
	if p.Layout == "" {
		p.Layout = "default"
	}
	if p.PageTemplate == "" {
		p.PageTemplate = "default"
	}
	if p.Status == "" {
		p.Status = "draft"
	}

	q := "INSERT INTO posts (uuid, slug, title, status, resource, page_template, layout, fields, codeinjection_head, codeinjection_foot, user_id, published_at, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	c, err := s.db.Exec(q, uuid.New().String(), p.Slug, p.Title, p.Status, p.Resource, p.PageTemplate, p.Layout, p.Fields, p.CodeInjectHead, p.CodeInjectFoot, p.UserId, p.PublishedAt)
	if err != nil {
		return domain.Post{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the post with the title: %v", p.Title), Operation: op, Err: err}
	}

	id, err := c.LastInsertId()
	if err != nil {
		return domain.Post{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created post ID with the title: %v", p.Title), Operation: op, Err: err}
	}

	post, err := s.GetById(int(id))
	if err != nil {
		return domain.Post{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created post with the title: %v", p.Title), Operation: op, Err: err}
	}

	// Convert the PostCreate type to type of Post to be returned
	// to the controller, used for binding & validation.
	if err := s.seoMetaModel.UpdateCreate(&post); err != nil {
		return domain.Post{},  err
	}

	return post, nil
}

// Update a post by Id
// Returns errors.NOTFOUND if the post was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *PostStore) Update(p *domain.PostCreate) (domain.Post, error) {
	const op = "PostsRepository.Update"

	_, err := s.GetById(p.Id)
	if err != nil {
		return domain.Post{}, err
	}

	// Check if the author is set assign to owner if not.
	p.Author = s.checkOwner(*p)
	p.UserId = p.Author

	// Update the posts table with data
	q := "UPDATE posts SET slug = ?, title = ?, status = ?, resource = ?, page_template = ?, layout = ?, fields = ?, codeinjection_head = ?, codeinjection_foot = ?, user_id = ?, published_at = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(q, p.Slug, p.Title, p.Status, p.Resource, p.PageTemplate, p.Layout, p.Fields, p.CodeInjectHead, p.CodeInjectFoot, p.Author, p.PublishedAt, p.Id)
	if err != nil {
		return domain.Post{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the post wuth the title: %v", p.Title), Operation: op, Err: err}
	}

	// Update the categories based on the array of integers that
	// are passed. If the categories
	if err := s.categoriesModel.InsertPostCategories(p.Id, p.Categories); err != nil {
		return domain.Post{}, err
	}

	// Convert the PostCreate type to type of Post to be returned
	// to the controller, used for binding & validation.
	convertedPost := s.convertToPost(*p)
	if err := s.seoMetaModel.UpdateCreate(&convertedPost); err != nil {
		return domain.Post{}, err
	}

	return convertedPost, nil
}

// Delete post
// Returns errors.NOTFOUND if the post was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *PostStore) Delete(id int) error {
	const op = "PostsRepository.Delete"

	_, err := s.GetById(id)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("DELETE FROM posts WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete post with the ID: %v", id), Operation: op, Err: err}
	}

	return nil
}

// Total gets the total number of posts
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *PostStore) Total() (int, error) {
	const op = "PostsRepository.Total"
	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM posts").Scan(&total); err != nil {
		return -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of posts", Operation: op, Err: err}
	}
	return total, nil
}

// Exists Checks if a post exists by the given slug
func (s *PostStore) Exists(slug string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM posts WHERE slug = ?)", slug).Scan(&exists)
	return exists
}

// convertToPost converts are post create into a standard post
func (s *PostStore) convertToPost(c domain.PostCreate) domain.Post {
	return domain.Post{
		Id:             c.Id,
		UUID:  			c.UUID,
		Slug:           c.Slug,
		Title:          c.Title,
		Status:         c.Status,
		Resource:       c.Resource,
		PageTemplate:   c.PageTemplate,
		Fields:         c.Fields,
		CodeInjectHead: c.CodeInjectHead,
		CodeInjectFoot: c.CodeInjectFoot,
		UserId:         c.UserId,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
		SeoMeta:        c.SeoMeta,
	}
}

// checkOwner Checks if the author is set or if the author does not exist.
// Returns the owner ID under circumstances.
func (s *PostStore) checkOwner(p domain.PostCreate) int {
	if p.UserId == 0 || !s.userModel.Exists(p.Author) {
		owner, err := s.userModel.GetOwner()
		if err != nil {
			log.Panic(err)
		}
		return owner.Id
	}
	return p.Id
}