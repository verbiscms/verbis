package templates

/*
 * Auth
 * Functions for templates for anything else
 */

// isAuth If the user is authenticated (logged in)
// Returns false if the cookie was not found or not authenticated.
func (t *TemplateFunctions) isAuth() bool {
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

// isAdmin If the user is authenticated (logged in) & an admin user.
// Returns false if the cookie was not found or not authenticated.
func (t *TemplateFunctions) isAdmin() bool {
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
