package encryption

import (
	"crypto/md5"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

// Generate user token
func GenerateUserToken(name string, email string) string {
	emailHash := MD5Hash(email)

	hash := MD5Hash(name + time.Now().String() + "3de" + strconv.Itoa(rand.Intn(143 - 0) + 0) + emailHash)

	token := strconv.Itoa(rand.Intn(143 - 0) + 0) + hash + strconv.Itoa(rand.Intn(143 - 0) + 0)

	return token
}

// Generate generic random token
func GenerateEmailToken(email string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)

	if err != nil {
		return "", err
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

