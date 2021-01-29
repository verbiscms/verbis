package auth

// Auth
//
// If the user is authenticated (logged in).
// Return false if the cookie was not found or not authenticated.
func (ns *Namespace) Auth() bool {
	cookie, err := ns.ctx.Cookie("verbis-session")

	if err != nil {
		return false
	}

	_, err = ns.deps.Store.User.GetByToken(cookie)
	if err != nil {
		return false
	}

	return true
}

// Admin
//
// If the user is authenticated (logged in) & an admin user.
// Returns false if the cookie was not found or not authenticated.
func (ns *Namespace) Admin() bool {
	cookie, err := ns.ctx.Cookie("verbis-session")

	if err != nil {
		return false
	}

	user, err := ns.deps.Store.User.GetByToken(cookie)
	if err != nil {
		return false
	}

	if user.Role.Id < 5 {
		return false
	}

	return true
}
