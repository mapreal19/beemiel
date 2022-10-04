package matcher

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"regexp"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/gomega/types"
)

type setFlashMatcher struct {
	level   interface{}
	message interface{}
}

func SetFlash(level interface{}, message interface{}) types.GomegaMatcher {
	return &setFlashMatcher{
		level:   level,
		message: message,
	}
}

func (matcher *setFlashMatcher) Match(actual interface{}) (success bool, err error) {
	response, ok := actual.(*httptest.ResponseRecorder)
	if !ok {
		return false, fmt.Errorf("SetFlash matcher expects an http.Response")
	}

	flashLevel := getFlashLevel(response)
	flashMessage := getFlashMessage(response)

	success = reflect.DeepEqual(flashLevel, matcher.level) &&
		reflect.DeepEqual(flashMessage, matcher.message)

	return success, nil
}

func (matcher *setFlashMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf(
		"Expected response to set flash level to %#v and message to %#v, recieved response with flash level: %#v and message: %#v",
		matcher.level,
		matcher.message,
		getFlashLevel(actual.(*httptest.ResponseRecorder)),
		getFlashMessage(actual.(*httptest.ResponseRecorder)),
	)
}

func (matcher *setFlashMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf(
		"Expected response not to set flash level to %#v and message to %#v, recieved response with flash level: %#v and message: %#v",
		matcher.level,
		matcher.message,
		getFlashLevel(actual.(*httptest.ResponseRecorder)),
		getFlashMessage(actual.(*httptest.ResponseRecorder)),
	)
}

func findFlashCookie(response *httptest.ResponseRecorder) *http.Cookie {
	for _, cookie := range response.Result().Cookies() {
		if cookie.Name == web.BConfig.WebConfig.FlashName {
			return cookie
		}
	}
	panic("can't find cookie flash")
}

func getFlashLevel(response *httptest.ResponseRecorder) string {
	cookie := findFlashCookie(response)
	regexExpression, _ := regexp.Compile("%00(.*)%23" + web.BConfig.WebConfig.FlashSeparator)
	level := regexExpression.FindStringSubmatch(cookie.Value)[1]

	return level
}

func getFlashMessage(response *httptest.ResponseRecorder) string {
	cookie := findFlashCookie(response)
	regexExpression, _ := regexp.Compile(web.BConfig.WebConfig.FlashSeparator + "%23(.*)%00")
	message := regexExpression.FindStringSubmatch(cookie.Value)[1]

	messageUnescaped, _ := url.QueryUnescape(message)

	return messageUnescaped
}
