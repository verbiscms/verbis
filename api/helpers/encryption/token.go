package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

// GenerateUserToken generates a new user token based on name & email.
func GenerateUserToken(name string, email string) string {
	emailHash := MD5Hash(email)
	hash := MD5Hash(name + time.Now().String() + "3de" + strconv.Itoa(rand.Intn(143 - 0) + 0) + emailHash)
	token := strconv.Itoa(rand.Intn(143 - 0) + 0) + hash + strconv.Itoa(rand.Intn(143 - 0) + 0)
	return token
}

// GenerateEmailToken generates a token based on the email given.
// Returns errors.INTERNAL if the bcrypt failed to generate
// from password.
func GenerateEmailToken(email string) (string, error) {
	const op = "encryption.GenerateEmailToken"
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not generate the email token with the email: %s", email), Operation: op, Err: err}
	}
	return MD5Hash(string(hash)), nil
}

// GenerateToken returns a unique token based on the provided email string
func GenerateSessionToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
	}
	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

