package models

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type PostsRepository interface {
	Get(meta http.Params) ([]domain.Post, error)
	GetById(id int) (domain.Post, error)
	GetBySlug(slug string) (domain.Post, error)
	Create(p *domain.PostCreate) (domain.Post, error)
	Update(p *domain.PostCreate) (domain.Post, error)
	Delete(id int) error
	Exists(slug string) bool
	Total() (int, error)
}

type PostStore struct {
	db *sqlx.DB
	seoMetaModel 	SeoMetaRepository
	userModel 	 	UserRepository
	categoriesModel CategoryRepository
}

//Construct
func newPosts(db *sqlx.DB, sm SeoMetaRepository, um UserRepository, cm CategoryRepository) *PostStore {
	ps := &PostStore{
		db: db,
		seoMetaModel: sm,
		userModel: um,
		categoriesModel: cm,
	}

	return ps
}

// Get all posts
func (s *PostStore) Get(meta http.Params) ([]domain.Post, error) {
	var p []domain.Post
	q := fmt.Sprintf("SELECT posts.*, seo_meta_options.seo 'options.seo', seo_meta_options.meta 'options.meta' FROM posts LEFT JOIN seo_meta_options ON posts.id = seo_meta_options.page_id ORDER BY posts.%s %s LIMIT %v OFFSET %v", meta.OrderBy, meta.OrderDirection, meta.Limit, meta.Page * meta.Limit)
	if err := s.db.Select(&p, q); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Could not get posts")
	}

	if len(p) == 0 {
		return []domain.Post{}, nil
	}

	return p, nil
}

// Get the post by ID
func (s *PostStore) GetById(id int) (domain.Post, error) {
	var p domain.Post
	if err := s.db.Get(&p, "SELECT posts.*, seo_meta_options.seo 'options.seo', seo_meta_options.meta 'options.meta' FROM posts LEFT JOIN seo_meta_options ON posts.id = seo_meta_options.page_id WHERE posts.id = ? LIMIT 1", id); err != nil {
		log.Info(err)
		return domain.Post{}, fmt.Errorf("Could not get post with the ID: %v", id)
	}
	return p, nil
}

// Get the post by slug
func (s *PostStore) GetBySlug(slug string) (domain.Post, error) {
	var p domain.Post
	if err := s.db.Get(&p, "SELECT * FROM posts WHERE slug = ?", slug); err != nil {
		log.Info(err)
		return domain.Post{}, fmt.Errorf("Could not get post with the slug %v", slug)
	}
	return p, nil
}

// Create post
func (s *PostStore) Create(p *domain.PostCreate) (domain.Post, error) {

	// See if the slug already exists within the database.
	// Bail if exists
	if s.Exists(p.Slug) {
		return domain.Post{}, fmt.Errorf("Could not create the post, the slug %v, already exists", p.Slug)
	}

	// Check if the author is set assign to owner if not.
	p.Id = s.checkOwner(*p)

	// Insert into posts table with data
	q := "INSERT INTO posts (uuid, slug, title, status, resource, page_template, fields, codeinjection_head, codeinjection_foot, user_id, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	c, err := s.db.Exec(q, uuid.New().String(), p.Slug, p.Title, p.Status, p.Resource, p.PageTemplate, p.Fields, p.CodeInjectHead, p.CodeInjectFoot, p.UserId)
	if err != nil {
		log.Error(err)
		return domain.Post{}, fmt.Errorf("Could not create the post with the title: %v", p.Title)
	}

	// Get the last inserted ID and assign to the newly
	// created post
	id, err := c.LastInsertId()
	if err != nil {
		log.Error(err)
		return domain.Post{}, fmt.Errorf("Could not get the newly created post ID with the title: %v", p.Title)
	}
	p.Id = int(id)

	// Convert the PostCreate type to type of Post to be returned
	// to the controller, used for binding & validation.
	convertedPost := s.convertToPost(*p)
	if err := s.seoMetaModel.UpdateCreate(&convertedPost); err != nil {
		return domain.Post{}, err
	}

	return convertedPost, nil
}

// Update post
func (s *PostStore) Update(p *domain.PostCreate) (domain.Post, error) {
	_, err := s.GetById(p.Id)
	if err != nil {
		log.Info(err)
		return domain.Post{}, err
	}

	// Check if the author is set assign to owner if not.
	p.Id = s.checkOwner(*p)

	// Update the posts table with data
	q := "UPDATE posts SET slug = ?, title = ?, status = ?, resource = ?, page_template = ?, fields = ?, codeinjection_head = ?, codeinjection_foot = ?, user_id = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(q, p.Slug, p.Title, p.Status, p.Resource, p.PageTemplate, p.Fields, p.CodeInjectHead, p.CodeInjectFoot, p.Author, p.Id)
	if err != nil {
		log.Error(err)
		return domain.Post{}, fmt.Errorf("Could not update the post wuth the title: %v", p.Title)
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
func (s *PostStore) Delete(id int) error {
	_, err := s.GetById(id)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("DELETE FROM posts WHERE id = ?", id); err != nil {
		log.Info(err)
		return fmt.Errorf("Could not delete post with the ID: %v", id)
	}
	return nil
}

// Get the total number of posts
func (s *PostStore) Total() (int, error) {
	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM posts").Scan(&total); err != nil {
		log.Error(err)
		return -1, fmt.Errorf("Could not get the total number of posts")
	}
	return total, nil
}

// Check if a post exists by slug
func (s *PostStore) Exists(slug string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM posts WHERE slug = ?)", slug).Scan(&exists)
	return exists
}

// Convert post create to post
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

// Check if the author is set or if the author does not exist.
// Return the owner ID under circumstances.
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