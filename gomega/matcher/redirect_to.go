package matcher

import (
	"fmt"
	"net/http/httptest"
	"reflect"

	"github.com/onsi/gomega/types"
)

type redirectToMatcher struct {
	expected interface{}
}

func RedirectTo(expected interface{}) types.GomegaMatcher {
	return &redirectToMatcher{
		expected: expected,
	}
}

func (matcher *redirectToMatcher) Match(actual interface{}) (success bool, err error) {
	response, ok := actual.(*httptest.ResponseRecorder)
	if !ok {
		return false, fmt.Errorf("RedirectTo matcher expects an http.Response")
	}

	location := getLocation(response)
	success = reflect.DeepEqual(location, matcher.expected)

	return success, nil
}

func (matcher *redirectToMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf(
		"Expected response to redirect to %#v, recieved: %#v",
		matcher.expected,
		getLocation(actual.(*httptest.ResponseRecorder)),
	)
}

func (matcher *redirectToMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf(
		"Expected response not to redirect to %#v, recieved: %#v",
		matcher.expected,
		getLocation(actual.(*httptest.ResponseRecorder)),
	)
}

func getLocation(response *httptest.ResponseRecorder) string {
	return response.Header().Get("Location")
}
