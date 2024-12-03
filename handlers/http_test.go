package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHandler_ServeHTTP(t *testing.T) {
	assertHTTP(t, "GET", "/add?a=3&b=4", http.StatusOK, "7")
	assertHTTP(t, "GET", "/sub?a=4&b=2", http.StatusOK, "2")
	assertHTTP(t, "GET", "/mul?a=3&b=4", http.StatusOK, "12")
	assertHTTP(t, "GET", "/div?a=6&b=3", http.StatusOK, "2")
	assertHTTP(t, "GET", "/add?a=NaN&b=3", http.StatusUnprocessableEntity, "a must be an integer value\n")
	assertHTTP(t, "GET", "/add?a=4&b=NaN", http.StatusUnprocessableEntity, "b must be an integer value\n")
	assertHTTP(t, "POST", "/add?a=4&b=3", http.StatusMethodNotAllowed, "Method Not Allowed\n")
	assertHTTP(t, "GET", "/ASDDF?a=4&b=3", http.StatusNotFound, "404 page not found\n")
}

func assertHTTP(t *testing.T, method string, target string, statusCode int, responseBody string) {
	t.Run(fmt.Sprintf("%s %s", method, target), func(t *testing.T) {
		router := NewHTTPRouter()
		response := httptest.NewRecorder()
		request := httptest.NewRequest(method, target, nil)
		router.ServeHTTP(response, request)
		AssertEquals(t, statusCode, response.Code)
		AssertEquals(t, responseBody, response.Body.String())
		AssertEquals(t, response.Header().Get("Content-Type"), "text/plain; charset=utf-8")
	})
}
