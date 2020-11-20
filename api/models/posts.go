package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"strings"
)

// PostsRepository defines methods for Posts to interact with the database
type PostsRepository interface {
	Get(meta http.Params, resource string) ([]domain.Post, int, error)
	Format(post domain.Post) (domain.PostData, error)
	FormatMultiple(posts []domain.Post) ([]domain.PostData, error)
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
	db              *sqlx.DB
	seoMetaModel    SeoMetaRepository
	userModel       UserRepository
	categoriesModel CategoryRepository
	fieldsModel FieldsRepository
}

// newPosts - Construct
func newPosts(db *sqlx.DB, config config.Configuration) *PostStore {
	return &PostStore{
		db:              db,
		seoMetaModel:    newSeoMeta(db),
		userModel:       newUser(db, config),
		categoriesModel: newCategories(db),
		fieldsModel:    newFields(db),
	}
}

// Get all posts
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no posts available.
func (s *PostStore) Get(meta http.Params, resource string) ([]domain.Post, int, error) {
	const op = "PostsRepository.Get"

	var p []domain.Post
	q := fmt.Sprintf("SELECT posts.*, post_options.seo 'options.seo', post_options.meta 'options.meta' FROM posts LEFT JOIN post_options ON posts.id = post_options.page_id")
	countQ := fmt.Sprintf("SELECT COUNT(*) FROM posts LEFT JOIN post_options ON posts.id = post_options.page_id")

	// Apply filters to total and original query
	filter, err := filterRows(s.db, meta.Filters, "posts")
	if err != nil {
		return nil, -1, err
	}
	q += filter
	countQ += filter

	// Get by resource
	if resource != "all" && resource != "" {
		if len(meta.Filters) > 0 {
			q += fmt.Sprintf(" AND")
			countQ += fmt.Sprintf(" AND")
		} else {
			q += fmt.Sprintf(" WHERE")
			countQ += fmt.Sprintf(" WHERE")
		}

		// If the resource is pages or a resource
		resourceQ := ""
		if resource == "pages" {
			resourceQ = fmt.Sprintf(" resource IS NULL")
		} else {
			resourceQ = fmt.Sprintf(" resource = '%s'", resource)
		}

		q += resourceQ
		countQ += resourceQ
	}

	// Apply pagination
	q += fmt.Sprintf(" ORDER BY posts.%s %s LIMIT %v OFFSET %v", meta.OrderBy, meta.OrderDirection, meta.Limit, (meta.Page-1)*meta.Limit)

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
//
// Returns errors.NOTFOUND if the post was not found by the given Id.
func (s *PostStore) GetById(id int) (domain.Post, error) {
	const op = "PostsRepository.GetById"
	var p domain.Post
	if err := s.db.Get(&p, "SELECT posts.*, post_options.seo 'options.seo', post_options.meta 'options.meta' FROM posts LEFT JOIN post_options ON posts.id = post_options.page_id WHERE posts.id = ? LIMIT 1", id); err != nil {
		return domain.Post{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the post with the ID: %d", id), Operation: op}
	}
	return p, nil
}

// GetBySlug returns a a post by slug
// Returns errors.NOTFOUND if the post was not found by the given slug.
func (s *PostStore) GetBySlug(slug string) (domain.Post, error) {
	const op = "PostsRepository.GetBySlug"
	var p domain.Post
	if err := s.db.Get(&p, "SELECT posts.*, post_options.seo 'options.seo', post_options.meta 'options.meta' FROM posts LEFT JOIN post_options ON posts.id = post_options.page_id WHERE posts.slug = ? LIMIT 1", slug); err != nil {
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

	if err := s.validateUrl(p.Slug); err != nil {
		return domain.Post{}, err
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

	// Update the categories based on the array of integers that
	// are passed.
	if err := s.categoriesModel.InsertPostCategory(int(id), p.Category); err != nil {
		return domain.Post{}, err
	}

	// Convert the PostCreate type to type of Post to be returned
	// to the controller, used for binding & validation.
	convertedPost := s.convertToPost(*p)
	convertedPost.Id = int(id)
	if err := s.seoMetaModel.UpdateCreate(&convertedPost); err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

// Update a post by Id
// Returns errors.NOTFOUND if the post was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *PostStore) Update(p *domain.PostCreate) (domain.Post, error) {
	const op = "PostsRepository.Update"

	oldPost, err := s.GetById(p.Id)
	if err != nil {
		return domain.Post{}, err
	}

	if oldPost.Slug != p.Slug {
		if err := s.validateUrl(p.Slug); err != nil {
			return domain.Post{}, err
		}
	}

	// Check if the author is set assign to owner if not.
	p.Author = s.checkOwner(*p)
	p.UserId = p.Author

	// Update the posts table with data
	q := "UPDATE posts SET slug = ?, title = ?, status = ?, resource = ?, page_template = ?, layout = ?, fields = ?, codeinjection_head = ?, codeinjection_foot = ?, user_id = ?, published_at = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(q, p.Slug, p.Title, p.Status, p.Resource, p.PageTemplate, p.Layout, p.Fields, p.CodeInjectHead, p.CodeInjectFoot, p.UserId, p.PublishedAt, p.Id)
	if err != nil {
		return domain.Post{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the post wuth the title: %v", p.Title), Operation: op, Err: err}
	}

	// Update the categories based on the array of integers that
	// are passed. If the categories
	if err := s.categoriesModel.InsertPostCategory(p.Id, p.Category); err != nil {
		return domain.Post{}, err
	}

	// Convert the PostCreate type to type of Post to be returned
	// to the controller, used for binding & validation.
	convertedPost := s.convertToPost(*p)
	if err := s.seoMetaModel.UpdateCreate(&convertedPost); err != nil {
		return domain.Post{}, err
	}

	// Clear the cache
	cache.Store.Delete(convertedPost.Slug)

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
		UUID:           c.UUID,
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
	if p.Author == 0 || !s.userModel.Exists(p.Author) {
		owner, err := s.userModel.GetOwner()
		if err != nil {
			log.Panic(err)
		}
		return owner.Id
	}
	return p.Author
}

// validateUrl checks if the url is valid for creating or updating a new
// post.
//
// Returns errors.CONFLICT if the post slug already exists
// Or the slug contains the admin path, .i.e /admin
func (s *PostStore) validateUrl(slug string) error {
	const op = "PostsRepository.validateUrl"

	if s.Exists(slug) {
		return &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the post, the slug %v, already exists", slug), Operation: op}
	}

	slugArr := strings.Split(slug, "/")
	if len(slugArr) > 1 {
		if strings.Contains(slugArr[1], "admin") {
			return &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the post, the path /admin is reserved"), Operation: op}
		}
	}

	return nil
}

// Format formats the post into as domain.PostData type which contains
// the category, author & fields associated with the post.
func (s *PostStore) Format(post domain.Post) (domain.PostData, error) {


	author, err := s.userModel.GetById(post.UserId)
	if err != nil {
		return domain.PostData{}, err
	}

	// Get the categories associated with the post
	category, err := s.categoriesModel.GetByPost(post.Id)

	// Get the layout associated with the post
	layout, err := s.fieldsModel.GetLayout(post, author, category)
	if err != nil {
		return domain.PostData{}, err
	}

	pd := domain.PostData{
		Post:   post,
		Layout: layout,
		Author: domain.PostAuthor(author),
	}

	if category != nil {
		pd.Categories = &domain.PostCategory{
			Id:          category.Id,
			Slug:        category.Slug,
			Name:        category.Name,
			Description: category.Description,
			Resource:    category.Resource,
			ParentId:    category.ParentId,
			UpdatedAt:   category.UpdatedAt,
			CreatedAt:   category.CreatedAt,
		}
	}

	return pd, nil
}

// FormatMultiple formats an array of posts to return categories,
// fields and author.
func (s *PostStore) FormatMultiple(posts []domain.Post) ([]domain.PostData, error) {
	var postData []domain.PostData
	for _, post := range posts {
		formatted, err := s.Format(post)
		if err != nil {
			return nil, err
		} else {
			postData = append(postData, formatted)
		}
	}
	return postData, nil
}