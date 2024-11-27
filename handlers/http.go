package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

func NewHTTPRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("GET /add", NewHttpHandler(&calc.Addition{}))
	router.Handle("GET /sub", NewHttpHandler(&calc.Subtraction{}))
	router.Handle("GET /mul", NewHttpHandler(&calc.Multiplication{}))
	router.Handle("GET /div", NewHttpHandler(&calc.Division{}))
	return router
}

type HttpHandler struct {
	calculator calc.Calculator
}

func NewHttpHandler(calculator calc.Calculator) http.Handler {
	return &HttpHandler{
		calculator: calculator,
	}
}

func (this *HttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	value1, err := strconv.Atoi(query.Get("a"))
	if err != nil {
		http.Error(response, "a must be an integer", http.StatusUnprocessableEntity)
		return
	}
	value2, err := strconv.Atoi(query.Get("b"))
	if err != nil {
		http.Error(response, "b must be an integer", http.StatusUnprocessableEntity)
		return
	}
	result := this.calculator.Calculate(value1, value2)
	response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(response, result)
	if err != nil {
		log.Println("response write err:", err)
	}

}
