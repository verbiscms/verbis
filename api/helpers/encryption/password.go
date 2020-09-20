package encryption

import "golang.org/x/crypto/bcrypt"

// Hash the users password
func HashPassword(password string) (string, error) {
	// Get current password in byte.
	bytePassword := []byte(password)

	// Hashing the password with the default cost of 10.
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}
