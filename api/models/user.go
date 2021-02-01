package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

// UserRepository defines methods for Users to interact with the database
type UserRepository interface {
	Get(meta params.Params) (domain.Users, int, error)
	GetById(id int) (domain.User, error)
	GetOwner() (domain.User, error)
	GetByToken(token string) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetRoles() ([]domain.UserRole, error)
	Create(u *domain.UserCreate) (domain.User, error)
	Update(u *domain.User) (domain.User, error)
	Delete(id int) error
	CheckSession(token string) error
	ResetPassword(id int, reset domain.UserPasswordReset) error
	CheckToken(token string) (domain.User, error)
	Exists(id int) bool
	ExistsByEmail(email string) bool
	Total() (int, error)
}

// UserStore defines the data layer for Users
type UserStore struct {
	db          *sqlx.DB
	config      config.Configuration
	optionsRepo domain.Options
}

// newUser - Construct
func newUser(db *sqlx.DB, config config.Configuration) *UserStore {
	return &UserStore{
		db:          db,
		config:      config,
		optionsRepo: newOptions(db, config).GetStruct(),
	}
}

// Get all users
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no users available.
func (s *UserStore) Get(meta params.Params) (domain.Users, int, error) {
	const op = "UserRepository.Get"

	var u domain.Users
	q := fmt.Sprintf("SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id INNER JOIN roles ON user_roles.role_id = roles.id")
	countQ := fmt.Sprintf("SELECT COUNT(*) FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id INNER JOIN roles ON user_roles.role_id = roles.id")

	// Check if there is a role filter, for example
	// roles.name and reorder meta.Filters
	table := "users"
	for k, v := range meta.Filters {
		if strings.Contains(k, "roles") {
			arr := strings.Split(k, ".")
			if len(arr) > 1 {
				meta.Filters[arr[1]] = v
				delete(meta.Filters, k)
				table = "roles"
			}
		}
	}

	// Apply filters to total and original query
	filter, err := filterRows(s.db, meta.Filters, table)
	if err != nil {
		return nil, -1, err
	}
	q += filter
	countQ += filter

	// Apply order
	q += fmt.Sprintf(" ORDER BY users.%s %s", meta.OrderBy, meta.OrderDirection)

	// Apply pagination
	if !meta.LimitAll {
		q += fmt.Sprintf(" LIMIT %v OFFSET %v", meta.Limit, (meta.Page-1)*meta.Limit)
	}

	// Select users
	if err := s.db.Select(&u, q); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get users", Operation: op, Err: err}
	}

	// Return not found error if no users are available
	if len(u) == 0 {
		return nil, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No users available", Operation: op}
	}

	// Count the total number of users
	var total int
	if err := s.db.QueryRow(countQ).Scan(&total); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of posts", Operation: op, Err: err}
	}

	return u, total, nil
}

// GetById returns a user by Id
// Returns errors.NOTFOUND if the user was not found by the given Id.
func (s *UserStore) GetById(id int) (domain.User, error) {
	const op = "UserRepository.GetById"
	var u domain.User
	if err := s.db.Get(&u, "SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id INNER JOIN roles ON user_roles.role_id = roles.id WHERE users.id = ?", id); err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the user with the ID: %d", id), Operation: op, Err: err}
	}
	return u, nil
}

// GetOwner gets the owner of the site with the Id of 6
// Returns errors.NOTFOUND if the owner was not found.
func (s *UserStore) GetOwner() (domain.User, error) {
	const op = "UserRepository.GetOwner"
	var u domain.User
	if err := s.db.Get(&u, "SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id INNER JOIN roles ON user_roles.role_id = roles.id WHERE roles.id = 6 LIMIT 1"); err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "Could not get the owner of the site", Operation: op, Err: err}
	}
	return u, nil
}

// GetByToken gets a user with the given token.
// Returns errors.NOTFOUND if the owner was not found.
func (s *UserStore) GetByToken(token string) (domain.User, error) {
	const op = "UserRepository.GetOwner"
	var u domain.User
	if err := s.db.Get(&u, "SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id INNER JOIN roles ON user_roles.role_id = roles.id WHERE token = ? LIMIT 1", token); err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the user with the token: %s", token), Operation: op, Err: err}
	}
	return u, nil
}

