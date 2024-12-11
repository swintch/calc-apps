package shuttle

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	InputModel interface {
		Bind(request *http.Request) error
	}
	InputError struct {
		Fields  []string `json:"fields,omitempty"`
		Name    string   `json:"name,omitempty"`
		Message string   `json:"message,omitempty"`
	}
)

func (this InputError) Error() string {
	return this.Message
}

type (
	Processor interface {
		Process(ctx context.Context, v any) any
	}
	SerializeResult struct {
		StatusCode int
		Content    any
	}
)

func NewHandler(input func() InputModel, processor func() Processor) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		inputModel := input()
		err := inputModel.Bind(request)
		if err != nil {
			http.Error(response, "Bad Request", http.StatusBadRequest)
			return
		}
		viewModel := processor().Process(context.Background(), inputModel)
		jsonBytes, err := json.Marshal(viewModel)
		if err != nil {
			http.Error(response, "JSON parse Error: ", http.StatusInternalServerError)
			return
		}
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.WriteHeader(http.StatusOK)
		_, err = fmt.Fprint(response, string(jsonBytes))
		if err != nil {
			log.Println("Response Error:", err)
		}
	})
}
