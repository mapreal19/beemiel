package passwords_test

import (
	"encoding/hex"
	"github.com/mapreal19/beemiel/passwords"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/ed25519"
	"os"
	"time"
)

var _ = Describe("passwords", func() {
	Describe("NewResetToken", func() {
		It("returns PASETO token", func() {
			_, privateKey, _ := ed25519.GenerateKey(nil)
			_ = os.Setenv("PASETO_PRIVATE_KEY", hex.EncodeToString(privateKey))
			userID := "1"
			expiration := time.Now().Add(2 * time.Hour)

			result := passwords.NewResetToken(userID, expiration)

			Expect(result).To(HavePrefix("v2.public."))
		})

		Context("when having a wrong PASETO key", func() {
			It("panics", func() {
				_ = os.Setenv("PASETO_PRIVATE_KEY", "wannabekey")
				userID := "1"
				expiration := time.Now().Add(2 * time.Hour)

				result := func() {
					passwords.NewResetToken(userID, expiration)
				}

				Expect(result).Should(Panic())
			})
		})
	})

	Describe("GetUserIDFromToken", func() {
		It("returns user id", func() {
			publicKey, privateKey, _ := ed25519.GenerateKey(nil)
			_ = os.Setenv("PASETO_PUBLIC_KEY", hex.EncodeToString(publicKey))
			_ = os.Setenv("PASETO_PRIVATE_KEY", hex.EncodeToString(privateKey))
			userID := "1"
			expiration := time.Now().Add(2 * time.Hour)
			token := passwords.NewResetToken(userID, expiration)

			result, err := passwords.GetUserIdFromToken(token)

			Expect(err).To(BeNil())
			Expect(result).To(Equal(userID))
		})

		Context("when token is expired", func() {
			It("returns error", func() {
				publicKey, privateKey, _ := ed25519.GenerateKey(nil)
				_ = os.Setenv("PASETO_PUBLIC_KEY", hex.EncodeToString(publicKey))
				_ = os.Setenv("PASETO_PRIVATE_KEY", hex.EncodeToString(privateKey))
				userID := "1"
				expiration := time.Now().Add(-2 * time.Hour)
				token := passwords.NewResetToken(userID, expiration)

				_, err := passwords.GetUserIdFromToken(token)

				Expect(err.Error()).To(Equal("token has expired: token validation error"))
			})
		})
	})
})