// GetByToken gets a user with the given token.
// Returns errors.NOTFOUND if the owner was not found.
func (s *UserStore) GetByEmail(email string) (domain.User, error) {
	const op = "UserRepository.GetByEmail"
	var u domain.User
	if err := s.db.Get(&u, "SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id INNER JOIN roles ON user_roles.role_id = roles.id WHERE email = ? LIMIT 1", email); err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the user with the email: %s", email), Operation: op, Err: err}
	}
	return u, nil
}

// GetRoles gets all of the roles in the roles table
// Returns errors.INTERNAL if the roles table was inaccessible.
func (s *UserStore) GetRoles() ([]domain.UserRole, error) {
	const op = "UserRepository.GetRoles"
	var r []domain.UserRole
	if err := s.db.Select(&r, "SELECT * FROM roles"); err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the user roles", Operation: op, Err: err}
	}
	return r, nil
}

// Create user
// Returns errors.CONFLICT if the the post slug already exists.
// Returns errors.INTERNAL if the SQL query was invalid, the function
// could not get the newly created ID or the user role failed to be inserted.
func (s *UserStore) Create(u *domain.UserCreate) (domain.User, error) {
	const op = "UserRepository.Create"

	if exists := s.ExistsByEmail(u.Email); exists {
		return domain.User{}, &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("Could not create the user, the email %v, already exists", u.Email), Operation: op, Err: fmt.Errorf("user already exists")}
	}

	hashedPassword, err := encryption.HashPassword(u.Password)
	if err != nil {
		return domain.User{}, err
	}

	token := encryption.GenerateUserToken(u.FirstName+u.LastName, u.Email)

	userQ := "INSERT INTO users (uuid, first_name, last_name, email, password, website, facebook, twitter, linked_in, instagram, biography, profile_picture_id, token, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"
	c, err := s.db.Exec(userQ, uuid.New().String(), u.FirstName, u.LastName, u.Email, hashedPassword, u.Website, u.Facebook, u.Twitter, u.Linkedin, u.Instagram, u.Biography, u.ProfilePictureID, token)
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the user with the email: %s", u.Email), Operation: op, Err: err}
	}

	id, err := c.LastInsertId()
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created user ID with the email: %v", u.Email), Operation: op, Err: err}
	}

	roleQ := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)"
	_, err = s.db.Exec(roleQ, id, u.Role.Id)
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the user role for user with the email: %s", u.Email), Operation: op, Err: err}
	}

	newUser, err := s.GetById(int(id))
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created user with the id: %v", u.Id), Operation: op, Err: err}
	}

	// TODO: Determine of email verified is turned on.
	//If the user is not the owner, send the verification email
	//if u.Role.Id != 6 {
	//	ve, err := events.NewVerifyEmail()
	//	if err != nil {
	//		return domain.User{}, err
	//	}
	//
	//	err = ve.Send(&newUser, s.optionsRepo.SiteTitle)
	//	if err != nil {
	//		return domain.User{}, err
	//	}
	//}

	return newUser, nil
}

// Update user
// Returns errors.NOTFOUND if the user was not found.
// Returns errors.INTERNAL if the SQL query was invalid for updating the user
// or user roles table.
func (s *UserStore) Update(u *domain.User) (domain.User, error) {
	const op = "UserRepository.Update"

	_, err := s.GetById(u.Id)
	if err != nil {
		return domain.User{}, err
	}

	userQ := "UPDATE users SET first_name = ?, last_name = ?, email = ?, website = ?, facebook = ?, twitter = ?, linked_in = ?, instagram = ?, biography = ?, profile_picture_id = ?, updated_at = NOW() WHERE id = ?"
	_, err = s.db.Exec(userQ, u.FirstName, u.LastName, u.Email, u.Website, u.Facebook, u.Twitter, u.Linkedin, u.Instagram, u.Biography, u.ProfilePictureID, u.Id)
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the user with the email: %s", u.Email), Operation: op, Err: err}
	}

	roleQ := "UPDATE user_roles SET role_id = ? WHERE user_id = ?"
	_, err = s.db.Exec(roleQ, u.Role.Id, u.Id)
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the user roles with the user ID: %d", u.Id), Operation: op, Err: err}
	}

	return *u, nil
}

