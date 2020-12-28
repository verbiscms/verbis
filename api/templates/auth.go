package templates

// auth
//
// If the user is authenticated (logged in).
// Return false if the cookie was not found or not authenticated.
func (t *TemplateManager) auth() bool {
	cookie, err := t.gin.Cookie("verbis-session")

	if err != nil {
		return false
	}

	_, err = t.store.User.GetByToken(cookie)
	if err != nil {
		return false
	}

	return true
}

// admin
//
// If the user is authenticated (logged in) & an admin user.
// Returns false if the cookie was not found or not authenticated.
func (t *TemplateManager) admin() bool {
	cookie, err := t.gin.Cookie("verbis-session")

	if err != nil {
		return false
	}

	user, err := t.store.User.GetByToken(cookie)
	if err != nil {
		return false
	}

	if user.Role.Id < 5 {
		return false
	}

	return true
}
