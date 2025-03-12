package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"go-backend-api/global"
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
func HashPassword(password string, salt string, secretKey string) string {
	saltedPassword := password + salt + secretKey
	hashPasword := sha256.Sum256(([]byte(saltedPassword)))
	return hex.EncodeToString(hashPasword[:])
}

func MatchingPassword(storeHash string, password string, salt string) bool {
	hashPassword := HashPassword(password, salt, global.Config.JWT.SECRET_KEY)
	return storeHash == hashPassword
}
