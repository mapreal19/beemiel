package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Response struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

const recaptchaURL = "https://www.google.com/recaptcha/api/siteverify"

// Validates if the token from the client is valid
func Validate(token string) Response {
	var recaptchaResponse Response

	if os.Getenv("RECAPTCHA_MOCK_SERVER") == "true" {
		recaptchaResponse.Success = true
		return recaptchaResponse
	}

	httpResponse, err := http.PostForm(
		recaptchaURL,
		url.Values{"secret": {os.Getenv("RECAPTCHA_SECRET_KEY")}, "response": {token}},
	)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &recaptchaResponse)
	if err != nil {
		panic(err)
	}

	return recaptchaResponse
}
