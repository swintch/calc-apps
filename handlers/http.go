package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/swintch/calc"
)

func NewHTTPRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /add", NewHTTPHandler(&calc.Addition{}))
	mux.Handle("GET /sub", NewHTTPHandler(&calc.Subtraction{}))
	mux.Handle("GET /mul", NewHTTPHandler(&calc.Multiplication{}))
	mux.Handle("GET /div", NewHTTPHandler(&calc.Division{}))
	return mux
}

type HTTPHandler struct {
	calculator calc.Calculator
}

func NewHTTPHandler(calculator calc.Calculator) http.Handler {
	return &HTTPHandler{calculator: calculator}
}

func (this *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	value1, err := strconv.Atoi(query.Get("a"))
	if err != nil {
		http.Error(response, "a must be an integer value", http.StatusUnprocessableEntity)
		return
	}
	value2, err := strconv.Atoi(query.Get("b"))
	if err != nil {
		http.Error(response, "b must be an integer value", http.StatusUnprocessableEntity)
		return
	}
	result := this.calculator.Calculate(value1, value2)
	response.Header().Set("Content-Type", "plain/text; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(response, result)
	if err != nil {
		log.Println("Response write error: ", err)
	}
}
