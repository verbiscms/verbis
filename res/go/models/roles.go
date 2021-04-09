// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

// RoleRepository defines methods for Posts to interact with the database
type RoleRepository interface {
	Get() ([]domain.Role, error)
	GetByID(id int) (domain.Role, error)
	Create(r *domain.Role) (domain.Role, error)
	Update(r *domain.Role) (domain.Role, error)
	Exists(name string) bool
}

// PostStore defines the data layer for Posts
type RoleStore struct {
	*StoreCfgOld
}

// newRoles - Construct
func newRoles(cfg *StoreCfgOld) *RoleStore {
	return &RoleStore{
		StoreCfgOld: cfg,
	}
}

// Get all roles
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no roles available.
func (s *RoleStore) Get() ([]domain.Role, error) {
	const op = "RoleRepository.Get"

	var r []domain.Role
	if err := s.DB.Select(&r, "SELECT * FROM roles"); err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not get roles", Operation: op, Err: err}
	}

	if len(r) == 0 {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "No roles available", Operation: op}
	}

	return r, nil
}

// Get the role by ID
// Returns errors.NOTFOUND if the post was not found by the given ID.
func (s *RoleStore) GetByID(id int) (domain.Role, error) {
	const op = "RoleRepository.GetByID"
	var r domain.Role
	if err := s.DB.Get(&r, "SELECT * FROM roles WHERE id = ? LIMIT 1", id); err != nil {
		return domain.Role{}, fmt.Errorf("could not get role with the ID: %v", id)
	}
	return r, nil
}

// Create a new role
// Returns errors.CONFLICT if the the post slug already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function
// could not get the newly created ID.
func (s *RoleStore) Create(r *domain.Role) (domain.Role, error) {
	const op = "RoleRepository.Create"

	if s.Exists(r.Name) {
		return domain.Role{}, &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the role, the name %v, already exists", r.Name), Operation: op}
	}

	q := "INSERT INTO roles (id, name, description) VALUES (?, ?, ?)"
	c, err := s.DB.Exec(q, r.Id, r.Name, r.Description)
	if err != nil {
		return domain.Role{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the role with the name: %v", r.Name), Operation: op, Err: err}
	}

	id, err := c.LastInsertId()
	if err != nil {
		return domain.Role{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created post Role ID with the name: %v", r.Name), Operation: op, Err: err}
	}
	r.Id = int(id)

	return *r, nil
}

// Update role
// Returns errors.NOTFOUND if the role was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *RoleStore) Update(r *domain.Role) (domain.Role, error) {
	const op = "RoleRepository.Update"

	_, err := s.GetByID(r.Id)
	if err != nil {
		return domain.Role{}, err
	}

	q := "UPDATE roles SET name = ?, description = ? WHERE id = ?"
	_, err = s.DB.Exec(q, r.Name, r.Description)
	if err != nil {
		return domain.Role{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the role with the name: %v", r.Name), Operation: op, Err: err}
	}

	return *r, nil
}

// Exists Checks if a role exists by the given name
func (s *RoleStore) Exists(name string) bool {
	var exists bool
	_ = s.DB.QueryRow("SELECT EXISTS (SELECT id FROM roles WHERE name = ?)", name).Scan(&exists)
	return exists
}
