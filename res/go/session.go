


// SessionRepository defines methods for Sessions to interact with the database
type SessionRepository interface {
	GetByKey(sessionKey string) (*domain.UserSession, error)
	Create(id int, email string) (string, error)
	Update(sessionKey string) error
	Delete(userId int) error
	Has(userId int) bool
	Check(userId int) error
}

// SessionRepository defines the data layer for Sessions
type SessionStore struct {
	db *sqlx.DB
}

// newSession - Construct
func newSession(db *sqlx.DB) *SessionStore {
	return &SessionStore{
		db: db,
	}
}

// GetByKey gets the user session by key
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no posts available.
func (s *SessionStore) GetByKey(sessionKey string) (*domain.UserSession, error) {
	const op = "SessionRepository.GetByKey"
	var us domain.UserSession
	if err := s.db.Get(&us, "SELECT * FROM user_sessions WHERE session_key = ? LIMIT 1", sessionKey); err != nil {
		return &domain.UserSession{}, &errors.Error{Code: errors.NOTFOUND, Message:  fmt.Sprintf("Could not get user session with the session key: %v", sessionKey), Operation: op}
	}
	return &us, nil
}

// Create the user session once logged in, will return a
// unique session string and create a session within
// the user_sessions table.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *SessionStore) Create(id int, email string) (string, error) {
	const op = "SessionRepository.Create"

	sessionToken := encryption.GenerateSessionToken(email)
	q := "INSERT INTO user_sessions (user_id, session_key, login_time, last_seen_time) VALUES (?, ?, NOW(), NOW())"
	_, err := s.db.Exec(q, id, sessionToken)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the user session with the ID: %v", id), Operation: op, Err: err}
	}

	return sessionToken, nil
}

// Update session and set last seen time
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *SessionStore) Update(sessionKey string) error {
	const op = "SessionRepository.Update"
	q := "UPDATE user_sessions SET last_seen_time = NOW() WHERE session_key = ?"
	_, err := s.db.Exec(q, sessionKey)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Could not update last seen time in the user's session", Operation: op, Err: err}
	}
	return nil
}

// Delete the session of not valid
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *SessionStore) Delete(userId int) error {
	const op = "SessionRepository.Delete"
	if _, err := s.db.Exec("DELETE FROM user_sessions WHERE user_id = ?", userId); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete user session with user ID of: %v", userId), Operation: op, Err: err}
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
// Returns errors.INTERNAL if the SQL query was invalid
// Returns errors.CONFLICT time has surpassed the configuration inactive session time variable.
func (s *SessionStore) Check(userId int) error {
	const op = "SessionRepository.Check"

	var inactiveFor int
	err := s.db.Get(&inactiveFor,"SELECT TIMESTAMPDIFF(MINUTE, last_seen_time, NOW()) AS valid FROM user_sessions WHERE user_id = ?", userId)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Unable to get user from sessions", Operation: op, Err: err}
	}

	if inactiveFor > config.Admin.InactiveSessionTime {
		return &errors.Error{Code: errors.CONFLICT, Message: fmt.Sprintf("User has been in active for %v minutes", inactiveFor), Operation: op, Err: err}
	}

	return nil
}