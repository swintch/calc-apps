package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func assertHTTP(this *HTTPHandlerFixture, method string, target string, statusCode int, responseBody string) {
	router := NewHTTPRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(method, target, nil)
	router.ServeHTTP(response, request)
	this.So(statusCode, should.Equal, response.Code)
	this.So(response.Body.String(), should.Equal, responseBody)
	this.So(response.Header().Get("Content-Type"), should.Equal, "text/plain; charset=utf-8")
}

func TestHTTPHandler(t *testing.T) {
	gunit.Run(new(HTTPHandlerFixture), t)
}

type HTTPHandlerFixture struct {
	*gunit.Fixture
}

func (this *HTTPHandlerFixture) TestServeHTTP() {
	assertHTTP(this, "GET", "/add?a=3&b=4", http.StatusOK, "7")
	assertHTTP(this, "GET", "/sub?a=4&b=2", http.StatusOK, "2")
	assertHTTP(this, "GET", "/mul?a=3&b=4", http.StatusOK, "12")
	assertHTTP(this, "GET", "/div?a=6&b=3", http.StatusOK, "2")
	assertHTTP(this, "GET", "/add?a=NaN&b=3", http.StatusUnprocessableEntity, "a must be an integer value\n")
	assertHTTP(this, "GET", "/add?a=4&b=NaN", http.StatusUnprocessableEntity, "b must be an integer value\n")
	assertHTTP(this, "POST", "/add?a=4&b=3", http.StatusMethodNotAllowed, "Method Not Allowed\n")
	assertHTTP(this, "GET", "/ASDDF?a=4&b=3", http.StatusNotFound, "404 page not found\n")
}
