package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/events"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// AuthRepository defines methods for for Users to gain
// Auth for interacting with the database.
type AuthRepository interface {
	Authenticate(email string, password string) (domain.User, error)
	Logout(token string) (int, error)
	ResetPassword(token string, password string) error
	SendResetPassword(email string) error
	VerifyEmail(md5String string) error
	VerifyPasswordToken(token string) error
	CleanPasswordResets() error
}

// AuthStore defines the store for Authentication
type AuthStore struct {
	db *sqlx.DB
}

// newAuth - Construct
func newAuth(db *sqlx.DB) *AuthStore {
	return &AuthStore{
		db: db,
	}
}

// Authenticate compares the email & password for a match in the DB.
// Returns errors.NOTFOUND if the user is not found.
func (s *AuthStore) Authenticate(email string, password string) (domain.User, error) {
	const op = "AuthRepository.Authenticate"

	var u domain.User
	if err := s.db.Get(&u, "SELECT * FROM users WHERE email = ? LIMIT 1", email); err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "These credentials don't match our records.", Operation: op, Err: err}
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "These credentials don't match our records.", Operation: op, Err: err}
	}

	return u, nil
}

// Get the user by Token
// Logout checks to see if see if the
func (s *AuthStore) Logout(token string) (int, error) {
	const op = "AuthRepository.Logout"
	var u domain.User
	if err := s.db.Get(&u, "SELECT * FROM users WHERE token = ? LIMIT 1", token); err != nil {
		log.Info(err)
		return -1, fmt.Errorf("Could not get user with token: %v", token)
	}

	newToken := encryption.GenerateUserToken(u.FirstName + u.LastName, u.Email)
	q := "UPDATE users SET token = ?, updated_at = NOW() WHERE token = token"
	_, err := s.db.Exec(q, newToken)
	if err != nil {
		log.Error(err)
		return -1, fmt.Errorf("Could not update the user's token with the name: %v", u.FirstName + " " + u.LastName)
	}

	return u.Id, nil
}

// Reset the password by comparing token
func (s *AuthStore) ResetPassword(token string, password string) error {
	const op = "AuthRepository.ResetPassword"
	var rp domain.PasswordReset
	if err := s.db.Get(&rp, "SELECT * FROM password_resets WHERE token = ? LIMIT 1", token); err != nil {
		log.Info(err)
		return fmt.Errorf("Could not get user with token: %v", token)
	}

	hashedPassword, err := encryption.HashPassword(password)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not create the new passsword, please try again")
	}

	updateQ := "UPDATE users SET password = ? WHERE email = ?"
	_, err = s.db.Exec(updateQ, rp.Email, hashedPassword)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not update password, please try again")
	}

	if _, err := s.db.Exec("DELETE FROM password_resets WHERE token = ?", token); err != nil {
		log.Error(err)
		return fmt.Errorf("Could not delete from the reset passwords table")
	}

	return nil
}

// Reset the users password email, send email verification link
// and insert into the password_resets table
func (s *AuthStore) SendResetPassword(email string) error {
	const op = "AuthRepository.SendResetPassword"
	var u domain.User
	if err := s.db.Get(&u, "SELECT * FROM users WHERE email = ? LIMIT 1", email); err != nil {
		log.Info(err.Error())
		return fmt.Errorf("We couldn't find a user with that email address.")
	}

	// Generate an email token
	token, err := encryption.GenerateEmailToken(email)
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("Could not generate the user token.")
	}

	// Insert the email & token into the password resets table.
	q := "INSERT INTO password_resets (email, token, created_at) VALUES (?, ?, NOW())"
	_, err = s.db.Exec(q, email, token)
	if err != nil {
		log.Info(err)
		return fmt.Errorf("Could not insert into password resets table with the email: %v", email)
	}

	// Create a new reset password event.
	rp, err := events.NewResetPassword()
	if err != nil {
		log.Error(err)
		return err
	}

	// Send the reset password email
	err = rp.Send(&u, token)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not sent the reset password email")
	}

	return nil
}

// Verify the users email address based on the encryption hash string passed
func (s *AuthStore) VerifyEmail(md5String string) error {
	const op = "AuthRepository.VerifyEmail"

	var userVerified = struct{
		ID   	int		`db:"id"`
		Hash 	string	`db:"hash"`
	}{}

	fmt.Println(md5String)

	if err := s.db.Get(&userVerified, "SELECT id AS id, MD5(CONCAT(id, email)) AS hash FROM users WHERE MD5(CONCAT(id, email)) = ?", md5String); err != nil {
		log.Info(err)
		return fmt.Errorf("Could not get user from database")
	}

	q := "UPDATE users SET email_verified_at = NOW() WHERE ID = ?"
	_, err := s.db.Exec(q, userVerified.ID)
	if err != nil {
		log.Info(err)
		return fmt.Errorf("Could not users verifiy email address")
	}

	return nil
}

// Verify the token is valid from the password resets table
func (s *AuthStore) VerifyPasswordToken(token string) error {
	const op = "AuthRepository.VerifyPasswordToken"
	var pr domain.PasswordReset
	if err := s.db.Get(&pr, "SELECT * FROM password_resets WHERE token = ? LIMIT 1", token); err != nil {
		log.Info(err)
		return fmt.Errorf("We couldn't find a email matching that token")
	}

	return nil
}

// Verify the token is valid from the password resets table
func (s *AuthStore) CleanPasswordResets() error {
	if _, err := s.db.Exec("DELETE FROM password_resets WHERE created_at < (NOW() - INTERVAL 2 HOUR)"); err != nil {
		log.Error(err)
		return fmt.Errorf("Could not delete from the reset passwords table")
	}
	return nil
}
