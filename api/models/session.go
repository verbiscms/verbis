package models

import (
	"cms/api/config"
	"cms/api/domain"
	"cms/api/helpers/encryption"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type SessionRepository interface {
	GetByKey(sessionKey string) (*domain.UserSession, error)
	Create(id int, email string) (string, error)
	Update(sessionKey string) error
	Delete(userId int) error
	Has(userId int) bool
	Check(userId int) error
}

type SessionStore struct {
	db *sqlx.DB
}

//Construct
func newSession(db *sqlx.DB) *SessionStore {
	return &SessionStore{
		db: db,
	}
}

// Get the user session by key
func (s *SessionStore) GetByKey(sessionKey string) (*domain.UserSession, error) {
	var us domain.UserSession
	if err := s.db.Get(&us, "SELECT * FROM user_sessions WHERE session_key = ? LIMIT 1", sessionKey); err != nil {
		log.Info(err)
		return &domain.UserSession{}, fmt.Errorf("Could not get user session with the session key: %v", sessionKey)
	}
	return &us, nil
}

// Create the user session once logged in, will return a
// unique session string and create a session within
// the user_sessions table.
func (s *SessionStore) Create(id int, email string) (string, error) {

	// Create the session token
	sessionToken := encryption.GenerateSessionToken(email)

	// Insert into the user_sessions table
	q := "INSERT INTO user_sessions (user_id, session_key, login_time, last_seen_time) VALUES (?, ?, NOW(), NOW())"
	_, err := s.db.Exec(q, id, sessionToken)
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Could not create the user session with the ID: %v", id)
	}

	return sessionToken, nil
}

// Update session and set last seen time
func (s *SessionStore) Update(sessionKey string) error {
	q := "UPDATE user_sessions SET last_seen_time = NOW() WHERE session_key = ?"
	_, err := s.db.Exec(q, sessionKey)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Could not update last seen time in the user's session")
	}

	return nil
}

// Delete the session of not valid
func (s *SessionStore) Delete(userId int) error {
	if _, err := s.db.Exec("DELETE FROM user_sessions WHERE user_id = ?", userId); err != nil {
		log.Error(err)
		return fmt.Errorf("Could not delete user session with user ID of: %v", userId)
	}
	return nil
}

// Check if the user has a session
func (s *SessionStore) Has(userId int) bool {
	var has bool
	_ = s.db.QueryRow("SELECT EXISTS (SELECT id FROM user_sessions WHERE user_id = ?)", userId).Scan(&has)
	return has
}

// Check if the session is valid
func (s *SessionStore) Check(userId int) error {
	var inactiveFor int
	err := s.db.Get(&inactiveFor,"SELECT TIMESTAMPDIFF(MINUTE, last_seen_time, NOW()) AS valid FROM user_sessions WHERE user_id = ?", userId)
	if err != nil {
		log.Error(err)
	}

	if inactiveFor > config.Admin.InactiveSessionTime {
		return fmt.Errorf("User has been in active for %v minutes", inactiveFor)
	}

	return nil
}