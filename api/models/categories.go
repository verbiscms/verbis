package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// CategoryRepository defines methods for Categories to interact with the database
type CategoryRepository interface {
	Get(meta http.Params) ([]domain.Category, error)
	GetById(id int) (domain.Category, error)
	GetByPost(pageId int) ([]domain.Category, error)
	Create(c *domain.Category) (domain.Category, error)
	Update(c *domain.Category) error
	InsertPostCategories(postId int, ids []int) error
	DeletePostCategories(id int) error
	Delete(id int) error
	Exists(id int) bool
	ExistsByName(name string) bool
	Total() (int, error)
}

// CategoryStore defines the data layer for Categories
type CategoryStore struct {
	db *sqlx.DB
}

// newCategories - Construct
func newCategories(db *sqlx.DB) *CategoryStore {
	return &CategoryStore{
		db: db,
	}
}

// Get all categories
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no categories available.
func (s *CategoryStore) Get(meta http.Params) ([]domain.Category, error) {
	const op = "CategoryRepository.Get"

	var c []domain.Category
	q := fmt.Sprintf("SELECT * FROM categories ORDER BY categories.%s %s LIMIT %v OFFSET %v", meta.OrderBy, meta.OrderDirection, meta.Limit, meta.Page * meta.Limit)
	if err := s.db.Select(&c, q); err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not get categories", Operation: op, Err: err}
	}

	if len(c) == 0 {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No categories available", Operation: op}
	}

	return c, nil
}

// Get the category by Id
// Returns errors.NOTFOUND if the category was not found by the given Id.
func (s *CategoryStore) GetById(id int) (domain.Category, error) {
	const op = "CategoryRepository.GetById"
	var c domain.Category
	if err := s.db.Get(&c, "SELECT * FROM categories WHERE id = ?", id); err != nil {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get category with the ID: %d", id), Operation: op}
	}
	return c, nil
}

// Get the category by post
// Returns errors.NOTFOUND if the category was not found by the given Post Id.
func (s *CategoryStore) GetByPost(postId int) ([]domain.Category, error) {
	const op = "CategoryRepository.GetByPost"
	var c []domain.Category
	if err := s.db.Select(&c, "SELECT * FROM categories c WHERE EXISTS (SELECT post_id FROM post_categories p WHERE p.post_id = ? AND c.id = p.category_id)", postId); err != nil {
		return []domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get category with the post ID: %d", postId), Operation: op}
	}
	return c, nil
}

// Create a new category
// Returns errors.CONFLICT if the the category (name) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *CategoryStore) Create(c *domain.Category) (domain.Category, error) {
	const op = "CategoryRepository.Create"

	if s.ExistsByName(c.Name) {
		return domain.Category{}, &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the post, the name %v, already exists", c.Name), Operation: op}
	}

	q := "INSERT INTO categories (uuid, slug, name, description, hidden, parent_id, page_template, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	e, err := s.db.Exec(q, uuid.New().String(), c.Slug, c.Name, c.Description, c.Hidden, c.ParentId, c.PageTemplate)
	if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the categort with the name: %v", c.Name), Operation: op, Err: err}
	}

	id, err := e.LastInsertId()
	if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created category ID with the name: %v", c.Name), Operation: op, Err: err}
	}

	nc, err := s.GetById(int(id))
	if err != nil {
		return domain.Category{}, err
	}

	return nc, nil
}

// Update category
// Returns errors.NOTFOUND if the category was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *CategoryStore) Update(c *domain.Category) error {
	const op = "CategoryRepository.Update"

	_, err := s.GetById(c.Id)
	if err != nil {
		return err
	}

	q := "UPDATE categories SET slug = ?, name = ? description = ?, hidden = ?, parent_id = ?, page_template = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(q, c.Slug, c.Name, c.Description, c.Hidden, c.ParentId, c.PageTemplate)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the category with the name: %s", c.Name), Operation: op, Err: err}
	}

	return nil
}

// Delete category from categories and post categories table
// Returns errors.NOTFOUND if the category was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *CategoryStore) Delete(id int) error {
	const op = "CategoryRepository.Delete"

	_, err := s.GetById(id)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("DELETE FROM categories WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the category with the ID: %v", id), Operation: op, Err: err}
	}

	if _, err := s.db.Exec("DELETE FROM post_categories WHERE category_id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete post category with the ID: %v", id), Operation: op, Err: err}
	}

	return nil
}

// Exists Checks if a category exists by the given Id
func (s *CategoryStore) Exists(id int) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM categories WHERE id = ?)", id).Scan(&exists)
	return exists
}

// Exists Checks if a category exists by the given name
func (s *CategoryStore) ExistsByName(name string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT name FROM categories WHERE name = ?)", name).Scan(&exists)
	return exists
}

// TODO, come back too! Maybe this should be in posts?!

// Insert into post categories with array of ID's
func (s *CategoryStore) InsertPostCategories(postId int, ids []int) error {
	const op = "CategoryRepository.InsertPostCategories"

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
	const op = "CategoryRepository.DeletePostCategories"
	if _, err := s.db.Exec("DELETE FROM post_categories WHERE category_id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the post categories with the ID of: %d", id), Operation: op, Err: err}
	}
	return nil
}

// Get the total number of categories
func (s *CategoryStore) Total() (int, error) {
	const op = "CategoryRepository.Total"
	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&total); err != nil {
		return -1, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the total number of categories"), Operation: op, Err: err}
	}
	return total, nil
}