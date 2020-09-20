package models

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository interface {
	Get(meta http.Params) ([]domain.Category, error)
	GetById(id int) (domain.Category, error)
	GetByPost(pageId int) ([]domain.Category, error)
	Create(c *domain.Category) (int, error)
	Update(c *domain.Category) error
	InsertPostCategories(postId int, ids []int) error
	DeletePostCategories(id int) error
	Delete(id int) error
	Exists(id int) bool
	Total() (int, error)
}

type CategoryStore struct {
	db *sqlx.DB
}

//Construct
func newCategories(db *sqlx.DB) *CategoryStore {
	return &CategoryStore{
		db: db,
	}
}

// Get all categories
func (s *CategoryStore) Get(meta http.Params) ([]domain.Category, error) {
	var c []domain.Category
	q := fmt.Sprintf("SELECT * FROM categories ORDER BY categories.%s %s LIMIT %v OFFSET %v", meta.OrderBy, meta.OrderDirection, meta.Limit, meta.Page * meta.Limit)
	if err := s.db.Select(&c, q); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Could not get categories")
	}

	if len(c) == 0 {
		var test = make([]domain.Category, 1)
		return test, nil
	}

	return c, nil
}

// Get the category by ID
func (s *CategoryStore) GetById(id int) (domain.Category, error) {
	var c domain.Category
	if err := s.db.Get(&c, "SELECT * FROM categories WHERE id = ?", id); err != nil {
		log.Info(err)
		return domain.Category{}, fmt.Errorf("Could not get category with ID: %v", id)
	}
	return c, nil
}

// Get the category by post
func (s *CategoryStore) GetByPost(postId int) ([]domain.Category, error) {
	var c []domain.Category
	if err := s.db.Select(&c, "SELECT * FROM categories c WHERE EXISTS (SELECT post_id FROM post_categories p WHERE p.post_id = ? AND c.id = p.category_id)", postId); err != nil {
		log.Error(err)
		return []domain.Category{}, fmt.Errorf("Could not get categories with post ID: %v", postId)
	}
	return c, nil
}

// Create category
func (s *CategoryStore) Create(c *domain.Category) (int, error) {
	q := "INSERT INTO categories (uuid, slug, name, description, hidden, parent_id, page_template, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	e, err := s.db.Exec(q, uuid.New().String(), c.Slug, c.Name, c.Description, c.Hidden, c.ParentId, c.PageTemplate)

	if err != nil {
		log.Error(err)
		return 0, fmt.Errorf("Could not create the category with the name: %v", c.Name)
	}

	id, err := e.LastInsertId()
	if err != nil {
		log.Error(err)
		return 0, fmt.Errorf("Could not get the newly created category ID with the name: %v", c.Name)
	}

	return int(id), nil
}

// Update category
func (s *CategoryStore) Update(c *domain.Category) error {
	_, err := s.GetById(c.Id)
	if err != nil {
		log.Info(err)
		return err
	}

	q := "UPDATE categories SET slug = ?, name = ? description = ?, hidden = ?, parent_id = ?, page_template = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(q, c.Slug, c.Name, c.Description, c.Hidden, c.ParentId, c.PageTemplate)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not update the category with the name: %v", c.Name)
	}

	return nil
}

// Delete category
func (s *CategoryStore) Delete(id int) error {
	c, err := s.GetById(id)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("DELETE FROM categories WHERE id = ?", id); err != nil {
		log.Info(err)
		return fmt.Errorf("Could not delete category with the ID: %v", c.Name)
	}

	if _, err := s.db.Exec("DELETE FROM post_categories WHERE category_id = ?", id); err != nil {
		log.Error(err)
		return fmt.Errorf("Could not delete the category from the post categories table with the ID: %v", id)
	}

	return nil
}

// Insert into post categories with array of ID's
func (s *CategoryStore) InsertPostCategories(postId int, ids []int) error {
	for _, id := range ids {

		// Check if the record in post categories already exists
		pcExists := false
		_ = s.db.QueryRow("SELECT EXISTS (SELECT category_id FROM post_categories WHERE category_id = ? AND post_id = ?)", id, postId).Scan(&pcExists)

		if !pcExists {
			// Check if the category in the categories table exists before inserting
			if s.Exists(id) {
				q := "INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)"
				_, err := s.db.Exec(q, postId, id)
				if err != nil {
					return fmt.Errorf("Could not insert to the post categories table with the ID: %v", id)
				}
			}
		}

		//
		// Get all the post ids with the category
	}

	return nil
}

// Delete from the post categories table
func (s *CategoryStore) DeletePostCategories(id int) error {
	if _, err := s.db.Exec("DELETE FROM post_categories WHERE category_id = ?", id); err != nil {
		log.Info(err)
		return fmt.Errorf("Could not delete post category with the ID: %v", id)
	}
	return nil
}

// Check if a category exists by ID
func (s *CategoryStore) Exists(id int) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM categories WHERE id = ?)", id).Scan(&exists)
	return exists
}

// Get the total number of categories
func (s *CategoryStore) Total() (int, error) {
	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&total); err != nil {
		log.Error(err)
		return -1, fmt.Errorf("Could not get the total number of categories")
	}
	return total, nil
}