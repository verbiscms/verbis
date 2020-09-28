package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// RoleRepository defines methods for Posts to interact with the database
type RoleRepository interface {
	Get() ([]domain.UserRole, error)
	GetById(id int) (domain.UserRole, error)
	Create(r *domain.UserRole) (domain.UserRole, error)
	Update(r *domain.UserRole) (domain.UserRole, error)
	Delete(id int) error
	Exists(name string) bool
}

// PostStore defines the data layer for Posts
type RoleStore struct {
	db *sqlx.DB
}

// newRoles - Construct
func newRoles(db *sqlx.DB) *RoleStore {
	return &RoleStore{
		db: db,
	}
}

// Get all roles
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no roles available.
func (s *RoleStore) Get() ([]domain.UserRole, error) {
	const op = "RoleRepository.Get"

	var r []domain.UserRole
	if err := s.db.Select(&r, "SELECT * FROM roles"); err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not get roles", Operation: op, Err: err}
	}

	if len(r) == 0 {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No roles available", Operation: op}
	}

	return r, nil
}

// Get the role by ID
// Returns errors.NOTFOUND if the post was not found by the given Id.
func (s *RoleStore) GetById(id int) (domain.UserRole, error) {
	const op = "RoleRepository.GetById"
	var r domain.UserRole
	if err := s.db.Get(&r, "SELECT * FROM roles WHERE id = ? LIMIT 1", id); err != nil {
		log.Info(err)
		return domain.UserRole{}, fmt.Errorf("Could not get role with the ID: %v", id)
	}
	return r, nil
}

// Create a new role
// Returns errors.CONFLICT if the the post slug already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function
// could not get the newly created ID.
func (s *RoleStore) Create(r *domain.UserRole) (domain.UserRole, error) {
	const op = "RoleRepository.Create"

	if s.Exists(r.Name) {
		return domain.UserRole{}, &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the role, the name %v, already exists", r.Name), Operation: op}
	}

	q := "INSERT INTO roles (id, name, description) VALUES (?, ?, ?)"
	c, err := s.db.Exec(q, r.Id, r.Name, r.Description)
	if err != nil {
		return domain.UserRole{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the role with the name: %v", r.Name), Operation: op, Err: err}
	}

	id, err := c.LastInsertId()
	if err != nil {
		log.Error(err)
		return domain.UserRole{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created post Role ID with the name: %v", r.Name), Operation: op, Err: err}
	}
	r.Id = int(id)

	return *r, nil
}

// Update role
// Returns errors.NOTFOUND if the role was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *RoleStore) Update(r *domain.UserRole) (domain.UserRole, error) {
	const op = "RoleRepository.Update"

	_, err := s.GetById(r.Id)
	if err != nil {
		return domain.UserRole{}, err
	}

	q := "UPDATE roles SET name = ?, description = ? WHERE id = ?"
	_, err = s.db.Exec(q, r.Name, r.Description)
	if err != nil {
		return domain.UserRole{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the role with the name: %v", r.Name), Operation: op, Err: err}
	}

	return *r, nil
}

// Delete role
// Returns errors.NOTFOUND if the role was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *RoleStore) Delete(id int) error {
	const op = "RoleRepository.Delete"

	_, err := s.GetById(id)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("DELETE FROM roles WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the role with the ID: %v", id), Operation: op, Err: err}
	}

	return nil
}

// Exists Checks if a role exists by the given name
func (s *RoleStore) Exists(name string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM roles WHERE name = ?)", name).Scan(&exists)
	return exists
}


