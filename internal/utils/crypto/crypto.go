package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}

// GeneraSalt a random salt
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

// hash password hash
func HashPassword(password string, salt string) string {
	saltedPassword := password + salt
	hashPasword := sha256.Sum256(([]byte(saltedPassword)))
	return hex.EncodeToString(hashPasword[:])
}

func MatchingPassword(storeHash string, password string, salt string) bool {
	hashPassword := HashPassword(password, salt)
	return storeHash == hashPassword
}
