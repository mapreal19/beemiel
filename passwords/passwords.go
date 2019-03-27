package passwords

import (
	"encoding/hex"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/ed25519"
	"os"
	"time"
)

func NewResetToken(userID string, expiration time.Time) string {
	now := time.Now()

	jsonToken := paseto.JSONToken{
		Subject:    userID,
		IssuedAt:   now,
		NotBefore:  now,
		Expiration: expiration,
	}

	v2 := paseto.NewV2()
	token, err := v2.Sign(privateKey(), jsonToken, nil)

	if err != nil {
		panic(err)
	}

	return token
}

func GetUserIdFromToken(token string) (userID string, err error) {
	v2 := paseto.NewV2()
	jsonToken := paseto.JSONToken{}

	err = v2.Verify(token, publicKey(), &jsonToken, nil)
	if err != nil {
		return
	}

	err = jsonToken.Validate()
	if err != nil {
		return
	}

	userID = jsonToken.Subject
	return
}

func publicKey() ed25519.PublicKey {
	bytes, _ := hex.DecodeString(os.Getenv("PASETO_PUBLIC_KEY"))
	return ed25519.PublicKey(bytes)
}

func privateKey() ed25519.PrivateKey {
	bytes, _ := hex.DecodeString(os.Getenv("PASETO_PRIVATE_KEY"))
	return ed25519.PrivateKey(bytes)
}
