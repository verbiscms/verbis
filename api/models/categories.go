// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/google/uuid"
	"github.com/gookit/color"
	"strconv"
)

// CategoryRepository defines methods for Categories to interact with the database
type CategoryRepository interface {
	Get(meta params.Params) (domain.Categories, int, error)
	GetByID(id int) (domain.Category, error)
	GetByPost(pageID int) (*domain.Category, error)
	GetBySlug(slug string) (domain.Category, error)
	GetByName(name string) (domain.Category, error)
	GetParent(id int) (domain.Category, error)
	Create(c *domain.Category) (domain.Category, error)
	Update(c *domain.Category) (domain.Category, error)
	Delete(id int) error
	Exists(id int) bool
	// ExistsByName(name string) bool
	// ExistsBySlug(slug string) bool
	InsertPostCategory(postID int, categoryID *int) error
	DeletePostCategories(id int) error
	// Total() (int, error)
}

// CategoryStore defines the data layer for Categories
type CategoryStore struct {
	*StoreCfgOld
}

// newCategories - Construct
func newCategories(cfg *StoreCfgOld) *CategoryStore {
	return &CategoryStore{
		StoreCfgOld: cfg,
	}
}

// Get all categories
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no categories available.
func (s *CategoryStore) Get(meta params.Params) (domain.Categories, int, error) {
	const op = "CategoryRepository.Get"

	var c domain.Categories
	q := "SELECT * FROM categories"
	countQ := "SELECT COUNT(*) FROM categories"

	// Apply filters to total and original query
	filter, err := filterRows(s.DB, meta.Filters, "categories")
	if err != nil {
		return nil, -1, err
	}
	q += filter
	countQ += filter

	// Apply order
	q += fmt.Sprintf(" ORDER BY categories.%s %s", meta.OrderBy, meta.OrderDirection)

	// Apply pagination
	if !meta.LimitAll {
		q += fmt.Sprintf(" LIMIT %v OFFSET %v", meta.Limit, (meta.Page-1)*meta.Limit)
	}

	color.Green.Println(q)

	// Select categories
	if err := s.DB.Select(&c, q); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get categories", Operation: op, Err: err}
	}

	// Return not found error if no posts are available
	if len(c) == 0 {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No categories available", Operation: op}
	}

	// Count the total number of media
	var total int
	if err := s.DB.QueryRow(countQ).Scan(&total); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of category items", Operation: op, Err: err}
	}

	return c, total, nil
}

// Get the category by Id
// Returns errors.NOTFOUND if the category was not found by the given Id.
func (s *CategoryStore) GetByID(id int) (domain.Category, error) {
	const op = "CategoryRepository.GetByID"
	var c domain.Category
	if err := s.DB.Get(&c, "SELECT * FROM categories WHERE id = ?", id); err != nil {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get category with the ID: %d", id), Operation: op, Err: err}
	}
	return c, nil
}

// Get the category by post
// Returns errors.NOTFOUND if the category was not found by the given Post Id.
func (s *CategoryStore) GetByPost(postID int) (*domain.Category, error) {
	const op = "CategoryRepository.GetByPost"
	var c domain.Category
	if err := s.DB.Get(&c, "SELECT * FROM categories c WHERE EXISTS (SELECT post_id FROM post_categories p WHERE p.post_id = ? AND c.id = p.category_id) LIMIT 1", postID); err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get category with the post ID: %d", postID), Operation: op, Err: err}
	}
	return &c, nil
}

// Get the category by slug
// Returns errors.NOTFOUND if the category was not found by the given slug.
func (s *CategoryStore) GetBySlug(slug string) (domain.Category, error) {
	const op = "CategoryRepository.GetBySlug"
	var c domain.Category
	if err := s.DB.Get(&c, "SELECT * FROM categories WHERE slug = ?", slug); err != nil {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get category with the slug: %v", slug), Operation: op, Err: err}
	}
	return c, nil
}

// Get the category by slug
// Returns errors.NOTFOUND if the category was not found by the given slug.
func (s *CategoryStore) GetByName(name string) (domain.Category, error) {
	const op = "CategoryRepository.GetByName"
	var c domain.Category
	if err := s.DB.Get(&c, "SELECT * FROM categories WHERE name = ?", name); err != nil {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get category with the name: %v", name), Operation: op, Err: err}
	}
	return c, nil
}

// Get the parent category by ID
// Returns errors.NOTFOUND if the category was not found by the given slug.
func (s *CategoryStore) GetParent(id int) (domain.Category, error) {
	const op = "CategoryRepository.GetByParent"
	var c domain.Category
	if err := s.DB.Get(&c, "SELECT * FROM categories WHERE parent_id = ?", id); err != nil {
		return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get category with the parent ID: %d", id), Operation: op, Err: err}
	}
	return c, nil
}

