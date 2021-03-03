package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
)

// RedirectRepository defines methods for Redirects to interact with the database
type RedirectRepository interface {
	Get(meta params.Params) (domain.Redirects, int, error)
	GetByID(id int64) (domain.Redirect, error)
	GetByFrom(from string) (domain.Redirect, error)
	Create(r *domain.Redirect) (domain.Redirect, error)
	Update(r *domain.Redirect) (domain.Redirect, error)
	Delete(id int64) error
	Exists(id int64) bool
	ExistsByFromPath(from string) bool
}

// RedirectStore defines the data layer for Categories
type RedirectStore struct {
	*StoreConfig
}

// newRedirects - Construct
func newRedirects(cfg *StoreConfig) *RedirectStore {
	return &RedirectStore{
		StoreConfig: cfg,
	}
}

// Get all categories
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no categories available.
func (s *RedirectStore) Get(meta params.Params) (domain.Redirects, int, error) {
	const op = "RedirectStore.Get"

	var r domain.Redirects
	q := "SELECT * FROM redirects"
	countQ := "SELECT COUNT(*) FROM redirects"

	// Apply filters to total and original query
	filter, err := filterRows(s.DB, meta.Filters, "redirects")
	if err != nil {
		return nil, -1, err
	}
	q += filter
	countQ += filter

	// Apply order
	q += fmt.Sprintf(" ORDER BY redirects.%s %s", meta.OrderBy, meta.OrderDirection)

	// Apply pagination
	if !meta.LimitAll {
		q += fmt.Sprintf(" LIMIT %v OFFSET %v", meta.Limit, (meta.Page-1)*meta.Limit)
	}

	// Select redirects
	if err := s.DB.Select(&r, q); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get redirects", Operation: op, Err: err}
	}

	// Return not found error if no redirects are available
	if len(r) == 0 {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No redirects available", Operation: op}
	}

	// Count the total number of redirects
	var total int
	if err := s.DB.QueryRow(countQ).Scan(&total); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of redirect items", Operation: op, Err: err}
	}

	return r, total, nil
}

// GetByID the redirect by ID.
// Returns errors.NOTFOUND if the redirect was not found by the given from path.
func (s *RedirectStore) GetByID(id int64) (domain.Redirect, error) {
	const op = "RedirectStore.GetByPost"
	var r domain.Redirect
	if err := s.DB.Get(&r, "SELECT * FROM redirects WHERE id = ?", id); err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get redirect with the ID: %d", id), Operation: op, Err: err}
	}
	return r, nil
}

// Get the redirect by from path
// Returns errors.NOTFOUND if the redirect was not found by the given from path.
func (s *RedirectStore) GetByFrom(from string) (domain.Redirect, error) {
	const op = "RedirectStore.GetByPost"
	var r domain.Redirect
	if err := s.DB.Get(&r, "SELECT * FROM redirects WHERE from_path = ?", from); err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get redirect with the path: %s", from), Operation: op, Err: err}
	}
	return r, nil
}

// Create a new redirect
// Returns errors.CONFLICT if the the redirect (from path) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *RedirectStore) Create(r *domain.Redirect) (domain.Redirect, error) {
	const op = "RedirectStore.Create"

	if s.ExistsByFromPath(r.From) {
		return domain.Redirect{}, &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the redirect, the from path %v, already exists", r.From), Operation: op, Err: fmt.Errorf("name already exists")}
	}

	q := "INSERT INTO redirects (from_path, to_path, code, updated_at, created_at) VALUES (?, ?, ?, NOW(), NOW())"
	e, err := s.DB.Exec(q, r.From, r.To, r.Code)
	if err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the redirect with the from path: %v", r.From), Operation: op, Err: err}
	}

	id, err := e.LastInsertId()
	if err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created redirect ID with the from path: %v", r.From), Operation: op, Err: err}
	}
	r.Id = id

	return *r, nil
}

// Update redirect
// Returns errors.NOTFOUND if the redirect was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *RedirectStore) Update(r *domain.Redirect) (domain.Redirect, error) {
	const op = "RedirectStore.Update"

	if !s.Exists(r.Id) {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("No redirect exists with the Id: %d", r.Id), Operation: op, Err: fmt.Errorf("no redirect exists")}
	}

	q := "UPDATE redirects SET from_path = ?, to_path = ?, code = ?, updated_at = NOW() WHERE id = ?"
	_, err := s.DB.Exec(q, r.From, r.To, r.Code, r.Id)
	if err != nil {
		return domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the redirect with the from path: %s", r.From), Operation: op, Err: err}
	}

	return *r, nil
}

// Delete redirect
// Returns errors.NOTFOUND if the redirect was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *RedirectStore) Delete(id int64) error {
	const op = "RedirectStore.Delete"

	if !s.Exists(id) {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("No redirect exists with the Id: %d", id), Operation: op, Err: fmt.Errorf("no redirect exists with the id: %d", id)}
	}

	if _, err := s.DB.Exec("DELETE FROM redirects WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the redirect with the ID: %v", id), Operation: op, Err: err}
	}

	return nil
}

// Exists Checks if a redirect exists by the given Id
func (s *RedirectStore) Exists(id int64) bool {
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT id FROM redirects WHERE id = ?)", id).Scan(&exists)
	return exists
}

// Exists Checks if a redirect exists by the given from path
func (s *RedirectStore) ExistsByFromPath(from string) bool {
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT from_path FROM redirects WHERE from_path = ?)", from).Scan(&exists)
	return exists
}