// Delete user
// Returns errors.NOTFOUND if user wanst found.
// Returns errors.CONFLICT if role ID is thE owner
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *UserStore) Delete(id int) error {
	const op = "UserRepository.Delete"

	u, err := s.GetById(id)
	if err != nil {
		return err
	}

	if u.Role.Name == "Owner" {
		return &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("The owner of the site cannot be deleted."), Operation: op, Err: err}
	}

	if _, err := s.db.Exec("DELETE FROM users WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete user with the ID: %d", id), Operation: op, Err: err}
	}

	if _, err := s.db.Exec("DELETE FROM user_roles WHERE user_id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not from the user roles with the ID: %d", id), Operation: op, Err: err}
	}

	return nil
}

// CheckSession
// Returns errors.NOTFOUND if the user was not found
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.CONFLICT if the user session expired.
func (s *UserStore) CheckSession(token string) error {
	const op = "UserRepository.UpdateSession"

	u, err := s.GetByToken(token)
	if err != nil {
		return err
	}

	// If not login
	if u.TokenLastUsed != nil {

		// Destroy the token and create a new one if session expired.
		inactiveFor := time.Now().Sub(*u.TokenLastUsed).Minutes()

		if int(inactiveFor) > s.config.Admin.InactiveSessionTime {
			newToken := encryption.GenerateUserToken(u.FirstName+u.LastName, u.Email)

			_, err := s.db.Exec("UPDATE users SET token = ?, updated_at = NOW() WHERE token = token", newToken)
			if err != nil {
				fmt.Println(err)
				return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the user's token with the name: %v", u.FirstName+" "+u.LastName), Operation: op, Err: err}
			}

			return &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("User seesion expiered, please login again."), Operation: op, Err: fmt.Errorf("user session expiered")}
		}
	}

	_, err = s.db.Exec("UPDATE users SET token_last_used = NOW() WHERE token = ?", token)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the user token last used."), Operation: op, Err: err}
	}

	return nil
}

// ResetPassword
// Returns errors.INVALID if the current password didn't match.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *UserStore) ResetPassword(id int, reset domain.UserPasswordReset) error {
	const op = "UserRepository.ResetPassword"

	hashedPassword, err := encryption.HashPassword(reset.NewPassword)
	if err != nil {
		return err
	}

	_, err = s.db.Exec("UPDATE users SET password = ? WHERE id = ?", hashedPassword, id)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not update the users table with the new password", Operation: op, Err: err}
	}

	return nil
}

// Get the user by Token
// Returns errors.NOTFOUND if there are the user is not found.
func (s *UserStore) CheckToken(token string) (domain.User, error) {
	const op = "UserRepository.CheckToken"
	var u domain.User
	if err := s.db.Get(&u, "SELECT users.*, roles.id 'roles.id', roles.name 'roles.name', roles.description 'roles.description' FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id LEFT JOIN roles ON user_roles.role_id = roles.id WHERE users.token = ?", token); err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get user with token: %v", token), Operation: op, Err: err}
	}
	return u, nil
}

// Exists checks if the user record exists by ID
func (s *UserStore) Exists(id int) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM users WHERE id = ?)", id).Scan(&exists)
	return exists
}

// ExistsByEmail checks if the user record exists by email
func (s *UserStore) ExistsByEmail(email string) bool {
	var exists bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM users WHERE email = ?)", email).Scan(&exists)
	return exists
}

// Get the total number of posts
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *UserStore) Total() (int, error) {
	const op = "UserRepository.Total"
	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&total); err != nil {
		return -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of users", Operation: op, Err: err}
	}
	return total, nil
}