// Create a new category
// Returns errors.CONFLICT if the the category (name) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *CategoryStore) Create(c *domain.Category) (domain.Category, error) {
	const op = "CategoryRepository.Create"

	if s.ExistsByName(c.Name) {
		return domain.Category{}, &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the post, the name %v, already exists", c.Name), Operation: op, Err: fmt.Errorf("name already exists")}
	}

	q := "INSERT INTO categories (uuid, slug, name, description, parent_id, resource, archive_id, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	e, err := s.DB.Exec(q, uuid.New().String(), c.Slug, c.Name, c.Description, c.ParentId, c.Resource, c.ArchiveId)
	if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the category with the name: %v", c.Name), Operation: op, Err: err}
	}

	id, err := e.LastInsertId()
	if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created category ID with the name: %v", c.Name), Operation: op, Err: err}
	}

	if c.ArchiveId != nil {
		err := s.changeArchivePostSlug(*c.ArchiveId, c.Slug, c.Resource)
		if err != nil {
			return domain.Category{}, err
		}
	}

	nc, err := s.GetByID(int(id))
	if err != nil {
		return domain.Category{}, err
	}

	return nc, nil
}

// Update category
// Returns errors.NOTFOUND if the category was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *CategoryStore) Update(c *domain.Category) (domain.Category, error) {
	const op = "CategoryRepository.Update"

	oldCategory, err := s.GetByID(c.Id)
	if err != nil {
		return domain.Category{}, err
	}

	q := "UPDATE categories SET slug = ?, name = ?, description = ?, resource = ?, parent_id = ?, archive_id = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.DB.Exec(q, c.Slug, c.Name, c.Description, c.Resource, c.ParentId, c.ArchiveId, c.Id)
	if err != nil {
		return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the category with the name: %s", c.Name), Operation: op, Err: err}
	}

	if oldCategory.ArchiveId != c.ArchiveId {
		s.resolveNewPostSlug(*oldCategory.ArchiveId, c.Resource)
	}

	if oldCategory.Slug != c.Slug {
		var posts []domain.Post
		if err := s.DB.Select(&posts, "SELECT * FROM posts WHERE slug LIKE '%"+oldCategory.Slug+"%'"); err != nil {
			return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: "Could not get categories", Operation: op, Err: err}
		}
	}

	if c.ArchiveId != nil {
		err := s.changeArchivePostSlug(*c.ArchiveId, c.Slug, c.Resource)
		if err != nil {
			return domain.Category{}, err
		}
	}

	return *c, nil
}

// Delete category from categories and post categories table
// Returns errors.NOTFOUND if the category was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *CategoryStore) Delete(id int) error {
	const op = "CategoryRepository.Delete"

	c, err := s.GetByID(id)
	if err != nil {
		return err
	}

	if _, err := s.DB.Exec("DELETE FROM categories WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the category with the ID: %v", id), Operation: op, Err: err}
	}

	if _, err := s.DB.Exec("DELETE FROM post_categories WHERE category_id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete post category with the ID: %v", id), Operation: op, Err: err}
	}

	if c.ArchiveId != nil {
		s.resolveNewPostSlug(*c.ArchiveId, c.Resource)
	}

	return nil
}

// Exists Checks if a category exists by the given Id
func (s *CategoryStore) Exists(id int) bool {
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT id FROM categories WHERE id = ?)", id).Scan(&exists)
	return exists
}

// Exists Checks if a category exists by the given name
func (s *CategoryStore) ExistsByName(name string) bool {
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT name FROM categories WHERE name = ?)", name).Scan(&exists)
	return exists
}

// Exists Checks if a category exists by the given slug
func (s *CategoryStore) ExistsBySlug(slug string) bool {
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT name FROM categories WHERE slug = ?)", slug).Scan(&exists)
	return exists
}

// InsertPostCategories - Insert into post categories with array of ID's.
// This function deletes all categories from the pivot before
// inserting again.
func (s *CategoryStore) InsertPostCategory(postID int, categoryID *int) error {
	const op = "CategoryRepository.InsertPostCategories"

	if _, err := s.DB.Exec("DELETE FROM post_categories WHERE post_id = ?", postID); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete from the post categories table with the ID: %v", postID), Operation: op, Err: err}
	}

	if categoryID != nil {
		q := "INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)"
		_, err := s.DB.Exec(q, postID, categoryID)
		if err != nil {
			return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not insert into the post categories table with the ID: %v", postID), Operation: op, Err: err}
		}
	}

	return nil
}

// Delete from the post categories table
func (s *CategoryStore) DeletePostCategories(id int) error {
	const op = "CategoryRepository.DeletePostCategories"
	if _, err := s.DB.Exec("DELETE FROM post_categories WHERE category_id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the post categories with the post ID of: %d", id), Operation: op, Err: err}
	}
	return nil
}

// changeArchivePostSlug changes the archive post slug when updating.
// Returns errors.INTERNAL if the SQL query was invalid or the new slug exists
func (s *CategoryStore) changeArchivePostSlug(id int, slug, resource string) error {
	const op = "CategoryRepository.ChangeArchivePostSlug"
	newSlug := ""
	if resource != "pages" {
		newSlug += "/" + resource
	}
	newSlug += "/" + slug
	if _, err := s.DB.Exec("UPDATE posts SET slug = ? WHERE id = ?", newSlug, id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the posts table with the new slug: %s", slug), Operation: op, Err: err}
	}
	return nil
}

// resolveNewPostSlug adds untitled to the new slug if it already exists.
func (s *CategoryStore) resolveNewPostSlug(id int, resource string) {
	slug := "untitled"
	counter := 1
	for {
		err := s.changeArchivePostSlug(id, slug, resource)
		if err != nil {
			slug = "untitled-" + strconv.Itoa(counter)
			counter++
			continue
		}
		break
	}
}
