package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"
)

// GenerateToken создает случайный токен и возвращает его вместе с временем истечения срока действия.
func GenerateToken(ttlMinutes int) (token string, expiresAt time.Time, err error) {
	b := make([]byte, 32) // 256 bits
	_, err = rand.Read(b)
	if err != nil {
		return "", time.Time{}, err
	}
	token = base64.URLEncoding.EncodeToString(b)
	expiresAt = time.Now().Add(time.Duration(ttlMinutes) * time.Minute)
	return token, expiresAt, nil
}

// HashToken -> sha256 hex
func HashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}
