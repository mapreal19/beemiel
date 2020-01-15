package tokens

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"github.com/o1egl/paseto"
)

func NewAccessToken(payload string, expiration time.Time) string {
	now := time.Now()

	jsonToken := paseto.JSONToken{
		Subject:    payload,
		IssuedAt:   now,
		NotBefore:  now,
		Expiration: expiration,
	}

	v2 := paseto.NewV2()
	token, err := v2.Encrypt(symmetricKey(), jsonToken, nil)

	if err != nil {
		panic(err)
	}

	return token
}

func GetPayloadFromToken(token string) (payload string, err error) {
	v2 := paseto.NewV2()
	jsonToken := paseto.JSONToken{}

	v2.Decrypt(token, symmetricKey(), &jsonToken, nil)
	if err != nil {
		return
	}

	err = jsonToken.Validate()
	if err != nil {
		return
	}

	payload = jsonToken.Subject
	return
}

func symmetricKey() (key []byte) {
	key, _ = hex.DecodeString(os.Getenv("PASETO_SYMMETRIC_KEY"))
	return
}

func GenerateSymmetricKey() string {
	key := make([]byte, 32)
	_, _ = rand.Read(key)
	return hex.EncodeToString(key)
}
