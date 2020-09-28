package encryption

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword gets the password in byte format and generates
// a hashed password with the default cost of 10.
// Returns errors.INTERNAL if the bcrypt failed to generate
// from password.
func HashPassword(password string) (string, error) {
	const op = "encryption.HashPassword"
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not hash the password with the string: %s", password), Operation: op, Err: err}
	}
	return string(hashedPassword), err
}
