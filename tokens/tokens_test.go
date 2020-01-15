package tokens_test

import (
	"encoding/json"
	"os"
	"time"

	"github.com/mapreal19/beemiel/tokens"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("tokens", func() {
	Describe("NewAccessToken", func() {
		It("returns PASETO token", func() {
			_ = os.Setenv("PASETO_SYMMETRIC_KEY", tokens.GenerateSymmetricKey())
			userID := "1"
			expiration := time.Now().Add(2 * time.Hour)

			result := tokens.NewAccessToken(userID, expiration)

			Expect(result).To(HavePrefix("v2.local."))
		})

		Context("when having a wrong PASETO key", func() {
			It("panics", func() {
				_ = os.Setenv("PASETO_SYMMETRIC_KEY", "wannabekey")
				userID := "1"
				expiration := time.Now().Add(2 * time.Hour)

				result := func() {
					tokens.NewAccessToken(userID, expiration)
				}

				Expect(result).Should(Panic())
			})
		})
	})

	Describe("GetPayloadFromToken", func() {
		It("returns json payload", func() {
			_ = os.Setenv("PASETO_SYMMETRIC_KEY", tokens.GenerateSymmetricKey())
			payload := `{"base":"BTC","currency":"USD","amount":12334.87}`

			expiration := time.Now().Add(2 * time.Hour)
			token := tokens.NewAccessToken(payload, expiration)

			result, err := tokens.GetPayloadFromToken(token)

			type money struct {
				Base     string  `json:"base"`
				Currency string  `json:"currency"`
				Amount   float32 `json:"amount"`
			}
			var response money

			err = json.Unmarshal([]byte(result), &response)

			Expect(err).To(BeNil())
			Expect("BTC").To(Equal(response.Base))
			Expect("USD").To(Equal(response.Currency))
		})

		Context("when token is expired", func() {
			It("returns error", func() {
				_ = os.Setenv("PASETO_SYMMETRIC_KEY", tokens.GenerateSymmetricKey())
				userID := "1"
				expiration := time.Now().Add(-2 * time.Hour)
				token := tokens.NewAccessToken(userID, expiration)

				_, err := tokens.GetPayloadFromToken(token)

				Expect(err.Error()).To(Equal("token has expired: token validation error"))
			})
		})
	})
})
