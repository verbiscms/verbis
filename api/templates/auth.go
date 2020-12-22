package templates

// auth
//
// If the user is authenticated (logged in) return false
// if the cookie was not found or not authenticated.
func (t *TemplateFunctions) auth() bool {
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
// If the user is authenticated (logged in) & an admin user. Returns
// false if the cookie was not found or not authenticated.
func (t *TemplateFunctions) admin() bool {
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
