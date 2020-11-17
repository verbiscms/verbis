package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash generates a a random MD% based on the string given.
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
