package models

import (
	"cms/api/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type RoleRepository interface {
	GetAll() ([]domain.UserRole, error)
	GetById(id int) (domain.UserRole, error)
	Create(r *domain.UserRole) (domain.UserRole, error)
	Update(r *domain.UserRole) (domain.UserRole, error)
	Delete(id int) error
	Exists(name string) bool
}

type RoleStore struct {
	db *sqlx.DB
}

//Construct
func newRoles(db *sqlx.DB) *RoleStore {
	return &RoleStore{
		db: db,
	}
}

// Get all roles
func (s *RoleStore) GetAll() ([]domain.UserRole, error) {
	var r []domain.UserRole
	if err := s.db.Select(&r, "SELECT * FROM roles"); err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Could not get roles")
	}
	if len(r) == 0 {
		return []domain.UserRole{}, nil
	}
	return r, nil
}

// Get the role by ID
func (s *RoleStore) GetById(id int) (domain.UserRole, error) {
	var r domain.UserRole
	if err := s.db.Get(&r, "SELECT * FROM roles WHERE id = ? LIMIT 1", id); err != nil {
		log.Info(err)
		return domain.UserRole{}, fmt.Errorf("Could not get role with the ID: %v", id)
	}
	return r, nil
}

// Create role
func (s *RoleStore) Create(r *domain.UserRole) (domain.UserRole, error) {
	q := "INSERT INTO roles (name, description) VALUES (?, ?)"
	c, err := s.db.Exec(q, r.Name, r.Description)
	if err != nil {
		log.Error(err)
		return domain.UserRole{}, fmt.Errorf("Could not create the role: %v", r.Name)
	}

	id, err := c.LastInsertId()
	if err != nil {
		log.Error(err)
		return domain.UserRole{}, fmt.Errorf("Could not get the newly created role ID: %v", r.Name)
	}
	r.Id = int(id)


	return *r, nil
}

// Update role
func (s *RoleStore) Update(r *domain.UserRole) (domain.UserRole, error) {

	_, err := s.GetById(r.Id)
	if err != nil {
		log.Info(err)
		return domain.UserRole{}, err
	}

	q := "UPDATE roles SET name = ?, description = ? WHERE id = ?"
	_, err = s.db.Exec(q, r.Name, r.Description)
	if err != nil {
		log.Error(err)
		return domain.UserRole{}, fmt.Errorf("Could not update the post: %v", r.Name)
	}

	return *r, nil
}

// Delete role
func (s *RoleStore) Delete(id int) error {
	_, err := s.GetById(id)

	if err != nil {
		log.Info(err)
		return err
	}

	if _, err := s.db.Exec("DELETE FROM roles WHERE id = ?", id); err != nil {
		log.Error(err)
		return fmt.Errorf("Could not delete role with the ID of: %v", id)
	}

	return nil
}

// Check if the role record exists
func (s *RoleStore) Exists(name string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM roles WHERE name = ?)", name).Scan(&exists)
	return exists
}


