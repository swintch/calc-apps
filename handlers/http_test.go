package handlers

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestHTTPHandler_ServeHTTP(t *testing.T) {

}

func assertHTTP(t *testing.T, method string, target string, statusCode int, responseBody string) {
	t.Run(fmt.Sprintf("%s %s", method, target), func(t *testing.T) {
		router := NewHTTPRouter()
		response := httptest.NewRecorder()
		request := httptest.NewRequest(method, target, nil)
		router.ServeHTTP(response, request)
		AssertEquals(t, statusCode, response.Code)
		AssertEquals(t, responseBody, response.Body.String())
		AssertEquals(t, response.Header().Get("Content-Type"), "plain/text; charset=utf-8")
	})
}
